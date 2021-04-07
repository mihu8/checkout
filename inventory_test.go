package checkout

import "testing"

func TestAddProductToEmptyInventory(t *testing.T) {
	inv := NewDummyInventory()
	if len(inv.list()) > 0 {
		t.Fatal("initial inventory should be empty")
	}

	inv.update(googleHome, 5)
	inv.update(macbookPro, 5)
	inv.update(googleHome, 5)

	if len(inv.list()) != 2 {
		t.Fatal("inventory should have two SKU")
	}

	if inv.list()[googleHome] != 10 {
		t.Fatal("There should be 10 google home in stock, got: ", inv.list()[googleHome])
	}
}
