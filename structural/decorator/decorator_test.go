package decorator

import "testing"

func TestPattern(t *testing.T) {
	pizza := &veggeMania{}

	otherPizza := &bigOne{}

	pizzaWithCheese := &cheeseTopping{
		pizza: pizza,
	}

	pizzaWithCheeseAndTomato := &tomatoTopping{
		pizza: pizzaWithCheese,
	}

	t.Log("Price of veggeMania with tomato and cheese topping is", pizzaWithCheeseAndTomato.getPrice())

	pizzaWithPepper := &pepperTopping{
		pizza: otherPizza,
	}

	t.Log("Price of bigOne with pepper topping is", pizzaWithPepper.getPrice())
}
