package main

import (
	"strings"
)

func ValidateUserInput(userName string, userEmail string, userTickets uint) (bool, bool, bool) {

	isValidUserName := len(userName) >= 2
	isValidEmail := strings.Contains(userEmail, "@") //strings.Contains return boolean
	isValidTicketsNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidEmail, isValidUserName, isValidTicketsNumber
}
