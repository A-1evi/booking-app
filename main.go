package main

import (
	"fmt"
	"strings"
)

var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = 50

var bookings []string // slice of strings

func main() {

	greetUser()
	for {
		userName, userEmail, userTickets := getUserInput()

		isValidEmail, isValidUserName, isValidTicketsNumber := validateUserInput(userName, userEmail, userTickets)
		if isValidEmail && isValidUserName && isValidTicketsNumber {
			bookingTickets(userTickets, userName, userEmail)
			// Print the first names of the users who booke the  tickets
			firstNames := getFirstName()
			fmt.Printf("List of firstnames: %v of the person booked the coference\n", firstNames)
			//var ticketRemaing bool = remainingTickets ==0
			if remainingTickets == 0 {
				//end the program
				fmt.Println("The conference tickets are over")
				break
			}
		} else {
			if !isValidUserName {
				fmt.Println("The user name is not valid")
			}
			if !isValidTicketsNumber {
				fmt.Println("The tickets no. are not valid")
			}
			if !isValidUserName {
				fmt.Println("The email is not valid")
			}
		}
	}

}

func greetUser() {
	fmt.Printf("Welcome to %v Booking Application\n", conferenceName)
	fmt.Printf("We hav total of %v ticker available and %v are still available\n", conferenceTickets, remainingTickets)

}

func getFirstName() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		var names = strings.Fields(booking) // strings.Fields seperate string based on whitespaces
		firstNames = append(firstNames, names[0])

	}
	return firstNames

}

func validateUserInput(userName string, userEmail string, userTickets uint) (bool, bool, bool) {

	isValidUserName := len(userName) >= 2
	isValidEmail := strings.Contains(userEmail, "@") //strings.Contains return boolean
	isValidTicketsNumber := userTickets > 0 && userTickets <= remainingTickets

	return isValidEmail, isValidUserName, isValidTicketsNumber
}

func getUserInput() (string, string, uint) {
	var userName string
	var userEmail string
	var userTickets uint
	fmt.Println("Enter your FirstName:")
	fmt.Scan(&userName)

	fmt.Println("Enter your Email:")
	fmt.Scan(&userEmail)

	fmt.Println("Enter the number of tickets you wnat to book:")
	fmt.Scan(&userTickets)

	return userName, userEmail, userTickets
}

func bookingTickets(userTickets uint, userName string, userEmail string) {
	remainingTickets = remainingTickets - userTickets
	bookings = append(bookings, userName+" "+userEmail)
	fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email at %v\n", userName, userTickets, userEmail)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}
