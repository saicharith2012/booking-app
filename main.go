package main

import "fmt" //Format Package

func main() { //entry point

	var conferenceName = "Go Conference"
	const conferenceTickets = 50 //cant be changed
	var remainingTickets = 50    //decreases as tickets get booked.

	// Println ends the output with a new line.
	// Printf is for formatted data
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are remaining.\n", conferenceTickets, remainingTickets)
	fmt.Print("Get your tickets here to attend.")
	fmt.Println("")
}
