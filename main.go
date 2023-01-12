package main

import (
	"fmt"
)

func main() {
	machine := CreateMachine()
	running := true
	for running {

		order := getOrder()
		switch order {
		case buy:
			var coffeeOrder string
			fmt.Println("What do you want to buy? 1 - espresso, 2 - latte, 3 - cappuccino, back - to menu: ")
			_, _ = fmt.Scanln(&coffeeOrder)
			if coffeeOrder == "back" {
				break
			}
			coffee := GetCoffee(CoffeeType(coffeeOrder))
			machine.ProcessOrder(coffee)
			break
		case fill:
			var water, milk, beans, cups int
			fmt.Println("Write how many ml of water you want to add: ")
			fmt.Scanln(&water)
			machine.water += water
			fmt.Println("Write how many ml of milk you want to add:")
			fmt.Scanln(&milk)
			machine.milk += milk
			fmt.Println("Write how many grams of coffee beans you want to add: ")
			fmt.Scanln(&beans)
			machine.beans += beans
			fmt.Println("Write how many disposable cups you want to add: ")
			fmt.Scanln(&cups)
			machine.cups += cups
			break
		case take:
			machine.takeMoney()
			break
		case remaining:
			machine.PrintMachineState()
			break
		case exit:
			running = false
		default:
			fmt.Println("Invalid command, please try again!")
		}
	}
}

func getOrder() Order {
	fmt.Println("Write action (buy, fill, take, remaining, exit): ")
	var order string
	_, err := fmt.Scanln(&order)
	if err != nil {
		return ""
	}
	return Order(order)
}
