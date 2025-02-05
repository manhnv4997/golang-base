package utils

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func Debug(rows *sql.Rows) {
	// Lấy danh sách các cột
	columns, err := rows.Columns()
	if err != nil {
		log.Fatal(err)
	}

	// Tạo một slice để chứa các giá trị của mỗi cột
	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))

	// Gán con trỏ cho mỗi giá trị
	for i := range values {
		valuePtrs[i] = &values[i]
	}

	// Duyệt qua từng hàng
	for rows.Next() {
		// Quét dữ liệu của hàng hiện tại vào các con trỏ
		err := rows.Scan(valuePtrs...)
		if err != nil {
			log.Fatal(err)
		}

		// Tạo một map để lưu thông tin của hàng hiện tại
		rowData := make(map[string]interface{})
		for i, colName := range columns {
			var v interface{}
			val := values[i]

			// Kiểm tra kiểu dữ liệu của cột
			switch val.(type) {
			case []byte:
				v = string(val.([]byte)) // Chuyển []byte thành string
			default:
				v = val
			}

			rowData[colName] = v
		}

		// In dữ liệu của hàng hiện tại
		fmt.Println(rowData)
	}

	// Kiểm tra lỗi khi duyệt qua rows
	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
}

func RoutePath(prefix string, path string) string {
	var sb strings.Builder

	sb.WriteString("/")
	sb.WriteString(prefix)
	if path != "" {
		sb.WriteString("/")
		sb.WriteString(path)
	}

	fmt.Println(sb.String())

	return sb.String()
}

func LoadEnv() {
	envPath := "env/.env"

	if err := godotenv.Load(envPath); err != nil {
		log.Println("Không tìm thấy file .env hoặc lỗi khi tải.")
	} else {
		_ = godotenv.Load(envPath)
	}
}

func GetEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}
