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

	fmt.Printf("Welcome to %s booking application\n", conferenceName)
	fmt.Printf("we have total of %v tickets, and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	for {

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

		bookings = append(bookings, firstName+" "+lastName)

		fmt.Printf("User %v booked %v tickets.\n", firstName, userTickets)

		remainingTickets = remainingTickets - userTickets

		fmt.Printf("Thank you %v %v for booking %v tickets. you will receive a confirmation email at %v.\n", firstName, lastName, userTickets, email)

		fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

		firstNames := []string{}

		for _, booking := range bookings {
			firstNames = append(firstNames, strings.Split(booking, " ")[0])
		}
		fmt.Printf("Booking  %v\n", firstNames)

	}

}
