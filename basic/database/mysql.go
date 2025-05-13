package database

var connector string

func init() {
	connector = "mysql"
}

func GetDatabase() string {
	return connector
}