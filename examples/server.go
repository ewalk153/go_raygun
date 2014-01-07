package main

import (
	"log"
	"raygun/goraygun"
)

func mimicPanic() {

	log.Println("In panicMethod")

	method2()

}

func method2() {
	log.Println("In method2")

	method3()
}

func method3() {
	log.Println("In method3")

	method4()
}

func method4() {
	log.Println("In method4")

	method5()
}

func method5() {
	log.Println("In method5")

	method6()
}

func method6() {

	log.Println("In method6")

	var temp []int

	_ = temp[1]
}

func main() {

	goraygun := new(goraygun.Goraygun)

	err := goraygun.LoadRaygunSettings()
	if err != nil {
		log.Printf("Error loading goraygun config: %v", err)
	}

	defer goraygun.RaygunRecovery()

	mimicPanic()

	log.Println("Have recovered and now can continue or fail")

}
