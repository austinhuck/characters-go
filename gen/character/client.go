// Code generated by goa v3.9.1, DO NOT EDIT.
//
// character client
//
// Command:
// $ goa gen github.com/austinhuck/characters-go/design

package character

import (
	"context"

	goa "goa.design/goa/v3/pkg"
)

// Client is the "character" service client.
type Client struct {
	GetCharactersEndpoint   goa.Endpoint
	GetCharacterEndpoint    goa.Endpoint
	CreateCharacterEndpoint goa.Endpoint
	UpdateCharacterEndpoint goa.Endpoint
	DeleteCharacterEndpoint goa.Endpoint
}

// NewClient initializes a "character" service client given the endpoints.
func NewClient(getCharacters, getCharacter, createCharacter, updateCharacter, deleteCharacter goa.Endpoint) *Client {
	return &Client{
		GetCharactersEndpoint:   getCharacters,
		GetCharacterEndpoint:    getCharacter,
		CreateCharacterEndpoint: createCharacter,
		UpdateCharacterEndpoint: updateCharacter,
		DeleteCharacterEndpoint: deleteCharacter,
	}
}

// GetCharacters calls the "getCharacters" endpoint of the "character" service.
// GetCharacters may return the following errors:
//   - "IdNotFound" (type *goa.ServiceError)
//   - "BadData" (type *goa.ServiceError)
//   - "IDsExhausted" (type *goa.ServiceError)
//   - error: internal error
func (c *Client) GetCharacters(ctx context.Context) (res []*Character, err error) {
	var ires interface{}
	ires, err = c.GetCharactersEndpoint(ctx, nil)
	if err != nil {
		return
	}
	return ires.([]*Character), nil
}

// GetCharacter calls the "getCharacter" endpoint of the "character" service.
// GetCharacter may return the following errors:
//   - "IdNotFound" (type *goa.ServiceError)
//   - "BadData" (type *goa.ServiceError)
//   - "IDsExhausted" (type *goa.ServiceError)
//   - error: internal error
func (c *Client) GetCharacter(ctx context.Context, p *CharacterID) (res *Character, err error) {
	var ires interface{}
	ires, err = c.GetCharacterEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Character), nil
}

// CreateCharacter calls the "createCharacter" endpoint of the "character"
// service.
// CreateCharacter may return the following errors:
//   - "BadData" (type *goa.ServiceError)
//   - "IDsExhausted" (type *goa.ServiceError)
//   - "IdNotFound" (type *goa.ServiceError)
//   - error: internal error
func (c *Client) CreateCharacter(ctx context.Context, p *CharacterNew) (res *Character, err error) {
	var ires interface{}
	ires, err = c.CreateCharacterEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Character), nil
}

// UpdateCharacter calls the "updateCharacter" endpoint of the "character"
// service.
// UpdateCharacter may return the following errors:
//   - "IdNotFound" (type *goa.ServiceError)
//   - "BadData" (type *goa.ServiceError)
//   - "IDsExhausted" (type *goa.ServiceError)
//   - error: internal error
func (c *Client) UpdateCharacter(ctx context.Context, p *Character) (res *Character, err error) {
	var ires interface{}
	ires, err = c.UpdateCharacterEndpoint(ctx, p)
	if err != nil {
		return
	}
	return ires.(*Character), nil
}

// DeleteCharacter calls the "deleteCharacter" endpoint of the "character"
// service.
// DeleteCharacter may return the following errors:
//   - "IdNotFound" (type *goa.ServiceError)
//   - "BadData" (type *goa.ServiceError)
//   - "IDsExhausted" (type *goa.ServiceError)
//   - error: internal error
func (c *Client) DeleteCharacter(ctx context.Context, p *CharacterID) (err error) {
	_, err = c.DeleteCharacterEndpoint(ctx, p)
	return
}
