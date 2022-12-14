// Code generated by goa v3.9.1, DO NOT EDIT.
//
// character client HTTP transport
//
// Command:
// $ goa gen github.com/austinhuck/characters-go/design

package client

import (
	"context"
	"net/http"

	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// Client lists the character service endpoint HTTP clients.
type Client struct {
	// GetCharacters Doer is the HTTP client used to make requests to the
	// getCharacters endpoint.
	GetCharactersDoer goahttp.Doer

	// GetCharacter Doer is the HTTP client used to make requests to the
	// getCharacter endpoint.
	GetCharacterDoer goahttp.Doer

	// CreateCharacter Doer is the HTTP client used to make requests to the
	// createCharacter endpoint.
	CreateCharacterDoer goahttp.Doer

	// UpdateCharacter Doer is the HTTP client used to make requests to the
	// updateCharacter endpoint.
	UpdateCharacterDoer goahttp.Doer

	// DeleteCharacter Doer is the HTTP client used to make requests to the
	// deleteCharacter endpoint.
	DeleteCharacterDoer goahttp.Doer

	// RestoreResponseBody controls whether the response bodies are reset after
	// decoding so they can be read again.
	RestoreResponseBody bool

	scheme  string
	host    string
	encoder func(*http.Request) goahttp.Encoder
	decoder func(*http.Response) goahttp.Decoder
}

// NewClient instantiates HTTP clients for all the character service servers.
func NewClient(
	scheme string,
	host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restoreBody bool,
) *Client {
	return &Client{
		GetCharactersDoer:   doer,
		GetCharacterDoer:    doer,
		CreateCharacterDoer: doer,
		UpdateCharacterDoer: doer,
		DeleteCharacterDoer: doer,
		RestoreResponseBody: restoreBody,
		scheme:              scheme,
		host:                host,
		decoder:             dec,
		encoder:             enc,
	}
}

// GetCharacters returns an endpoint that makes HTTP requests to the character
// service getCharacters server.
func (c *Client) GetCharacters() goa.Endpoint {
	var (
		decodeResponse = DecodeGetCharactersResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetCharactersRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetCharactersDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("character", "getCharacters", err)
		}
		return decodeResponse(resp)
	}
}

// GetCharacter returns an endpoint that makes HTTP requests to the character
// service getCharacter server.
func (c *Client) GetCharacter() goa.Endpoint {
	var (
		decodeResponse = DecodeGetCharacterResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildGetCharacterRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.GetCharacterDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("character", "getCharacter", err)
		}
		return decodeResponse(resp)
	}
}

// CreateCharacter returns an endpoint that makes HTTP requests to the
// character service createCharacter server.
func (c *Client) CreateCharacter() goa.Endpoint {
	var (
		encodeRequest  = EncodeCreateCharacterRequest(c.encoder)
		decodeResponse = DecodeCreateCharacterResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildCreateCharacterRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.CreateCharacterDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("character", "createCharacter", err)
		}
		return decodeResponse(resp)
	}
}

// UpdateCharacter returns an endpoint that makes HTTP requests to the
// character service updateCharacter server.
func (c *Client) UpdateCharacter() goa.Endpoint {
	var (
		encodeRequest  = EncodeUpdateCharacterRequest(c.encoder)
		decodeResponse = DecodeUpdateCharacterResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildUpdateCharacterRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		err = encodeRequest(req, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.UpdateCharacterDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("character", "updateCharacter", err)
		}
		return decodeResponse(resp)
	}
}

// DeleteCharacter returns an endpoint that makes HTTP requests to the
// character service deleteCharacter server.
func (c *Client) DeleteCharacter() goa.Endpoint {
	var (
		decodeResponse = DecodeDeleteCharacterResponse(c.decoder, c.RestoreResponseBody)
	)
	return func(ctx context.Context, v interface{}) (interface{}, error) {
		req, err := c.BuildDeleteCharacterRequest(ctx, v)
		if err != nil {
			return nil, err
		}
		resp, err := c.DeleteCharacterDoer.Do(req)
		if err != nil {
			return nil, goahttp.ErrRequestError("character", "deleteCharacter", err)
		}
		return decodeResponse(resp)
	}
}
