// Code generated by goa v3.9.1, DO NOT EDIT.
//
// item HTTP client encoders and decoders
//
// Command:
// $ goa gen github.com/austinhuck/characters-go/design

package client

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"

	item "github.com/austinhuck/characters-go/gen/item"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// BuildGetItemsRequest instantiates a HTTP request object with method and path
// set to call the "item" service "getItems" endpoint
func (c *Client) BuildGetItemsRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetItemsItemPath()}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("item", "getItems", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeGetItemsResponse returns a decoder for responses returned by the item
// getItems endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeGetItemsResponse may return the following errors:
//   - "IdNotFound" (type *goa.ServiceError): http.StatusNotFound
//   - "BadData" (type *goa.ServiceError): http.StatusBadRequest
//   - "IDsExhausted" (type *goa.ServiceError): http.StatusInsufficientStorage
//   - error: internal error
func DecodeGetItemsResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetItemsResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("item", "getItems", err)
			}
			for _, e := range body {
				if e != nil {
					if err2 := ValidateItemResponse(e); err2 != nil {
						err = goa.MergeErrors(err, err2)
					}
				}
			}
			if err != nil {
				return nil, goahttp.ErrValidationError("item", "getItems", err)
			}
			res := NewGetItemsItemOK(body)
			return res, nil
		case http.StatusNotFound:
			var (
				body GetItemsIDNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("item", "getItems", err)
			}
			err = ValidateGetItemsIDNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("item", "getItems", err)
			}
			return nil, NewGetItemsIDNotFound(&body)
		case http.StatusBadRequest:
			var (
				body GetItemsBadDataResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("item", "getItems", err)
			}
			err = ValidateGetItemsBadDataResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("item", "getItems", err)
			}
			return nil, NewGetItemsBadData(&body)
		case http.StatusInsufficientStorage:
			var (
				body GetItemsIDsExhaustedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("item", "getItems", err)
			}
			err = ValidateGetItemsIDsExhaustedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("item", "getItems", err)
			}
			return nil, NewGetItemsIDsExhausted(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("item", "getItems", resp.StatusCode, string(body))
		}
	}
}

// BuildGetItemRequest instantiates a HTTP request object with method and path
// set to call the "item" service "getItem" endpoint
func (c *Client) BuildGetItemRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id int
	)
	{
		p, ok := v.(*item.ItemID)
		if !ok {
			return nil, goahttp.ErrInvalidType("item", "getItem", "*item.ItemID", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: GetItemItemPath(id)}
	req, err := http.NewRequest("GET", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("item", "getItem", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeGetItemResponse returns a decoder for responses returned by the item
// getItem endpoint. restoreBody controls whether the response body should be
// restored after having been read.
// DecodeGetItemResponse may return the following errors:
//   - "IdNotFound" (type *goa.ServiceError): http.StatusNotFound
//   - "BadData" (type *goa.ServiceError): http.StatusBadRequest
//   - "IDsExhausted" (type *goa.ServiceError): http.StatusInsufficientStorage
//   - error: internal error
func DecodeGetItemResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body GetItemResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("item", "getItem", err)
			}
			err = ValidateGetItemResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("item", "getItem", err)
			}
			res := NewGetItemItemOK(&body)
			return res, nil
		case http.StatusNotFound:
			var (
				body GetItemIDNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("item", "getItem", err)
			}
			err = ValidateGetItemIDNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("item", "getItem", err)
			}
			return nil, NewGetItemIDNotFound(&body)
		case http.StatusBadRequest:
			var (
				body GetItemBadDataResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("item", "getItem", err)
			}
			err = ValidateGetItemBadDataResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("item", "getItem", err)
			}
			return nil, NewGetItemBadData(&body)
		case http.StatusInsufficientStorage:
			var (
				body GetItemIDsExhaustedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("item", "getItem", err)
			}
			err = ValidateGetItemIDsExhaustedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("item", "getItem", err)
			}
			return nil, NewGetItemIDsExhausted(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("item", "getItem", resp.StatusCode, string(body))
		}
	}
}

