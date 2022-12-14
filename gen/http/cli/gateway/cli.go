// Code generated by goa v3.9.1, DO NOT EDIT.
//
// gateway HTTP client CLI support package
//
// Command:
// $ goa gen github.com/austinhuck/characters-go/design

package cli

import (
	"flag"
	"fmt"
	"net/http"
	"os"

	characterc "github.com/austinhuck/characters-go/gen/http/character/client"
	inventoryc "github.com/austinhuck/characters-go/gen/http/inventory/client"
	itemc "github.com/austinhuck/characters-go/gen/http/item/client"
	goahttp "goa.design/goa/v3/http"
	goa "goa.design/goa/v3/pkg"
)

// UsageCommands returns the set of commands and sub-commands using the format
//
//	command (subcommand1|subcommand2|...)
func UsageCommands() string {
	return `character (get-characters|get-character|create-character|update-character|delete-character)
item (get-items|get-item|create-item|update-item|delete-item)
inventory (get-inventory|update-inventory|clear-inventory)
`
}

// UsageExamples produces an example of a valid invocation of the CLI tool.
func UsageExamples() string {
	return os.Args[0] + ` character get-characters` + "\n" +
		os.Args[0] + ` item get-items` + "\n" +
		os.Args[0] + ` inventory get-inventory --id 4162138413151734252` + "\n" +
		""
}

