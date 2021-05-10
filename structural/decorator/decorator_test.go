package decorator

import "testing"

func TestPattern(t *testing.T) {
	pizza := &veggeMania{}

	otherPizza := &bigOne{}

	// add cheese topping
	pizzaWithCheese := &cheeseTopping{
		pizza: pizza,
	}

	// add tomato topping
	pizzaWithCheeseAndTomato := &tomatoTopping{
		pizza: pizzaWithCheese,
	}

	t.Log("Price of veggeMania with tomato and cheese topping is", pizzaWithCheeseAndTomato.getPrice())

	// add pepper topping
	pizzaWithPepper := &pepperTopping{
		pizza: otherPizza,
	}

	t.Log("Price of bigOne with pepper topping is", pizzaWithPepper.getPrice())
}