// BuildCreateItemRequest instantiates a HTTP request object with method and
// path set to call the "item" service "createItem" endpoint
func (c *Client) BuildCreateItemRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: CreateItemItemPath()}
	req, err := http.NewRequest("POST", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("item", "createItem", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeCreateItemRequest returns an encoder for requests sent to the item
// createItem server.
func EncodeCreateItemRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*item.ItemNew)
		if !ok {
			return goahttp.ErrInvalidType("item", "createItem", "*item.ItemNew", v)
		}
		body := NewCreateItemRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("item", "createItem", err)
		}
		return nil
	}
}

// DecodeCreateItemResponse returns a decoder for responses returned by the
// item createItem endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeCreateItemResponse may return the following errors:
//   - "BadData" (type *goa.ServiceError): http.StatusBadRequest
//   - "IDsExhausted" (type *goa.ServiceError): http.StatusInsufficientStorage
//   - "IdNotFound" (type *goa.ServiceError): http.StatusNotFound
//   - error: internal error
func DecodeCreateItemResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body CreateItemResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("item", "createItem", err)
			}
			err = ValidateCreateItemResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("item", "createItem", err)
			}
			res := NewCreateItemItemOK(&body)
			return res, nil
		case http.StatusBadRequest:
			var (
				body CreateItemBadDataResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("item", "createItem", err)
			}
			err = ValidateCreateItemBadDataResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("item", "createItem", err)
			}
			return nil, NewCreateItemBadData(&body)
		case http.StatusInsufficientStorage:
			var (
				body CreateItemIDsExhaustedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("item", "createItem", err)
			}
			err = ValidateCreateItemIDsExhaustedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("item", "createItem", err)
			}
			return nil, NewCreateItemIDsExhausted(&body)
		case http.StatusNotFound:
			var (
				body CreateItemIDNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("item", "createItem", err)
			}
			err = ValidateCreateItemIDNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("item", "createItem", err)
			}
			return nil, NewCreateItemIDNotFound(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("item", "createItem", resp.StatusCode, string(body))
		}
	}
}

// BuildUpdateItemRequest instantiates a HTTP request object with method and
// path set to call the "item" service "updateItem" endpoint
func (c *Client) BuildUpdateItemRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id int
	)
	{
		p, ok := v.(*item.Item)
		if !ok {
			return nil, goahttp.ErrInvalidType("item", "updateItem", "*item.Item", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: UpdateItemItemPath(id)}
	req, err := http.NewRequest("PUT", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("item", "updateItem", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// EncodeUpdateItemRequest returns an encoder for requests sent to the item
// updateItem server.
func EncodeUpdateItemRequest(encoder func(*http.Request) goahttp.Encoder) func(*http.Request, interface{}) error {
	return func(req *http.Request, v interface{}) error {
		p, ok := v.(*item.Item)
		if !ok {
			return goahttp.ErrInvalidType("item", "updateItem", "*item.Item", v)
		}
		body := NewUpdateItemRequestBody(p)
		if err := encoder(req).Encode(&body); err != nil {
			return goahttp.ErrEncodingError("item", "updateItem", err)
		}
		return nil
	}
}

// DecodeUpdateItemResponse returns a decoder for responses returned by the
// item updateItem endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeUpdateItemResponse may return the following errors:
//   - "IdNotFound" (type *goa.ServiceError): http.StatusNotFound
//   - "BadData" (type *goa.ServiceError): http.StatusBadRequest
//   - "IDsExhausted" (type *goa.ServiceError): http.StatusInsufficientStorage
//   - error: internal error
func DecodeUpdateItemResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusOK:
			var (
				body UpdateItemResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("item", "updateItem", err)
			}
			err = ValidateUpdateItemResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("item", "updateItem", err)
			}
			res := NewUpdateItemItemOK(&body)
			return res, nil
		case http.StatusNotFound:
			var (
				body UpdateItemIDNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("item", "updateItem", err)
			}
			err = ValidateUpdateItemIDNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("item", "updateItem", err)
			}
			return nil, NewUpdateItemIDNotFound(&body)
		case http.StatusBadRequest:
			var (
				body UpdateItemBadDataResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("item", "updateItem", err)
			}
			err = ValidateUpdateItemBadDataResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("item", "updateItem", err)
			}
			return nil, NewUpdateItemBadData(&body)
		case http.StatusInsufficientStorage:
			var (
				body UpdateItemIDsExhaustedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("item", "updateItem", err)
			}
			err = ValidateUpdateItemIDsExhaustedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("item", "updateItem", err)
			}
			return nil, NewUpdateItemIDsExhausted(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("item", "updateItem", resp.StatusCode, string(body))
		}
	}
}

