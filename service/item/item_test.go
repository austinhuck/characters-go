package item

import (
	"context"
	"github.com/austinhuck/characters-go/gen/item"
	"log"
	"math"
	"os"
	"reflect"
	"sync"
	"testing"
)

var logger = log.New(os.Stderr, "[item_test] ", log.Ltime)

func intPtr(literal int) *int {
	return &literal
}

func Test_itemsrvc_CreateItem(t *testing.T) {
	type fields struct {
		logger        *log.Logger
		mtx           sync.Mutex
		items         map[int]*item.Item
		itemNameIndex map[string]*int
		currentId     int
	}
	type args struct {
		ctx     context.Context
		newItem *item.ItemNew
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *item.Item
		wantErr bool
	}{
		{
			name: "Normal",
			fields: fields{
				logger: logger,
				items: map[int]*item.Item{
					0: {
						ID:          0,
						Name:        "Bag of Holding",
						Description: "This bag has an interior space considerably larger than its outside dimensions.",
						Damage:      0,
						Healing:     0,
						Protection:  0,
					},
					1: {
						ID:          1,
						Name:        "Greatsword",
						Description: "A very large sword.",
						Damage:      8,
						Healing:     0,
						Protection:  0,
					}},
				itemNameIndex: map[string]*int{
					"bag of holding": intPtr(0),
					"greatsword":     intPtr(1),
				},
				currentId: 1,
			},
			args: args{nil, &item.ItemNew{
				Name:        "Potion of Healing",
				Description: "Regain hit points when you drink this potion ",
				Damage:      0,
				Healing:     10,
				Protection:  0,
			}},
			want: &item.Item{
				ID:          2,
				Name:        "Potion of Healing",
				Description: "Regain hit points when you drink this potion ",
				Damage:      0,
				Healing:     10,
				Protection:  0,
			},
			wantErr: false,
		},
		{
			name: "Duplicate_Name",
			fields: fields{
				logger: logger,
				items: map[int]*item.Item{
					0: {
						ID:          0,
						Name:        "Bag of Holding",
						Description: "This bag has an interior space considerably larger than its outside dimensions.",
						Damage:      0,
						Healing:     0,
						Protection:  0,
					},
					1: {
						ID:          1,
						Name:        "Greatsword",
						Description: "A very large sword.",
						Damage:      8,
						Healing:     0,
						Protection:  0,
					}},
				itemNameIndex: map[string]*int{
					"bag of holding": intPtr(0),
					"greatsword":     intPtr(1),
				},
				currentId: 1,
			},
			args: args{nil, &item.ItemNew{
				Name:        "Bag of Holding",
				Description: "This bag has an interior space considerably larger than its outside dimensions.",
				Damage:      0,
				Healing:     0,
				Protection:  0,
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "IDs_Max",
			fields: fields{
				logger: logger,
				items: map[int]*item.Item{
					math.MaxInt - 1: {
						ID:          0,
						Name:        "Bag of Holding",
						Description: "This bag has an interior space considerably larger than its outside dimensions.",
						Damage:      0,
						Healing:     0,
						Protection:  0,
					},
					math.MaxInt: {
						ID:          1,
						Name:        "Greatsword",
						Description: "A very large sword.",
						Damage:      8,
						Healing:     0,
						Protection:  0,
					}},
				itemNameIndex: map[string]*int{
					"grog":  intPtr(math.MaxInt - 1),
					"fjord": intPtr(math.MaxInt),
				},
				currentId: math.MaxInt,
			},
			args: args{nil, &item.ItemNew{
				Name:        "Bag of Holding",
				Description: "This bag has an interior space considerably larger than its outside dimensions.",
				Damage:      0,
				Healing:     0,
				Protection:  0,
			}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &itemsrvc{
				logger:        tt.fields.logger,
				mtx:           tt.fields.mtx,
				items:         tt.fields.items,
				itemNameIndex: tt.fields.itemNameIndex,
				currentId:     tt.fields.currentId,
			}
			got, err := s.CreateItem(tt.args.ctx, tt.args.newItem)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateItem() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_itemsrvc_DeleteItem(t *testing.T) {
	type fields struct {
		logger        *log.Logger
		mtx           sync.Mutex
		items         map[int]*item.Item
		itemNameIndex map[string]*int
		currentId     int
	}
	type args struct {
		ctx context.Context
		p   *item.ItemID
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
				items: map[int]*item.Item{
					0: {
						ID:          0,
						Name:        "Bag of Holding",
						Description: "This bag has an interior space considerably larger than its outside dimensions.",
						Damage:      0,
						Healing:     0,
						Protection:  0,
					},
					1: {
						ID:          1,
						Name:        "Greatsword",
						Description: "A very large sword.",
						Damage:      8,
						Healing:     0,
						Protection:  0,
					}},
				itemNameIndex: map[string]*int{
					"bag of holding": intPtr(0),
					"greatsword":     intPtr(1),
				},
				currentId: 1,
			},
			args: args{nil, &item.ItemID{
				ID: 0,
			}},
			wantErr: false,
		},
		{
			name: "Unknown_ID",
			fields: fields{
				logger: logger,
				items: map[int]*item.Item{
					0: {
						ID:          0,
						Name:        "Bag of Holding",
						Description: "This bag has an interior space considerably larger than its outside dimensions.",
						Damage:      0,
						Healing:     0,
						Protection:  0,
					},
					1: {
						ID:          1,
						Name:        "Greatsword",
						Description: "A very large sword.",
						Damage:      8,
						Healing:     0,
						Protection:  0,
					}},
				itemNameIndex: map[string]*int{
					"bag of holding": intPtr(0),
					"greatsword":     intPtr(1),
				},
				currentId: 1,
			},
			args: args{nil, &item.ItemID{
				ID: 23,
			}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &itemsrvc{
				logger:        tt.fields.logger,
				mtx:           tt.fields.mtx,
				items:         tt.fields.items,
				itemNameIndex: tt.fields.itemNameIndex,
				currentId:     tt.fields.currentId,
			}
			if err := s.DeleteItem(tt.args.ctx, tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("DeleteItem() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_itemsrvc_GetItem(t *testing.T) {
	type fields struct {
		logger        *log.Logger
		mtx           sync.Mutex
		items         map[int]*item.Item
		itemNameIndex map[string]*int
		currentId     int
	}
	type args struct {
		ctx context.Context
		p   *item.ItemID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *item.Item
		wantErr bool
	}{
		{
			name: "Normal",
			fields: fields{
				logger: logger,
				items: map[int]*item.Item{
					0: {
						ID:          0,
						Name:        "Bag of Holding",
						Description: "This bag has an interior space considerably larger than its outside dimensions.",
						Damage:      0,
						Healing:     0,
						Protection:  0,
					},
					1: {
						ID:          1,
						Name:        "Greatsword",
						Description: "A very large sword.",
						Damage:      8,
						Healing:     0,
						Protection:  0,
					}},
				itemNameIndex: map[string]*int{
					"bag of holding": intPtr(0),
					"greatsword":     intPtr(1),
				},
				currentId: 1,
			},
			args: args{nil, &item.ItemID{
				ID: 1,
			}},
			want: &item.Item{
				ID:          1,
				Name:        "Greatsword",
				Description: "A very large sword.",
				Damage:      8,
				Healing:     0,
				Protection:  0,
			},
			wantErr: false,
		},
		{
			name: "Unknown_ID",
			fields: fields{
				logger: logger,
				items: map[int]*item.Item{
					0: {
						ID:          0,
						Name:        "Bag of Holding",
						Description: "This bag has an interior space considerably larger than its outside dimensions.",
						Damage:      0,
						Healing:     0,
						Protection:  0,
					},
					1: {
						ID:          1,
						Name:        "Greatsword",
						Description: "A very large sword.",
						Damage:      8,
						Healing:     0,
						Protection:  0,
					}},
				itemNameIndex: map[string]*int{
					"bag of holding": intPtr(0),
					"greatsword":     intPtr(1),
				},
				currentId: 1,
			},
			args: args{nil, &item.ItemID{
				ID: 23,
			}},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &itemsrvc{
				logger:        tt.fields.logger,
				mtx:           tt.fields.mtx,
				items:         tt.fields.items,
				itemNameIndex: tt.fields.itemNameIndex,
				currentId:     tt.fields.currentId,
			}
			got, err := s.GetItem(tt.args.ctx, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetItem() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_itemsrvc_GetItems(t *testing.T) {
	type fields struct {
		logger        *log.Logger
		mtx           sync.Mutex
		items         map[int]*item.Item
		itemNameIndex map[string]*int
		currentId     int
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*item.Item
		wantErr bool
	}{
		{
			name: "Normal",
			fields: fields{
				logger: logger,
				items: map[int]*item.Item{
					0: {
						ID:          0,
						Name:        "Bag of Holding",
						Description: "This bag has an interior space considerably larger than its outside dimensions.",
						Damage:      0,
						Healing:     0,
						Protection:  0,
					},
					1: {
						ID:          1,
						Name:        "Greatsword",
						Description: "A very large sword.",
						Damage:      8,
						Healing:     0,
						Protection:  0,
					}},
				itemNameIndex: map[string]*int{
					"bag of holding": intPtr(0),
					"greatsword":     intPtr(1),
				},
				currentId: 1,
			},
			args: args{nil},
			want: []*item.Item{
				{
					ID:          0,
					Name:        "Bag of Holding",
					Description: "This bag has an interior space considerably larger than its outside dimensions.",
					Damage:      0,
					Healing:     0,
					Protection:  0,
				},
				{
					ID:          1,
					Name:        "Greatsword",
					Description: "A very large sword.",
					Damage:      8,
					Healing:     0,
					Protection:  0,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &itemsrvc{
				logger:        tt.fields.logger,
				mtx:           tt.fields.mtx,
				items:         tt.fields.items,
				itemNameIndex: tt.fields.itemNameIndex,
				currentId:     tt.fields.currentId,
			}
			got, err := s.GetItems(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetItems() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetItems() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_itemsrvc_UpdateItem(t *testing.T) {
	type fields struct {
		logger        *log.Logger
		mtx           sync.Mutex
		items         map[int]*item.Item
		itemNameIndex map[string]*int
		currentId     int
	}
	type args struct {
		ctx        context.Context
		updateItem *item.Item
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *item.Item
		wantErr bool
	}{
		{
			name: "Normal",
			fields: fields{
				logger: logger,
				items: map[int]*item.Item{
					0: {
						ID:          0,
						Name:        "Bag of Holding",
						Description: "This bag has an interior space considerably larger than its outside dimensions.",
						Damage:      0,
						Healing:     0,
						Protection:  0,
					},
					1: {
						ID:          1,
						Name:        "Greatsword",
						Description: "A very large sword.",
						Damage:      8,
						Healing:     0,
						Protection:  0,
					}},
				itemNameIndex: map[string]*int{
					"bag of holding": intPtr(0),
					"greatsword":     intPtr(1),
				},
				currentId: 1,
			},
			args: args{nil, &item.Item{
				ID:          1,
				Name:        "Greatsword",
				Description: "A very large 2 hand sword.",
				Damage:      8,
				Healing:     0,
				Protection:  0,
			}},
			want: &item.Item{
				ID:          1,
				Name:        "Greatsword",
				Description: "A very large 2 hand sword.",
				Damage:      8,
				Healing:     0,
				Protection:  0,
			},
			wantErr: false,
		},
		{
			name: "Normal_Name_Change",
			fields: fields{
				logger: logger,
				items: map[int]*item.Item{
					0: {
						ID:          0,
						Name:        "Bag of Holding",
						Description: "This bag has an interior space considerably larger than its outside dimensions.",
						Damage:      0,
						Healing:     0,
						Protection:  0,
					},
					1: {
						ID:          1,
						Name:        "Greatsword",
						Description: "A very large sword.",
						Damage:      8,
						Healing:     0,
						Protection:  0,
					}},
				itemNameIndex: map[string]*int{
					"bag of holding": intPtr(0),
					"greatsword":     intPtr(1),
				},
				currentId: 1,
			},
			args: args{nil, &item.Item{
				ID:          1,
				Name:        "Greatsword +2",
				Description: "A very large sword.",
				Damage:      10,
				Healing:     0,
				Protection:  0,
			}},
			want: &item.Item{
				ID:          1,
				Name:        "Greatsword +2",
				Description: "A very large sword.",
				Damage:      10,
				Healing:     0,
				Protection:  0,
			},
			wantErr: false,
		},
		{
			name: "Duplicate_Name",
			fields: fields{
				logger: logger,
				items: map[int]*item.Item{
					0: {
						ID:          0,
						Name:        "Bag of Holding",
						Description: "This bag has an interior space considerably larger than its outside dimensions.",
						Damage:      0,
						Healing:     0,
						Protection:  0,
					},
					1: {
						ID:          1,
						Name:        "Greatsword",
						Description: "A very large sword.",
						Damage:      8,
						Healing:     0,
						Protection:  0,
					}},
				itemNameIndex: map[string]*int{
					"bag of holding": intPtr(0),
					"greatsword":     intPtr(1),
				},
				currentId: 1,
			},
			args: args{nil, &item.Item{
				ID:          0,
				Name:        "GREATSWORD",
				Description: "This bag has an interior space considerably larger than its outside dimensions.",
				Damage:      0,
				Healing:     0,
				Protection:  0,
			}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &itemsrvc{
				logger:        tt.fields.logger,
				mtx:           tt.fields.mtx,
				items:         tt.fields.items,
				itemNameIndex: tt.fields.itemNameIndex,
				currentId:     tt.fields.currentId,
			}
			got, err := s.UpdateItem(tt.args.ctx, tt.args.updateItem)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateItem() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateItem() got = %v, want %v", got, tt.want)
			}
		})
	}
}
