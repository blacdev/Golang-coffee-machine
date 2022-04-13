package main

import (
	F "fmt"
	"math"
	"sort"
)

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
	var water, coffee, milk, machinewater, machinecoffee, machinemilk, machinecups int

	water, milk, coffee, machinewater, machinecoffee, machinemilk, machinecups = machineIngredient()

	F.Println("machineIngredient", water, milk, coffee, machinewater, machinecoffee, machinemilk, machinecups)
	F.Println("Write how many cups of coffee you will need:")
	F.Scan(&numOfCups)

	rwater = (actualWater * numOfCups)
	rmilk = (actualMilk * numOfCups)
	rcoffee = (actualCoffee * numOfCups)

	F.Printf("response %d %d %d \n", rwater, rmilk, rcoffee)
	if rmilk == 0 && rwater == 0 && rcoffee == 0 && machinemilk == 0 && machinecoffee == 0 && machinewater == 0 && numOfCups == 0 {
		F.Println("Yes, I can make that amount of coffee.")
	} else {
		fwater := round(float64(machinewater) / float64(rwater))
		fmilk := round(float64(machinemilk) / float64(rmilk))
		fcoffee := round(float64(machinecoffee) / float64(rcoffee))

		F.Printf("makeable cups %d %d %d \n", fwater, fmilk, fcoffee)

		array := []int{fwater, fmilk, fcoffee}
		sort.Ints(array)
		cup := array[0]
		F.Println("array index one", cup)
		//F.Printf("For %v cups of coffee you will need:\n", numOfCups)
		//F.Printf("%v ml of water\n", rwater)
		//F.Printf("%v ml of milk\n", rcoffee)
		//F.Printf("%v g of coffee beans\n", rmilk)

		if machinecups <= cup && numOfCups > cup {
			cup = cup - machinecups
			if cup < 0 {
				cup = 0
				F.Printf("No, I can only make %d cups of coffee\n", cup)
			} else {
				F.Printf("No, I can only make %d cups of coffee\n", cup)
			}

		} else if machinecups == cup && numOfCups == machinecups {
			// can give the exact amount of cups requested for
			F.Println("Yes, I can make that amount of coffee ")
		} else if machinecups > numOfCups {
			// subtract cups from cup
			// the answer is the amount of cups that can be made more than the amount  requested for
			F.Println(machinecups)
			F.Println(cup)
			remainderofMachinecups := machinecups - numOfCups
			F.Println(remainderofMachinecups)
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

	water = machinewater / actualWater    // Tells us how many cups of coffee can be made from what is in the machine
	milk = machinemilk / actualMilk       // Tells us how many cups of coffee can be made from what is in the machine
	coffee = machinecoffee / actualCoffee // Tells us how many cups of coffee can be made from what is in the machine
	array := []int{water, milk, coffee}
	sort.Ints(array)
	machinecups = array[0]
	F.Println(array)
	F.Println(machinecups)
	return
}

func round(x float64) (D int) {
	F.Println("value of x", x)
	d := math.Ceil(x)
	s := math.Floor(x)
	g := d - x
	F.Println("weting dey sup", g)
	if g <= 0.1 {
		D = int(d)
	} else {
		D = int(s)
	}
	return
}
