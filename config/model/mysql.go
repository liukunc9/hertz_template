package model

type Mysql struct {
	Username string `json:"username" mapstructure:"username"`
	Password string `json:"password" mapstructure:"password"`
	Dbname   string `json:"db-name" mapstructure:"db-name"`
	Host     string `json:"host" mapstructure:"host"`
	Port     int32  `json:"port" mapstructure:"port"`

	TablePrefix string `json:"table-prefix" mapstructure:"table-prefix"`
	LogMode     string `json:"log-mode" mapstructure:"log-mode"`
}
