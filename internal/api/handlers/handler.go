package handlers

import (
	"fmt"
	"go-http/internal/constants"
	"go-http/internal/models"
)

func FetchInventory(limit int, page int) (items models.InventoryItems) {

	currentLimit := page * limit
	offset := (page - 1) * limit
	err := constants.DbConn.QueryRow("SELECT * FROM items where isactive = true ORDER BY binid LIMIT $1 OFFSET $2", currentLimit, offset).Scan(currentLimit)
	if err != nil {
		fmt.Printf(err.Error())
	}
	return
}

func AddToInventory(item *models.InventoryItem) (err error) {
	if item.Quantity < 0 {
		return fmt.Errorf("Quantity cant be negative")
	}
	var exists any
	err = constants.DbConn.QueryRow("select Exists (select 1 from binids where id = $1)", item.BinId).Scan(exists)
	if err != nil {
		return fmt.Errorf("Bin Doesn't exists")
	}
	_, err = constants.DbConn.Exec("INSERT INTO items(sku,quantity,binid) values($1,$2,$3)", item.Sku, item.Quantity, item.BinId)
	return
}
