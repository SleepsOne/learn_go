package main

import (
	"fmt"
)

// 1. Tạo custom error struct
type DivideError struct {
	Dividend int
	Divisor  int
	Message  string
}

// 2. Implement method Error() để thỏa interface error
func (e DivideError) Error() string {
	return fmt.Sprintf("Lỗi chia: %d / %d - %s", e.Dividend, e.Divisor, e.Message)
}

// 3. Hàm chia, trả về custom error nếu chia cho 0
func divide(a, b int) (int, error) {
	if b == 0 {
		return 0, DivideError{
			Dividend: a,
			Divisor:  b,
			Message:  "Không được chia cho 0",
		}
	}
	return a / b, nil
}

func main() {
	result, err := divide(10, 2)

	if err != nil {
		// 4. Kiểm tra kiểu lỗi cụ thể (type assertion)
		if de, ok := err.(DivideError); ok {
			fmt.Println("Chi tiết lỗi:")
			fmt.Println("  -> Số bị chia:", de.Dividend)
			fmt.Println("  -> Số chia   :", de.Divisor)
			fmt.Println("  -> Lý do     :", de.Message)
		} else {
			fmt.Println("Lỗi:", err)
		}
		return
	}

	fmt.Println("Kết quả:", result)
}
