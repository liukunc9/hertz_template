package initialize

import (
	"github.com/bwmarrin/snowflake"
	"github.com/liukunc9/hertz_template/global"
)

func initSnowflake() {
	// Create a new Node with a Node number of 1
	node, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}
	global.Snowflake = node
}
