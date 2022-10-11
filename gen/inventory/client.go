// Code generated by goa v3.9.1, DO NOT EDIT.
//
// inventory client
//
// Command:
// $ goa gen github.com/austinhuck/characters-go/design

package inventory

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "inventory" service client.
type Client struct {
	GetInventoryEndpoint    goa.Endpoint
	UpdateInventoryEndpoint goa.Endpoint
	ClearInventoryEndpoint  goa.Endpoint
}

// NewClient initializes a "inventory" service client given the endpoints.
func NewClient(getInventory, updateInventory, clearInventory goa.Endpoint) *Client {
	return &Client{
		GetInventoryEndpoint:    getInventory,
		UpdateInventoryEndpoint: updateInventory,
		ClearInventoryEndpoint:  clearInventory,
	}
}

// GetInventory calls the "getInventory" endpoint of the "inventory" service.
// GetInventory may return the following errors:
//   - "IdNotFound" (type *goa.ServiceError)
//   - "BadData" (type *goa.ServiceError)
//   - "IDsExhausted" (type *goa.ServiceError)
//   - error: internal error
func (c *Client) GetInventory(ctx context.Context, p *CharacterID) (res []*InventoryEntry, err error) {
	var ires interface{}
	ires, err = c.GetInventoryEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.([]*InventoryEntry), nil
}

// UpdateInventory calls the "updateInventory" endpoint of the "inventory"
// service.
// UpdateInventory may return the following errors:
//   - "IdNotFound" (type *goa.ServiceError)
//   - "BadData" (type *goa.ServiceError)
//   - "IDsExhausted" (type *goa.ServiceError)
//   - error: internal error
func (c *Client) UpdateInventory(ctx context.Context, p *InventoryUpdate) (err error) {
	_, err = c.UpdateInventoryEndpoint(ctx, p)
	return
}

// ClearInventory calls the "clearInventory" endpoint of the "inventory"
// service.
// ClearInventory may return the following errors:
//   - "IdNotFound" (type *goa.ServiceError)
//   - "BadData" (type *goa.ServiceError)
//   - "IDsExhausted" (type *goa.ServiceError)
//   - error: internal error
func (c *Client) ClearInventory(ctx context.Context, p *CharacterID) (err error) {
	_, err = c.ClearInventoryEndpoint(ctx, p)
	return
}
