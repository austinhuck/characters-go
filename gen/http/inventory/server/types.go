// Code generated by goa v3.9.1, DO NOT EDIT.
//
// inventory HTTP server types
//
// Command:
// $ goa gen github.com/austinhuck/characters-go/design

package server

import (
	inventory "github.com/austinhuck/characters-go/gen/inventory"
	goa "goa.design/goa/v3/pkg"
)

// UpdateInventoryRequestBody is the type of the "inventory" service
// "updateInventory" endpoint HTTP request body.
type UpdateInventoryRequestBody struct {
	// Item ID
	Item *int `form:"item,omitempty" json:"item,omitempty" xml:"item,omitempty"`
	// Quantity of the item
	Quantity *uint `form:"quantity,omitempty" json:"quantity,omitempty" xml:"quantity,omitempty"`
}

// GetInventoryResponseBody is the type of the "inventory" service
// "getInventory" endpoint HTTP response body.
type GetInventoryResponseBody []*InventoryEntryResponse

// GetInventoryIDNotFoundResponseBody is the type of the "inventory" service
// "getInventory" endpoint HTTP response body for the "IdNotFound" error.
type GetInventoryIDNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// GetInventoryBadDataResponseBody is the type of the "inventory" service
// "getInventory" endpoint HTTP response body for the "BadData" error.
type GetInventoryBadDataResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// GetInventoryIDsExhaustedResponseBody is the type of the "inventory" service
// "getInventory" endpoint HTTP response body for the "IDsExhausted" error.
type GetInventoryIDsExhaustedResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UpdateInventoryIDNotFoundResponseBody is the type of the "inventory" service
// "updateInventory" endpoint HTTP response body for the "IdNotFound" error.
type UpdateInventoryIDNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UpdateInventoryBadDataResponseBody is the type of the "inventory" service
// "updateInventory" endpoint HTTP response body for the "BadData" error.
type UpdateInventoryBadDataResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// UpdateInventoryIDsExhaustedResponseBody is the type of the "inventory"
// service "updateInventory" endpoint HTTP response body for the "IDsExhausted"
// error.
type UpdateInventoryIDsExhaustedResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ClearInventoryIDNotFoundResponseBody is the type of the "inventory" service
// "clearInventory" endpoint HTTP response body for the "IdNotFound" error.
type ClearInventoryIDNotFoundResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ClearInventoryBadDataResponseBody is the type of the "inventory" service
// "clearInventory" endpoint HTTP response body for the "BadData" error.
type ClearInventoryBadDataResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// ClearInventoryIDsExhaustedResponseBody is the type of the "inventory"
// service "clearInventory" endpoint HTTP response body for the "IDsExhausted"
// error.
type ClearInventoryIDsExhaustedResponseBody struct {
	// Name is the name of this class of errors.
	Name string `form:"name" json:"name" xml:"name"`
	// ID is a unique identifier for this particular occurrence of the problem.
	ID string `form:"id" json:"id" xml:"id"`
	// Message is a human-readable explanation specific to this occurrence of the
	// problem.
	Message string `form:"message" json:"message" xml:"message"`
	// Is the error temporary?
	Temporary bool `form:"temporary" json:"temporary" xml:"temporary"`
	// Is the error a timeout?
	Timeout bool `form:"timeout" json:"timeout" xml:"timeout"`
	// Is the error a server-side fault?
	Fault bool `form:"fault" json:"fault" xml:"fault"`
}

// InventoryEntryResponse is used to define fields on response body types.
type InventoryEntryResponse struct {
	// Item ID
	Item int `form:"item" json:"item" xml:"item"`
	// Quantity of the item
	Quantity uint `form:"quantity" json:"quantity" xml:"quantity"`
}

// NewGetInventoryResponseBody builds the HTTP response body from the result of
// the "getInventory" endpoint of the "inventory" service.
func NewGetInventoryResponseBody(res []*inventory.InventoryEntry) GetInventoryResponseBody {
	body := make([]*InventoryEntryResponse, len(res))
	for i, val := range res {
		body[i] = marshalInventoryInventoryEntryToInventoryEntryResponse(val)
	}
	return body
}

