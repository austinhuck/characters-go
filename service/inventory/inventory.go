package inventory

import (
	"context"
	"github.com/austinhuck/characters-go/gen/character"
	"github.com/austinhuck/characters-go/gen/item"
	"log"
	"sync"

	"github.com/austinhuck/characters-go/gen/inventory"
)

// inventory service implementation.
type inventorysrvc struct {
	logger          *log.Logger
	mtx             sync.Mutex
	inventory       map[int]map[int]*inventory.InventoryEntry
	characterClient character.Service
	itemClient      item.Service
}

// NewInventory returns the inventory service implementation.
func NewInventory(logger *log.Logger, characterClient character.Service, itemClient item.Service) inventory.Service {
	return &inventorysrvc{
		logger:          logger,
		inventory:       make(map[int]map[int]*inventory.InventoryEntry),
		characterClient: characterClient,
		itemClient:      itemClient,
	}
}

// GetInventory implements getInventory.
func (s *inventorysrvc) GetInventory(ctx context.Context, p *inventory.CharacterID) ([]*inventory.InventoryEntry, error) {
	s.logger.Print("inventory.getInventory")
	// Check with the character service that the ID exists
	_, err := s.characterClient.GetCharacter(ctx, &character.CharacterID{ID: p.ID})
	if err != nil {
		return nil, err
	}

	s.mtx.Lock()
	defer s.mtx.Unlock()

	// Try to find inventory for the character ID
	characterInventory := s.inventory[p.ID]
	if characterInventory == nil {
		// Character has no inventory, return empty slice
		return make([]*inventory.InventoryEntry, 0), nil
	}

	// Character does have an inventory, collect the map values into a slice
	result := make([]*inventory.InventoryEntry, 0, len(characterInventory))
	for _, value := range characterInventory {
		result = append(result, value)
	}

	return result, nil
}

// UpdateInventory implements updateInventory.
func (s *inventorysrvc) UpdateInventory(ctx context.Context, p *inventory.InventoryUpdate) (err error) {
	s.logger.Print("inventory.updateInventory")

	// Check with the character service that the ID exists
	_, err = s.characterClient.GetCharacter(ctx, &character.CharacterID{ID: p.ID})
	if err != nil {
		return err
	}

	// Check with the item service that the ID exists
	_, err = s.itemClient.GetItem(ctx, &item.ItemID{ID: p.Item})
	if err != nil {
		return err
	}

	// Go ahead and update inventory
	s.mtx.Lock()
	defer s.mtx.Unlock()

	if p.Quantity == 0 {
		// Remove the inventory entry from the character's inventory map
		delete(s.inventory[p.ID], p.Item)
		return nil
	}

	if s.inventory[p.ID] == nil {
		// Populate the inventory map for the character
		s.inventory[p.ID] = make(map[int]*inventory.InventoryEntry)
	}

	// Add or update the inventory entry in the character's inventory map
	s.inventory[p.ID][p.Item] = &inventory.InventoryEntry{
		Item:     p.Item,
		Quantity: p.Quantity,
	}

	return nil
}

// ClearInventory implements clearInventory.
func (s *inventorysrvc) ClearInventory(ctx context.Context, p *inventory.CharacterID) (err error) {
	s.logger.Print("inventory.clearInventory")

	// Check with the character service that the ID exists
	_, err = s.characterClient.GetCharacter(ctx, &character.CharacterID{ID: p.ID})
	if err != nil {
		return err
	}

	// Go ahead and clear inventory
	s.mtx.Lock()
	defer s.mtx.Unlock()

	delete(s.inventory, p.ID)

	return nil
}
