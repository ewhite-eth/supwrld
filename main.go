package main

import (
	"fmt"
	"math"
)

func main() {
	 var user_height float64 
	 var user_kg float64
	 fmt.Println("--------------Calculation of body mass index--------------")
	 fmt.Print("Enter your height: ")
	 fmt.Scan(&user_height)
	 fmt.Print("Enter your weight: ")
	 fmt.Scan(&user_kg)
	 BMI := user_kg / math.Pow(user_height,2) 
	 fmt.Printf("Your body mass index: %v" ,BMI)
}