// NewGetInventoryIDNotFoundResponseBody builds the HTTP response body from the
// result of the "getInventory" endpoint of the "inventory" service.
func NewGetInventoryIDNotFoundResponseBody(res *goa.ServiceError) *GetInventoryIDNotFoundResponseBody {
	body := &GetInventoryIDNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetInventoryBadDataResponseBody builds the HTTP response body from the
// result of the "getInventory" endpoint of the "inventory" service.
func NewGetInventoryBadDataResponseBody(res *goa.ServiceError) *GetInventoryBadDataResponseBody {
	body := &GetInventoryBadDataResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetInventoryIDsExhaustedResponseBody builds the HTTP response body from
// the result of the "getInventory" endpoint of the "inventory" service.
func NewGetInventoryIDsExhaustedResponseBody(res *goa.ServiceError) *GetInventoryIDsExhaustedResponseBody {
	body := &GetInventoryIDsExhaustedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateInventoryIDNotFoundResponseBody builds the HTTP response body from
// the result of the "updateInventory" endpoint of the "inventory" service.
func NewUpdateInventoryIDNotFoundResponseBody(res *goa.ServiceError) *UpdateInventoryIDNotFoundResponseBody {
	body := &UpdateInventoryIDNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateInventoryBadDataResponseBody builds the HTTP response body from the
// result of the "updateInventory" endpoint of the "inventory" service.
func NewUpdateInventoryBadDataResponseBody(res *goa.ServiceError) *UpdateInventoryBadDataResponseBody {
	body := &UpdateInventoryBadDataResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateInventoryIDsExhaustedResponseBody builds the HTTP response body
// from the result of the "updateInventory" endpoint of the "inventory" service.
func NewUpdateInventoryIDsExhaustedResponseBody(res *goa.ServiceError) *UpdateInventoryIDsExhaustedResponseBody {
	body := &UpdateInventoryIDsExhaustedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewClearInventoryIDNotFoundResponseBody builds the HTTP response body from
// the result of the "clearInventory" endpoint of the "inventory" service.
func NewClearInventoryIDNotFoundResponseBody(res *goa.ServiceError) *ClearInventoryIDNotFoundResponseBody {
	body := &ClearInventoryIDNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewClearInventoryBadDataResponseBody builds the HTTP response body from the
// result of the "clearInventory" endpoint of the "inventory" service.
func NewClearInventoryBadDataResponseBody(res *goa.ServiceError) *ClearInventoryBadDataResponseBody {
	body := &ClearInventoryBadDataResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewClearInventoryIDsExhaustedResponseBody builds the HTTP response body from
// the result of the "clearInventory" endpoint of the "inventory" service.
func NewClearInventoryIDsExhaustedResponseBody(res *goa.ServiceError) *ClearInventoryIDsExhaustedResponseBody {
	body := &ClearInventoryIDsExhaustedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetInventoryCharacterID builds a inventory service getInventory endpoint
// payload.
func NewGetInventoryCharacterID(id int) *inventory.CharacterID {
	v := &inventory.CharacterID{}
	v.ID = id

	return v
}

// NewUpdateInventoryInventoryUpdate builds a inventory service updateInventory
// endpoint payload.
func NewUpdateInventoryInventoryUpdate(body *UpdateInventoryRequestBody, id int) *inventory.InventoryUpdate {
	v := &inventory.InventoryUpdate{
		Item:     *body.Item,
		Quantity: *body.Quantity,
	}
	v.ID = id

	return v
}

// NewClearInventoryCharacterID builds a inventory service clearInventory
// endpoint payload.
func NewClearInventoryCharacterID(id int) *inventory.CharacterID {
	v := &inventory.CharacterID{}
	v.ID = id

	return v
}

// ValidateUpdateInventoryRequestBody runs the validations defined on
// UpdateInventoryRequestBody
func ValidateUpdateInventoryRequestBody(body *UpdateInventoryRequestBody) (err error) {
	if body.Item == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("item", "body"))
	}
	if body.Quantity == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("quantity", "body"))
	}
	return
}