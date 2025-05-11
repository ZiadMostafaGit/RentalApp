package main

import (
	"database/sql"
	"fmt"
	"log"
	"net"

	"github.com/ZiadMostafaGit/rental-app/internal/config"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	cfg := config.LoadConfig() // Call the exported function

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true",
		cfg.DBUser, cfg.DBPassword, cfg.DBHost, cfg.DBPort, cfg.DBName,
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("failed to connect to the database:", err)
	}
	defer db.Close()

}
