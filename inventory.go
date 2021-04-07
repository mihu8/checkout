package checkout

import "sync"

type Id int64
type Cents int64 // FIXME: some currency uses 1/1000 as "cents"

type Quantity int64 // FIXME: iPhone discrete, apples are not (1.1kg)

type Product struct {
	Id
	Sku string
	Name string
	StickerPrice Cents
}

type Inventory interface {
	list() map[Product]Quantity

	update(Product, delta Quantity)
}

type DummyInventory struct {
	// FIXME: let's consider mutex later.
	m sync.Mutex

	Stock map[Product]Quantity // TODO: move to a better inventory system asap...
}

func NewDummyInventory() DummyInventory {
	return DummyInventory{
		m:     sync.Mutex{},
		Stock: nil,
	}
}

type ProductBook []Product

