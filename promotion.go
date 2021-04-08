package checkout

type Promotion interface {
	Apply(cart *Cart) error
}

type PromotionBook struct {
	promotions []Promotion
}

func (b PromotionBook) ApplyAll(cart *Cart) error {
	for _, promotion := range b.promotions {
		if err := promotion.Apply(cart); err != nil {
			return err
		}
	}

	return nil
}

// DiscountOneForOne means once an item is in the shopping cart, give discount on target for every condition product purchased.
type DiscountOneForOne struct {
	condition Product
	target    Product

	// maxDiscount stands for the maximum possible discount. If the target product is free, make maxDiscount = target.StickerPrice
	maxDiscount Cents
}

type DiscountOnEveryX struct {
	product Product

	// give discount on Nth item. e.g. 5 means give discount on the 5th and 10th item.
	everyN      int
	maxDiscount Cents
}

// DiscountTriggerOnQuantity will be activated once the product has reached triggerQuantity. And it applies discount on every products in cart
type DiscountTriggerOnQuantity struct {
	product         Product
	triggerQuantity Quantity // FIXME: this is coupled with line items' quantity
	maxDiscount     Cents
}

func (d DiscountTriggerOnQuantity) Apply(cart *Cart) error {
	n := Quantity(0)
	activated := false

	for _, lineItem := range cart.lineItems {
		if lineItem.product == d.product {
			n++
			if n >= d.triggerQuantity {
				activated = true
				break
			}
		}
	}

	if activated {
		var lineItems []LineItem
		for _, lineItem := range cart.lineItems {
			if lineItem.product == d.product {
				discount := d.maxDiscount
				if lineItem.GetActualPrice() <= discount {
					discount = lineItem.GetActualPrice()
				}

				lineItem.discount += discount
			}

			lineItems = append(lineItems, lineItem)
		}

		cart.lineItems = lineItems
	}

	return nil
}

func (d DiscountOnEveryX) Apply(cart *Cart) error {
	n := d.everyN

	var lineItems []LineItem
	for _, lineItem := range cart.lineItems {
		if lineItem.product == d.product {
			if n == 0 {
				discount := d.maxDiscount
				if lineItem.GetActualPrice() <= discount {
					discount = lineItem.GetActualPrice()
				}
				lineItem.discount += discount

				n = d.everyN
			} else {
				n = n - 1
			}
		}

		lineItems = append(lineItems, lineItem)
	}
	cart.lineItems = lineItems

	return nil
}

func (d DiscountOneForOne) Apply(cart *Cart) error {
	discountOpportunities := 0

	for _, lineItem := range cart.lineItems {
		if lineItem.product == d.condition {
			discountOpportunities++
		}
	}

	var lineItems []LineItem
	for _, lineItem := range cart.lineItems {
		if (discountOpportunities > 0) && (lineItem.product == d.target) {
			discount := d.maxDiscount
			if lineItem.GetActualPrice() > discount {
				discount = lineItem.GetActualPrice()
			}
			lineItem.discount += discount
			discountOpportunities--
		}
		lineItems = append(lineItems, lineItem)
	}
	cart.lineItems = lineItems

	return nil
}
