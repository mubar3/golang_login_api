package Learn

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func setting_db() {
	envFileLocation := "../.env"
	err := godotenv.Load(envFileLocation)
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	// fmt.Println(os.Getenv("DB_PORT"))

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
	)

	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Failed to connect to MySQL:", err)
	}
	defer db.Close()

	rows, err := db.Query("use " + os.Getenv("DB_NAME"))
	if err != nil {
		// log.Fatal("Failed to query MySQL:", err)
	}
	defer rows.Close()

}
