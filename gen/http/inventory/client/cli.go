// Code generated by goa v3.9.1, DO NOT EDIT.
//
// inventory HTTP client CLI support package
//
// Command:
// $ goa gen github.com/austinhuck/characters-go/design

package client

import (
	"encoding/json"
	"fmt"
	"strconv"

	inventory "github.com/austinhuck/characters-go/gen/inventory"
)

// BuildGetInventoryPayload builds the payload for the inventory getInventory
// endpoint from CLI flags.
func BuildGetInventoryPayload(inventoryGetInventoryID string) (*inventory.CharacterID, error) {
	var err error
	var id int
	{
		var v int64
		v, err = strconv.ParseInt(inventoryGetInventoryID, 10, strconv.IntSize)
		id = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be INT")
		}
	}
	v := &inventory.CharacterID{}
	v.ID = id

	return v, nil
}

// BuildUpdateInventoryPayload builds the payload for the inventory
// updateInventory endpoint from CLI flags.
func BuildUpdateInventoryPayload(inventoryUpdateInventoryBody string, inventoryUpdateInventoryID string) (*inventory.InventoryUpdate, error) {
	var err error
	var body UpdateInventoryRequestBody
	{
		err = json.Unmarshal([]byte(inventoryUpdateInventoryBody), &body)
		if err != nil {
			return nil, fmt.Errorf("invalid JSON for body, \nerror: %s, \nexample of valid JSON:\n%s", err, "'{\n      \"item\": 1766420538716923054,\n      \"quantity\": 10997400203762602967\n   }'")
		}
	}
	var id int
	{
		var v int64
		v, err = strconv.ParseInt(inventoryUpdateInventoryID, 10, strconv.IntSize)
		id = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be INT")
		}
	}
	v := &inventory.InventoryUpdate{
		Item:     body.Item,
		Quantity: body.Quantity,
	}
	v.ID = id

	return v, nil
}

// BuildClearInventoryPayload builds the payload for the inventory
// clearInventory endpoint from CLI flags.
func BuildClearInventoryPayload(inventoryClearInventoryID string) (*inventory.CharacterID, error) {
	var err error
	var id int
	{
		var v int64
		v, err = strconv.ParseInt(inventoryClearInventoryID, 10, strconv.IntSize)
		id = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for id, must be INT")
		}
	}
	v := &inventory.CharacterID{}
	v.ID = id

	return v, nil
}
