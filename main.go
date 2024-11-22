package main

import (
	"fmt"

)

func main(){
	fmt.Println("Welcome to aircraft records keeper\n")
	var aircraftno int
	var flownhours int
	var dailylandings int
	
	fmt.Print("Aircaft number: ")
	fmt.Scanln(&aircraftno)

	fmt.Print("Daily flown hours: ")
	fmt.Scanln(&flownhours)

	fmt.Print("Daily landings: ")
	fmt.Scanln(&dailylandings)


	fmt.Printf("Aircraft no: %d, Daily flown hours: %d, Daily landings: %d\n", aircraftno, flownhours, dailylandings)
}