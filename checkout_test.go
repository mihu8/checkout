package checkout

import (
	"log"
	"testing"
)

func TestStandardPricing(t *testing.T) {
	emptyInv := NewDummyInventory()
	inv := getStandardInventory()

	cart := NewCart()
	cart.Add(macbookPro)
	cart.Add(raspberryPiB)

	if cart.IsStockAvailable(emptyInv) {
		log.Fatal("We don't expect stock to be available for empty inventory")
	}

	if !cart.IsStockAvailable(inv) {
		log.Fatal("Standard inventory should have enough stock for this cart")
	}

	if cart.GetTotalPrice() != macbookPro.StickerPrice+raspberryPiB.StickerPrice {
		log.Fatal("Total price should be", macbookPro.StickerPrice+raspberryPiB.StickerPrice, "but got", cart.GetTotalPrice())
	}
}
