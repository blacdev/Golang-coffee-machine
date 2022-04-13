package main

import (
	F "fmt"
	"math"
	"sort"
)

// these are the actual values required in a making a cup of coffee
const (
	actualWater  int = 200
	actualCoffee int = 15
	actualMilk   int = 50
)

func main() {
	// write your code here
	Greeting()
	ingredientCalculator()
}

// First thing to show when the machine is turned on
func Greeting() {
	F.Println("Starting to make a coffee")
	F.Println("Grinding coffee beans")
	F.Println("Boiling water")
	F.Println("Mixing boiled water with crushed coffee beans")
	F.Println("Pouring coffee into the cup")
	F.Println("Pouring some milk into the cup")
	F.Println("Coffee is ready!")
}

func ingredientCalculator() {
	var numOfCups int
	var rwater, rcoffee, rmilk int
	var machinewater, machinecoffee, machinemilk, machinecups int

	// varaibles returned from machine ingredient function
	_, _, _, machinewater, machinecoffee, machinemilk, machinecups = machineIngredient()

	F.Println("Write how many cups of coffee you will need:")
	// take in put from the user to know how many cups they need
	F.Scan(&numOfCups)

	// calculate how much ingredient is required to make the number of cups requested for
	rwater = (actualWater * numOfCups)
	rmilk = (actualMilk * numOfCups)
	rcoffee = (actualCoffee * numOfCups)

	// check if all values return zero
	if rmilk == 0 && rwater == 0 && rcoffee == 0 && machinemilk == 0 && machinecoffee == 0 && machinewater == 0 && numOfCups == 0 {
		F.Println("Yes, I can make that amount of coffee.")
	} else {
		// round up or down the the values afeter dividing the ingredients in the machine required amount of ingredient
		fwater := round(float64(machinewater) / float64(rwater))
		fmilk := round(float64(machinemilk) / float64(rmilk))
		fcoffee := round(float64(machinecoffee) / float64(rcoffee))

		// sort the the values to know the smallest value which would be the amount of cups of coffee that can be made
		array := []int{fwater, fmilk, fcoffee}
		sort.Ints(array)
		cup := array[0]

		//F.Printf("For %v cups of coffee you will need:\n", numOfCups)
		//F.Printf("%v ml of water\n", rwater)
		//F.Printf("%v ml of milk\n", rcoffee)
		//F.Printf("%v g of coffee beans\n", rmilk)

		// check if the number of cups the machine can make is less or equal to the numbercups that can be made based on the cups requested for
		if machinecups <= cup && numOfCups > cup {
			cup = cup - machinecups

			// check if cup is less than 0
			if cup < 0 {
				cup = 0
				F.Printf("No, I can only make %d cups of coffee\n", cup)
			} else {
				F.Printf("No, I can only make %d cups of coffee\n", cup)
			}

			// check if all cups are of the same value
		} else if machinecups == cup && numOfCups == machinecups {
			// can give the exact amount of cups requested for
			F.Println("Yes, I can make that amount of coffee ")

			// check if machine cups is greater than number of cups requested for
		} else if machinecups > numOfCups {
			// subtract cups from cup
			// the answer is the amount of cups that can be made more than the amount  requested for
			remainderofMachinecups := machinecups - numOfCups
			F.Printf("Yes, I can make that amount of coffee (and even %d more than that)\n", remainderofMachinecups)
		} else {
			F.Printf("No, I can make only %d cups of coffee\n", machinecups)
		}
	}
}
func machineIngredient() (water, milk, coffee, machinewater, machinecoffee, machinemilk, machinecups int) {
	F.Println("Write how ml of water the coffee machine has:")
	F.Scan(&machinewater)
	F.Println("Write how many ml of milk the coffee machine has:")
	F.Scan(&machinemilk)
	F.Println("Write how many grams of coffee beans the machine has:")
	F.Scan(&machinecoffee)

	// divide the amount of ingredients in the machine by the actual amount of ingredient neede to make a cup of coffee
	water = machinewater / actualWater    // Tells us how many cups of coffee can be made from what is in the machine
	milk = machinemilk / actualMilk       // Tells us how many cups of coffee can be made from what is in the machine
	coffee = machinecoffee / actualCoffee // Tells us how many cups of coffee can be made from what is in the machine
	// sort the the values to know the smallest value which would be the amount of cups of coffee that can be made
	array := []int{water, milk, coffee}
	sort.Ints(array)
	machinecups = array[0]
	return
}

func round(x float64) (D int) {
	// round up x
	d := math.Ceil(x)

	// round down x
	s := math.Floor(x)

	// subtract x from d
	g := d - x

	// check if the value is less than or equal to .004%
	if g <= 0.004 {
		D = int(d)
	} else {
		D = int(s)
	}
	return
}
