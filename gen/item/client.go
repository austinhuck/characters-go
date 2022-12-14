// Code generated by goa v3.9.1, DO NOT EDIT.
//
// item client
//
// Command:
// $ goa gen github.com/austinhuck/characters-go/design

package item

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "item" service client.
type Client struct {
	GetItemsEndpoint   goa.Endpoint
	GetItemEndpoint    goa.Endpoint
	CreateItemEndpoint goa.Endpoint
	UpdateItemEndpoint goa.Endpoint
	DeleteItemEndpoint goa.Endpoint
}

// NewClient initializes a "item" service client given the endpoints.
func NewClient(getItems, getItem, createItem, updateItem, deleteItem goa.Endpoint) *Client {
	return &Client{
		GetItemsEndpoint:   getItems,
		GetItemEndpoint:    getItem,
		CreateItemEndpoint: createItem,
		UpdateItemEndpoint: updateItem,
		DeleteItemEndpoint: deleteItem,
	}
}

// GetItems calls the "getItems" endpoint of the "item" service.
// GetItems may return the following errors:
//   - "IdNotFound" (type *goa.ServiceError)
//   - "BadData" (type *goa.ServiceError)
//   - "IDsExhausted" (type *goa.ServiceError)
//   - error: internal error
func (c *Client) GetItems(ctx context.Context) (res []*Item, err error) {
	var ires interface{}
	ires, err = c.GetItemsEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.([]*Item), nil
}

// GetItem calls the "getItem" endpoint of the "item" service.
// GetItem may return the following errors:
//   - "IdNotFound" (type *goa.ServiceError)
//   - "BadData" (type *goa.ServiceError)
//   - "IDsExhausted" (type *goa.ServiceError)
//   - error: internal error
func (c *Client) GetItem(ctx context.Context, p *ItemID) (res *Item, err error) {
	var ires interface{}
	ires, err = c.GetItemEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Item), nil
}

// CreateItem calls the "createItem" endpoint of the "item" service.
// CreateItem may return the following errors:
//   - "BadData" (type *goa.ServiceError)
//   - "IDsExhausted" (type *goa.ServiceError)
//   - "IdNotFound" (type *goa.ServiceError)
//   - error: internal error
func (c *Client) CreateItem(ctx context.Context, p *ItemNew) (res *Item, err error) {
	var ires interface{}
	ires, err = c.CreateItemEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Item), nil
}

// UpdateItem calls the "updateItem" endpoint of the "item" service.
// UpdateItem may return the following errors:
//   - "IdNotFound" (type *goa.ServiceError)
//   - "BadData" (type *goa.ServiceError)
//   - "IDsExhausted" (type *goa.ServiceError)
//   - error: internal error
func (c *Client) UpdateItem(ctx context.Context, p *Item) (res *Item, err error) {
	var ires interface{}
	ires, err = c.UpdateItemEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Item), nil
}

// DeleteItem calls the "deleteItem" endpoint of the "item" service.
// DeleteItem may return the following errors:
//   - "IdNotFound" (type *goa.ServiceError)
//   - "BadData" (type *goa.ServiceError)
//   - "IDsExhausted" (type *goa.ServiceError)
//   - error: internal error
func (c *Client) DeleteItem(ctx context.Context, p *ItemID) (err error) {
	_, err = c.DeleteItemEndpoint(ctx, p)
	return
}
