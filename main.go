package main

import (
	"fmt"
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

	}

}
