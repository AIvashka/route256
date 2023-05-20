package model

type Warehouse struct {
	ID    int64
	Items []WarehouseItem
}

type WarehouseItem struct {
	ItemID uint32
	Count  uint64
}
