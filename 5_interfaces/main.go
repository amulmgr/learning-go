package main

import (
	"fmt"
	"reflect"
)

//Some code to get familiar with interfaces
//Based off of : http://go-book.appspot.com/interfaces.html

//Create some interfaces
//-----------------------------------------------------------
type Mover interface {
	Move()
}

type TrickSter interface {
	MakeTrick()
}

//-----------------------------------------------------------

//Create a few structs : a Vehicle, a Bicycle, a Skateboard and a Car
type Vehicle struct {
	WheelCount  int
	VehicleType string
}

type Bicycle struct {
	Vehicle //Anonymous field type of Vehicle
	weight  int
}

type Skateboard struct {
	Vehicle       //Anonymous field type of Vehicle
	gripTapeBrand string
}

type Car struct {
	Vehicle //Anonymous field type of Vehicle
	mpg     int
}

//-----------------------------------------------------------

//Movers is a slice of Mover interfaces
type Movers []Mover

//TrickSters is a slice fo TrickSter interfaces
type TrickSters []TrickSter

//-----------------------------------------------------------

//Vehicle will implement the Mover interface
func (v *Vehicle) Move() {
	fmt.Println("Moving " + v.VehicleType)
}

//Bicycle will override the Mover interface
func (v *Bicycle) Move() {
	fmt.Println("Moving " + v.VehicleType + " Merrily because I'm on a bike!!!")
}

//Bicycle will implement the Trickster interface
func (v *Bicycle) MakeTrick() {
	fmt.Println("Tabletop")
}

//Skateboard will implement the TrickSter interface
func (v *Skateboard) MakeTrick() {
	fmt.Println("Flip Kick")
}

func main() {

	//Make a bike
	bike := new(Bicycle)
	bike.VehicleType = "Bike"
	bike.WheelCount = 2
	bike.weight = 8

	//Make a car
	car := new(Car)
	car.VehicleType = "Car"
	car.WheelCount = 4
	car.mpg = 45

	//Make a Skateboard
	skateboard := new(Skateboard)
	skateboard.VehicleType = "Skateboard"
	skateboard.WheelCount = 4
	skateboard.gripTapeBrand = "Flypaper"

	//Slice of Mover interface .. add bike, car and skateboard
	movers := Movers{bike, car, skateboard}

	fmt.Println("Calling Move() ... ")

	//Range over movers and invoke the Move() method
	for _, item := range movers {
		fmt.Println("--")
		//Print out the type
		fmt.Print("Type : ")
		fmt.Println(reflect.TypeOf(item))
		fmt.Println("--")
		item.Move()
		fmt.Println("")
	}

	fmt.Println("\n\nCalling MakeTrick() ...")

	//Slice of Trickster interfaces.. add bike and skateboard
	tricksters := TrickSters{bike, skateboard}
	//Range over movers and invoke the Move() method
	for _, item := range tricksters {
		fmt.Println("--")
		//Print out the type
		fmt.Print("Type : ")
		fmt.Println(reflect.TypeOf(item))
		fmt.Println("--")
		item.MakeTrick()
		fmt.Println("")
	}

}
