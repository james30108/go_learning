package main

import "fmt"

// ตัวอย่างการทำงานของ Structure
type Product struct {
	name     string
	price    float64
	discount int
}

type Thing struct {
	Name string
	Num  int
}

func (t *Thing) Init(name string, num int) {
	t.Name = name
	t.Num = num
}

func main() {

	// การเรียกใช้งาน Structure
	product1 := Product{name: "ปากกา", price: 40, discount: 10}
	product2 := Product{name: "ยางลบ", price: 20, discount: 5}

	// การเปลี่ยนแปลงข้อมูล
	product1.price = 100
	fmt.Println(product1)
	fmt.Println(product2)

	// การเรียกใช้งาน Structure ผ่าน Function
	t := new(Thing)
	t.Init("Hello", 5)
	fmt.Printf("%s: %d\n", t.Name, t.Num)
}
