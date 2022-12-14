// Code generated by goa v3.9.1, DO NOT EDIT.
//
// character HTTP server types
//
// Command:
// $ goa gen github.com/austinhuck/characters-go/design

package server

import (
	character "github.com/austinhuck/characters-go/gen/character"
	goa "goa.design/goa/v3/pkg"
)

// CreateCharacterRequestBody is the type of the "character" service
// "createCharacter" endpoint HTTP request body.
type CreateCharacterRequestBody struct {
	// Character name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Character description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Character health
	Health *uint `form:"health,omitempty" json:"health,omitempty" xml:"health,omitempty"`
	// Character experience
	Experience *uint `form:"experience,omitempty" json:"experience,omitempty" xml:"experience,omitempty"`
}

// UpdateCharacterRequestBody is the type of the "character" service
// "updateCharacter" endpoint HTTP request body.
type UpdateCharacterRequestBody struct {
	// Character name
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// Character description
	Description *string `form:"description,omitempty" json:"description,omitempty" xml:"description,omitempty"`
	// Character health
	Health *uint `form:"health,omitempty" json:"health,omitempty" xml:"health,omitempty"`
	// Character experience
	Experience *uint `form:"experience,omitempty" json:"experience,omitempty" xml:"experience,omitempty"`
}

// GetCharactersResponseBody is the type of the "character" service
// "getCharacters" endpoint HTTP response body.
type GetCharactersResponseBody []*CharacterResponse

// GetCharacterResponseBody is the type of the "character" service
// "getCharacter" endpoint HTTP response body.
type GetCharacterResponseBody struct {
	// Character identifier
	ID int `form:"id" json:"id" xml:"id"`
	// Character name
	Name string `form:"name" json:"name" xml:"name"`
	// Character description
	Description string `form:"description" json:"description" xml:"description"`
	// Character health
	Health uint `form:"health" json:"health" xml:"health"`
	// Character experience
	Experience uint `form:"experience" json:"experience" xml:"experience"`
}

// CreateCharacterResponseBody is the type of the "character" service
// "createCharacter" endpoint HTTP response body.
type CreateCharacterResponseBody struct {
	// Character identifier
	ID int `form:"id" json:"id" xml:"id"`
	// Character name
	Name string `form:"name" json:"name" xml:"name"`
	// Character description
	Description string `form:"description" json:"description" xml:"description"`
	// Character health
	Health uint `form:"health" json:"health" xml:"health"`
	// Character experience
	Experience uint `form:"experience" json:"experience" xml:"experience"`
}

// UpdateCharacterResponseBody is the type of the "character" service
// "updateCharacter" endpoint HTTP response body.
type UpdateCharacterResponseBody struct {
	// Character identifier
	ID int `form:"id" json:"id" xml:"id"`
	// Character name
	Name string `form:"name" json:"name" xml:"name"`
	// Character description
	Description string `form:"description" json:"description" xml:"description"`
	// Character health
	Health uint `form:"health" json:"health" xml:"health"`
	// Character experience
	Experience uint `form:"experience" json:"experience" xml:"experience"`
}