// BuildDeleteItemRequest instantiates a HTTP request object with method and
// path set to call the "item" service "deleteItem" endpoint
func (c *Client) BuildDeleteItemRequest(ctx context.Context, v interface{}) (*http.Request, error) {
	var (
		id int
	)
	{
		p, ok := v.(*item.ItemID)
		if !ok {
			return nil, goahttp.ErrInvalidType("item", "deleteItem", "*item.ItemID", v)
		}
		id = p.ID
	}
	u := &url.URL{Scheme: c.scheme, Host: c.host, Path: DeleteItemItemPath(id)}
	req, err := http.NewRequest("DELETE", u.String(), nil)
	if err != nil {
		return nil, goahttp.ErrInvalidURL("item", "deleteItem", u.String(), err)
	}
	if ctx != nil {
		req = req.WithContext(ctx)
	}

	return req, nil
}

// DecodeDeleteItemResponse returns a decoder for responses returned by the
// item deleteItem endpoint. restoreBody controls whether the response body
// should be restored after having been read.
// DecodeDeleteItemResponse may return the following errors:
//   - "IdNotFound" (type *goa.ServiceError): http.StatusNotFound
//   - "BadData" (type *goa.ServiceError): http.StatusBadRequest
//   - "IDsExhausted" (type *goa.ServiceError): http.StatusInsufficientStorage
//   - error: internal error
func DecodeDeleteItemResponse(decoder func(*http.Response) goahttp.Decoder, restoreBody bool) func(*http.Response) (interface{}, error) {
	return func(resp *http.Response) (interface{}, error) {
		if restoreBody {
			b, err := io.ReadAll(resp.Body)
			if err != nil {
				return nil, err
			}
			resp.Body = io.NopCloser(bytes.NewBuffer(b))
			defer func() {
				resp.Body = io.NopCloser(bytes.NewBuffer(b))
			}()
		} else {
			defer resp.Body.Close()
		}
		switch resp.StatusCode {
		case http.StatusNoContent:
			return nil, nil
		case http.StatusNotFound:
			var (
				body DeleteItemIDNotFoundResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("item", "deleteItem", err)
			}
			err = ValidateDeleteItemIDNotFoundResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("item", "deleteItem", err)
			}
			return nil, NewDeleteItemIDNotFound(&body)
		case http.StatusBadRequest:
			var (
				body DeleteItemBadDataResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("item", "deleteItem", err)
			}
			err = ValidateDeleteItemBadDataResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("item", "deleteItem", err)
			}
			return nil, NewDeleteItemBadData(&body)
		case http.StatusInsufficientStorage:
			var (
				body DeleteItemIDsExhaustedResponseBody
				err  error
			)
			err = decoder(resp).Decode(&body)
			if err != nil {
				return nil, goahttp.ErrDecodingError("item", "deleteItem", err)
			}
			err = ValidateDeleteItemIDsExhaustedResponseBody(&body)
			if err != nil {
				return nil, goahttp.ErrValidationError("item", "deleteItem", err)
			}
			return nil, NewDeleteItemIDsExhausted(&body)
		default:
			body, _ := io.ReadAll(resp.Body)
			return nil, goahttp.ErrInvalidResponse("item", "deleteItem", resp.StatusCode, string(body))
		}
	}
}

// unmarshalItemResponseToItemItem builds a value of type *item.Item from a
// value of type *ItemResponse.
func unmarshalItemResponseToItemItem(v *ItemResponse) *item.Item {
	res := &item.Item{
		ID:          *v.ID,
		Name:        *v.Name,
		Description: *v.Description,
		Damage:      *v.Damage,
		Healing:     *v.Healing,
		Protection:  *v.Protection,
	}

	return res
}
