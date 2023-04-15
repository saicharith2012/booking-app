package main

import "fmt" //Format Package

func main() { //entry point

	var conferenceName = "Go Conference"
	const conferenceTickets = 50 //cant be changed
	var remainingTickets = 50    //decreases as tickets get booked.

	fmt.Println("Welcome to", conferenceName, "booking application") //Println ends the output with a new line.
	fmt.Println("We have a total of", conferenceTickets, "tickets and", remainingTickets, "are remaining.")
	fmt.Print("Get your tickets here to attend.")
	fmt.Println("")
}
