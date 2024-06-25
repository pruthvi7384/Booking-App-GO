package services

import (
	"bookingApp/utils"
	"fmt"
)

var totalTicket int = 50
var bookedTicket []Booker
var remaningTicket int = 50

// Main Service
func TicketBookService() {
	appIntroduction()
	for i := 1; i <= remaningTicket; i++ {
		ticketDetails()
		bookTicket()
	}
}

// CLI App Introduction
func appIntroduction() {
	fmt.Printf("\n************ Welcome to the Ticket Booking App! *************\n")

	fmt.Printf("\n---------------- Conference Details ------------------\n")
	fmt.Printf("\nName : %v\n", utils.ConferanceName)
	fmt.Printf("\nCity : %v\n", utils.ConferanceCity)
	fmt.Printf("\nSponsor : %v\n", utils.Sponser)
	fmt.Printf("\n-----------X---- Conference Details ----X--------------\n")
}

// Ticket Details
func ticketDetails() {
	fmt.Printf("\n---------------- Ticket Details -----------------\n")
	fmt.Printf("\nTotal Tickets : %v\n", totalTicket)
	fmt.Printf("\nRemaining Tickets : %v\n", remaningTicket)
	fmt.Printf("\nBooked Tickets : %v\n", totalTicket-remaningTicket)
	fmt.Printf("\n-----------X---- Ticket Details ----X-------------\n")
}

// Book Ticket
func bookTicket() {
	var booker Booker
	fmt.Printf("\n---------------- Ticket Booking Form -----------------\n")
	fmt.Printf("\nWhat is Your First Name ?\n")
	fmt.Scan(&booker.firstName)
	fmt.Printf("\nWhat is your Last Name ?\n")
	fmt.Scan(&booker.lastName)
	fmt.Printf("\nWhat is your Email ?\n")
	fmt.Scan(&booker.email)
	fmt.Println("\nHow Many Ticket You What To Book ?")
	fmt.Scan(&booker.ticket)

	// Ticket Booking Validation
	if booker.ticket <= remaningTicket {
		// Book The Ticket
		bookedTicket = append(bookedTicket, booker)

		// Ticket Logic
		ticketDetailsManage()

		// Booking Details
		bookingDetails(booker)
	} else {
		fmt.Printf("\nInsufficient Tickets to Book!\nWe apologize for the inconvenience\n")
	}
}

// Booking Details
func bookingDetails(bookedTicket Booker) {
	fmt.Printf("\nYou're Welcome! Thank You for Booking with Us!\n")
	fmt.Printf("\nYour Booking Deatails : \n")
	fmt.Printf("\nName : %v\nEmail : %v\nTickets : %v\n", bookedTicket.firstName+" "+bookedTicket.lastName, bookedTicket.email, bookedTicket.ticket)
}

// Ticket Details Managing
func ticketDetailsManage() {
	var ticketBooked int
	for _, value := range bookedTicket {
		ticketBooked = ticketBooked + value.ticket
	}
	remaningTicket = remaningTicket - ticketBooked
}

// Booker Details
type Booker struct {
	firstName string
	lastName  string
	email     string
	ticket    int
}
