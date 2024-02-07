package main

import (
	"bookingapp/helper"
	"fmt"
	"time"
	// "strconv"
	// "strings"
)

// Package level Variavles
// accesesd inside any of function

const conferenceTickets int = 50

var conferenceName = "Go conferenece"
var remainingTickets uint = 50

// slice of strings
// var bookings = []string{}
// slice of map= lists of map          initial size
// var bookings =make([]map[string]string, 0)

// lists of Userdata structure
var bookings = make([]UserData, 0)

type UserData struct {
	firstName       string
	lastName        string
	email           string
	numberOfTickets uint
}

func main() {
	// var conferenceName string = "Go conferenece"
	//we can use this notation as well but only for vaiables not for constatnst
	// conferenceName := "Go conferenece"
	// fmt.Println(conferenceName)

	// const conferenceTickets int = 50

	// remainingTickets cant be negative so we use unsignedint which will only take whole no.
	// var remainingTickets uint = 50

	// array
	// var bookings = [50]string{"NAma", "Nicole", "Peter"}

	// var bookings [50]string
	// bookings[0]="Nana"

	//    slice(same as vector)
	// var bookings []string
	// alternative symbol
	// var bookings =[]string{}
	// bookings:=[]string{}

	// greetUsers(conferenceName, conferenceTickets, remainingTickets)
	greetUsers()

	// %T is placeholder for the type of variable we are referencing to
	// fmt.Printf("conferenceTickets is %T, remainingTickets is %T,conferenceName is %T\n ", conferenceTickets, remainingTickets, conferenceName)

	// fmt.Println("Welcome to", conferenceName, "booking application")
	// fmt.Println("we have total of",conferenceTickets, "tickets and" ,remainingTickets, "are still available")
	// fmt.Println("Get your tickets here to attend")

	// printf and placeholders
	// fmt.Printf("Welcome to %v booking application\n", conferenceName)
	// fmt.Printf("we have total of %v tickets and %v are still available\n",conferenceTickets, remainingTickets)
	// fmt.Println("Get your tickets here to attend")

	// for remainingTickets>0 && len(bookings)< 50{
	// or
	for {

		firstName, lastName, email, userTickets := getUserInput()

		isValidName, isValidEamil, isValidTicketNumeber := helper.ValidateUserInput(firstName, lastName, email, userTickets, remainingTickets)

		if isValidName && isValidEamil && isValidTicketNumeber {

			// bookTicket(remainingTickets , userTickets, bookings, firstName , lastName ,email , conferenceName)
			bookTicket(userTickets, firstName, lastName, email)

			// to make apllication asynchronous
			// go ... starts a new goroutine, a goroutine is a lightweight thread managed by Go runtime
			// sending and generating tickets works simultaneoulsy in background
			go sendTicket(userTickets, firstName, lastName, email)

			firstNames := getFirstNames()
			fmt.Printf("First anems of bookings: %v\n", firstNames)

			if remainingTickets == 0 {
				// end program
				fmt.Printf("Your conference is booked out, come back next yr")
				break
			}

		} else {
			if !isValidName {
				fmt.Printf("first naem or last name is too short\n")
			}
			if !isValidEamil {
				fmt.Printf("Email adress doesnt contain @ sign\n")
			}
			if !isValidTicketNumeber {
				fmt.Printf("no. of tickets is invalid\n")
			}

			// fmt.Printf("Your input data is invalid try it again\n")
			// 	// continue agar if me hhota tho
		}

	}
}

// func greetUsers(confName string, confTickets int, remainingTickets uint) {
// 	fmt.Printf("Welcome to %v bookiing application\n", confName)
// 	fmt.Printf("we have total of %v tickets and %v are still available\n", confTickets, remainingTickets)
// 	fmt.Println("Get your tickets here to attend")
// }

func greetUsers() {
	fmt.Printf("Welcome to %v bookiing application\n", conferenceName)
	fmt.Printf("we have total of %v tickets and %v are still available\n", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

func getFirstNames() []string {

	firstNames := []string{}
	for _, booking := range bookings {
		// when booking was slice
		// var names = strings.Fields(booking)
		// var firstName = names[0]

		// when booking is map
		// var firstName=booking["firstName"]
		// firstNames = append(firstNames, firstName)

		// when booking is struct
		firstNames = append(firstNames, booking.firstName)
	}
	return firstNames
}

func getUserInput() (string, string, string, uint) {
	// if you dont assign value immediately Go will not know the datatype
	// so in that case use dattatype
	var firstName string
	var lastName string
	var email string
	var userTickets uint

	// ask user for their name
	fmt.Println("enter your first name")
	fmt.Scan(&firstName) //use pointer

	fmt.Println("enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("enter your email")
	fmt.Scan(&email)

	fmt.Println("enter no. of tickets")
	fmt.Scan(&userTickets)

	return firstName, lastName, email, userTickets
}

func bookTicket(userTickets uint, firstName string, lastName string, email string) {
	remainingTickets = remainingTickets - userTickets

	// create a map for user
	// var userData = make(map[string]string)
	//   data type for key     value

	// userData["firstName"]=firstName
	// userData["lastName"]=lastName
	// userData["email"]=email
	// userData["numberOfTickets"]=strconv.FormatUint(uint64(userTickets),10)

	// struct
	var userData = UserData{
		firstName:       firstName,
		lastName:        lastName,
		email:           email,
		numberOfTickets: userTickets,
	}

	// bookings[0]=firstName + " " + lastName for aaray
	// for slice
	// bookings = append(bookings, firstName+" "+lastName)

	bookings = append(bookings, userData)
	fmt.Printf("List of bookings is %v\n", bookings)

	// fmt.Printf("The whole array: %v\n", bookings)
	// fmt.Printf("The first value: %v\n", bookings[0])
	// fmt.Printf("aray type:%T\n", bookings);
	// fmt.Printf("araay lenght:%v\n", len(bookings))

	// fmt.Printf("The whole slice: %v\n", bookings)
	// fmt.Printf("The first value: %v\n", bookings[0])
	// fmt.Printf("slice type:%T\n", bookings);
	// fmt.Printf("slice lenght:%v\n", len(bookings))

	// userTickets=2
	fmt.Printf("Thank You %v %v for booking %v tickets, you will recieve confirmtaion email at %v\n", firstName, lastName, userTickets, email)

	fmt.Printf("%v tickets remaing for %v\n", remainingTickets, conferenceName)

}

func sendTicket(userTickets uint, firstName string, lastName string, email string) {
	time.Sleep(10 * time.Second) //stop execution for 10 secs
	var ticket = fmt.Sprintf("%v tickets for %v %v", userTickets, firstName, lastName)
	fmt.Printf("#########")
	fmt.Printf("Sending ticket:\n %v\n to email adderess %v\n", ticket, email)
	fmt.Println("#########")
}

// for index, booking
// but index not used so we used _ which is Blak idetifier to ignore varialbe w dont use
