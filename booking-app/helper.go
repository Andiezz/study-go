package main

import "strings"

// function return multiple values
// Export -> name of the function starts with capital letter
func validateUserInput(userName string, email string, userTickets uint) (bool, bool, bool) {
	isValidName := len(userName) >= 2
	isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
	isValidTicket := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEmail, isValidTicket
}
