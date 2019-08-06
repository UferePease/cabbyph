package model

import (
	"fmt"
	"math"
	"math/rand"
	"time"
)

type Ride struct {
    PickupLoc, DropoffLoc string
	Fare, ReceivedPayment, Tip float64
	Passenger Customer
	Completed bool
	StartTime, EndTime int64
}

// display message
func showMessage(msg string)  {
	fmt.Println(msg)
}

func (ride *Ride) CalculateFare()  {
	showMessage("\n**************************** FARE ESTIMATION *********************************")
	// Generating a random integer between 800 to 5000
	rand.Seed(time.Now().UnixNano())
	randFare := 800 + rand.Float64() * (5000 - 800)
	estimatedRate := math.Round(randFare * 100) / 100
	ride.Fare = estimatedRate
	fmt.Printf("This ride %v --> %v will cost %v\n", ride.PickupLoc, ride.DropoffLoc, estimatedRate)
}

func (ride *Ride) Ride()  {
	// Epoch time in milliseconds
	showMessage("\n****************************  RIDING  *********************************")
	started := time.Now().Unix()
	ride.StartTime = started
	fmt.Printf("Ride started on %v\n", started)
	fmt.Printf("Ride %v -->> %v in progress...\n", ride.PickupLoc, ride.DropoffLoc)

	// sleep for 10 secs
	time.Sleep(10 * time.Second)

	ended := time.Now().Unix()
	ride.EndTime = ended
	fmt.Printf("Ride successfully ended on %v\n", ended)
}
