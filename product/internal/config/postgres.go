package config

import "fmt"

func GetPostgresConnectionString(DBType, DBUser, DBPassword, DBHost, DBPort, DBName string) string {
	dsn := fmt.Sprintf("%s://%s:%s@%s:%s/%s",
		DBType, DBUser, DBPassword, DBHost, DBPort, DBName)
	return dsn
}
