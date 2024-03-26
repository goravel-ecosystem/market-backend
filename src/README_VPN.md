# VPN

We use WireGuard VPN to build a bridge between local and staging environment. The document will guide you to use 
WireGuard VPN to connect to the staging environment.

## Install WireGuard

### Debian 11 / Ubuntu 20

```shell
sudo apt update
sudo apt install wireguard wireguard-tools
```

### MacOS

```shell
brew install wireguard-tools
```

### Other OS

Please check the [official document](https://www.wireguard.com/install/).

## Configure WireGuard VPN Server

We use Ubuntu 20 as the VPN server, you can use other OS as the VPN server, but the configuration may be different.

If you only focus on the client configuration, you can skip this section and jump to [the next section](#configure-wireguard-vpn-client).

```shell
# Create a configuration folder
sudo mkdir /etc/wireguard

# Generate a private key and a public key
wg genkey | sudo tee /etc/wireguard/privatekey | wg pubkey | sudo tee /etc/wireguard/publickey

# Create a network configuration file
sudo vim /etc/wireguard/wg0.conf
```

The content of `wg0.conf`, please remove the comments if you want to copy the content to your configuration file:

```text
[Interface]
# The IP address of server, will be used by the client to connect to the server 
Address = 10.0.0.1/24
# The port of server, will be used by the client to connect to the server
ListenPort = 51820
# The private key of server, please replace SERVER_PRIVATE_KEY with the content of /etc/wireguard/privatekey
PrivateKey = SERVER_PRIVATE_KEY
SaveConfig = true
```

> If the firewall is enabled, please open the port `51820`. If you are using a cloud server, you need to modify the 
> security group to open the port `UDP/51820`.

Modify the permission of the private key and the configuration file:

```shell
sudo chmod 600 /etc/wireguard/{privatekey,wg0.conf}
```

Start the WireGuard service

```shell
sudo wg-quick up wg0
```

The console will print:

```text
[#] ip link add wg0 type wireguard
[#] wg setconf wg0 /dev/fd/63
[#] ip -4 address add 10.0.0.1/24 dev wg0
[#] ip link set mtu 1420 up dev wg0
```

You can use `wg` command to check the status of the VPN server:

```shell
sudo wg show wg0
```

Configure the WireGuard service to start automatically:

```shell
sudo systemctl enable wg-quick@wg0
```

### Errors

1. `/usr/bin/wg-quick: line 32: resolvconf: command not found`

```shell
sudo apt install openresolv
```

## Configure WireGuard VPN Client

It's very simple to configure the WireGuard VPN client, you only need to generate a private key and a public key, 
add your public key and IP address to the server, and add the server's public key and IP address to your 
configuration file.

The following is an example of MacOS, you can use the same way to configure the client on other OS.

```shell
# Install WireGuard VPN as shown above:
brew install wireguard-tools

# Create a configuration folder
sudo mkdir /usr/local/etc/wireguard

# Generate a private key and a public key
wg genkey | tee /usr/local/etc/wireguard/privatekey | wg pubkey | tee /usr/local/etc/wireguard/publickey

# Create a network configuration file
sudo vim /usr/local/etc/wireguard/wg0.conf
```

The content of `wg0.conf`, send your public key to the server administrator ,and he will give you the server public 
key, server IP, server port, and client IP. Please remove the comments if you want to copy the content to your configuration file:

```text
[Interface]
# The private key of generated above, please replace CLIENT_PRIVATE_KEY with the content of /usr/local/etc/wireguard/privatekey
PrivateKey = CLIENT_PRIVATE_KEY
# The port of client
ListenPort = 51820
# The IP address of client, will be used by the server to connect to the client
Address = CLIENT_IP/32
DNS = 8.8.8.8
MTU = 1420

[Peer]
# The public key of server
PublicKey = SERVER_PUBLIC_KEY
# The IP address and port of server
Endpoint = SERVER_IP:SERVER_PORT
# The allowed IP address of server
AllowedIPs = 10.0.0.1/24
PersistentKeepalive = 25
```

Once the serve administrator has added your public key to the server, you can start the WireGuard service:

```shell
sudo wg-quick up wg0
````

The console will print:

```text
[#] wireguard-go utun
[+] Interface for wg0 is utun6
[#] wg setconf utun6 /dev/fd/63
[#] ifconfig utun6 inet 10.0.0.2/24 10.0.0.2 alias
[#] ifconfig utun6 mtu 1420
[#] ifconfig utun6 up
[#] route -q -n add -inet 10.0.0.0/24 -interface utun6
[#] networksetup -getdnsservers Wi-Fi
[#] networksetup -getsearchdomains Wi-Fi
[#] networksetup -getdnsservers Thunderbolt Bridge
[#] networksetup -getsearchdomains Thunderbolt Bridge
[#] networksetup -setdnsservers Wi-Fi 8.8.8.8
[#] networksetup -setsearchdomains Wi-Fi Empty
[#] networksetup -setdnsservers Thunderbolt Bridge 8.8.8.8
[#] networksetup -setsearchdomains Thunderbolt Bridge Empty
[+] Backgrounding route monitor
```

Then, you can test if you can connect to the server:

```shell
nc -zv 10.0.0.1 4000
```

If success, you will get the following output:

```text
Connection to 10.0.0.1 port 4000 [tcp/terabase] succeeded!
```

## Add the client to the server

The following content is handled by the server administrator. If you only focus on the client configuration, you can 
skip this section.

```shell
# Create a new folder for clients
mkdir /etc/wireguard/clients/

# Add the client public key to the server, please replace CLIENT_PUBLIC_KEY and CLIENT_NAME with your client public 
# key and name
echo CLIENT_PUBLIC_KEY > /etc/wireguard/clients/CLIENT_NAME

# Add the client to wg0.conf, please replace CLIENT_NAME with your client name
sudo wg set wg0 peer $(cat /etc/wireguard/clients/CLIENT_NAME) allowed-ips 10.0.0.0/24

# Check the server and the clients status
sudo wg

# If no problems, save the configuration
sudo wg-quick save wg0
```

### Remove the client from the server

```shell
sudo wg set wg0 peer $(cat /etc/wireguard/clients/CLIENT_NAME) remove

sudo wg-quick save wg0
```

### How to check the wireguard logo

```shell
# Check support debug, should print: debugfs on /sys/kernel/debug type debugfs (rw,relatime)
mount | grep debug

# Enable the debug
modprobe wireguard
echo module wireguard +p > /sys/kernel/debug/dynamic_debug/control

# Get log
dmesg -wH
```
