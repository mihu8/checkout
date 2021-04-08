package checkout

import (
	"sync"
)

type Id int64
type Cents int64 // FIXME: some currency systems use 1/1000 as "cents". And consider there is 10% discount on $9.99, we should really use big/decimal instead.

type Quantity int64 // FIXME: in terms of Unit, iPhone is discrete, apples are not (1.1kg)

type Product struct {
	Id
	Sku          string
	Name         string
	StickerPrice Cents
}

type Inventory interface {
	list() map[Product]Quantity

	update(Product Product, delta Quantity)
}

type DummyInventory struct {
	// FIXME: let's consider mutex later.
	m sync.Mutex

	stock map[Product]Quantity // TODO: move to a better inventory system asap...
}

func (inv *DummyInventory) list() map[Product]Quantity {
	// FIXME: not ideal
	return inv.stock
}

func (inv *DummyInventory) update(Product Product, delta Quantity) {
	// FIXME: we should use atomic or lock

	inv.m.Lock()
	defer inv.m.Unlock()

	inv.stock[Product] += delta
}

func NewDummyInventory() *DummyInventory {
	return &DummyInventory{
		m:     sync.Mutex{},
		stock: map[Product]Quantity{},
	}
}

// ProductBook holds all products
type ProductBook []Product
