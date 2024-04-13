package models

import (
	"log"
	"math/rand"

	"github.com/bwmarrin/snowflake"
	"github.com/goravel/framework/database/orm"
)

var snowflakeInstance *snowflake.Node

func init() {
	var err error
	snowflakeInstance, err = snowflake.NewNode(int64(rand.Intn(1023)))
	if err != nil {
		log.Panicf("snowflake.NewNode err: %v", err)
	}
}

type UUIDModel struct {
	ID uint64 `gorm:"primaryKey"`
	orm.Timestamps
}

func (r UUIDModel) GetID() uint64 {
	return uint64(snowflakeInstance.Generate())
}
