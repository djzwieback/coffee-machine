package main

import "fmt"

type CoffeeMachine struct {
	cups  int
	money int
	water int
	milk  int
	beans int
}

func CreateMachine() CoffeeMachine {
	return CoffeeMachine{
		cups:  9,
		money: 550,
		water: 400,
		milk:  540,
		beans: 120,
	}
}

func (machine *CoffeeMachine) PrintMachineState() {
	fmt.Println("The coffee machine has:")
	fmt.Printf("%d ml of water\n", machine.water)
	fmt.Printf("%d ml of milk\n", machine.milk)
	fmt.Printf("%d g of coffee beans\n", machine.beans)
	fmt.Printf("%d disposable cups\n", machine.cups)
	fmt.Printf("$%d of money\n", machine.money)

}

func (machine *CoffeeMachine) takeMoney() {
	fmt.Printf("I gave you $%d\n", machine.money)
	machine.money = 0
}

func (machine *CoffeeMachine) ProcessOrder(beverage Beverage) {
	if beverage.water > machine.water || beverage.milk > machine.milk || beverage.beans > machine.beans || machine.cups <= 0 {
		return
	}
	machine.water -= beverage.water
	machine.milk -= beverage.milk
	machine.beans -= beverage.beans
	machine.money += beverage.price
	machine.cups--
}
