package design

import (
	. "goa.design/goa/v3/dsl"
)

var Character = Type("Character", func() {
	Description("Character and attributes")
	Field(1, "id", Int, "Character identifier")
	Field(2, "name", String, "Character name")
	Field(3, "description", String, "Character description")
	Field(4, "health", UInt, "Character health")
	Field(5, "experience", UInt, "Character experience")
	Required("id", "name", "description", "health", "experience")
})

var CharacterNew = Type("CharacterNew", func() {
	Description("Character and attributes payload")
	Field(1, "name", String, "Character name")
	Field(2, "description", String, "Character description")
	Field(3, "health", UInt, "Character health")
	Field(4, "experience", UInt, "Character experience")
	Required("name", "description", "health", "experience")
})

var CharacterID = Type("CharacterID", func() {
	Description("Character identifier")
	Field(1, "id", Int, "Character identifier")
	Required("id")
})

var Item = Type("Item", func() {
	Description("Item and attributes")
	Field(1, "id", Int, "Item ID")
	Field(2, "name", String, "Item name")
	Field(3, "description", String, "Item description")
	Field(4, "damage", UInt, "Item damage")
	Field(5, "healing", UInt, "Item healing")
	Field(6, "protection", Int, "Item protection")
	Required("id", "name", "description", "damage", "healing", "protection")
})

var ItemNew = Type("ItemNew", func() {
	Description("Item and attributes payload")
	Field(1, "name", String, "Item name")
	Field(2, "description", String, "Item description")
	Field(3, "damage", UInt, "Item damage")
	Field(4, "healing", UInt, "Item healing")
	Field(5, "protection", Int, "Item protection")
	Required("name", "description", "damage", "healing", "protection")
})

var ItemID = Type("ItemID", func() {
	Description("Item identifier")
	Field(1, "id", Int, "Item ID")
	Required("id")
})

var InventoryEntry = Type("InventoryEntry", func() {
	Description("Character inventory entry")
	Field(1, "item", Int, "Item ID")
	Field(2, "quantity", UInt, "Quantity of the item")
	Required("item", "quantity")
})

var InventoryUpdate = Type("InventoryUpdate", func() {
	Description("Inventory entry")
	Field(1, "id", Int, "Character ID")
	Field(2, "item", Int, "Item ID")
	Field(3, "quantity", UInt, "Quantity of the item")
	Required("id", "item", "quantity")
})

var _ = API("characterapi", func() {
	Title("Character API")
	Description("HTTP service for managing characters.")
	Error("IdNotFound")
	Error("BadData")
	Error("IDsExhausted")
	HTTP(func() {
		Response("IdNotFound", StatusNotFound)
		Response("BadData", StatusBadRequest)
		Response("IDsExhausted", StatusInsufficientStorage)
	})
	Server("gateway", func() {
		Services("character", "item", "inventory")
		Host("localhost", func() {
			URI("http://localhost:8000")
		})
	})
})

var _ = Service("character", func() {
	Description("Character service.")
	Error("IdNotFound")
	Error("BadData")
	Error("IDsExhausted")
	GRPC(func() {
		Response("IdNotFound", CodeNotFound)
		Response("BadData", CodeInvalidArgument)
		Response("IDsExhausted", CodeResourceExhausted)
	})
	Method("getCharacters", func() {
		Result(ArrayOf(Character))
		HTTP(func() {
			GET("/characters")
		})
	})

	Method("getCharacter", func() {
		Payload(CharacterID)
		Result(Character)
		Error("IdNotFound")
		HTTP(func() {
			GET("/characters/{id}")
		})
	})

	Method("createCharacter", func() {
		Payload(CharacterNew)
		Result(Character)
		Error("BadData")
		Error("IDsExhausted")
		HTTP(func() {
			POST("/characters")
		})
	})

	Method("updateCharacter", func() {
		Payload(Character)
		Result(Character)
		Error("IdNotFound")
		Error("BadData")
		HTTP(func() {
			PUT("/characters/{id}")
		})
	})

	Method("deleteCharacter", func() {
		Payload(CharacterID)
		Result(Empty)
		Error("IdNotFound")
		HTTP(func() {
			DELETE("/characters/{id}")
		})
	})
})

var _ = Service("item", func() {
	Description("Item service.")
	Error("IdNotFound")
	Error("BadData")
	Error("IDsExhausted")
	GRPC(func() {
		Response("IdNotFound", CodeNotFound)
		Response("BadData", CodeInvalidArgument)
		Response("IDsExhausted", CodeResourceExhausted)
	})
	Method("getItems", func() {
		Result(ArrayOf(Item))
		HTTP(func() {
			GET("/items")
		})
	})

	Method("getItem", func() {
		Payload(ItemID)
		Result(Item)
		Error("IdNotFound")
		HTTP(func() {
			GET("/items/{id}")
		})
	})

	Method("createItem", func() {
		Payload(ItemNew)
		Result(Item)
		Error("BadData")
		Error("IDsExhausted")
		HTTP(func() {
			POST("/items")
		})
	})

	Method("updateItem", func() {
		Payload(Item)
		Result(Item)
		Error("IdNotFound")
		Error("BadData")
		HTTP(func() {
			PUT("/items/{id}")
		})
	})

	Method("deleteItem", func() {
		Payload(ItemID)
		Result(Empty)
		Error("IdNotFound")
		HTTP(func() {
			DELETE("/items/{id}")
		})
	})
})

var _ = Service("inventory", func() {
	Description("Inventory service.")
	Error("IdNotFound")
	Error("BadData")
	Error("IDsExhausted")
	GRPC(func() {
		Response("IdNotFound", CodeNotFound)
		Response("BadData", CodeInvalidArgument)
		Response("IDsExhausted", CodeResourceExhausted)
	})
	Method("getInventory", func() {
		Payload(CharacterID)
		Result(ArrayOf(InventoryEntry))
		HTTP(func() {
			GET("/character/{id}/inventory")
		})
	})

	Method("updateInventory", func() {
		Payload(InventoryUpdate)
		Result(Empty)
		Error("IdNotFound")
		Error("BadData")
		HTTP(func() {
			PATCH("/character/{id}/inventory")
		})
	})

	Method("clearInventory", func() {
		Payload(CharacterID)
		Result(Empty)
		Error("IdNotFound")
		HTTP(func() {
			DELETE("/character/{id}/inventory")
		})
	})
})
