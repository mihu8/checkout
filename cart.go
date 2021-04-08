package checkout

import "fmt"

// LineItem is for product in Cart.
type LineItem struct {
	product Product

	// discount is 0 or positive. If it's 5 cents, it means it gives discount of 5c, so actual price would be product.Price - discount
	discount Cents

	// FIXME: we should aggregate same product together
}

func (LineItem LineItem) GetActualPrice() Cents {
	return LineItem.product.StickerPrice - LineItem.discount
}

type Cart struct {
	// FIXME: add sync

	lineItems []LineItem
}

func NewCart() *Cart {
	return &Cart{lineItems: nil}
}

func (c *Cart) ResetPromotion() {
	var items []LineItem
	for _, item := range c.lineItems {
		item.discount = 0
		items = append(items, item)
	}

	c.lineItems = items
}

// Add puts a product into the shopping cart. Note: it will reset all promotion in the cart.
func (c *Cart) Add(product Product) {
	lineItem := LineItem{
		product:  product,
		discount: 0,
	}

	c.lineItems = append(c.lineItems, lineItem)

	c.ResetPromotion()
}

func (c *Cart) IsStockAvailable(inv Inventory) bool {
	total := make(map[Product]Quantity)

	for _, item := range c.lineItems {
		product := item.product
		total[product] += 1
	}

	for product, totalQuantity := range total {
		if totalQuantity > inv.list()[product] {
			return false
		}
	}

	return true
}

func (c *Cart) GetTotalPrice() Cents {
	totalPrice := Cents(0)

	for _, item := range c.lineItems {
		totalPrice = totalPrice + item.GetActualPrice()
	}

	return totalPrice
}

func (c *Cart) Dump() string {
	line := "Cart:\n"

	for idx, item := range c.lineItems {
		line = line + fmt.Sprintf("%02d. %s [%s] %dC - %dC\n", idx, item.product.Name, item.product.Sku, item.product.StickerPrice, item.discount)
	}

	return line
}
