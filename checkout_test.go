package checkout

import (
	"math"
	"testing"
)

func TestStandardPricing(t *testing.T) {
	emptyInv := NewDummyInventory()
	inv := getStandardInventory()

	cart := NewCart()
	cart.Add(macbookPro)
	cart.Add(raspberryPiB)

	if cart.IsStockAvailable(emptyInv) {
		t.Fatal("We don't expect stock to be available for empty inventory")
	}

	if !cart.IsStockAvailable(inv) {
		t.Fatal("Standard inventory should have enough stock for this cart")
	}

	if cart.GetTotalPrice() != macbookPro.StickerPrice+raspberryPiB.StickerPrice {
		t.Fatal("Total price should be", macbookPro.StickerPrice+raspberryPiB.StickerPrice, "but got", cart.GetTotalPrice())
	}
}

func TestFreeRaspberryPiPromotion(t *testing.T) {

	cart := NewCart()
	cart.Add(macbookPro)
	cart.Add(raspberryPiB)

	if cart.GetTotalPrice() != macbookPro.StickerPrice+raspberryPiB.StickerPrice {
		t.Fatal("Total price should be ", macbookPro.StickerPrice+raspberryPiB.StickerPrice, " but got ", cart.GetTotalPrice())
	}

	if err := freeRaspberryPiPromotion.Apply(cart); err != nil {
		t.Fatal("Failed to apply raspberry pi promotion due to ", err)
	}

	if cart.GetTotalPrice() != macbookPro.StickerPrice {
		t.Fatal("Total price should be ", macbookPro.StickerPrice, " after free raspberry pi discount, but got ", cart.GetTotalPrice())
	}
}

func TestGoogleHomePromosionThreeForTwo(t *testing.T) {
	cart := NewCart()
	cart.Add(googleHome)
	cart.Add(googleHome)

	if err := googleHomesThreeForTwoPromotion.Apply(cart); err != nil {
		t.Fatal("Failed to apply google home promotion due to ", err)
	}

	if cart.GetTotalPrice() != 2*googleHome.StickerPrice {
		t.Fatal("Total price should be ", 2*googleHome.StickerPrice, " after google home 3 for 2 discount, but got ", cart.GetTotalPrice())
	}

	cart.Add(googleHome)

	if cart.GetTotalPrice() != 3*googleHome.StickerPrice {
		t.Fatal("Total price should be ", 3*googleHome.StickerPrice, " but got ", cart.GetTotalPrice())
	}

	if err := freeRaspberryPiPromotion.Apply(cart); err != nil {
		t.Fatal("Failed to apply raspberry pi promotion due to ", err)
	}

	if err := googleHomesThreeForTwoPromotion.Apply(cart); err != nil {
		t.Fatal("Failed to apply google home promotion due to ", err)
	}

	if cart.GetTotalPrice() != 2*googleHome.StickerPrice {
		t.Fatal("Total price should be ", 2*googleHome.StickerPrice, " after google home 3 for 2 discount, but got ", cart.GetTotalPrice())
	}

	// try 6 units

	cart.ResetPromotion()
	cart.Add(googleHome)
	cart.Add(googleHome)
	cart.Add(googleHome)

	if cart.GetTotalPrice() != 6*googleHome.StickerPrice {
		t.Fatal("Total price should be ", 6*googleHome.StickerPrice, " but got ", cart.GetTotalPrice())
	}

	if err := freeRaspberryPiPromotion.Apply(cart); err != nil {
		t.Fatal("Failed to apply raspberry pi promotion due to ", err)
	}

	if err := googleHomesThreeForTwoPromotion.Apply(cart); err != nil {
		t.Fatal("Failed to apply google home promotion due to ", err)
	}

	if cart.GetTotalPrice() != 4*googleHome.StickerPrice {
		t.Fatal("Total price should be ", 4*googleHome.StickerPrice, " after google home 3 for 2 discount, but got ", cart.GetTotalPrice())
	}
}

func TestAlexaPromotion(t *testing.T) {
	cart := NewCart()
	cart.Add(googleHome)
	cart.Add(alexaSpeaker)
	cart.Add(alexaSpeaker)

	alexaSpeakerPromotion.Apply(cart)
	if cart.GetTotalPrice() != googleHome.StickerPrice+alexaSpeaker.StickerPrice*2 {
		t.Fatal("Alexa speaker promotion should not be activated for two speakers")
	}

	cart.ResetPromotion()
	cart.Add(alexaSpeaker)
	cart.Add(alexaSpeaker)

	alexaSpeakerPromotion.Apply(cart)

	// t.Log(cart.Dump())

	if cart.GetTotalPrice() != googleHome.StickerPrice+(4*alexaSpeaker.StickerPrice-4*Cents(math.Floor(0.1*float64(alexaSpeaker.StickerPrice)))) {
		t.Fatal("Alexa speaker promotion should be activated for 4 speakers, expecting ", googleHome.StickerPrice+(4*alexaSpeaker.StickerPrice-4*Cents(math.Floor(0.1*float64(alexaSpeaker.StickerPrice)))), " got: ", cart.GetTotalPrice())
	}
}

func TestScenario1(t *testing.T) {
	cart := NewCart()
	cart.Add(macbookPro)
	cart.Add(raspberryPiB)
	promotionBook := GetPromotionBook()
	promotionBook.ApplyAll(cart)

	if cart.GetTotalPrice() != 539999 {
		t.Fatal("Expecting scenario to be 539999, got: ", cart.GetTotalPrice())
	}
}

func TestScenario2(t *testing.T) {
	cart := NewCart()
	cart.Add(googleHome)
	cart.Add(googleHome)
	cart.Add(googleHome)
	promotionBook := GetPromotionBook()
	promotionBook.ApplyAll(cart)

	if cart.GetTotalPrice() != 9998 {
		t.Fatal("Expecting scenario to be 9998, got: ", cart.GetTotalPrice())
	}
}

func TestScenario3(t *testing.T) {
	cart := NewCart()
	cart.Add(alexaSpeaker)
	cart.Add(alexaSpeaker)
	cart.Add(alexaSpeaker)

	promotionBook := GetPromotionBook()
	promotionBook.ApplyAll(cart)

	t.Log(cart.Dump())

	if cart.GetTotalPrice() != 29565 {
		t.Fatal("Expecting scenario to be 29565, got: ", cart.GetTotalPrice())
	}
}