// GetCharactersIDNotFoundResponseBody is the type of the "character" service
// "getCharacters" endpoint HTTP response body for the "IdNotFound" error.
type GetCharactersIDNotFoundResponseBody struct {
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

// GetCharactersBadDataResponseBody is the type of the "character" service
// "getCharacters" endpoint HTTP response body for the "BadData" error.
type GetCharactersBadDataResponseBody struct {
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

// GetCharactersIDsExhaustedResponseBody is the type of the "character" service
// "getCharacters" endpoint HTTP response body for the "IDsExhausted" error.
type GetCharactersIDsExhaustedResponseBody struct {
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

// GetCharacterIDNotFoundResponseBody is the type of the "character" service
// "getCharacter" endpoint HTTP response body for the "IdNotFound" error.
type GetCharacterIDNotFoundResponseBody struct {
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

// GetCharacterBadDataResponseBody is the type of the "character" service
// "getCharacter" endpoint HTTP response body for the "BadData" error.
type GetCharacterBadDataResponseBody struct {
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

// GetCharacterIDsExhaustedResponseBody is the type of the "character" service
// "getCharacter" endpoint HTTP response body for the "IDsExhausted" error.
type GetCharacterIDsExhaustedResponseBody struct {
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

// CreateCharacterBadDataResponseBody is the type of the "character" service
// "createCharacter" endpoint HTTP response body for the "BadData" error.
type CreateCharacterBadDataResponseBody struct {
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

// CreateCharacterIDsExhaustedResponseBody is the type of the "character"
// service "createCharacter" endpoint HTTP response body for the "IDsExhausted"
// error.
type CreateCharacterIDsExhaustedResponseBody struct {
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

// CreateCharacterIDNotFoundResponseBody is the type of the "character" service
// "createCharacter" endpoint HTTP response body for the "IdNotFound" error.
type CreateCharacterIDNotFoundResponseBody struct {
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

// UpdateCharacterIDNotFoundResponseBody is the type of the "character" service
// "updateCharacter" endpoint HTTP response body for the "IdNotFound" error.
type UpdateCharacterIDNotFoundResponseBody struct {
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

// UpdateCharacterBadDataResponseBody is the type of the "character" service
// "updateCharacter" endpoint HTTP response body for the "BadData" error.
type UpdateCharacterBadDataResponseBody struct {
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

// UpdateCharacterIDsExhaustedResponseBody is the type of the "character"
// service "updateCharacter" endpoint HTTP response body for the "IDsExhausted"
// error.
type UpdateCharacterIDsExhaustedResponseBody struct {
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

// DeleteCharacterIDNotFoundResponseBody is the type of the "character" service
// "deleteCharacter" endpoint HTTP response body for the "IdNotFound" error.
type DeleteCharacterIDNotFoundResponseBody struct {
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

// DeleteCharacterBadDataResponseBody is the type of the "character" service
// "deleteCharacter" endpoint HTTP response body for the "BadData" error.
type DeleteCharacterBadDataResponseBody struct {
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

// DeleteCharacterIDsExhaustedResponseBody is the type of the "character"
// service "deleteCharacter" endpoint HTTP response body for the "IDsExhausted"
// error.
type DeleteCharacterIDsExhaustedResponseBody struct {
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

// CharacterResponse is used to define fields on response body types.
type CharacterResponse struct {
	// Character identifier
	ID int `form:"id" json:"id" xml:"id"`
	// Character name
	Name string `form:"name" json:"name" xml:"name"`
	// Character description
	Description string `form:"description" json:"description" xml:"description"`
	// Character health
	Health uint `form:"health" json:"health" xml:"health"`
	// Character experience
	Experience uint `form:"experience" json:"experience" xml:"experience"`
}

// NewGetCharactersResponseBody builds the HTTP response body from the result
// of the "getCharacters" endpoint of the "character" service.
func NewGetCharactersResponseBody(res []*character.Character) GetCharactersResponseBody {
	body := make([]*CharacterResponse, len(res))
	for i, val := range res {
		body[i] = marshalCharacterCharacterToCharacterResponse(val)
	}
	return body
}

// NewGetCharacterResponseBody builds the HTTP response body from the result of
// the "getCharacter" endpoint of the "character" service.
func NewGetCharacterResponseBody(res *character.Character) *GetCharacterResponseBody {
	body := &GetCharacterResponseBody{
		ID:          res.ID,
		Name:        res.Name,
		Description: res.Description,
		Health:      res.Health,
		Experience:  res.Experience,
	}
	return body
}

// NewCreateCharacterResponseBody builds the HTTP response body from the result
// of the "createCharacter" endpoint of the "character" service.
func NewCreateCharacterResponseBody(res *character.Character) *CreateCharacterResponseBody {
	body := &CreateCharacterResponseBody{
		ID:          res.ID,
		Name:        res.Name,
		Description: res.Description,
		Health:      res.Health,
		Experience:  res.Experience,
	}
	return body
}

// NewUpdateCharacterResponseBody builds the HTTP response body from the result
// of the "updateCharacter" endpoint of the "character" service.
func NewUpdateCharacterResponseBody(res *character.Character) *UpdateCharacterResponseBody {
	body := &UpdateCharacterResponseBody{
		ID:          res.ID,
		Name:        res.Name,
		Description: res.Description,
		Health:      res.Health,
		Experience:  res.Experience,
	}
	return body
}

// NewGetCharactersIDNotFoundResponseBody builds the HTTP response body from
// the result of the "getCharacters" endpoint of the "character" service.
func NewGetCharactersIDNotFoundResponseBody(res *goa.ServiceError) *GetCharactersIDNotFoundResponseBody {
	body := &GetCharactersIDNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetCharactersBadDataResponseBody builds the HTTP response body from the
// result of the "getCharacters" endpoint of the "character" service.
func NewGetCharactersBadDataResponseBody(res *goa.ServiceError) *GetCharactersBadDataResponseBody {
	body := &GetCharactersBadDataResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetCharactersIDsExhaustedResponseBody builds the HTTP response body from
// the result of the "getCharacters" endpoint of the "character" service.
func NewGetCharactersIDsExhaustedResponseBody(res *goa.ServiceError) *GetCharactersIDsExhaustedResponseBody {
	body := &GetCharactersIDsExhaustedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetCharacterIDNotFoundResponseBody builds the HTTP response body from the
// result of the "getCharacter" endpoint of the "character" service.
func NewGetCharacterIDNotFoundResponseBody(res *goa.ServiceError) *GetCharacterIDNotFoundResponseBody {
	body := &GetCharacterIDNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetCharacterBadDataResponseBody builds the HTTP response body from the
// result of the "getCharacter" endpoint of the "character" service.
func NewGetCharacterBadDataResponseBody(res *goa.ServiceError) *GetCharacterBadDataResponseBody {
	body := &GetCharacterBadDataResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetCharacterIDsExhaustedResponseBody builds the HTTP response body from
// the result of the "getCharacter" endpoint of the "character" service.
func NewGetCharacterIDsExhaustedResponseBody(res *goa.ServiceError) *GetCharacterIDsExhaustedResponseBody {
	body := &GetCharacterIDsExhaustedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateCharacterBadDataResponseBody builds the HTTP response body from the
// result of the "createCharacter" endpoint of the "character" service.
func NewCreateCharacterBadDataResponseBody(res *goa.ServiceError) *CreateCharacterBadDataResponseBody {
	body := &CreateCharacterBadDataResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateCharacterIDsExhaustedResponseBody builds the HTTP response body
// from the result of the "createCharacter" endpoint of the "character" service.
func NewCreateCharacterIDsExhaustedResponseBody(res *goa.ServiceError) *CreateCharacterIDsExhaustedResponseBody {
	body := &CreateCharacterIDsExhaustedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewCreateCharacterIDNotFoundResponseBody builds the HTTP response body from
// the result of the "createCharacter" endpoint of the "character" service.
func NewCreateCharacterIDNotFoundResponseBody(res *goa.ServiceError) *CreateCharacterIDNotFoundResponseBody {
	body := &CreateCharacterIDNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateCharacterIDNotFoundResponseBody builds the HTTP response body from
// the result of the "updateCharacter" endpoint of the "character" service.
func NewUpdateCharacterIDNotFoundResponseBody(res *goa.ServiceError) *UpdateCharacterIDNotFoundResponseBody {
	body := &UpdateCharacterIDNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateCharacterBadDataResponseBody builds the HTTP response body from the
// result of the "updateCharacter" endpoint of the "character" service.
func NewUpdateCharacterBadDataResponseBody(res *goa.ServiceError) *UpdateCharacterBadDataResponseBody {
	body := &UpdateCharacterBadDataResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewUpdateCharacterIDsExhaustedResponseBody builds the HTTP response body
// from the result of the "updateCharacter" endpoint of the "character" service.
func NewUpdateCharacterIDsExhaustedResponseBody(res *goa.ServiceError) *UpdateCharacterIDsExhaustedResponseBody {
	body := &UpdateCharacterIDsExhaustedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewDeleteCharacterIDNotFoundResponseBody builds the HTTP response body from
// the result of the "deleteCharacter" endpoint of the "character" service.
func NewDeleteCharacterIDNotFoundResponseBody(res *goa.ServiceError) *DeleteCharacterIDNotFoundResponseBody {
	body := &DeleteCharacterIDNotFoundResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewDeleteCharacterBadDataResponseBody builds the HTTP response body from the
// result of the "deleteCharacter" endpoint of the "character" service.
func NewDeleteCharacterBadDataResponseBody(res *goa.ServiceError) *DeleteCharacterBadDataResponseBody {
	body := &DeleteCharacterBadDataResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewDeleteCharacterIDsExhaustedResponseBody builds the HTTP response body
// from the result of the "deleteCharacter" endpoint of the "character" service.
func NewDeleteCharacterIDsExhaustedResponseBody(res *goa.ServiceError) *DeleteCharacterIDsExhaustedResponseBody {
	body := &DeleteCharacterIDsExhaustedResponseBody{
		Name:      res.Name,
		ID:        res.ID,
		Message:   res.Message,
		Temporary: res.Temporary,
		Timeout:   res.Timeout,
		Fault:     res.Fault,
	}
	return body
}

// NewGetCharacterCharacterID builds a character service getCharacter endpoint
// payload.
func NewGetCharacterCharacterID(id int) *character.CharacterID {
	v := &character.CharacterID{}
	v.ID = id

	return v
}

// NewCreateCharacterCharacterNew builds a character service createCharacter
// endpoint payload.
func NewCreateCharacterCharacterNew(body *CreateCharacterRequestBody) *character.CharacterNew {
	v := &character.CharacterNew{
		Name:        *body.Name,
		Description: *body.Description,
		Health:      *body.Health,
		Experience:  *body.Experience,
	}

	return v
}

// NewUpdateCharacterCharacter builds a character service updateCharacter
// endpoint payload.
func NewUpdateCharacterCharacter(body *UpdateCharacterRequestBody, id int) *character.Character {
	v := &character.Character{
		Name:        *body.Name,
		Description: *body.Description,
		Health:      *body.Health,
		Experience:  *body.Experience,
	}
	v.ID = id

	return v
}

// NewDeleteCharacterCharacterID builds a character service deleteCharacter
// endpoint payload.
func NewDeleteCharacterCharacterID(id int) *character.CharacterID {
	v := &character.CharacterID{}
	v.ID = id

	return v
}

// ValidateCreateCharacterRequestBody runs the validations defined on
// CreateCharacterRequestBody
func ValidateCreateCharacterRequestBody(body *CreateCharacterRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Description == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("description", "body"))
	}
	if body.Health == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("health", "body"))
	}
	if body.Experience == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("experience", "body"))
	}
	return
}

// ValidateUpdateCharacterRequestBody runs the validations defined on
// UpdateCharacterRequestBody
func ValidateUpdateCharacterRequestBody(body *UpdateCharacterRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Description == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("description", "body"))
	}
	if body.Health == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("health", "body"))
	}
	if body.Experience == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("experience", "body"))
	}
	return
}
