package main

import (
	"fmt"
	"strings"
)

func main() {
	conferenceName := "Go conference"
	const conferenceTickets int = 50
	var remainingTickets uint = 50

	type UserData struct {
		firstName       string
		lastName        string
		email           string
		numberOfTickets uint
	}

	bookings := make([]UserData, 0)

	greetUsers(conferenceName)

	fmt.Println(conferenceName)

	fmt.Println("Welcome to", conferenceName, "booking application")
	fmt.Printf("We have a total of %v tickets and %v are still available.\n", conferenceTickets, remainingTickets)
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

		fmt.Println("Enter number of tickets ")
		fmt.Scan(&userTickets)

		isValidName := len(firstName) >= 2 && len(lastName) >= 2
		isValidEmail := strings.Contains(email, "@")
		isValidTicketNumber := userTickets > 0 && userTickets <= remainingTickets

		if isValidName && isValidEmail && isValidTicketNumber {
			remainingTickets = remainingTickets - userTickets

			var userData = UserData{
				firstName:       firstName,
				lastName:        lastName,
				email:           email,
				numberOfTickets: userTickets,
			}

			bookings = append(bookings, userData)
			fmt.Printf("List of bookings is %v\n", bookings)

			fmt.Printf("The whole slice: %v\n", bookings)
			fmt.Printf("The first value: %v\n", bookings[0])
			fmt.Printf("Slice type:  %T\n", bookings)
			fmt.Printf("Slice length: %v\n", len(bookings))

			fmt.Printf("Thank %v %v for booking %v tickets. You will receive a confirmation email at %v\n", firstName, lastName, userTickets, email)
			fmt.Printf("%v tickets remaining for %v\n", remainingTickets, conferenceName)

			firstNames := []string{}
			for _, booking := range bookings {

				firstNames = append(firstNames, booking.firstName)
			}
			fmt.Printf("The first names of bookings are: %v\n", firstNames)

			fmt.Printf("These are all our bookings: %v\n", bookings)

			if remainingTickets == 0 {
				// end program
				fmt.Println("Our conference is booked out. Come back next year.")
				break
			}

			sendTicket(userTickets, firstName, lastName, email)

		} else {
			if !isValidName {
				fmt.Println("First name or last name is too short")
			}
			if !isValidEmail {
				fmt.Println("Email address you entered does not contain @")
			}
			if !isValidTicketNumber {
				fmt.Println("Number of tickets you entered is invalid")
			}
			fmt.Println("Your input data is invalid")

		}

	}
}

func greetUsers(conferenceName string) {
	fmt.Printf("Welcome to %v booking application\n", conferenceName)
}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Println()
	fmt.Printf("Sending ticket:\n %v to email address %v\n", ticket, email)
	fmt.Println()
}
