package decorator

// wiki: https://en.wikipedia.org/wiki/Decorator_pattern

/**
 * Decorator is a structural design pattern that attch new behaviors to objects by
 * placing thses objects inside special wrapper objects that contain the behaviors.
 */

type pizza interface {
	getPrice() int
}

type veggeMania struct{}

func (p *veggeMania) getPrice() int {
	return 15
}

type bigOne struct{}

func (b *bigOne) getPrice() int {
	return 50
}

// concrete decorator
type tomatoTopping struct {
	pizza pizza
}

func (c *tomatoTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 7
}

// concrete decorator
type cheeseTopping struct {
	pizza pizza
}

func (c *cheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 10
}

// concrete decorator
type pepperTopping struct {
	pizza pizza
}

func (p *pepperTopping) getPrice() int {
	pizzaPrice := p.pizza.getPrice()
	return pizzaPrice + 20
}
