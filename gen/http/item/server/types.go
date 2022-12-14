// Code generated by goa v3.9.1, DO NOT EDIT.
//
// item HTTP server types
//
// Command:
// $ goa gen github.com/austinhuck/characters-go/design

package server

import (
	item "github.com/austinhuck/characters-go/gen/item"
	goa "goa.design/goa/v3/pkg"
)

// CreateItemRequestBody is the type of the "item" service "createItem"
// endpoint HTTP request body.
type CreateItemRequestBody struct {
	// Item name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Item description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Item damage
	Damage *uint `form:"damage,omitempty" json:"damage,omitempty" xml:"damage,omitempty"`
	// Item healing
	Healing *uint `form:"healing,omitempty" json:"healing,omitempty" xml:"healing,omitempty"`
	// Item protection
	Protection *int `form:"protection,omitempty" json:"protection,omitempty" xml:"protection,omitempty"`
}

// UpdateItemRequestBody is the type of the "item" service "updateItem"
// endpoint HTTP request body.
type UpdateItemRequestBody struct {
	// Item name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Item description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Item damage
	Damage *uint `form:"damage,omitempty" json:"damage,omitempty" xml:"damage,omitempty"`
	// Item healing
	Healing *uint `form:"healing,omitempty" json:"healing,omitempty" xml:"healing,omitempty"`
	// Item protection
	Protection *int `form:"protection,omitempty" json:"protection,omitempty" xml:"protection,omitempty"`
}

// GetItemsResponseBody is the type of the "item" service "getItems" endpoint
// HTTP response body.
type GetItemsResponseBody []*ItemResponse

// GetItemResponseBody is the type of the "item" service "getItem" endpoint
// HTTP response body.
type GetItemResponseBody struct {
	// Item ID
	ID int `form:"id" json:"id" xml:"id"`
	// Item name
	Name string `form:"name" json:"name" xml:"name"`
	// Item description
	Description string `form:"description" json:"description" xml:"description"`
	// Item damage
	Damage uint `form:"damage" json:"damage" xml:"damage"`
	// Item healing
	Healing uint `form:"healing" json:"healing" xml:"healing"`
	// Item protection
	Protection int `form:"protection" json:"protection" xml:"protection"`
}

// CreateItemResponseBody is the type of the "item" service "createItem"
// endpoint HTTP response body.
type CreateItemResponseBody struct {
	// Item ID
	ID int `form:"id" json:"id" xml:"id"`
	// Item name
	Name string `form:"name" json:"name" xml:"name"`
	// Item description
	Description string `form:"description" json:"description" xml:"description"`
	// Item damage
	Damage uint `form:"damage" json:"damage" xml:"damage"`
	// Item healing
	Healing uint `form:"healing" json:"healing" xml:"healing"`
	// Item protection
	Protection int `form:"protection" json:"protection" xml:"protection"`
}

// UpdateItemResponseBody is the type of the "item" service "updateItem"
// endpoint HTTP response body.
type UpdateItemResponseBody struct {
	// Item ID
	ID int `form:"id" json:"id" xml:"id"`
	// Item name
	Name string `form:"name" json:"name" xml:"name"`
	// Item description
	Description string `form:"description" json:"description" xml:"description"`
	// Item damage
	Damage uint `form:"damage" json:"damage" xml:"damage"`
	// Item healing
	Healing uint `form:"healing" json:"healing" xml:"healing"`
	// Item protection
	Protection int `form:"protection" json:"protection" xml:"protection"`
}

