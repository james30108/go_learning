package main

import (
	"fmt"
)

// ในภาษา Go จะมีแค่ For Loop
func main() {

	// ตัวอย่าง
	for count := 1; count <= 10; count++ {
		fmt.Println("บรรทัดที่ : ", count)

		// ตัวอย่าง continnue
		if count == 2 {
			continue
		}

		// ตัวอย่าง brack
		if count == 5 {
			break
		}

	}
	fmt.Println("จบโปรแกรม")

	// ตัวอย่างการใช้ For Loop มาแสดงผลข้อมูลใน Array
	numbers := []int{10, 20, 30, 40, 50, 60}

	for index := 0; index < len(numbers); index++ {
		fmt.Println(numbers[index])
	}

	// ตัวอย่างใช่งาน for lenge ในการแสดงผลข้อมูลใน Array
	for index, value := range numbers {
		fmt.Println("Index ที่ : ", index, " = Value ที่ : ", value)
	}

	// การ Loop ค่าใน Map
	language := map[string]string{"TH": "thai", "EN": "english"}
	for index, value := range language {
		fmt.Println("Index ที่ : ", index, " = Value ที่ : ", value)
	}
}
