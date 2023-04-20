package main

import (
	"fmt" //Format Package
)

type userData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

// package level variables
// variable cant be declared with short declaration at package levels.
var conferenceName = "Go Conference" // short declaration...const cannot be declared like this and type declaration cannot be done explicitly
const conferenceTickets uint = 50    //cant be changed
var remainingTickets uint = 50       //decreases as tickets get booked.
var bookings = make([]userData, 0)

func main() { //entry point

	fmt.Printf("conferenceName is %T, conferenceTickets is %T, remainingTickets is %T\n", conferenceName, conferenceTickets, remainingTickets)

	// Println ends the output with a new line.
	// Printf is for formatted data

	//adding an Infinite For loop to execute the code of bookings multiple times.
	for {

		greetUsers()

		// asking user for inputs
		firstName, lastName, email, userTickets := getUserInputs()

		// user input validation
		isValidName, isValidEmail, isTicketNumberPositive, isValidTicketNumber :=
			validateUserInputs(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isTicketNumberPositive && isValidTicketNumber {

			// calling bookTickets
			bookTickets(userTickets, firstName, lastName, email)

			// calling the function for printing the first names
			var firstNames = getFirstNames(bookings) //return value of getFirstNames
			fmt.Printf("The first names of all our bookings are: %v\n", firstNames)

			fmt.Printf(" __________________________________________________________\n\n")

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our Conference is booked out. Come back next year.")
				break //exits the program
			}
		} else { // only ifs are used instead of "else if" since each conditon must be checked once.
			if !isValidName {
				fmt.Println("The firstname or lastname you entered is invalid. Atleast 2 characters long.")
			}
			if !isValidEmail {
				fmt.Println("The email you have entered is invalid.")
			}
			if !isTicketNumberPositive {
				fmt.Println("The number of tickets you entered is invalid.")
			}
			if !isValidTicketNumber {
				fmt.Printf("We have only %v tickets available. So you can't book %v tickets. Come back next time.\n", remainingTickets, userTickets)
				// continue; not needed.
			}
			fmt.Printf(" __________________________________________________________\n\n")
		}
	}

	// city := "New York"

	// switch city {
	// case "New York":
	// 	// execute code for New York conference tickets.
	// case "Singapore", "Hong Kong":
	// 	// execute code for New York conference tickets.
	// case "Malaysia":
	// 	// execute code for New York conference tickets.
	// case "London", "Melbourne":
	// 	// execute code for New York conference tickets.

	// default:
	// 	fmt.Println("No valid city selected.")
	// }
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are remaining.\n", conferenceTickets, remainingTickets)
	fmt.Print("Get your tickets here to attend.\n")
}

func getFirstNames([]userData) []string {
	firstNames := []string{}
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName+",")
	}
	return firstNames
}

func getUserInputs() (string, string, string, uint) {
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

	return firstName, lastName, email, userTickets

}

func bookTickets(userTickets uint, firstName string, lastName string, email string) ([]userData, uint) {
	remainingTickets = remainingTickets - userTickets //remaining tickets after booking.

	var userData = userData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	bookings = append(bookings, userData)
	fmt.Printf("The list of bookings are %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email soon at %v.\n",
		firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for the %v.\n", remainingTickets, conferenceName)

	return bookings, remainingTickets

}
