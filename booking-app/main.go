package main

import (
	"fmt"
	"strings"
)


func main() {
	conferenceName := "Go Conference"
	const conferenceTickets int = 50
	var remainingTickets int = 50
	// slice
	var bookings []string // or: bookings := []string{}

	fmt.Printf("conference ticket is %T, remaining ticket is %T, conference name is %T\n", conferenceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to %v booking app!\n", conferenceName)
	fmt.Printf("We have total of %v tickets and %v tickets remaining.\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")

	// infinite loop -> for {}
	// conditional loop -> for condition {}
	for remainingTickets > 0  && len(bookings) < 50 {	
		var userName string
		var email string
		var userTickets int
		// ask user to enter their name
		fmt.Print("Enter your name: ")
		fmt.Scan(&userName)

		fmt.Print("Enter your email: ")
		fmt.Scan(&email)
		
		// ask user to enter number of tickets
		fmt.Print("Enter number of tickets you want to book: ")
		fmt.Scan(&userTickets)

		isValidName := len(userName) >= 2
		isValidEmail := strings.Contains(email, "@") && strings.Contains(email, ".")
		isValidTicket := userTickets > 0 && userTickets <= remainingTickets
		
		if !isValidName || !isValidEmail || !isValidTicket {
			if !isValidName {
				fmt.Println("Invalid name")
			} 
			if !isValidEmail {
				fmt.Println("Invalid email")
			} 
			if !isValidTicket {
				fmt.Println("Invalid ticket number")
			}
			continue
		} 
		// else if userTickets == remainingTickets {
		// 	// do sth
		// }

		remainingTickets = remainingTickets - userTickets
		bookings = append(bookings, userName + " Vu")

		fmt.Printf("The whole bookings: %v\n", bookings)
		fmt.Printf("The first booking: %v\n", bookings[0])
		fmt.Printf("Array type: %T\n", bookings)
		fmt.Printf("Array length: %v\n", len(bookings))

		fmt.Printf("User %v with the %v books %v tickets.\n", userName, email, userTickets)
		fmt.Printf("Remaining tickets: %v\n", remainingTickets)

		firstNames := []string{}

		// for each loop -> index, element := range array {}
		for _, booking := range bookings {
			var names = strings.Fields(booking)
			firstNames = append(firstNames, names[0])
		}
		fmt.Printf("These are all the bookings: %v\n", bookings)
		fmt.Printf("The first names of booking are: %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("All tickets are sold out!")
			break
		}
	}
}