package main

import (
	"fmt"
	"strings"
	"time"
	"sync"
	// "strconv"
	// "booking-app/helper"
)

// package level variables
const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50

// slice
// var bookings = []string{} // or: bookings := []string{}
// var bookings = make([]map[string]string, 0)
var bookings = make([]UserData, 0)

type UserData struct {
	userName        string
	email           string
	numberOfTickets uint
}

// func main() {
// 	greetUsers()

// 	// infinite loop -> for {}
// 	// conditional loop -> for condition {}
// 	for remainingTickets > 0 && len(bookings) < 50 {
// 		userName, email, userTickets := getUserInput()
// 		// isValidName, isValidEmail, isValidTicket := helper.ValidateUserInput(userName, email, userTickets, remainingTickets)
// 		isValidName, isValidEmail, isValidTicket := validateUserInput(userName, email, userTickets)

// 		if !isValidName || !isValidEmail || !isValidTicket {
// 			if !isValidName {
// 				fmt.Println("Invalid name")
// 			}
// 			if !isValidEmail {
// 				fmt.Println("Invalid email")
// 			}
// 			if !isValidTicket {
// 				fmt.Println("Invalid ticket number")
// 			}
// 			continue
// 		}

// 		bookTicket(userName, email, userTickets)
// 		go sendTicket(userTickets, userName, email)

// 		// call function print first name
// 		firstNames := getFirstName()
// 		fmt.Printf("The first names of booking are: %v\n", firstNames)

// 		if remainingTickets == 0 {
// 			fmt.Println("All tickets are sold out!")
// 			break
// 		}
// 	}

// 	city := "London"

// 	switch city {
// 	case "London", "Berlin":
// 		fmt.Println("The conference is in London")
// 	case "Singapore":
// 		fmt.Println("The conference is in Singapore")
// 	case "Paris":
// 		fmt.Println("The conference is in Paris")
// 	default:
// 		fmt.Println("Invalid city is entered!")
// 	}
// }

var wg = sync.WaitGroup {}

// synchronize Goroutines
func main() {
	greetUsers()

	// infinite loop -> for {}
	// conditional loop -> for condition {}

	userName, email, userTickets := getUserInput()
	// isValidName, isValidEmail, isValidTicket := helper.ValidateUserInput(userName, email, userTickets, remainingTickets)
	isValidName, isValidEmail, isValidTicket := validateUserInput(userName, email, userTickets)

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
	}

	bookTicket(userName, email, userTickets)

	wg.Add(1)
	go sendTicket(userTickets, userName, email)

	// call function print first name
	firstNames := getFirstName()
	fmt.Printf("The first names of booking are: %v\n", firstNames)

	if remainingTickets == 0 {
		fmt.Println("All tickets are sold out!")
	}

	city := "London"

	switch city {
	case "London", "Berlin":
		fmt.Println("The conference is in London")
	case "Singapore":
		fmt.Println("The conference is in Singapore")
	case "Paris":
		fmt.Println("The conference is in Paris")
	default:
		fmt.Println("Invalid city is entered!")
	}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("conference ticket is %T, remaining ticket is %T, conference name is %T\n", conferenceTickets, remainingTickets, conferenceName)
	fmt.Printf("Welcome to %v!\n", conferenceName)

	fmt.Printf("We have total of %v tickets and %v tickets remaining.\n", conferenceName, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstName() []string {
	firstNames := []string{}

	// for each loop -> index, element := range array {}
	for _, booking := range bookings {
		var names = strings.Fields(booking.userName)
		firstNames = append(firstNames, names[0])
	}
	fmt.Printf("These are all the bookings: %v\n", bookings)
	return firstNames
}

func getUserInput() (string, string, uint) {
	var userName string
	var email string
	var userTickets uint

	fmt.Println("Enter your name:")
	fmt.Scanln(&userName)

	fmt.Println("Enter your email:")
	fmt.Scanln(&email)

	fmt.Println("Enter the number of tickets you want to book:")
	fmt.Scanln(&userTickets)

	return userName, email, userTickets
}

func bookTicket(userName string, email string, userTickets uint) {
	remainingTickets = remainingTickets - userTickets

	// create a map for a user
	// make empty map -> make(map[keyType]valueType)
	// var userdata = make(map[string]string)
	var userdata = UserData{
		userName:        userName,
		email:           email,
		numberOfTickets: userTickets,
	}
	// userdata["numberOfTickets"] = strconv.FormatUint(uint64(userTickets), 10)

	bookings = append(bookings, userdata)

	fmt.Printf("The whole bookings: %v\n", bookings)
	fmt.Printf("The first booking: %v\n", bookings[0])
	fmt.Printf("Array type: %T\n", bookings)
	fmt.Printf("Array length: %v\n", len(bookings))

	fmt.Printf("User %v with the %v books %v tickets.\n", userName, email, userTickets)
	fmt.Printf("Remaining tickets: %v\n", remainingTickets)
}

func sendTicket(userTickets uint, userName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickers for %v", userTickets, userName)
	fmt.Println("################")
	fmt.Printf("Sending ticket:\n %v \n to email addess %v\n", ticket, email)
	fmt.Println("################")
	wg.Done()
}
