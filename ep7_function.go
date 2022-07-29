package main

import (
	"fmt"
)

// ใตัวอย่างการใช้งาน Function
func main() {

	// การเรียกใช้งาน
	showMessage("james")

	// รับค่าแบบ multiple return
	total, status := multiReturn(10, 20)
	fmt.Println("รลรวมเท่ากับ : ", total, " สถานะเท่ากับ : ", status)

}

// การสร้างฟังก์ชั่น
func showMessage(name string) {
	fmt.Println("Hello : ", name)
}

// แบบที่ 2
func showMessage2(number1, number2 int) {
	fmt.Println("Sum : ", number1+number2)
}

// แบบที่ 3 แบบมีการ return ค่า ให้ระบุชนิดข้อมูลที่ต้องการ return ค่าออกไปด้วย
func getSalary(salary int) int {
	return salary
}

// แบบที่ 4 แบบที่ return ค่ามากกว่า 1 ค่าออกไป
func multiReturn(number1, number2 int) (int, string) {
	total := number1 + number2
	status := ""
	if total > 100 {
		status = "มากกว่า 100"
	} else {
		status = "น้อยกว่า 100"
	}
	return total, status
}

// การนิยาม parameter แบบไม่ได้กำหนดขนาด
func summation(numbers ...int) int {
	total := 0
	for _, value := range numbers {
		total += value
	}
	return total
}
