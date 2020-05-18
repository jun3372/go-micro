package tool

import (
	"log"
	"math/rand"
	"strconv"

	"github.com/bwmarrin/snowflake"
)

var node *snowflake.Node

func init() {
	var err error
	node, err = snowflake.NewNode(int64(rand.Intn(1022) + 1))
	if err != nil {
		log.Println(err)
		return
	}
}

func GenerateUUID() string {
	// Generate a snowflake ID.
	id := node.Generate()
	return id.String()
}

func GenerateUUIDInt() int64 {
	id := GenerateUUID()
	uuid, _ := strconv.ParseInt(id, 10, 64)
	return uuid
}
