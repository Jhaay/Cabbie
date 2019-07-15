package main

import (
	"fmt"
	"time"
	
)

func main() {

	fmt.Println("Welcome to Cabbie")
	fmt.Printf("\n")
	fmt. Println("Our Operational areas include: \n Alakahia \n Aluu \n Choba \n Rumosi \n Rumola \n Mgbuoba")
	//Get users name.
	var name string
	fmt.Println("Enter your name: ")
	fmt.Scanf("%s", &name)
	fmt.Printf("\n")

	// Get users current location.
	fmt.Println(
		"Hello " + name + " please choose your pick up location.")

	var Location string
	fmt.Scanf("%s", &Location)
	fmt.Printf("\n")
	fmt.Println("Your pickup point is " + Location + ", Choose your drop off location.")

	//Get users drop off location.
	var Drop string
	fmt.Scanf("%s", &Drop)
	fmt.Printf("\n")

	var fare int


	// var Aluu string
	// var Choba string
	// var Mgbuoba string
	// var Rumosi string
	// var Rumokoro string
	// var Rumola string

	switch Location {

		case "Aluu":

			switch Drop {
				case "Alakahia":
					fare = 1000
				case "Choba":
					fare = 500
				case "Mgbuoba":
					fare = 2000
				case "Rumola":
					fare=2500
				case "Rumosi":
					fare=1000
				case "Rumokoro":
					fare=1500
			default:
					fmt.Println("Location false")
					fmt.Println("Bye Laters")
			return
		}
		case "Alakahia":
		
			switch Drop {
			case "Aluu":
				fare=1000
			case "Choba":
				fare=500
			case "Mgbuoba":
				fare=1500
			case "Rumola":
				fare=2000
			case "Rumosi":
				fare=500
			case "Rumokoro":
				fare=1000	
			default:
				fmt.Println("Location false")
				fmt.Println("Bye Laters")
			return
		}
		case "Choba":

		switch Drop {
			case "Alakahia":
				fare=500
			case "Aluu":
				fare=500
			case "Mgbuoba":
				fare=1500
			case "Rumola":
				fare=2000
			case "Rumosi":
				fare=1000
			case "Rumokoro":
				fare=1500	
			default:
				fmt.Println("Location false")
				fmt.Println("Bye Laters")
			return
		}
	case "Mgbuoba":
		
		switch Drop {
			case "Alakahia":
				fare=1500
			case "Aluu":
				fare=2000
			case "Choba":
				fare=1500
			case "Rumola":
				fare=1500
			case "Rumosi":
				fare=500
			case "Rumokoro":
				fare=1000	
			default:
				fmt.Println("Location false")
				fmt.Println("Bye Laters")
			return
		}
	case "Rumokoro":

		switch Drop {
			case "Alakahia":
				fare=1500
			case "Aluu":
				fare=1500
			case "Choba":
				fare=1500
			case "Mgbuoba":
				fare=1000
			case "Rumola":
				fare=1000
			case "Rumosi":
				fare=500	
			default:
				fmt.Println("Location false")
				fmt.Println("Bye Laters")
			return
		}
	case "Rumola":

		switch Drop {
			case "Alakahia":
				fare=2000
			case "Aluu":
				fare=2500
			case "Choba":
				fare=2000
			case "Mgbuoba":
				fare=1500
			case "Rumosi":
				fare=500
			case "Rumokoro":
				fare=1000	
			default:
				fmt.Println("Location false")
				fmt.Println("Bye Laters")
			return
		

		}
	case "Rumosi":

		switch Drop {
			case "Alakahia":
				fare=1000
			case "Aluu":
				fare=1500
			case "Choba":
				fare=1500
			case "Mgbuoba":
				fare=2000
			case "Rumokoro":
				fare=1000
			case "Rumola":
				fare=500	
		default:
			fmt.Println("Location false")
			fmt.Println("Bye Laters")
		return

		}

	default:
		fmt.Println("Location false")
		fmt.Println("Bye Laters")
		return

	}
	
fmt.Println("Your fare from " + Location + " to " + Drop + " is: " )
fmt.Print(fare)

	fmt.Printf("\n")
tick := time.Tick(100* time.Millisecond)
boom := time.After(500 * time.Millisecond)
	for{
		select{
			case<-tick:
			fmt.Println("Anticipating Arrival.............................")
			case<-boom:
		fmt.Println("We have arrived")
		return
		default:
		fmt.Println("..........................................................................................................................")
		
		time.Sleep(10000 * time.Millisecond)
break
		}
	break
	}
fmt. Println("We have arrived")
fmt.Printf("\n")

	var Change int
	var Fare int
	var Correct int
	fmt.Println("Enter your Fare")
	fmt.Scanf("%d", &Fare)
	Change = Fare - fare
	if Fare > fare {

		fmt.Println("Your Change is:")
		fmt.Print(Change)
		fmt.Printf("\n")

	}else if Fare < fare {
		
		for i:=1; i<fare; i++{
		fmt. Println("Your fare is incorrect, Please input correct amount")
			
			fmt.Println("Enter correct fare")
			fmt.Scanf("%d", &Correct)
			Correct = Change

		if i%4==0{
			 fmt. Println("You have been reported to the police") 
			 break
		  }
		
		if i == fare{
			break
		}
		
		}
		return
		} 
		
	var tip int
	fmt.Println("Enter your tip: ")
	fmt.Scanf("%d", &tip)
	fmt.Printf("\n")
	if tip > fare{
		fmt.Println("Gracias Mucho")
		fmt.Printf("\n")

	}else if tip < 100 {
		fmt. Println("You are a Stingy ass bastard")
		fmt.Printf("\n")

	}

fmt.Printf("\n")
fmt.Println("THANKS FOR USING CABBIE")
}