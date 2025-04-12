package vending_machine

type Inventory struct {
	products map[*Product]int
}

func NewInventory() *Inventory {
	return &Inventory{products: make(map[*Product]int)}
}

func (inv *Inventory) AddProduct(product *Product, quantity int) {
	inv.products[product] = quantity
}

func (inv *Inventory) RemoveProduct(product *Product) {
	if qty, exists := inv.products[product]; exists && qty > 0 {
		inv.products[product] = qty - 1
	}
}

func (inv *Inventory) IsAvailable(product *Product) bool {
	qty, exists := inv.products[product]
	return exists && qty > 0
}
