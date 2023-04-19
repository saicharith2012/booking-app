package main

import "strings"

func validateUserInputs(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2 // firstName and lastName should contain atleast two letters.
	isValidEmail := strings.Contains(email, "@")             // email must contain @ symbol
	isTicketNumberPositive := userTickets > 0
	isValidTicketNumber := userTickets <= remainingTickets // userTickets must be positive and less or equal to remaining tickets.

	return isValidName, isValidEmail, isTicketNumberPositive, isValidTicketNumber
}
