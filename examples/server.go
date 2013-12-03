package main

import (
	"github.com/DigitalInnovation/go_raygun/goraygun"
	"log"
)

func mimicPanic() {

	log.Println("In panicMethod")

	var temp []int

	_ = temp[1]

}

func main() {

	log.Panicln("In testapp main")

	goraygun := new(goraygun.Goraygun)

	err := goraygun.LoadRaygunSettings()
	if err != nil {
		log.Printf("Error loading goraygun config: %v", err)
	}

	defer goraygun.RaygunRecovery()

	mimicPanic()

	log.Println("Have recovered and now can continue or fail")

}
