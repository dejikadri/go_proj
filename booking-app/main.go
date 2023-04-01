package main 
import "fmt"

func main(){
	
	var conferenceName = "GopherCon"
	const conferenceTickets = 60
	var remainingTickets uint = 60

	var firstName string
	var lastName string
	var email string
	var userTickets uint
	var arrBookings [60]string

	
	// get user input
	fmt.Println("Please Enter your first name")
	fmt.Scan(&firstName)

	fmt.Println("Please Enter your last name")
	fmt.Scan(&lastName)

	fmt.Println("Please Enter your email")
	fmt.Scan(&email)

	fmt.Println("Please Enter  number of tickets")
	fmt.Scan(&userTickets)

	remainingTickets = remainingTickets - userTickets
	arrBookings[0] = firstName +" " + lastName

	fmt.Printf("Entire array: %v \n", arrBookings)
	fmt.Printf("First element: %v \n", arrBookings[0])
	fmt.Printf("Array type: %T \n", arrBookings)
	fmt.Printf("Array length: %v \n", len(arrBookings))

	fmt.Printf("Thank you %v %v for buying %v tickets for our conference( %v), you will receive an email at %v \n", firstName, lastName, userTickets, conferenceName, email)
	fmt.Printf("We have a total of %v tickets and are still %v available\n", conferenceTickets, remainingTickets)




	// var conferenceName = "GopherCon"
	// const conferenceTickets = 60
	// var remainingTickets = 60

	// fmt.Printf("Welcome to our %v conference booking app\n", conferenceName)
	// fmt.Printf("We have a total of %v tickets and are still %v available\n", conferenceTickets, remainingTickets)
	// fmt.Println("Get your tickets here to attend")

}

