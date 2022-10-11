package character

import (
	"context"
	"fmt"
	"log"
	"math"
	"strings"
	"sync"

	"github.com/austinhuck/characters-go/gen/character"
)

// Character service implementation.
type charactersrvc struct {
	logger             *log.Logger
	mtx                sync.Mutex
	characters         map[int]*character.Character
	characterNameIndex map[string]*int
	currentId          int
}

// NewCharacter returns the character service implementation.
func NewCharacter(logger *log.Logger) character.Service {
	svc := charactersrvc{
		logger:             logger,
		characters:         make(map[int]*character.Character),
		characterNameIndex: make(map[string]*int),
		currentId:          0,
	}

	_, _ = svc.addCharacter(&character.CharacterNew{Name: "Grog", Description: "Goliath barbarian", Health: 92, Experience: 4100})
	_, _ = svc.addCharacter(&character.CharacterNew{Name: "Fjord", Description: "Half-orc warlock", Health: 86, Experience: 3500})

	return &svc
}

// GetCharacters implements getCharacters.
func (s *charactersrvc) GetCharacters(ctx context.Context) ([]*character.Character, error) {
	s.logger.Print("character.getCharacters")
	s.mtx.Lock()
	defer s.mtx.Unlock()

	res := make([]*character.Character, 0, len(s.characters))
	for _, value := range s.characters {
		res = append(res, value)
	}

	return res, nil
}

// GetCharacter implements getCharacter.
func (s *charactersrvc) GetCharacter(ctx context.Context, p *character.CharacterID) (*character.Character, error) {
	s.logger.Print("character.getCharacter")
	s.mtx.Lock()
	defer s.mtx.Unlock()

	result := s.characters[p.ID]
	if result == nil {
		return nil, character.MakeIDNotFound(fmt.Errorf("character id '%d' does not exist", p.ID))
	}

	return result, nil
}

// CreateCharacter implements createCharacter.
func (s *charactersrvc) CreateCharacter(ctx context.Context, newCharacter *character.CharacterNew) (*character.Character, error) {
	s.logger.Print("character.createCharacter")
	return s.addCharacter(newCharacter)
}

// UpdateCharacter implements updateCharacter.
func (s *charactersrvc) UpdateCharacter(ctx context.Context, updateCharacter *character.Character) (*character.Character, error) {
	s.logger.Print("character.updateCharacter")
	s.mtx.Lock()
	defer s.mtx.Unlock()

	existingCharacter := s.characters[updateCharacter.ID]
	if existingCharacter == nil {
		return nil, character.MakeIDNotFound(fmt.Errorf("character id '%d' does not exist", updateCharacter.ID))
	}

	// Check for duplicate name
	nameMatchedCharacterID := s.characterNameIndex[strings.ToLower(updateCharacter.Name)]
	// If a character with the same new name is found it must be the existing character, otherwise it's a duplicate!
	if nameMatchedCharacterID != nil && *nameMatchedCharacterID != existingCharacter.ID {
		return nil, character.MakeBadData(fmt.Errorf("character name '%s' already exists", updateCharacter.Name))
	}

	// replace character and update name index
	s.characters[updateCharacter.ID] = updateCharacter
	delete(s.characterNameIndex, strings.ToLower(existingCharacter.Name))
	s.characterNameIndex[strings.ToLower(updateCharacter.Name)] = &updateCharacter.ID

	return updateCharacter, nil
}

// DeleteCharacter implements deleteCharacter.
func (s *charactersrvc) DeleteCharacter(ctx context.Context, p *character.CharacterID) error {
	s.logger.Print("character.deleteCharacter")
	s.mtx.Lock()
	defer s.mtx.Unlock()

	deleteCharacter := s.characters[p.ID]
	if deleteCharacter == nil {
		return character.MakeIDNotFound(fmt.Errorf("character id '%d' does not exist", p.ID))
	}

	delete(s.characters, deleteCharacter.ID)
	delete(s.characterNameIndex, strings.ToLower(deleteCharacter.Name))

	return nil
}

func (s *charactersrvc) addCharacter(p *character.CharacterNew) (*character.Character, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	// Find if character name already exists, case-insensitive
	if s.characterNameIndex[strings.ToLower(p.Name)] != nil {
		return nil, character.MakeBadData(fmt.Errorf("character name '%s' already exists", p.Name))
	}
	if s.currentId == math.MaxInt {
		// out of IDs
		return nil, character.MakeIDsExhausted(fmt.Errorf("no more IDs available"))
	}
	s.currentId++

	newCharacter := &character.Character{
		ID:          s.currentId,
		Name:        p.Name,
		Description: p.Description,
		Health:      p.Health,
		Experience:  p.Experience,
	}

	s.characters[newCharacter.ID] = newCharacter
	// Add to character name index, case-insensitive
	s.characterNameIndex[strings.ToLower(newCharacter.Name)] = &(newCharacter.ID)

	return newCharacter, nil
}
