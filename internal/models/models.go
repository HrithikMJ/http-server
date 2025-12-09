package models

import "github.com/google/uuid"

type InventoryItem struct {
	Sku      string
	Quantity int
	BinId    uuid.UUID
}

type InventoryItems []InventoryItem
