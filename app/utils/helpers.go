package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"database/sql"
	"encoding/hex"
	"fmt"
	"log"
	"net/url"
	"os"
	"sort"
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
	if path != "" {
		sb.WriteString(prefix)
	}
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
	LoadEnv()

	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue
}

func ValidateHMAC(query url.Values, shopifyApiSecret, receivedHMAC string) bool {
	// 1. Tạo chuỗi truy vấn chuẩn hóa (canonical query string)
	var keys []string
	for key := range query {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	var params []string
	for _, key := range keys {
		if key == "hmac" { // Loại bỏ tham số hmac khỏi chuỗi truy vấn
			continue
		}
		params = append(params, key+"="+query.Get(key))
	}
	queryString := strings.Join(params, "&")

	// 2. Tính toán HMAC sử dụng SHA256 và API secret
	h := hmac.New(sha256.New, []byte(shopifyApiSecret))

	h.Write([]byte(queryString))
	calculatedHMAC := hex.EncodeToString(h.Sum(nil))

	// 3. So sánh HMAC đã tính toán với HMAC nhận được (chú ý đến so sánh an toàn về thời gian)
	// Decode HMAC nhận được từ hex string
	decodedReceivedHMAC, err := hex.DecodeString(receivedHMAC)
	if err != nil {
		return false // Nếu HMAC nhận được không phải là hex hợp lệ, trả về false
	}
	decodedCalculatedHMAC, err := hex.DecodeString(calculatedHMAC)
	if err != nil {
		return false // Nếu HMAC tính toán được không phải là hex hợp lệ (về lý thuyết không xảy ra), trả về false
	}

	// Sử dụng subtle.ConstantTimeCompare để so sánh an toàn về thời gian, tránh tấn công thời gian
	if subtle.ConstantTimeCompare(decodedCalculatedHMAC, decodedReceivedHMAC) == 1 {
		return true
	}
	return false
}
