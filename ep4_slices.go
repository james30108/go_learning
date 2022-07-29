package main

import "fmt"

// Slices
// คล้ายกับ Array แต่สามารถปรับเปลี่ยนขนาดสมาชิกได้
func main() {

	// ตัวอย่างการนิยาม Slices
	// จะไม่ทำการกำหนดสมาชิกในก้ามปู
	numbers1 := []int{100, 200, 300}

	// การเปลี่ยนแปลงค่า
	numbers1[0] = 1000
	numbers1[1] = 2000

	// การเพิ่มจำนวน
	numbers1 = append(numbers1, 400)
	numbers1 = append(numbers1, 500)

	// การแสดงผล
	fmt.Println(numbers1)
	fmt.Println(numbers1[0])
	fmt.Println(numbers1[1])

	// การแสดงผลแบบจาก index ที่ 1 จนถึงสุดท้าย
	fmt.Println(numbers1[1:])

	// การแสดงผล index ที่ 1 ถึง 2
	fmt.Println(numbers1[1:3])

	// การแสดงผล index เริ่มต้นถึง 2
	fmt.Println(numbers1[:3])
}
