package main

import (
	"fmt"
	"go_learning/calculator"
)

// ใตัวอย่างการใช้งาน Package
func main() {

	// การเรียกใช้งาน package ที่สสร้าง
	fmt.Println("ผลรวมของการบวกคะแนน = ", calculator.Add(50, 100))
}