// ParseEndpoint returns the endpoint and payload as specified on the command
// line.
func ParseEndpoint(
	scheme, host string,
	doer goahttp.Doer,
	enc func(*http.Request) goahttp.Encoder,
	dec func(*http.Response) goahttp.Decoder,
	restore bool,
) (goa.Endpoint, interface{}, error) {
	var (
		characterFlags = flag.NewFlagSet("character", flag.ContinueOnError)

		characterGetCharactersFlags = flag.NewFlagSet("get-characters", flag.ExitOnError)

		characterGetCharacterFlags  = flag.NewFlagSet("get-character", flag.ExitOnError)
		characterGetCharacterIDFlag = characterGetCharacterFlags.String("id", "REQUIRED", "Character identifier")

		characterCreateCharacterFlags    = flag.NewFlagSet("create-character", flag.ExitOnError)
		characterCreateCharacterBodyFlag = characterCreateCharacterFlags.String("body", "REQUIRED", "")

		characterUpdateCharacterFlags    = flag.NewFlagSet("update-character", flag.ExitOnError)
		characterUpdateCharacterBodyFlag = characterUpdateCharacterFlags.String("body", "REQUIRED", "")
		characterUpdateCharacterIDFlag   = characterUpdateCharacterFlags.String("id", "REQUIRED", "Character identifier")

		characterDeleteCharacterFlags  = flag.NewFlagSet("delete-character", flag.ExitOnError)
		characterDeleteCharacterIDFlag = characterDeleteCharacterFlags.String("id", "REQUIRED", "Character identifier")

		itemFlags = flag.NewFlagSet("item", flag.ContinueOnError)

		itemGetItemsFlags = flag.NewFlagSet("get-items", flag.ExitOnError)

		itemGetItemFlags  = flag.NewFlagSet("get-item", flag.ExitOnError)
		itemGetItemIDFlag = itemGetItemFlags.String("id", "REQUIRED", "Item ID")

		itemCreateItemFlags    = flag.NewFlagSet("create-item", flag.ExitOnError)
		itemCreateItemBodyFlag = itemCreateItemFlags.String("body", "REQUIRED", "")

		itemUpdateItemFlags    = flag.NewFlagSet("update-item", flag.ExitOnError)
		itemUpdateItemBodyFlag = itemUpdateItemFlags.String("body", "REQUIRED", "")
		itemUpdateItemIDFlag   = itemUpdateItemFlags.String("id", "REQUIRED", "Item ID")

		itemDeleteItemFlags  = flag.NewFlagSet("delete-item", flag.ExitOnError)
		itemDeleteItemIDFlag = itemDeleteItemFlags.String("id", "REQUIRED", "Item ID")

		inventoryFlags = flag.NewFlagSet("inventory", flag.ContinueOnError)

		inventoryGetInventoryFlags  = flag.NewFlagSet("get-inventory", flag.ExitOnError)
		inventoryGetInventoryIDFlag = inventoryGetInventoryFlags.String("id", "REQUIRED", "Character identifier")

		inventoryUpdateInventoryFlags    = flag.NewFlagSet("update-inventory", flag.ExitOnError)
		inventoryUpdateInventoryBodyFlag = inventoryUpdateInventoryFlags.String("body", "REQUIRED", "")
		inventoryUpdateInventoryIDFlag   = inventoryUpdateInventoryFlags.String("id", "REQUIRED", "Character ID")

		inventoryClearInventoryFlags  = flag.NewFlagSet("clear-inventory", flag.ExitOnError)
		inventoryClearInventoryIDFlag = inventoryClearInventoryFlags.String("id", "REQUIRED", "Character identifier")
	)
	characterFlags.Usage = characterUsage
	characterGetCharactersFlags.Usage = characterGetCharactersUsage
	characterGetCharacterFlags.Usage = characterGetCharacterUsage
	characterCreateCharacterFlags.Usage = characterCreateCharacterUsage
	characterUpdateCharacterFlags.Usage = characterUpdateCharacterUsage
	characterDeleteCharacterFlags.Usage = characterDeleteCharacterUsage

	itemFlags.Usage = itemUsage
	itemGetItemsFlags.Usage = itemGetItemsUsage
	itemGetItemFlags.Usage = itemGetItemUsage
	itemCreateItemFlags.Usage = itemCreateItemUsage
	itemUpdateItemFlags.Usage = itemUpdateItemUsage
	itemDeleteItemFlags.Usage = itemDeleteItemUsage

	inventoryFlags.Usage = inventoryUsage
	inventoryGetInventoryFlags.Usage = inventoryGetInventoryUsage
	inventoryUpdateInventoryFlags.Usage = inventoryUpdateInventoryUsage
	inventoryClearInventoryFlags.Usage = inventoryClearInventoryUsage

	if err := flag.CommandLine.Parse(os.Args[1:]); err != nil {
		return nil, nil, err
	}

	if flag.NArg() < 2 { // two non flag args are required: SERVICE and ENDPOINT (aka COMMAND)
		return nil, nil, fmt.Errorf("not enough arguments")
	}

	var (
		svcn string
		svcf *flag.FlagSet
	)
	{
		svcn = flag.Arg(0)
		switch svcn {
		case "character":
			svcf = characterFlags
		case "item":
			svcf = itemFlags
		case "inventory":
			svcf = inventoryFlags
		default:
			return nil, nil, fmt.Errorf("unknown service %q", svcn)
		}
	}
	if err := svcf.Parse(flag.Args()[1:]); err != nil {
		return nil, nil, err
	}

	var (
		epn string
		epf *flag.FlagSet
	)
	{
		epn = svcf.Arg(0)
		switch svcn {
		case "character":
			switch epn {
			case "get-characters":
				epf = characterGetCharactersFlags

			case "get-character":
				epf = characterGetCharacterFlags

			case "create-character":
				epf = characterCreateCharacterFlags

			case "update-character":
				epf = characterUpdateCharacterFlags

			case "delete-character":
				epf = characterDeleteCharacterFlags

			}

		case "item":
			switch epn {
			case "get-items":
				epf = itemGetItemsFlags

			case "get-item":
				epf = itemGetItemFlags

			case "create-item":
				epf = itemCreateItemFlags

			case "update-item":
				epf = itemUpdateItemFlags

			case "delete-item":
				epf = itemDeleteItemFlags

			}

		case "inventory":
			switch epn {
			case "get-inventory":
				epf = inventoryGetInventoryFlags

			case "update-inventory":
				epf = inventoryUpdateInventoryFlags

			case "clear-inventory":
				epf = inventoryClearInventoryFlags

			}

		}
	}
	if epf == nil {
		return nil, nil, fmt.Errorf("unknown %q endpoint %q", svcn, epn)
	}

	// Parse endpoint flags if any
	if svcf.NArg() > 1 {
		if err := epf.Parse(svcf.Args()[1:]); err != nil {
			return nil, nil, err
		}
	}

	var (
		data     interface{}
		endpoint goa.Endpoint
		err      error
	)
	{
		switch svcn {
		case "character":
			c := characterc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "get-characters":
				endpoint = c.GetCharacters()
				data = nil
			case "get-character":
				endpoint = c.GetCharacter()
				data, err = characterc.BuildGetCharacterPayload(*characterGetCharacterIDFlag)
			case "create-character":
				endpoint = c.CreateCharacter()
				data, err = characterc.BuildCreateCharacterPayload(*characterCreateCharacterBodyFlag)
			case "update-character":
				endpoint = c.UpdateCharacter()
				data, err = characterc.BuildUpdateCharacterPayload(*characterUpdateCharacterBodyFlag, *characterUpdateCharacterIDFlag)
			case "delete-character":
				endpoint = c.DeleteCharacter()
				data, err = characterc.BuildDeleteCharacterPayload(*characterDeleteCharacterIDFlag)
			}
		case "item":
			c := itemc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "get-items":
				endpoint = c.GetItems()
				data = nil
			case "get-item":
				endpoint = c.GetItem()
				data, err = itemc.BuildGetItemPayload(*itemGetItemIDFlag)
			case "create-item":
				endpoint = c.CreateItem()
				data, err = itemc.BuildCreateItemPayload(*itemCreateItemBodyFlag)
			case "update-item":
				endpoint = c.UpdateItem()
				data, err = itemc.BuildUpdateItemPayload(*itemUpdateItemBodyFlag, *itemUpdateItemIDFlag)
			case "delete-item":
				endpoint = c.DeleteItem()
				data, err = itemc.BuildDeleteItemPayload(*itemDeleteItemIDFlag)
			}
		case "inventory":
			c := inventoryc.NewClient(scheme, host, doer, enc, dec, restore)
			switch epn {
			case "get-inventory":
				endpoint = c.GetInventory()
				data, err = inventoryc.BuildGetInventoryPayload(*inventoryGetInventoryIDFlag)
			case "update-inventory":
				endpoint = c.UpdateInventory()
				data, err = inventoryc.BuildUpdateInventoryPayload(*inventoryUpdateInventoryBodyFlag, *inventoryUpdateInventoryIDFlag)
			case "clear-inventory":
				endpoint = c.ClearInventory()
				data, err = inventoryc.BuildClearInventoryPayload(*inventoryClearInventoryIDFlag)
			}
		}
	}
	if err != nil {
		return nil, nil, err
	}

	return endpoint, data, nil
}

