package Learn

// set CGO_ENABLED=1

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func sqllite() {
	// Buka koneksi ke database
	db, err := sql.Open("sqlite3", "C:/Users/HAKIM/login_api.db")
	if err != nil {
		// Tangani kesalahan pembukaan koneksi
		fmt.Print("koneksi gagal: %v", err)
		return // Hentikan eksekusi handler HTTP
	}
	defer db.Close()

	// Query database
	rows, err := db.Query("SELECT id FROM user")
	if err != nil {
		fmt.Print("cek")
		fmt.Println("Error executing query:", err)
		return
	}
	defer rows.Close()

	// Mengambil dan menampilkan data
	// for rows.Next() {
	// 	var id int
	// 	if err := rows.Scan(&id); err != nil {
	// 		fmt.Print("cek")
	// 	}
	// 	// fmt.Print("ID:%s", id)
	// 	fmt.Println("ID", id)
	// }

	fmt.Print("cek")
}
