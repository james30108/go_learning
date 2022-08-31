package main

import (
	"encoding/json"
	"fmt"
	"strings"
)

func upperAllLetter(str *string) {
	// แปลงค่า str ที่รับเข้ามาเป็นตัวพิมพ์ใหญ่
	*str = strings.ToUpper(*str)
}

// Struct หรือ Constructor
type Profile struct {
	Name string
	Age  int
}

// function สำหรับเปลี่ยนแปลงค่าใน Struct ที่ชื่อ Profile
func upperProfileName(p *Profile) {
	p.Name = strings.ToUpper(p.Name)
}

// function สำหรับเปลี่ยนแปลงค่าใน Slice (Array)
func doubleAllElement(nums []int) {
	for i := range nums {
		nums[i] *= 2
	}
}

// สำหรับแปลงข้อมูลที่ return ออกไป
func show() *int {

	lv := 100
	return &lv
}

// ตัวอย่างการใช้งาน Pointer
func main() {

	// * ดอกจัน กำหนดให้กับค่าที่ต้องการกระทำให้มีการเปลี่ยนแปลง
	// & กำหนดให้กับค่าที่ถูกกระทำให้เปลี่ยนแปลง
	// -----------------------

	// ตัวอย่างที่ 1 การใช้งาน Pointer เพื่อเปลี่ยนค่า name โดยที่ไม่ต้องกำหนดผ่านตัวแปร name
	name := "Por"
	// ถ้าประกาศโดยใช้ data type ด้วยก็คือใช้ (ตัวเต็ม)
	// var pointerToName *string = &name
	pointerToName := &name

	*pointerToName = "Weerasak"
	// ผลที่ออกมาคือค่าของ name จะเปแลี่ยนเป็น Weerasak
	fmt.Println(name)

	// -----------------------
	// ตัวอย่างที่ 2 เปลี่ยนค่า name2 ผ่าน function
	name2 := "Weerasak"
	upperAllLetter(&name2)
	fmt.Println(name2)

	// -----------------------
	// ตัวอย่างที่ 3 เปลี่ยนค่า name ผ่าน Struct
	p := Profile{
		Name: "Por",
		Age:  35,
	}
	// เปลี่ยนค่า name เป็นตัวพิมพ์ใหญ่ผ่าน function pointer
	upperProfileName(&p)
	fmt.Println(p.Name)

	// -----------------------
	// ตัวอย่างที่ 4 เปลี่ยนค่าใน JSON
	jsonText := []byte(`[1, 2, 3]`)
	var nums []int
	// ใช้ & เพื่อเรียกใช้งาน pointer (ถ้าไม่ใส่จะแปลงข้อมูลไม่ได้)
	// Unmarshal คือ fuction ที่ใช้แปลงข้อมูลเป็น json
	json.Unmarshal(jsonText, &nums)
	fmt.Println(nums)

	// -----------------------
	// ตัวอย่างที่ 5 ใช้ Pointer เพื่อ Retuen ค่าออกมา
	number := show()
	fmt.Println("Value of number is: ", *number)
}
