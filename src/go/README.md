# Go

We move each go.mod of microservices to the `/src/go` folder, there are some benefits:

1. We can use `go mod tidy` to manage the dependencies of all microservices;
2. We can upgrade goravel/framework uniformly, imagine that we have 10 microservices, they are using different 
   versions of goravel/framework and there are some break changes, we have to use different interface in different 
   microservices, it's a nightmare to maintain them;
3. We can run Github Action to check the code quality of all microservices, it's very convenient to find the 
   problems;
4. It's easy to use the interface in the `proto` folder;