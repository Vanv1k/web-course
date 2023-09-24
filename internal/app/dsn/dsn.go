package dsn

import (
	"encoding/json"
	"fmt"
	"os"
)

type Configuration struct {
	DBHost string `json:"DBHost"`
	DBPort string `json:"DBPort"`
	DBUser string `json:"DBUser"`
	DBPass string `json:"DBPass"`
	DBName string `json:"DBName"`
}

// SetConnectionString собирает DSN строку
func SetConnectionString() string {
	// Открываем файл конфигурации
	file, err := os.Open("./internal/app/dsn/db_config.json")
	if err != nil {
		fmt.Println("Error opening config file:", err)
		return fmt.Sprintf("")
	}
	defer file.Close()

	// Декодируем JSON из файла в структуру Configuration
	decoder := json.NewDecoder(file)
	config := Configuration{}
	err = decoder.Decode(&config)
	if err != nil {
		fmt.Println("Error decoding config JSON:", err)
		return fmt.Sprintf("")
	}
	// connectionString := "user=postgres password=123987 dbname=IT_services host=localhost port=5432 sslmode=disable"
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", config.DBHost, config.DBPort, config.DBUser, config.DBPass, config.DBName)
}
