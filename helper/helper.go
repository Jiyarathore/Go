package helper

import "strings"

// exporting a varialble= make it available for all packages in app = capitalize first letter

func ValidateUserInput(firstName string, lastName string, email string, userTickets uint, remainingTickets uint) (bool, bool, bool) {

	isValidName := len(firstName) >= 2 && len(lastName) >= 2
	isValidEamil := strings.Contains(email, "@")
	isValidTicketNumeber := userTickets > 0 && userTickets <= remainingTickets

	return isValidName, isValidEamil, isValidTicketNumeber

}
// for multiple files run = go run main.go helper.go
// or go gun .