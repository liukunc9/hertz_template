package initialize

func DoInit() {
	initViper()
	initDB()
	initSnowflake()
}
