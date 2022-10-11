package item

import (
	"context"
	"fmt"
	"log"
	"math"
	"strings"
	"sync"

	"github.com/austinhuck/characters-go/gen/item"
)

// Item service implementation.
type itemsrvc struct {
	logger        *log.Logger
	mtx           sync.Mutex
	items         map[int]*item.Item
	itemNameIndex map[string]*int
	currentId     int
}

// NewItem returns the item service implementation.
func NewItem(logger *log.Logger) item.Service {
	svc := itemsrvc{
		logger:        logger,
		items:         make(map[int]*item.Item),
		itemNameIndex: make(map[string]*int),
		currentId:     0,
	}

	_, _ = svc.addItem(&item.ItemNew{
		Name:        "Bag of Holding",
		Description: "This bag has an interior space considerably larger than its outside dimensions.",
		Damage:      0,
		Healing:     0,
		Protection:  0,
	})
	_, _ = svc.addItem(&item.ItemNew{
		Name:        "Greatsword",
		Description: "A very large sword.",
		Damage:      8,
		Healing:     0,
		Protection:  0,
	})
	_, _ = svc.addItem(&item.ItemNew{
		Name:        "Chain Shirt",
		Description: "Medium armor",
		Damage:      0,
		Healing:     0,
		Protection:  2,
	})
	_, _ = svc.addItem(&item.ItemNew{
		Name:        "Potion of Healing",
		Description: "Regain hit points when you drink this potion ",
		Damage:      0,
		Healing:     10,
		Protection:  0,
	})

	return &svc
}

// GetItems implements getItems.
func (s *itemsrvc) GetItems(ctx context.Context) ([]*item.Item, error) {
	s.logger.Print("item.getItems")
	s.mtx.Lock()
	defer s.mtx.Unlock()

	res := make([]*item.Item, 0, len(s.items))
	for _, value := range s.items {
		res = append(res, value)
	}

	return res, nil
}

// GetItem implements getItem.
func (s *itemsrvc) GetItem(ctx context.Context, p *item.ItemID) (*item.Item, error) {
	s.logger.Print("item.getItem")
	s.mtx.Lock()
	defer s.mtx.Unlock()

	result := s.items[p.ID]
	if result == nil {
		return nil, item.MakeIDNotFound(fmt.Errorf("item id '%d' does not exist", p.ID))
	}

	return result, nil
}

// CreateItem implements createItem.
func (s *itemsrvc) CreateItem(ctx context.Context, newItem *item.ItemNew) (*item.Item, error) {
	s.logger.Print("item.createItem")
	return s.addItem(newItem)
}

// UpdateItem implements updateItem.
func (s *itemsrvc) UpdateItem(ctx context.Context, updateItem *item.Item) (*item.Item, error) {
	s.logger.Print("item.updateItem")
	s.mtx.Lock()
	defer s.mtx.Unlock()

	existingItem := s.items[updateItem.ID]
	if existingItem == nil {
		return nil, item.MakeIDNotFound(fmt.Errorf("item id '%d' does not exist", updateItem.ID))
	}

	// Check for duplicate name
	nameMatchedItemID := s.itemNameIndex[strings.ToLower(updateItem.Name)]
	// If an item with the same new name is found it must be the existing item, otherwise it's a duplicate!
	if nameMatchedItemID != nil && *nameMatchedItemID != existingItem.ID {
		return nil, item.MakeBadData(fmt.Errorf("item name '%s' already exists", updateItem.Name))
	}

	// replace item and update name index
	s.items[updateItem.ID] = updateItem
	delete(s.itemNameIndex, strings.ToLower(existingItem.Name))
	s.itemNameIndex[strings.ToLower(updateItem.Name)] = &updateItem.ID

	return updateItem, nil
}

// DeleteItem implements deleteItem.
func (s *itemsrvc) DeleteItem(ctx context.Context, p *item.ItemID) error {
	s.logger.Print("item.deleteItem")
	s.mtx.Lock()
	defer s.mtx.Unlock()

	deleteItem := s.items[p.ID]
	if deleteItem == nil {
		return item.MakeIDNotFound(fmt.Errorf("item id '%d' does not exist", p.ID))
	}

	delete(s.items, deleteItem.ID)
	delete(s.itemNameIndex, strings.ToLower(deleteItem.Name))

	return nil
}

func (s *itemsrvc) addItem(p *item.ItemNew) (*item.Item, error) {
	s.mtx.Lock()
	defer s.mtx.Unlock()
	// Find if item name already exists, case-insensitive
	if s.itemNameIndex[strings.ToLower(p.Name)] != nil {
		return nil, item.MakeBadData(fmt.Errorf("item name '%s' already exists", p.Name))
	}
	if s.currentId == math.MaxInt {
		// out of IDs
		return nil, item.MakeIDsExhausted(fmt.Errorf("no more IDs available"))
	}
	s.currentId++

	newItem := &item.Item{
		ID:          s.currentId,
		Name:        p.Name,
		Description: p.Description,
		Damage:      p.Damage,
		Healing:     p.Healing,
		Protection:  p.Protection,
	}

	s.items[newItem.ID] = newItem
	// Add to item name index, case-insensitive
	s.itemNameIndex[strings.ToLower(newItem.Name)] = &(newItem.ID)

	return newItem, nil
}
