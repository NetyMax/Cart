package main

import "fmt"

type Product struct {
	ID    int
	Name  string
	Price float64
}

type Cart struct {
	Items map[int]Product
}

func NewCart() *Cart {
	return &Cart{
		Items: make(map[int]Product),
	}
}

func (c *Cart) Add(p Product) {
	c.Items[p.ID] = p
	fmt.Println("Items added in cart", p.Name)
}

func (c *Cart) Remove(p Product) {
	if _, ok := c.Items[p.ID]; ok {
		fmt.Printf("Items %s deleted", c.Items[p.ID].Name)
		delete(c.Items, p.ID)
	} else {
		fmt.Println("товар не найден")
	}
}
