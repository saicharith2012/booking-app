package main

import (
	"fmt" //Format Package
	"strings"
)

func main() { //entry point

	conferenceName := "Go Conference" // short declaration...const cannot be declared like this and type declaration cannot be done explicitly
	const conferenceTickets uint = 50 //cant be changed
	var remainingTickets uint = 50    //decreases as tickets get booked.
	var bookings []string

	fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)

	// Println ends the output with a new line.
	// Printf is for formatted data

	//adding an Infinite For loop to execute the code of bookings multiple times.
	for {
		fmt.Printf("Welcome to %v booking application\n", conferenceName)
		fmt.Printf("We have a total of %v tickets and %v are remaining.\n", conferenceTickets, remainingTickets)
		fmt.Print("Get your tickets here to attend.\n")

		var firstName string
		var lastName string
		var email string
		var userTickets uint //uint since tickets cant be negative

		// ask the user for name and no.of tickets
		fmt.Println("Enter your first name: ")
		fmt.Scan(&firstName) //scan for input ...takes memory address(pointer) of the variables as the argument.

		fmt.Println("Enter your last name: ")
		fmt.Scan(&lastName)

		fmt.Println("Enter your email: ")
		fmt.Scan(&email)

		fmt.Println("Enter number of tickets: ")
		fmt.Scan(&userTickets)

		if userTickets <= remainingTickets {

			remainingTickets = remainingTickets - userTickets //remaining tickets after booking.
			bookings = append(bookings, firstName+" "+lastName+",")

			fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email soon at %v.\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for the %v.\n", remainingTickets, conferenceName)

			firstNames := []string{}

			for _, booking := range bookings {
				var names = strings.Fields(booking)
				firstNames = append(firstNames, names[0]+",")

			}

			fmt.Printf("The first names of all our bookings are: %v\n", firstNames)

			fmt.Printf(" __________________________________________________________\n\n")

			if remainingTickets == 0 {
				//end program
				fmt.Println("Our Conference is booked out. Come back next year.")
				break //exits the program
			}
		} else {
			fmt.Printf("We have only %v tickets available. So you can't book %v tickets. Come back next time.\n", remainingTickets, userTickets)
			// continue; not needed.
			fmt.Printf(" __________________________________________________________\n\n")
		}
	}
}
