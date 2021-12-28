package decorator

// wiki: https://en.wikipedia.org/wiki/Decorator_pattern
//
// Adds behavior to an object, statically or dynamically.

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

type tomatoTopping struct {
	pizza pizza
}

func (c *tomatoTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 7
}

type cheeseTopping struct {
	pizza pizza
}

func (c *cheeseTopping) getPrice() int {
	pizzaPrice := c.pizza.getPrice()
	return pizzaPrice + 10
}

type pepperTopping struct {
	pizza pizza
}

func (p *pepperTopping) getPrice() int {
	pizzaPrice := p.pizza.getPrice()
	return pizzaPrice + 20
}
