package main

import (
	"fmt"
	"goPractice/helper"
	"time"
)
const conferenceTickets int = 50

var conferenceName = "Go Conference"
var remainingTickets uint = 50
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {

	greetUsers()

	for {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEmail, isValidTicketNumber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEmail && isValidTicketNumber {

			bookTicket(userTickets, firstName, lastName, email)
			go sendTicket(userTickets, firstName, lastName, email)
			//Call function to print first names
			firstNames := getFirstNames()
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			if remainingTickets == 0 {
				fmt.Println("Our conference is booked out, come back next year.")
				break
			}
		} else {
			if !isValidName {
				fmt.Println("First name or last name entered is too short!")
			}
			if !isValidEmail {
				fmt.Println("Entered email address missing @ sign!")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid!")
			}
		}
	}
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application!\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
	fmt.Println("Click here to get your tickets")
}

func getFirstNames() []string {
	firstNames := []string{}

	//for-each loop
	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	//Request user their name
	fmt.Println("Enter your first name: ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your last name: ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your email: ")
	fmt.Scan(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	//Create a map for the user key-value pairs
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	//slice >> Dynamic array
	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	fmt.Printf("Thank you %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
	fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println("#######################")
	fmt.Printf("Sending ticket: %v\n to email address %v\n", ticket, email)
	fmt.Println("#######################")
}
