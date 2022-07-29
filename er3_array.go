package main

import "fmt"

// Array
func main() {

	// ตัวอย่างการประกาศตัวแปร
	var numbers [3]int

	// การกำหนดค่า
	numbers[0] = 100
	numbers[1] = 200
	numbers[2] = 300

	// การประกาศแบบลดรูป (Short hand)
	numbers2 := [3]int{400, 500, 600}

	// การประกาศแบบไม่กำหนดขนาด
	numbers3 := [...]int{400, 500, 600, 700, 800, 900}

	// การหาขนาดสมาชิกใน Array
	size := len(numbers)

	// การแสดงผล Array
	fmt.Println(numbers)
	fmt.Println(numbers2)
	fmt.Println(numbers3)
	fmt.Println(numbers2[0])
	fmt.Println(numbers2[1])
	fmt.Println(size)

}
