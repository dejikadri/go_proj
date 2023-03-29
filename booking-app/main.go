package main 
import "fmt"

func main(){
	var conferenceName = "GopherCon"
	const conferenceTickets = 60
	var remainingTickets = 60

	fmt.Printf("Welcome to our %v conference booking app\n", conferenceName)
	fmt.Printf("We have a total of %v tickets and are still %v available", conferenceTickets, remainingTickets)
	fmt.Println("Get your tickets here to attend")
}

