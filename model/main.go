package Model

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var DB *sql.DB

func Init() {
	// scan env
	envFileLocation := ".env"
	godotenv.Load(envFileLocation)

	// Inisialisasi koneksi database
	// dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
	// 	os.Getenv("DB_USER"),
	// 	os.Getenv("DB_PASSWORD"),
	// 	os.Getenv("DB_HOST"),
	// 	os.Getenv("DB_PORT"),
	// 	os.Getenv("DB_NAME"),
	// )

	dsn := fmt.Sprintf("root:@tcp(localhost:3306)/golang_login_api")

	var err error
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	// Tes koneksi ke database
	// err = db.Ping()
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// Set konfigurasi koneksi
	DB.SetMaxOpenConns(10)
	DB.SetMaxIdleConns(5)
	DB.SetConnMaxLifetime(5 * time.Minute)
}

func IsUsernameExists(username string, password string) bool {
	var count int

	row := DB.QueryRow("SELECT COUNT(*) FROM user WHERE username=? and pass=?", username, password)
	err := row.Scan(&count)
	if err != nil {
		log.Println("Error checking username:", err)
		return false
	}

	return count > 0
}
