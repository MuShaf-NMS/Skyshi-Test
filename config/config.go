package config

type Config struct {
	DB_User string
	DB_Pass string
	DB_Name string
	DB_Host string
	DB_Port string
}

func GetConfig() *Config {
	return &Config{
		DB_User: getVariable("MYSQL_USER"),
		DB_Pass: getVariable("MYSQL_PASSWORD"),
		DB_Name: getVariable("MYSQL_DBNAME"),
		DB_Host: getVariable("MYSQL_HOST"),
		DB_Port: getVariable("MYSQL_PORT"),
	}
}
