package main

import "fmt" //Format Package

func main() { //entry point

	var conferenceName string = "Go Conference"
	const conferenceTickets int = 50 //cant be changed
	var remainingTickets int = 50    //decreases as tickets get booked.

	fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)

	// Println ends the output with a new line.
	// Printf is for formatted data
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are remaining.\n", conferenceTickets, remainingTickets)
	fmt.Print("Get your tickets here to attend.")
	fmt.Println("")

	var userName string
	var userTickets int
	// ask the user for name and no.of tickets

	userName = "Dameon"
	userTickets = 3
	fmt.Printf("%v has booked %v tickets.\n", userName, userTickets)

}
