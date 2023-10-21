package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
)

func main() {
	g := gen.NewGenerator(gen.Config{
		OutPath:      "../dao",                                                           // 定义 dao 文件输出目录
		ModelPkgPath: "../dao/entity",                                                    // 定义 model 文件输出目录
		Mode:         gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	db, _ := gorm.Open(
		mysql.Open("root:mysql123@(127.0.0.1:3306)/go_demo?charset=utf8mb4&parseTime=True&loc=Local"),
		&gorm.Config{
			NamingStrategy: schema.NamingStrategy{
				SingularTable: true, // 单数表名
			},
		})
	g.UseDB(db) // reuse your gorm db

	userModel := g.GenerateModel("sys_user")
	g.ApplyBasic(userModel)

	g.Execute()
}
