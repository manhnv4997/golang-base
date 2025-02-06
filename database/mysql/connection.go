package mysql

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

// Connect mở kết nối đến MySQL
func Connect() {
	DBConfig := LoadDBConfig()

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=True", DBConfig.User, DBConfig.Password, DBConfig.Host, DBConfig.Port, DBConfig.DBName)

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
