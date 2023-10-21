package global

import (
	"github.com/bwmarrin/snowflake"
	"github.com/liukunc9/hertz_template/config/model"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var (
	DB        *gorm.DB
	Viper     *viper.Viper
	Config    model.Server
	Snowflake *snowflake.Node
)
