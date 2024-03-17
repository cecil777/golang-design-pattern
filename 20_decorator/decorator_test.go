package decorator

import "fmt"

func ExampleDecorator() {
	var c Component = &ConcreteComponent{}
	c = WrapMulDecorator(c, 10)
	c = WrapAddDecorator(c, 8)
	res := c.Calc()

	fmt.Printf("res %d\n", res)
}
