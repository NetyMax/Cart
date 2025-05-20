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

func (c *Cart) Remove(id int) {
	if _, ok := c.Items[p.ID]; ok {
		fmt.Printf("Items %s deleted", c.Items[p.ID].Name)
		delete(c.Items, p.ID)
	} else {
		fmt.Println("товар не найден")
	}
}

func (c *Cart) Show() {
	fmt.Println("Товар в корзине")
	if len(c.Items) == 0 {
		fmt.Println("Корзина пуста")
		return
	}
	for _, p := range c.Items {
		fmt.Println("ID: %d, Название: %s, Цена: %2f\n", p.ID, p.Name, p.Price)
	}
}

func main() {
	cart := NewCart()
	product1 := Product{ID: 1, Name: "Ноутбук", Price: 59999.99}
	product2 := Product{ID: 2, Name: "Смартфон", Price: 29999.50}
	product3 := Product{ID: 3, Name: "Наушники", Price: 3999.90}

	cart.Add(product1)
	cart.Add(product2)
	cart.Add(product3)

	cart.Show()
	cart.Remove(3)
	cart.Show()
}
