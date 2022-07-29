package main // package หลักที่จะต้องรันก่อนทุกครั้ง
import (
	"fmt" // กลุ่มคำสั่งพื้นฐาน เช่น สำหรับแสดงผลทางหน้าจอภาพ
	"go/constant"
)

// การนิยาม function main ซึ่งเป็นฟังก์ชั่นพิเศษที่จะทำการรันเป็นอันดับแรกเสมอเมื่อใช้งานโปรแกรม
func main() {

	// ตัวอย่าง การประกาศตัวแปร
	var name string = "Nathasophon" // แบบ String
	var age int = 28                // แบบ Integer
	var score float32 = 95.8        // แบบ float
	var isPass bool = true          // แบบ boolean

	// สามารถเขียนแบบนี้ก็ได้
	name2 := "James"
	age2 := 29
	score2 := 50.5
	isPass2 := false

	// ตัวแปรที่ไม่สามารถเปลี่ยนแปลงค่าได้
	const salary := 15500

	fmt.Println("My name is ", name) // แสดงข้อความ

	// ตรวจสอบชนิดข้อมูล
	fmt.Printf("My name is %v \n", name) // แสดง value
	fmt.Printf("My name is %T \n", name) // แสดงชนิดข้อมูล

}
