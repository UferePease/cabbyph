package main

import (
	"bufio"
	"fmt"
	"github.com/uferepease/phcabby/model"
	"math"
	"os"
	"strings"
	"strconv"
)

type LocationList []string // a slice of locations

var availLocations = LocationList {
	"Choba",
	"Alakahia",
	"Rumuosi",
	"Rumuokoro",
	"Mgbuoba",
	"Aluu",
	"Rumuola",
}

// Create a single reader which can be called multiple times
var scanner = bufio.NewScanner(os.Stdin)

// display message
func showMessage(msg string)  {
	fmt.Println(msg)
}

// read user input
func getInput() string  {
	scanner.Scan()
	userInput := scanner.Text()
	return userInput
}

// read user float input
func getFloatInput() (float64, error)  {
	scanner.Scan()
	userInput, err := strconv.ParseFloat(scanner.Text(), 64)

	if err != nil {
		// fmt.Println("Invalid Input")
		return 0, err
	}
	return userInput, nil
}

// confirm that user location is valid
func isLocationValid (input string) bool  {
	for i := 0; i < len(availLocations); i++ {
		if strings.ToLower(input) == strings.ToLower(availLocations[i]) {
			return true
		}
	}
	return false
}

// process user payment
func processPayment(ride *model.Ride)  {
	showMessage("\n****************************	PAYMENT PROCESSING	*********************************")
	fmt.Printf("This Ride Fare is %v. Please enter the amount you are paying:\n", ride.Fare)
	isPayAccepted := false
	for i := 0; i < 5; i++ {
		userPay, err := getFloatInput()
		if err == nil {		// ensure that user inputs valid numbers
			if userPay < ride.Fare {
				fmt.Printf("Your payment is not correct. Your fare is %v. Enter a valid number that at least equals the ride fare: \n", ride.Fare)
			} else {
				isPayAccepted = true
				showMessage("Thank you for riding with us.")
				if userPay > ride.Fare {
					change := userPay - ride.Fare
					fmt.Printf("Your CHANGE is %v\n", math.Round(change * 100) / 100)
				}
				break
			}
		} else {
			showMessage("Invalid Input. Please enter a valid number")
		}
	}
	
	if isPayAccepted {		
		processTip(ride)	// ask for tip
	} else {
		showMessage("You will be reported to the police for failing to pay valid amount for this ride.")
	}
	return
}

// for tipping the driver
func processTip(ride *model.Ride)  {
	showMessage("\n****************************  TIPPING  *********************************")
	showMessage("Any tip for your driver. Please enter your tip here: ")
	userTip, err := getFloatInput()
	if err == nil {		// ensure that user inputs valid numbers
		if userTip <= 0 {
			ride.Tip = 0
			showMessage("You are a stingy fool")
		} else if userTip > 0 && userTip <= ride.Fare {
			ride.Tip = userTip
			fmt.Printf("Thank you, %v!!!\n", ride.Passenger.Name)
		} else if userTip > ride.Fare {
			ride.Tip = userTip
			showMessage("Gracias Mucho")
		}
	} else {
		showMessage("Invalid Input. You are a stingy fool")
	}
}

func main() {
	showMessage("**************************** CABBY *********************************")
	showMessage("Welcome to CABBY. Your convenient ride around town!!!")
	showMessage("********************************************************************")

	// declare new customer and ride objects/containers/structs
	var thisCustomer model.Customer
	var thisRide model.Ride

	showMessage("Please tell us your name")

	// get user's name
	thisCustomer.Name = getInput()

	thisRide.Passenger = thisCustomer

	showMessage("\n**************************** CHOOSE LOCATIONS *********************************")
	showMessage("Tell us your PICK-UP location")
	fmt.Printf("Kindly choose from any of these  available locations %v\n", availLocations)
	pickupInp := getInput()

	showMessage("Tell us your DROP-OFF location")
	fmt.Printf("Kindly choose from any of these  available locations %v\n", availLocations)
	dropoffInp := getInput()

	if isLocationValid(pickupInp) && isLocationValid(dropoffInp) {
		if pickupInp != dropoffInp {
			thisRide.PickupLoc = pickupInp
			thisRide.DropoffLoc = dropoffInp
	
			// calculate ride fare
			thisRide.CalculateFare()
	
			// All is well, lets do the ride
			thisRide.Ride()
	
			// ask for payment
			processPayment(&thisRide)
		} else {
			showMessage("Sorry, Your pickup and drop off locations cannot be the same")
		}

	} else {
		showMessage("Sorry, we do not cover your chosen route yet")
		showMessage("Bye, Laters!!!")
	}
	
}

