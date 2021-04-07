package checkout

import (
	"sync"
)

type Id int64
type Cents int64 // FIXME: some currency uses 1/1000 as "cents"

type Quantity int64 // FIXME: iPhone discrete, apples are not (1.1kg)

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
	// FIXME: we should use atomic operations or lock

	inv.m.Lock()
	defer inv.m.Unlock()

	inv.stock[Product] += delta
}

func NewDummyInventory() DummyInventory {
	return DummyInventory{
		m:     sync.Mutex{},
		stock: map[Product]Quantity{},
	}
}

type ProductBook []Product
