package main

import "fmt"

// ตัวอย่างการทำงานของ Structure
type Product struct {
	name     string
	price    float64
	discount int
}

func main() {

	// การเรียกใช้งาน Structure
	product1 := Product{name: "ปากกา", price: 40, discount: 10}
	product2 := Product{name: "ยางลบ", price: 20, discount: 5}

	// การเปลี่ยนแปลงข้อมูล
	product1.price = 100
	fmt.Println(product1)
	fmt.Println(product2)
}
