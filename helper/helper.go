package helper

import (
	"fmt"
	"strings"
)

/**
  Validate input
*/

func ValidateInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets > remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

// Show Greeting messages
func Greeting(conferenceName string, conferenceTickets int, remainingTickets int) {

	fmt.Printf("Welcome to %s booking application\n", conferenceName)
	fmt.Printf("we have total of %v tickets, and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}



// get firstName from slice
func GetName(bookings []string) []string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, strings.Split(booking, " ")[0])
	}

	return firstNames
}

// Get user input
func GetUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)
	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)
	fmt.Println("Enter your email ")
	fmt.Scan(&email)

	fmt.Println("Number of Tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}
