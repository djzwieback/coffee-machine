package main

type Beverage struct {
	price int
	water int
	milk  int
	beans int
}

func getLatte() Beverage {
	return Beverage{
		price: 7,
		water: 350,
		milk:  75,
		beans: 20,
	}
}
func getCappuccino() Beverage {
	return Beverage{
		price: 6,
		water: 200,
		milk:  100,
		beans: 12,
	}
}
func getEspresso() Beverage {
	return Beverage{
		price: 4,
		water: 250,
		milk:  0,
		beans: 16,
	}
}

func GetCoffee(coffeeType CoffeeType) Beverage {
	switch coffeeType {
	case cappuccino:
		return getCappuccino()
	case espresso:
		return getEspresso()
	case latte:
		return getLatte()
	}
	return Beverage{
		price: 0,
		water: 0,
		milk:  0,
		beans: 0,
	}
}