// characterUsage displays the usage of the character command and its
// subcommands.
func characterUsage() {
	fmt.Fprintf(os.Stderr, `Character service.
Usage:
    %[1]s [globalflags] character COMMAND [flags]

COMMAND:
    get-characters: GetCharacters implements getCharacters.
    get-character: GetCharacter implements getCharacter.
    create-character: CreateCharacter implements createCharacter.
    update-character: UpdateCharacter implements updateCharacter.
    delete-character: DeleteCharacter implements deleteCharacter.

Additional help:
    %[1]s character COMMAND --help
`, os.Args[0])
}
func characterGetCharactersUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] character get-characters

GetCharacters implements getCharacters.

Example:
    %[1]s character get-characters
`, os.Args[0])
}

func characterGetCharacterUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] character get-character -id INT

GetCharacter implements getCharacter.
    -id INT: Character identifier

Example:
    %[1]s character get-character --id 7427596992036772911
`, os.Args[0])
}

func characterCreateCharacterUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] character create-character -body JSON

CreateCharacter implements createCharacter.
    -body JSON: 

Example:
    %[1]s character create-character --body '{
      "description": "Maxime non error dignissimos.",
      "experience": 11895619109444395021,
      "health": 13825686865421341647,
      "name": "Voluptates accusamus voluptas."
   }'
