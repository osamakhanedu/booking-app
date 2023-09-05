package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go Conference"
	const conferenceTickets uint = 50
	var remainingTickets uint = 50
	var bookings []string

	greeting(conferenceName, int(conferenceTickets), int(remainingTickets))

	for {

		if remainingTickets == 0 {
			fmt.Println("Our conference is booked out. Come back next year.")
			break

		}

		firstName, lastName, email, userTickets := getUserInput()
		isValidName, isValidEmail, isValidTicketNumber := validateInput(firstName, lastName, email, userTickets, remainingTickets)

		if !isValidName || !isValidEmail || !isValidTicketNumber {
			fmt.Println("You have enter invalid input")
			continue
		}

		if userTickets > remainingTickets {
			fmt.Printf("We only have %v ticket remaining, so you can not book %v tickets\n", remainingTickets, userTickets)
			continue
		}

		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("User %v booked %v tickets.\n", firstName, userTickets)

		remainingTickets = remainingTickets - userTickets

		fmt.Printf("Thank you %v %v for booking %v tickets. you will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)

		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := getName(bookings)

		fmt.Printf("Booking  %v\n", firstNames)

		// firstNames := []string{}

		// for _, booking := range bookings {
		// 	firstNames = append(firstNames, strings.Split(booking, " ")[0])
		// }
		// fmt.Printf("Booking  %v\n", firstNames)

	}

}

func greeting(conferenceName string, conferenceTickets int, remainingTickets int) {

	fmt.Printf("Welcome to %s booking application\n", conferenceName)
	fmt.Printf("we have total of %v tickets, and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getName(bookings []string) []string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, strings.Split(booking, " ")[0])
	}

	return firstNames
}

func validateInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {
	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEmail := strings.Contains(email, "@")
	isValidTicketNumber := userTickets > 0 && userTickets > remainingTickets

	return isValidName, isValidEmail, isValidTicketNumber
}

func getUserInput() (string, string, string, uint) {
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
