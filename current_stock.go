package checkout

import "math"

var googleHome = Product{
	Id:           0,
	Sku:          "120P90",
	Name:         "Google Home",
	StickerPrice: 4999,
}

var macbookPro = Product{
	Id:           1,
	Sku:          "43N23P",
	Name:         "MacBook Pro",
	StickerPrice: 539999,
}

var alexaSpeaker = Product{
	Id:           2,
	Sku:          "A304SD",
	Name:         "Alexa Speaker",
	StickerPrice: 10950,
}

var raspberryPiB = Product{
	Id:           3,
	Sku:          "234234",
	Name:         "Raspberry Pi B",
	StickerPrice: 3000,
}

func getStandardInventory() Inventory {
	inv := NewDummyInventory()
	inv.update(googleHome, 10)
	inv.update(macbookPro, 5)
	inv.update(alexaSpeaker, 10)
	inv.update(raspberryPiB, 2)
	return inv
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
