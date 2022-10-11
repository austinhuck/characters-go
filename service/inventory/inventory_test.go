package inventory

import (
	"context"
	"github.com/austinhuck/characters-go/gen/character"
	"github.com/austinhuck/characters-go/gen/inventory"
	"github.com/austinhuck/characters-go/gen/item"
	"log"
	"os"
	"reflect"
	"sync"
	"testing"
)

var logger = log.New(os.Stderr, "[inventory_test] ", log.Ltime)

type mockCharacterService struct{}

func (mockCharacterService) GetCharacters(ctx context.Context) (res []*character.Character, err error) {
	return nil, nil
}

func (mockCharacterService) GetCharacter(ctx context.Context, id *character.CharacterID) (res *character.Character, err error) {
	return nil, nil
}

func (mockCharacterService) CreateCharacter(ctx context.Context, characterNew *character.CharacterNew) (res *character.Character, err error) {
	return nil, nil
}

func (mockCharacterService) UpdateCharacter(ctx context.Context, c *character.Character) (res *character.Character, err error) {
	return nil, nil
}

func (mockCharacterService) DeleteCharacter(ctx context.Context, id *character.CharacterID) (err error) {
	return nil
}

type mockItemService struct{}

func (mockItemService) GetItems(ctx context.Context) (res []*item.Item, err error) {
	return nil, nil
}

func (mockItemService) GetItem(ctx context.Context, id *item.ItemID) (res *item.Item, err error) {
	return nil, nil
}

func (mockItemService) CreateItem(ctx context.Context, itemNew *item.ItemNew) (res *item.Item, err error) {
	return nil, nil
}

func (mockItemService) UpdateItem(ctx context.Context, i *item.Item) (res *item.Item, err error) {
	return nil, nil
}

func (mockItemService) DeleteItem(ctx context.Context, id *item.ItemID) (err error) {
	return nil
}

func Test_inventorysrvc_ClearInventory(t *testing.T) {
	type fields struct {
		logger          *log.Logger
		mtx             sync.Mutex
		inventory       map[int]map[int]*inventory.InventoryEntry
		characterClient character.Service
		itemClient      item.Service
	}
	type args struct {
		ctx context.Context
		p   *inventory.CharacterID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Normal",
			fields: fields{
				logger: logger,
				inventory: map[int]map[int]*inventory.InventoryEntry{
					1: {
						1: {
							Item:     1,
							Quantity: 5,
						},
						2: {
							Item:     2,
							Quantity: 7,
						},
					},
					2: {
						1: {
							Item:     1,
							Quantity: 1,
						},
					},
					3: {
						2: {
							Item:     2,
							Quantity: 23,
						},
						3: {
							Item:     3,
							Quantity: 1,
						},
					},
				},
				characterClient: mockCharacterService{},
				itemClient:      mockItemService{},
			},
			args: args{
				ctx: nil,
				p: &inventory.CharacterID{
					ID: 2,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &inventorysrvc{
				logger:          tt.fields.logger,
				mtx:             tt.fields.mtx,
				inventory:       tt.fields.inventory,
				characterClient: tt.fields.characterClient,
				itemClient:      tt.fields.itemClient,
			}
			if err := s.ClearInventory(tt.args.ctx, tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("ClearInventory() error = %v, wantErr %v", err, tt.wantErr)
			}
			res, err := s.GetInventory(tt.args.ctx, tt.args.p)
			if err != nil {
				t.Errorf("GetInventory() error = %v", err)
			}
			if len(res) != 0 {
				t.Errorf("GetInventory() returned %d inventory entries, expected 0", len(res))
			}
		})
	}
}

func Test_inventorysrvc_GetInventory(t *testing.T) {
	type fields struct {
		logger          *log.Logger
		mtx             sync.Mutex
		inventory       map[int]map[int]*inventory.InventoryEntry
		characterClient character.Service
		itemClient      item.Service
	}
	type args struct {
		ctx context.Context
		p   *inventory.CharacterID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*inventory.InventoryEntry
		wantErr bool
	}{
		{
			name: "Normal",
			fields: fields{
				logger: logger,
				inventory: map[int]map[int]*inventory.InventoryEntry{
					1: {
						1: {
							Item:     1,
							Quantity: 5,
						},
						2: {
							Item:     2,
							Quantity: 7,
						},
					},
					2: {
						1: {
							Item:     1,
							Quantity: 1,
						},
					},
					3: {
						2: {
							Item:     2,
							Quantity: 23,
						},
						3: {
							Item:     3,
							Quantity: 1,
						},
					},
				},
				characterClient: mockCharacterService{},
				itemClient:      mockItemService{},
			},
			args: args{
				ctx: nil,
				p: &inventory.CharacterID{
					ID: 3,
				},
			},
			want: []*inventory.InventoryEntry{
				{
					Item:     2,
					Quantity: 23,
				},
				{
					Item:     3,
					Quantity: 1,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &inventorysrvc{
				logger:          tt.fields.logger,
				mtx:             tt.fields.mtx,
				inventory:       tt.fields.inventory,
				characterClient: tt.fields.characterClient,
				itemClient:      tt.fields.itemClient,
			}
			got, err := s.GetInventory(tt.args.ctx, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetInventory() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetInventory() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_inventorysrvc_UpdateInventory(t *testing.T) {
	type fields struct {
		logger          *log.Logger
		mtx             sync.Mutex
		inventory       map[int]map[int]*inventory.InventoryEntry
		characterClient character.Service
		itemClient      item.Service
	}
	type args struct {
		ctx context.Context
		p   *inventory.InventoryUpdate
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		{
			name: "Normal",
			fields: fields{
				logger: logger,
				inventory: map[int]map[int]*inventory.InventoryEntry{
					1: {
						1: {
							Item:     1,
							Quantity: 5,
						},
						2: {
							Item:     2,
							Quantity: 7,
						},
					},
					2: {
						1: {
							Item:     1,
							Quantity: 1,
						},
					},
					3: {
						2: {
							Item:     2,
							Quantity: 23,
						},
						3: {
							Item:     3,
							Quantity: 1,
						},
					},
				},
				characterClient: mockCharacterService{},
				itemClient:      mockItemService{},
			},
			args: args{
				ctx: nil,
				p: &inventory.InventoryUpdate{
					ID:       3,
					Item:     4,
					Quantity: 1,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &inventorysrvc{
				logger:          tt.fields.logger,
				mtx:             tt.fields.mtx,
				inventory:       tt.fields.inventory,
				characterClient: tt.fields.characterClient,
				itemClient:      tt.fields.itemClient,
			}
			if err := s.UpdateInventory(tt.args.ctx, tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("UpdateInventory() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
