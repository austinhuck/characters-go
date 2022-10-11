package character

import (
	"context"
	"github.com/austinhuck/characters-go/gen/character"
	"log"
	"math"
	"os"
	"reflect"
	"sync"
	"testing"
)

var logger = log.New(os.Stderr, "[character_test] ", log.Ltime)

func intPtr(literal int) *int {
	return &literal
}

func Test_charactersrvc_CreateCharacter(t *testing.T) {
	type fields struct {
		logger             *log.Logger
		mtx                sync.Mutex
		characters         map[int]*character.Character
		characterNameIndex map[string]*int
		currentId          int
	}
	type args struct {
		ctx          context.Context
		newCharacter *character.CharacterNew
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *character.Character
		wantErr bool
	}{
		{
			name: "Normal",
			fields: fields{
				logger: logger,
				characters: map[int]*character.Character{
					0: {
						ID:          0,
						Name:        "Grog",
						Description: "Goliath barbarian",
						Health:      92,
						Experience:  4100,
					},
					1: {
						ID:          1,
						Name:        "Fjord",
						Description: "Half-orc warlock",
						Health:      86,
						Experience:  3400,
					}},
				characterNameIndex: map[string]*int{
					"grog":  intPtr(0),
					"fjord": intPtr(1),
				},
				currentId: 1,
			},
			args: args{nil, &character.CharacterNew{
				Name:        "Ildrex",
				Description: "Dragonborn druid",
				Health:      75,
				Experience:  2400,
			}},
			want: &character.Character{
				ID:          2,
				Name:        "Ildrex",
				Description: "Dragonborn druid",
				Health:      75,
				Experience:  2400,
			},
			wantErr: false,
		},
		{
			name: "Duplicate_Name",
			fields: fields{
				logger: logger,
				characters: map[int]*character.Character{
					0: {
						ID:          0,
						Name:        "Grog",
						Description: "Goliath barbarian",
						Health:      92,
						Experience:  4100,
					},
					1: {
						ID:          1,
						Name:        "Fjord",
						Description: "Half-orc warlock",
						Health:      86,
						Experience:  3400,
					}},
				characterNameIndex: map[string]*int{
					"grog":  intPtr(0),
					"fjord": intPtr(1),
				},
				currentId: 1,
			},
			args: args{nil, &character.CharacterNew{
				Name:        "Grog",
				Description: "Dragonborn druid",
				Health:      75,
				Experience:  2400,
			}},
			want:    nil,
			wantErr: true,
		},
		{
			name: "IDs_Max",
			fields: fields{
				logger: logger,
				characters: map[int]*character.Character{
					math.MaxInt - 1: {
						ID:          0,
						Name:        "Grog",
						Description: "Goliath barbarian",
						Health:      92,
						Experience:  4100,
					},
					math.MaxInt: {
						ID:          1,
						Name:        "Fjord",
						Description: "Half-orc warlock",
						Health:      86,
						Experience:  3400,
					}},
				characterNameIndex: map[string]*int{
					"grog":  intPtr(math.MaxInt - 1),
					"fjord": intPtr(math.MaxInt),
				},
				currentId: math.MaxInt,
			},
			args: args{nil, &character.CharacterNew{
				Name:        "Ildrex",
				Description: "Dragonborn druid",
				Health:      75,
				Experience:  2400,
			}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &charactersrvc{
				logger:             tt.fields.logger,
				mtx:                tt.fields.mtx,
				characters:         tt.fields.characters,
				characterNameIndex: tt.fields.characterNameIndex,
				currentId:          tt.fields.currentId,
			}
			got, err := s.CreateCharacter(tt.args.ctx, tt.args.newCharacter)
			if (err != nil) != tt.wantErr {
				t.Errorf("CreateCharacter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CreateCharacter() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_charactersrvc_DeleteCharacter(t *testing.T) {
	type fields struct {
		logger             *log.Logger
		mtx                sync.Mutex
		characters         map[int]*character.Character
		characterNameIndex map[string]*int
		currentId          int
	}
	type args struct {
		ctx context.Context
		p   *character.CharacterID
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
				characters: map[int]*character.Character{
					0: {
						ID:          0,
						Name:        "Grog",
						Description: "Goliath barbarian",
						Health:      92,
						Experience:  4100,
					},
					1: {
						ID:          1,
						Name:        "Fjord",
						Description: "Half-orc warlock",
						Health:      86,
						Experience:  3400,
					}},
				characterNameIndex: map[string]*int{
					"grog":  intPtr(0),
					"fjord": intPtr(1),
				},
				currentId: 1,
			},
			args: args{nil, &character.CharacterID{
				ID: 0,
			}},
			wantErr: false,
		},
		{
			name: "Unknown_ID",
			fields: fields{
				logger: logger,
				characters: map[int]*character.Character{
					0: {
						ID:          0,
						Name:        "Grog",
						Description: "Goliath barbarian",
						Health:      92,
						Experience:  4100,
					},
					1: {
						ID:          1,
						Name:        "Fjord",
						Description: "Half-orc warlock",
						Health:      86,
						Experience:  3400,
					}},
				characterNameIndex: map[string]*int{
					"grog":  intPtr(0),
					"fjord": intPtr(1),
				},
				currentId: 1,
			},
			args: args{nil, &character.CharacterID{
				ID: 23,
			}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &charactersrvc{
				logger:             tt.fields.logger,
				mtx:                tt.fields.mtx,
				characters:         tt.fields.characters,
				characterNameIndex: tt.fields.characterNameIndex,
				currentId:          tt.fields.currentId,
			}
			if err := s.DeleteCharacter(tt.args.ctx, tt.args.p); (err != nil) != tt.wantErr {
				t.Errorf("DeleteCharacter() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_charactersrvc_GetCharacter(t *testing.T) {
	type fields struct {
		logger             *log.Logger
		mtx                sync.Mutex
		characters         map[int]*character.Character
		characterNameIndex map[string]*int
		currentId          int
	}
	type args struct {
		ctx context.Context
		p   *character.CharacterID
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *character.Character
		wantErr bool
	}{
		{
			name: "Normal",
			fields: fields{
				logger: logger,
				characters: map[int]*character.Character{
					0: {
						ID:          0,
						Name:        "Grog",
						Description: "Goliath barbarian",
						Health:      92,
						Experience:  4100,
					},
					1: {
						ID:          1,
						Name:        "Fjord",
						Description: "Half-orc warlock",
						Health:      86,
						Experience:  3400,
					}},
				characterNameIndex: map[string]*int{
					"grog":  intPtr(0),
					"fjord": intPtr(1),
				},
				currentId: 1,
			},
			args: args{nil, &character.CharacterID{
				ID: 1,
			}},
			want: &character.Character{
				ID:          1,
				Name:        "Fjord",
				Description: "Half-orc warlock",
				Health:      86,
				Experience:  3400,
			},
			wantErr: false,
		},
		{
			name: "Unknown_ID",
			fields: fields{
				logger: logger,
				characters: map[int]*character.Character{
					0: {
						ID:          0,
						Name:        "Grog",
						Description: "Goliath barbarian",
						Health:      92,
						Experience:  4100,
					},
					1: {
						ID:          1,
						Name:        "Fjord",
						Description: "Half-orc warlock",
						Health:      86,
						Experience:  3400,
					}},
				characterNameIndex: map[string]*int{
					"grog":  intPtr(0),
					"fjord": intPtr(1),
				},
				currentId: 1,
			},
			args: args{nil, &character.CharacterID{
				ID: 23,
			}},
			want:    nil,
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &charactersrvc{
				logger:             tt.fields.logger,
				mtx:                tt.fields.mtx,
				characters:         tt.fields.characters,
				characterNameIndex: tt.fields.characterNameIndex,
				currentId:          tt.fields.currentId,
			}
			got, err := s.GetCharacter(tt.args.ctx, tt.args.p)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCharacter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCharacter() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_charactersrvc_GetCharacters(t *testing.T) {
	type fields struct {
		logger             *log.Logger
		mtx                sync.Mutex
		characters         map[int]*character.Character
		characterNameIndex map[string]*int
		currentId          int
	}
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    []*character.Character
		wantErr bool
	}{
		{
			name: "Normal",
			fields: fields{
				logger: logger,
				characters: map[int]*character.Character{
					0: {
						ID:          0,
						Name:        "Grog",
						Description: "Goliath barbarian",
						Health:      92,
						Experience:  4100,
					},
					1: {
						ID:          1,
						Name:        "Fjord",
						Description: "Half-orc warlock",
						Health:      86,
						Experience:  3400,
					}},
				characterNameIndex: map[string]*int{
					"grog":  intPtr(0),
					"fjord": intPtr(1),
				},
				currentId: 1,
			},
			args: args{nil},
			want: []*character.Character{
				{
					ID:          0,
					Name:        "Grog",
					Description: "Goliath barbarian",
					Health:      92,
					Experience:  4100,
				},
				{
					ID:          1,
					Name:        "Fjord",
					Description: "Half-orc warlock",
					Health:      86,
					Experience:  3400,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &charactersrvc{
				logger:             tt.fields.logger,
				mtx:                tt.fields.mtx,
				characters:         tt.fields.characters,
				characterNameIndex: tt.fields.characterNameIndex,
				currentId:          tt.fields.currentId,
			}
			got, err := s.GetCharacters(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetCharacters() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetCharacters() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_charactersrvc_UpdateCharacter(t *testing.T) {
	type fields struct {
		logger             *log.Logger
		mtx                sync.Mutex
		characters         map[int]*character.Character
		characterNameIndex map[string]*int
		currentId          int
	}
	type args struct {
		ctx             context.Context
		updateCharacter *character.Character
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *character.Character
		wantErr bool
	}{
		{
			name: "Normal",
			fields: fields{
				logger: logger,
				characters: map[int]*character.Character{
					0: {
						ID:          0,
						Name:        "Grog",
						Description: "Goliath barbarian",
						Health:      92,
						Experience:  4100,
					},
					1: {
						ID:          1,
						Name:        "Fjord",
						Description: "Half-orc warlock",
						Health:      86,
						Experience:  3400,
					}},
				characterNameIndex: map[string]*int{
					"grog":  intPtr(0),
					"fjord": intPtr(1),
				},
				currentId: 1,
			},
			args: args{nil, &character.Character{
				ID:          0,
				Name:        "Grog",
				Description: "Goliath barbarian/fighter",
				Health:      120,
				Experience:  4600,
			}},
			want: &character.Character{
				ID:          0,
				Name:        "Grog",
				Description: "Goliath barbarian/fighter",
				Health:      120,
				Experience:  4600,
			},
			wantErr: false,
		},
		{
			name: "Normal_Name_Change",
			fields: fields{
				logger: logger,
				characters: map[int]*character.Character{
					0: {
						ID:          0,
						Name:        "Grog",
						Description: "Goliath barbarian",
						Health:      92,
						Experience:  4100,
					},
					1: {
						ID:          1,
						Name:        "Fjord",
						Description: "Half-orc warlock",
						Health:      86,
						Experience:  3400,
					}},
				characterNameIndex: map[string]*int{
					"grog":  intPtr(0),
					"fjord": intPtr(1),
				},
				currentId: 1,
			},
			args: args{nil, &character.Character{
				ID:          0,
				Name:        "Grog Jr",
				Description: "Goliath barbarian",
				Health:      92,
				Experience:  4100,
			}},
			want: &character.Character{
				ID:          0,
				Name:        "Grog Jr",
				Description: "Goliath barbarian",
				Health:      92,
				Experience:  4100,
			},
			wantErr: false,
		},
		{
			name: "Duplicate_Name",
			fields: fields{
				logger: logger,
				characters: map[int]*character.Character{
					0: {
						ID:          0,
						Name:        "Grog",
						Description: "Goliath barbarian",
						Health:      92,
						Experience:  4100,
					},
					1: {
						ID:          1,
						Name:        "Fjord",
						Description: "Half-orc warlock",
						Health:      86,
						Experience:  3400,
					}},
				characterNameIndex: map[string]*int{
					"grog":  intPtr(0),
					"fjord": intPtr(1),
				},
				currentId: 1,
			},
			args: args{nil, &character.Character{
				ID:          0,
				Name:        "FJORD",
				Description: "Goliath barbarian",
				Health:      92,
				Experience:  4100,
			}},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &charactersrvc{
				logger:             tt.fields.logger,
				mtx:                tt.fields.mtx,
				characters:         tt.fields.characters,
				characterNameIndex: tt.fields.characterNameIndex,
				currentId:          tt.fields.currentId,
			}
			got, err := s.UpdateCharacter(tt.args.ctx, tt.args.updateCharacter)
			if (err != nil) != tt.wantErr {
				t.Errorf("UpdateCharacter() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("UpdateCharacter() got = %v, want %v", got, tt.want)
			}
		})
	}
}
