package main

import "fmt"

// Map
// คล้ายกับ Object
func main() {

	// การนิยาม Map
	// โครงสร้าง = map[key_type]value_type

	county1 := map[string]string{"TH": "thailand", "ENG": "england"}

	// การเพิ่มข้อมูล
	county1["JP"] = "japan"

	//การเข้าถึงข้อมูล
	fmt.Println(county1)
	fmt.Println(county1["TH"])

	// การตรวจสอบค่า
	// โครงสร้าง ค่าที่ส่งกลับออกมาเมื่อพบข้อมูล, ค่า boolean ว่ามีหรือไม่มี := county1["JP"]
	value, check := county1["JP"]

	if check {
		fmt.Println(value)
	} else {
		fmt.Println("ไม่พบข้อมูล")
	}

}
