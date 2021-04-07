package checkout

type LineItem struct {
	product Product

	// discount is 0 or positive. If it's 5 cents, it means it gives discount of 5c, so actual price would be product.Price - discount
	discount Cents
}

type Cart struct {
	// FIXME: add sync

	lineItems []LineItem
}

func NewCart() *Cart {
	return &Cart{lineItems: nil}
}

func (c *Cart) Add(product Product) {
	lineItem := LineItem{
		product:  product,
		discount: 0,
	}

	c.lineItems = append(c.lineItems, lineItem)
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
		totalPrice = totalPrice + (item.product.StickerPrice - item.discount)
	}

	return totalPrice
}
