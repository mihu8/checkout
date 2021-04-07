package checkout

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
	StickerPrice: 4999,
}

var raspberryPiB = Product{
	Id:           3,
	Sku:          "3000",
	Name:         "Raspberry Pi B",
	StickerPrice: 3000,
}

func getStandardInventory() Inventory {
	inv := NewDummyInventory()
	inv.update(googleHome, 10)
	inv.update(macbookPro, 5)
	inv.update(alexaSpeaker, 10)
	inv.update(raspberryPiB, 2)
	return &inv
}
