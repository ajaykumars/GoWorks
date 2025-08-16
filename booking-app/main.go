package main

import "fmt"

func main() {

	var conferenceName = "Go Confernce"
	const confernceTickets = 50
	var remainingTickets uint

	fmt.Printf("Total Tickets is of type : %T, Remainaing Tickets is of type : %T and Conference Name is of type : %T \n", confernceTickets, remainingTickets, conferenceName)

	fmt.Printf("Welcome to our %s booking application.\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and %v are remaining.", confernceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend.")


	var bookings [50] string	
		
	var firstName string
	var lastName string
	var emailId string

	var userTickets uint
	// Ask for User Name
	fmt.Println("Enter your First Name for booking : ")
	fmt.Scan(&firstName)

	fmt.Println("Enter your Last Name for booking : ")
	fmt.Scan(&lastName)

	fmt.Println("Enter your Email ID for booking : ")
	fmt.Scan(&emailId)

	fmt.Println("Enter number of tickets needed : ")
	fmt.Scan(&userTickets)

	remainingTickets = confernceTickets - userTickets
	bookings[0]  = firstName + " " + lastName

	fmt.Printf("Bookings : %v \n", bookings)
	fmt.Printf("First Booking : %v \n", bookings[0])
	fmt.Printf("Type of Booking : %T \n", bookings)
	fmt.Printf("Length of Bookings : %v \n", len(bookings))

	fmt.Printf("User %s %s with Email ID : %s booked %v tickets \n ", firstName, lastName, emailId, userTickets)
	fmt.Println("Thanks you! You will receive confirmation on you email id shortly...\n")

	fmt.Printf("Remaining Tickets : %v", remainingTickets)

}
