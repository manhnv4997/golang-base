package mysql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// Connect mở kết nối đến MySQL
func Connect(user, password, host, port, dbname string) {
	config := LoadConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", config.User, config.Password, config.Host, config.Port, config.DBName)

	var err error
	DB, err = sql.Open("mysql", dsn)

	if err != nil {
		log.Fatal("Lỗi kết nối MYSQL:", err)
	}

	// Kiểm tra kết nối

	if err = DB.Ping(); err != nil {
		log.Fatal("Không thể ping MySQL:", err)
	}

	fmt.Println("✅ Kết nối MySQL thành công!")
}

func Close() {
	if DB != nil {
		DB.Close()
		fmt.Println("🔴 Đã đóng kết nối MySQL")
	}
}
