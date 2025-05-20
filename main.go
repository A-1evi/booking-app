package main

import (
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Go Conference"

const conferenceTickets = 50

var remainingTickets uint = 50

var bookings = make([]UserData, 0) // slice of strings

type UserData struct {
	userName    string
	userEmail   string
	userTickets uint
}

var wg = sync.WaitGroup{}

func main() {

	greetUser()

	userName, userEmail, userTickets := getUserInput()

	isValidEmail, isValidUserName, isValidTicketsNumber := ValidateUserInput(userName, userEmail, userTickets)
	if isValidEmail && isValidUserName && isValidTicketsNumber {
		bookingTickets(userTickets, userName, userEmail)
		wg.Add(1)
		go sendTicket(userName, userEmail, userTickets)
		// Print the first names of the users who booke the  tickets
		var userData = UserData{
			userName:    userName,
			userEmail:   userEmail,
			userTickets: userTickets,
		}

		bookings = append(bookings, userData)
		fmt.Printf("List of booking %v\n", bookings)
		firstNames := getFirstName()
		fmt.Printf("List of firstnames: %v of the person booked the coference\n", firstNames)
		//var ticketRemaing bool = remainingTickets ==0
		if remainingTickets == 0 {
			//end the program
			fmt.Println("The conference tickets are over")
			//break
		}
	} else {
		if !isValidUserName {
			fmt.Println("The user name is not valid")
		}
		if !isValidTicketsNumber {
			fmt.Println("The tickets no. are not valid")
		}
		if !isValidEmail {
			fmt.Println("The email is doesnt contain @ in it ")
		}
	}
	wg.Wait()
}

func greetUser() {
	fmt.Printf("Welcome to %v Booking Application\n", conferenceName)
	fmt.Printf("We have total of %v ticket available and %v are still available\n", conferenceTickets, remainingTickets)

}

func getFirstName() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		// strings.Fields seperate string based on whitespaces

		//we are using map now so string.Fields logic no loger needed
		firstNames = append(firstNames, booking.userName)

	}
	return firstNames

}

func getUserInput() (string, string, uint) {
	var userName string
	var userEmail string
	var userTickets uint
	fmt.Println("Enter your username:")
	fmt.Scan(&userName)

	fmt.Println("Enter your Email:")
	fmt.Scan(&userEmail)

	fmt.Println("Enter the number of tickets you wnat to book:")
	fmt.Scan(&userTickets)

	return userName, userEmail, userTickets
}

func bookingTickets(userTickets uint, userName string, userEmail string) {
	remainingTickets = remainingTickets - userTickets

	fmt.Printf("Thank you %v for booking %v tickets. You will receive a confirmation email at %v\n", userName, userTickets, userEmail)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userName string, userEmail string, userTickets uint) {
	time.Sleep(10 * time.Second)

	println("\n****************")
	var ticket = fmt.Sprintf("\n%v are booked for %v ", userTickets, userName)
	fmt.Printf("%v, relavent info will be sent on %v\n ", ticket, userEmail)
	println("****************")
	wg.Done()
}
