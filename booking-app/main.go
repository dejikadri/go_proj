package main 
import "fmt"

func main(){
	

	var userName string
	var numTickets int
	userName ="Deji"
	numTickets = 15


	fmt.Printf("User %v bought %v tickets\n", userName, numTickets)
	fmt.Printf("the data type of username is %T\n", userName)
	fmt.Printf("the data type of numTickets is %T\n", numTickets)


	// var conferenceName = "GopherCon"
	// const conferenceTickets = 60
	// var remainingTickets = 60

	// fmt.Printf("Welcome to our %v conference booking app\n", conferenceName)
	// fmt.Printf("We have a total of %v tickets and are still %v available\n", conferenceTickets, remainingTickets)
	// fmt.Println("Get your tickets here to attend")

}

