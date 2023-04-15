package main

import "fmt" //Format Package

func main() { //entry point

	conferenceName := "Go Conference" // short declaration...const cannot be declared like this and type declaration cannot be done explicitly
	const conferenceTickets int = 50  //cant be changed
	var remainingTickets int = 50     //decreases as tickets get booked.

	fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)

	// Println ends the output with a new line.
	// Printf is for formatted data
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are remaining.\n", conferenceTickets, remainingTickets)
	fmt.Print("Get your tickets here to attend.")
	fmt.Println("")

	var firstName string
	var lastName string
	var email string
	var userTickets int

	// ask the user for name and no.of tickets
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName) //scan for input ...takes memory address(pointer) of the variables as the argument.

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email soon at %v\n", firstName, lastName, userTickets, email)

}
