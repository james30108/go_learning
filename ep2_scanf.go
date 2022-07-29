package main

import (
	"fmt"
)

func main() {

	var name string
	var score int
	var score_float float32

	// การป้อนข้อความและเก็บลงในตัวแปร

	fmt.Print("ป้อนชื่อของท่าน : ")
	fmt.Scanf("%s", &name) // เก็บตัวแปรเป็น String

	fmt.Print("คะแนน : ")
	fmt.Scanf("%f", &score_float) // เก็บตัวแปรเป็น float
	fmt.Scanf("%d", &score)       // เก็บตัวแปรเป็น integer

	// ---------------------------

	// Condition

	// if ... else
	if score >= 50 {
		fmt.Println("สอบผ่าน")
	} else if score == 50 {
		fmt.Println("ผ่านแบบคาบเส้น")
	} else {
		fmt.Println("สอบไม่ผ่าน")
	}

	// switch ...case
	// ใช้งานการเปรียบเทียบคำสั่งไม่ได้ ถ้าต้องการใช้งานเครื่องหมายมากกว่าน้อยกว่า ควรใช้ if ... else จะดีกว่า
	switch score {
	case 50:
		fmt.Println("เกรด F")
	case 70:
		fmt.Println("เกรด D")
	case 80:
		fmt.Println("เกรด C")
	case 90:
		fmt.Println("เกรด B")
	case 100:
		fmt.Println("เกรด A")
	}

}
