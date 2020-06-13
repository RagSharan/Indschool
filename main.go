package main

import "fmt"

func main(){
	fmt.Println("animal properties");
	animal := Animal{"tiger", "yellow", "hunter"}
	animal.animalInfo(animal);
	design = "design my bsg"
	fmt.Println(design)
}
var design string
//func receiver identifier parameters return
func add(num1 int, num2 int) float32{
	return float32(num1 + num2)

}

type Animal struct{
	animalName string
	color string
	animalType string
}

func (a Animal) animalInfo(animal Animal){
	
	println("Animal Name",animal.animalName)
	println("Animal Color", animal.color)
	println("Animal Type", animal.animalType)
}