`, os.Args[0])
}

func characterUpdateCharacterUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] character update-character -body JSON -id INT

UpdateCharacter implements updateCharacter.
    -body JSON: 
    -id INT: Character identifier

Example:
    %[1]s character update-character --body '{
      "description": "Qui ut.",
      "experience": 6143131330256813255,
      "health": 16114915148797695919,
      "name": "Reiciendis eius a totam."
   }' --id 4588229449694285417
`, os.Args[0])
}

func characterDeleteCharacterUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] character delete-character -id INT

DeleteCharacter implements deleteCharacter.
    -id INT: Character identifier

Example:
    %[1]s character delete-character --id 9095054343805998409
`, os.Args[0])
}

// itemUsage displays the usage of the item command and its subcommands.
func itemUsage() {
	fmt.Fprintf(os.Stderr, `Item service.
Usage:
    %[1]s [globalflags] item COMMAND [flags]

COMMAND:
    get-items: GetItems implements getItems.
    get-item: GetItem implements getItem.
    create-item: CreateItem implements createItem.
    update-item: UpdateItem implements updateItem.
    delete-item: DeleteItem implements deleteItem.

Additional help:
    %[1]s item COMMAND --help
`, os.Args[0])
}
func itemGetItemsUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] item get-items

GetItems implements getItems.

Example:
    %[1]s item get-items
`, os.Args[0])
}

func itemGetItemUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] item get-item -id INT

GetItem implements getItem.
    -id INT: Item ID

Example:
    %[1]s item get-item --id 7615456859102312540
`, os.Args[0])
}

func itemCreateItemUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] item create-item -body JSON

CreateItem implements createItem.
    -body JSON: 

Example:
    %[1]s item create-item --body '{
      "damage": 605133029443553342,
      "description": "Quae ea velit iste deserunt occaecati.",
      "healing": 9434446331462297666,
      "name": "Iste laborum ut ullam perferendis.",
      "protection": 6811333696036567306
   }'
`, os.Args[0])
}

func itemUpdateItemUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] item update-item -body JSON -id INT

UpdateItem implements updateItem.
    -body JSON: 
    -id INT: Item ID

Example:
    %[1]s item update-item --body '{
      "damage": 8836063704909434334,
      "description": "Illo ipsam voluptate magni minima.",
      "healing": 7974778561697912772,
      "name": "Sit id modi ut dolor ipsum.",
      "protection": 267106854576953551
   }' --id 3565726362524344167
`, os.Args[0])
}

func itemDeleteItemUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] item delete-item -id INT

DeleteItem implements deleteItem.
    -id INT: Item ID

Example:
    %[1]s item delete-item --id 1473388180512697621
`, os.Args[0])
}

// inventoryUsage displays the usage of the inventory command and its
// subcommands.
func inventoryUsage() {
	fmt.Fprintf(os.Stderr, `Inventory service.
Usage:
    %[1]s [globalflags] inventory COMMAND [flags]

COMMAND:
    get-inventory: GetInventory implements getInventory.
    update-inventory: UpdateInventory implements updateInventory.
    clear-inventory: ClearInventory implements clearInventory.

Additional help:
    %[1]s inventory COMMAND --help
`, os.Args[0])
}
func inventoryGetInventoryUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] inventory get-inventory -id INT

GetInventory implements getInventory.
    -id INT: Character identifier

Example:
    %[1]s inventory get-inventory --id 4162138413151734252
`, os.Args[0])
}

func inventoryUpdateInventoryUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] inventory update-inventory -body JSON -id INT

UpdateInventory implements updateInventory.
    -body JSON: 
    -id INT: Character ID

Example:
    %[1]s inventory update-inventory --body '{
      "item": 1766420538716923054,
      "quantity": 10997400203762602967
   }' --id 3681289626091285004
`, os.Args[0])
}

func inventoryClearInventoryUsage() {
	fmt.Fprintf(os.Stderr, `%[1]s [flags] inventory clear-inventory -id INT

ClearInventory implements clearInventory.
    -id INT: Character identifier

Example:
    %[1]s inventory clear-inventory --id 737855043031604889
`, os.Args[0])
}
