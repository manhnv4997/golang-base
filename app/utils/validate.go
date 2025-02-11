package utils

import (
	"crypto/hmac"
	"crypto/sha256"
	"crypto/subtle"
	"encoding/hex"
	"net/url"
	"sort"
	"strings"
)

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
