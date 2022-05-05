package main

import (
	"fmt"
	"os"
)

func main() {
	var component = []int{400, 540, 120, 9, 550}

	actionChoice(component)

}
func actionChoice(component []int) []int {
	var action string
	fmt.Printf("Write action (buy, fill, take, remaining, exit):")
	fmt.Scan(&action)

	switch action {
	case "buy":
		{
			buy(component)
			actionChoice(component)
		}
	case "fill":
		{
			fill(component)
			actionChoice(component)
		}
	case "take":
		{
			take(component)
			actionChoice(component)
		}
	case "remaining":
		{
			remaining(component)
			actionChoice(component)
		}
	case "exit":
		{
			os.Exit(0)
		}
	default:
		{
			actionChoice(component)
		}
	}
	return component
}

func espresso(component []int) {
	component[0] -= 250
	component[1] -= 0
	component[2] -= 16
	component[3] -= 1
	component[4] += 4
	return
}

func latte(component []int) {
	component[0] -= 350
	component[1] -= 75
	component[2] -= 20
	component[3] -= 1
	component[4] += 7
	return
}

func capuccino(component []int) {
	component[0] -= 200
	component[1] -= 100
	component[2] -= 12
	component[3] -= 1
	component[4] += 6
	return
}
func resourcesEspresso(component []int) {
	if component[0] >= 250 && component[2] >= 16 && component[3] >= 1 {
		fmt.Println("I have enough resources, making you a coffee!")
		espresso(component)
	} else if component[0] < 250 {
		fmt.Println("Sorry, not enough water!")
	} else if component[2] < 16 {
		fmt.Println("Sorry, not enough coffee beans!")
	} else if component[3] < 1 {
		fmt.Println("Sorry, not enough cups!")
	}
}
func resourcesLatte(component []int) {
	if component[0] >= 350 && component[1] >= 75 && component[2] >= 20 && component[3] >= 1 {
		fmt.Println("I have enough resources, making you a latte!")
		latte(component)
	} else if component[0] < 350 {
		fmt.Println("Sorry, not enough water!")
	} else if component[1] < 75 {
		fmt.Println("Sorry, not enough milk!")
	} else if component[2] < 20 {
		fmt.Println("Sorry, not enough coffee beans!")
	} else if component[3] < 1 {
		fmt.Println("Sorry, not enough cups!")
	}
}
func resourcesCapuccino(component []int) {
	if component[0] >= 200 && component[1] >= 100 && component[2] >= 12 && component[3] >= 1 {
		fmt.Println("I have enough resources, making you a capuccino!")
		capuccino(component)
	} else if component[0] < 200 {
		fmt.Println("Sorry, not enough water!")
	} else if component[1] < 100 {
		fmt.Println("Sorry, not enough milk!")
	} else if component[2] < 12 {
		fmt.Println("Sorry, not enough coffee beans!")
	} else if component[3] < 1 {
		fmt.Println("Sorry, not enough cups!")
	}
}
func buy(component []int) []int {
	var drink string
	fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, back - to main menu:")
	fmt.Scan(&drink)
	switch drink {
	case "1":
		{
			resourcesEspresso(component)
		}
	case "2":
		{
			resourcesLatte(component)
		}
	case "3":
		{
			resourcesCapuccino(component)
		}
	case "back":
		{
			actionChoice(component)
		}
	default:
		{
			buy(component)
		}
	}
	return component
}

func fill(component []int) []int {
	var addWater int
	var addMilk int
	var addBeans int
	var addCups int
	fmt.Println("Write how many ml of water you want to add: ")
	fmt.Scan(&addWater)
	fmt.Println("Write how many ml of milk you want to add: ")
	fmt.Scan(&addMilk)
	fmt.Println("Write how many grams of coffee beans you want to add: ")
	fmt.Scan(&addBeans)
	fmt.Println("Write how many disposable cups of coffee you want to add: ")
	fmt.Scan(&addCups)
	component[0] += addWater
	component[1] += addMilk
	component[2] += addBeans
	component[3] += addCups
	return component
}
func take(component []int) []int {
	fmt.Printf("%s%d%s", "I gave you $", component[4], "\n")
	component[4] -= component[4]
	return component
}

func remaining(component []int) {

	fmt.Println("The coffee machine has:")
	fmt.Println(component[0], "ml of water")
	fmt.Println(component[1], "ml of milk")
	fmt.Println(component[2], "g of coffee beans")
	fmt.Println(component[3], "disposable cups")
	fmt.Printf("$%d%s", component[4], " of money\n\n")
}
