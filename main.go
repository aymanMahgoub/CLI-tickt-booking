package main

import (
	"booking-app/user_validator"
	"fmt"
	"sync"
	"time"
)

var conferenceName = "Gaming Event"

const conferenceTickets int = 50

var remainingTickets uint = 50

var bookings = make([]User, 0)

type User struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

var wg = sync.WaitGroup{}

func main() {
	greetUsers()

	//for {
	firstName, lastName, email, numberOfTickets := getUserInput()
	isValidName, isValidEmail, isValidTicketNumber := user_validator.ValidateUserInput(firstName, lastName, email, numberOfTickets, remainingTickets)

	if isValidName && isValidEmail && isValidTicketNumber {
		remainingTickets = remainingTickets - numberOfTickets
		bookTicket(numberOfTickets, firstName, lastName, email)
		wg.Add(1)
		//goroutine thread
		go sendTicket(numberOfTickets, firstName, lastName, email)

		firstNames := printFirstNames()
		fmt.Printf("The first names %v\n", firstNames)

		if remainingTickets == 0 {
			fmt.Println("Our conference is currently booked out. Come back later.")
			//break
		}
	} else {
		if !isValidName {
			fmt.Println("firt name or last name you entered is too short")
		}
		if !isValidEmail {
			fmt.Println("email address you entered doesn't contain @ sign")
		}
		if !isValidTicketNumber {
			fmt.Println("number of tickets you entered is invalid")
		}
		//continue
	}
	//}
	wg.Wait()
}

func greetUsers() {
	fmt.Printf("Welcome to %v booking application.\nWe have total of %v tickets and %v are still available.\nGet your tickets here to attend\n", conferenceName, conferenceTickets, remainingTickets)
}

func getUserInput() (string, string, string, uint) {
	var firstName string
	var lastName string
	var email string
	var numberOfTickets uint

	fmt.Println("Enter Your First Name: ")
	fmt.Scanln(&firstName)

	fmt.Println("Enter Your Last Name: ")
	fmt.Scanln(&lastName)

	fmt.Println("Enter Your Email: ")
	fmt.Scanln(&email)

	fmt.Println("Enter number of tickets: ")
	fmt.Scanln(&numberOfTickets)

	return firstName, lastName, email, numberOfTickets
}

func bookTicket(numberOfTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - numberOfTickets

	// create user map
	var user = User{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: numberOfTickets,
	}

	bookings = append(bookings, user)

	fmt.Printf("Thank you %v %v for booking %v tickets. you will recive a confirmation email at %v\n", firstName, lastName, numberOfTickets, email)
	fmt.Printf("%v tickets is remaining for %v\n", remainingTickets, conferenceName)
}

func printFirstNames() []string {
	firstNames := []string{}

	for _, booking := range bookings {
		firstNames = append(firstNames, booking.firstName)
	}

	return firstNames
}

func sendTicket(numberOfTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second)
	var ticket = fmt.Sprintf("%v tickets for %v %v", numberOfTickets, firstName, lastName)
	fmt.Println("#################")
	fmt.Printf("Sending ticket:\n %v \nto email address %v\n", ticket, email)
	fmt.Println("#################")
	wg.Done()
}
