package checkout

import (
	"math"
)

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

type DiscountOneForOne struct {
	condition   Product
	target      Product
	maxDiscount Cents
}

type DiscountOnEveryX struct {
	product     Product
	everyN      int
	maxDiscount Cents
}

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

var freeRaspberryPiPromotion Promotion = DiscountOneForOne{
	condition:   macbookPro,
	target:      raspberryPiB,
	maxDiscount: raspberryPiB.StickerPrice,
}

var googleHomesThreeForTwoPromotion Promotion = DiscountOnEveryX{
	product:     googleHome,
	everyN:      2,
	maxDiscount: googleHome.StickerPrice,
}

var alexaSpeakerPromotion Promotion = DiscountTriggerOnQuantity{
	product:         alexaSpeaker,
	triggerQuantity: 3,
	maxDiscount:     Cents(math.Floor(0.1 * float64(alexaSpeaker.StickerPrice))),
}

// GetPromotionBook returns the promotion book based on a set of rules, like the current user, and so on
func GetPromotionBook() PromotionBook {
	// dummy implementation

	return PromotionBook{
		promotions: []Promotion{googleHomesThreeForTwoPromotion, alexaSpeakerPromotion, freeRaspberryPiPromotion},
	}
}