// GetItemsIDNotFoundResponseBody is the type of the "item" service "getItems"
// endpoint HTTP response body for the "IdNotFound" error.
type GetItemsIDNotFoundResponseBody struct {
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

// GetItemsBadDataResponseBody is the type of the "item" service "getItems"
// endpoint HTTP response body for the "BadData" error.
type GetItemsBadDataResponseBody struct {
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

// GetItemsIDsExhaustedResponseBody is the type of the "item" service
// "getItems" endpoint HTTP response body for the "IDsExhausted" error.
type GetItemsIDsExhaustedResponseBody struct {
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

// GetItemIDNotFoundResponseBody is the type of the "item" service "getItem"
// endpoint HTTP response body for the "IdNotFound" error.
type GetItemIDNotFoundResponseBody struct {
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

// GetItemBadDataResponseBody is the type of the "item" service "getItem"
// endpoint HTTP response body for the "BadData" error.
type GetItemBadDataResponseBody struct {
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

// GetItemIDsExhaustedResponseBody is the type of the "item" service "getItem"
// endpoint HTTP response body for the "IDsExhausted" error.
type GetItemIDsExhaustedResponseBody struct {
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

// CreateItemBadDataResponseBody is the type of the "item" service "createItem"
// endpoint HTTP response body for the "BadData" error.
type CreateItemBadDataResponseBody struct {
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

// CreateItemIDsExhaustedResponseBody is the type of the "item" service
// "createItem" endpoint HTTP response body for the "IDsExhausted" error.
type CreateItemIDsExhaustedResponseBody struct {
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

// CreateItemIDNotFoundResponseBody is the type of the "item" service
// "createItem" endpoint HTTP response body for the "IdNotFound" error.
type CreateItemIDNotFoundResponseBody struct {
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

// UpdateItemIDNotFoundResponseBody is the type of the "item" service
// "updateItem" endpoint HTTP response body for the "IdNotFound" error.
type UpdateItemIDNotFoundResponseBody struct {
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

// UpdateItemBadDataResponseBody is the type of the "item" service "updateItem"
// endpoint HTTP response body for the "BadData" error.
type UpdateItemBadDataResponseBody struct {
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

// UpdateItemIDsExhaustedResponseBody is the type of the "item" service
// "updateItem" endpoint HTTP response body for the "IDsExhausted" error.
type UpdateItemIDsExhaustedResponseBody struct {
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

// DeleteItemIDNotFoundResponseBody is the type of the "item" service
// "deleteItem" endpoint HTTP response body for the "IdNotFound" error.
type DeleteItemIDNotFoundResponseBody struct {
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

// DeleteItemBadDataResponseBody is the type of the "item" service "deleteItem"
// endpoint HTTP response body for the "BadData" error.
type DeleteItemBadDataResponseBody struct {
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

// DeleteItemIDsExhaustedResponseBody is the type of the "item" service
// "deleteItem" endpoint HTTP response body for the "IDsExhausted" error.
type DeleteItemIDsExhaustedResponseBody struct {
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

// ItemResponse is used to define fields on response body types.
type ItemResponse struct {
	// Item ID
	ID int `form:"id" json:"id" xml:"id"`
	// Item name
	Name string `form:"name" json:"name" xml:"name"`
	// Item description
	Description string `form:"description" json:"description" xml:"description"`
	// Item damage
	Damage uint `form:"damage" json:"damage" xml:"damage"`
	// Item healing
	Healing uint `form:"healing" json:"healing" xml:"healing"`
	// Item protection
	Protection int `form:"protection" json:"protection" xml:"protection"`
}

// NewGetItemsResponseBody builds the HTTP response body from the result of the
// "getItems" endpoint of the "item" service.
func NewGetItemsResponseBody(res []*item.Item) GetItemsResponseBody {
	body := make([]*ItemResponse, len(res))
	for i, val := range res {
		body[i] = marshalItemItemToItemResponse(val)
	}
	return body
}

// NewGetItemResponseBody builds the HTTP response body from the result of the
// "getItem" endpoint of the "item" service.
func NewGetItemResponseBody(res *item.Item) *GetItemResponseBody {
	body := &GetItemResponseBody{
		ID:          res.ID,
		Name:        res.Name,
		Description: res.Description,
		Damage:      res.Damage,
		Healing:     res.Healing,
		Protection:  res.Protection,
	}
	return body
}

// NewCreateItemResponseBody builds the HTTP response body from the result of
// the "createItem" endpoint of the "item" service.
func NewCreateItemResponseBody(res *item.Item) *CreateItemResponseBody {
	body := &CreateItemResponseBody{
		ID:          res.ID,
		Name:        res.Name,
		Description: res.Description,
		Damage:      res.Damage,
		Healing:     res.Healing,
		Protection:  res.Protection,
	}
	return body
}

// NewUpdateItemResponseBody builds the HTTP response body from the result of
// the "updateItem" endpoint of the "item" service.
func NewUpdateItemResponseBody(res *item.Item) *UpdateItemResponseBody {
	body := &UpdateItemResponseBody{
		ID:          res.ID,
		Name:        res.Name,
		Description: res.Description,
		Damage:      res.Damage,
		Healing:     res.Healing,
		Protection:  res.Protection,
	}
	return body
}

// NewGetItemsIDNotFoundResponseBody builds the HTTP response body from the
// result of the "getItems" endpoint of the "item" service.
func NewGetItemsIDNotFoundResponseBody(res *goa.ServiceError) *GetItemsIDNotFoundResponseBody {
	body := &GetItemsIDNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetItemsBadDataResponseBody builds the HTTP response body from the result
// of the "getItems" endpoint of the "item" service.
func NewGetItemsBadDataResponseBody(res *goa.ServiceError) *GetItemsBadDataResponseBody {
	body := &GetItemsBadDataResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetItemsIDsExhaustedResponseBody builds the HTTP response body from the
// result of the "getItems" endpoint of the "item" service.
func NewGetItemsIDsExhaustedResponseBody(res *goa.ServiceError) *GetItemsIDsExhaustedResponseBody {
	body := &GetItemsIDsExhaustedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetItemIDNotFoundResponseBody builds the HTTP response body from the
// result of the "getItem" endpoint of the "item" service.
func NewGetItemIDNotFoundResponseBody(res *goa.ServiceError) *GetItemIDNotFoundResponseBody {
	body := &GetItemIDNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetItemBadDataResponseBody builds the HTTP response body from the result
// of the "getItem" endpoint of the "item" service.
func NewGetItemBadDataResponseBody(res *goa.ServiceError) *GetItemBadDataResponseBody {
	body := &GetItemBadDataResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetItemIDsExhaustedResponseBody builds the HTTP response body from the
// result of the "getItem" endpoint of the "item" service.
func NewGetItemIDsExhaustedResponseBody(res *goa.ServiceError) *GetItemIDsExhaustedResponseBody {
	body := &GetItemIDsExhaustedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateItemBadDataResponseBody builds the HTTP response body from the
// result of the "createItem" endpoint of the "item" service.
func NewCreateItemBadDataResponseBody(res *goa.ServiceError) *CreateItemBadDataResponseBody {
	body := &CreateItemBadDataResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateItemIDsExhaustedResponseBody builds the HTTP response body from the
// result of the "createItem" endpoint of the "item" service.
func NewCreateItemIDsExhaustedResponseBody(res *goa.ServiceError) *CreateItemIDsExhaustedResponseBody {
	body := &CreateItemIDsExhaustedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateItemIDNotFoundResponseBody builds the HTTP response body from the
// result of the "createItem" endpoint of the "item" service.
func NewCreateItemIDNotFoundResponseBody(res *goa.ServiceError) *CreateItemIDNotFoundResponseBody {
	body := &CreateItemIDNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateItemIDNotFoundResponseBody builds the HTTP response body from the
// result of the "updateItem" endpoint of the "item" service.
func NewUpdateItemIDNotFoundResponseBody(res *goa.ServiceError) *UpdateItemIDNotFoundResponseBody {
	body := &UpdateItemIDNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateItemBadDataResponseBody builds the HTTP response body from the
// result of the "updateItem" endpoint of the "item" service.
func NewUpdateItemBadDataResponseBody(res *goa.ServiceError) *UpdateItemBadDataResponseBody {
	body := &UpdateItemBadDataResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateItemIDsExhaustedResponseBody builds the HTTP response body from the
// result of the "updateItem" endpoint of the "item" service.
func NewUpdateItemIDsExhaustedResponseBody(res *goa.ServiceError) *UpdateItemIDsExhaustedResponseBody {
	body := &UpdateItemIDsExhaustedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewDeleteItemIDNotFoundResponseBody builds the HTTP response body from the
// result of the "deleteItem" endpoint of the "item" service.
func NewDeleteItemIDNotFoundResponseBody(res *goa.ServiceError) *DeleteItemIDNotFoundResponseBody {
	body := &DeleteItemIDNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewDeleteItemBadDataResponseBody builds the HTTP response body from the
// result of the "deleteItem" endpoint of the "item" service.
func NewDeleteItemBadDataResponseBody(res *goa.ServiceError) *DeleteItemBadDataResponseBody {
	body := &DeleteItemBadDataResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewDeleteItemIDsExhaustedResponseBody builds the HTTP response body from the
// result of the "deleteItem" endpoint of the "item" service.
func NewDeleteItemIDsExhaustedResponseBody(res *goa.ServiceError) *DeleteItemIDsExhaustedResponseBody {
	body := &DeleteItemIDsExhaustedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetItemItemID builds a item service getItem endpoint payload.
func NewGetItemItemID(id int) *item.ItemID {
	v := &item.ItemID{}
	v.ID = id

	return v
}

// NewCreateItemItemNew builds a item service createItem endpoint payload.
func NewCreateItemItemNew(body *CreateItemRequestBody) *item.ItemNew {
	v := &item.ItemNew{
		Name:        *body.Name,
		Description: *body.Description,
		Damage:      *body.Damage,
		Healing:     *body.Healing,
		Protection:  *body.Protection,
	}

	return v
}

// NewUpdateItemItem builds a item service updateItem endpoint payload.
func NewUpdateItemItem(body *UpdateItemRequestBody, id int) *item.Item {
	v := &item.Item{
		Name:        *body.Name,
		Description: *body.Description,
		Damage:      *body.Damage,
		Healing:     *body.Healing,
		Protection:  *body.Protection,
	}
	v.ID = id

	return v
}

// NewDeleteItemItemID builds a item service deleteItem endpoint payload.
func NewDeleteItemItemID(id int) *item.ItemID {
	v := &item.ItemID{}
	v.ID = id

	return v
}

// ValidateCreateItemRequestBody runs the validations defined on
// CreateItemRequestBody
func ValidateCreateItemRequestBody(body *CreateItemRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Description == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("description", "body"))
	}
	if body.Damage == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("damage", "body"))
	}
	if body.Healing == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("healing", "body"))
	}
	if body.Protection == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("protection", "body"))
	}
	return
}

// ValidateUpdateItemRequestBody runs the validations defined on
// UpdateItemRequestBody
func ValidateUpdateItemRequestBody(body *UpdateItemRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Description == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("description", "body"))
	}
	if body.Damage == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("damage", "body"))
	}
	if body.Healing == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("healing", "body"))
	}
	if body.Protection == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("protection", "body"))
	}
	return
}
