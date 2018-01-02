package api

import (
	"encoding/json"
)

// Since Abyss, this is bool (false)
type SocketAttribute string

func (a *SocketAttribute) UnmarshalJSON(data []byte) error {
	// Try to unmarshall as bool
	var boolValue bool
	err := json.Unmarshal(data, &boolValue)
	if err == nil {
		if boolValue {
			*a = "true"
		} else {
			*a = "false"
		}
		return nil
	}
	// If that fails, try to unmarshall as string
	var stringValue string
	err = json.Unmarshal(data, &stringValue)
	if err != nil {
		*a = SocketAttribute(stringValue)
	}
	return err
}

type Socket struct {
	GroupId   int             `json:"group"`
	Attribute SocketAttribute `json:"attr"`
	Colour    string          `json:"sColour"`
}

type ItemProperty struct {
	Name        string        `json:"name"`
	Values      []interface{} `json:"values"`
	DisplayMode int           `json:"displayMode"`
}

type FrameType int

const (
	NormalItemFrameType FrameType = iota
	MagicItemFrameType
	RareItemFrameType
	UniqueItemFrameType
	GemFrameType
	CurrencyFrameType
	DivinationCardFrameType
	QuestItemFrameType
	ProphecyFrameType
	RelicFrameType
)

type Item struct {
	// Names for some items may include markup. For example: <<set:MS>><<set:M>><<set:S>>Roth's Reach
	Name string `json:"name"`
	Type string `json:"typeLine"`

	Properties   []ItemProperty `json:"properties"`
	Requirements []ItemProperty `json:"requirements"`

	Sockets []Socket `json:"sockets"`

	ExplicitMods []string `json:"explicitMods"`
	ImplicitMods []string `json:"implicitMods"`
	UtilityMods  []string `json:"utilityMods"`
	EnchantMods  []string `json:"enchantMods"`
	CraftedMods  []string `json:"craftedMods"`
	CosmeticMods []string `json:"cosmeticMods"`

	Note string `json:"note"`

	IsVerified             bool      `json:"verified"`
	Width                  int       `json:"w"`
	Height                 int       `json:"h"`
	ItemLevel              int       `json:"ilvl"`
	Icon                   string    `json:"icon"`
	League                 string    `json:"league"`
	Id                     string    `json:"id"`
	IsIdentified           bool      `json:"identified"`
	IsCorrupted            bool      `json:"corrupted"`
	IsLockedToCharacter    bool      `json:"lockedToCharacter"`
	IsSupport              bool      `json:"support"`
	DescriptionText        string    `json:"descrText"`
	SecondDescriptionText  string    `json:"secDescrText"`
	FlavorText             []string  `json:"flavourText"`
	ArtFilename            string    `json:"artFilename"`
	FrameType              FrameType `json:"frameType"`
	StackSize              int       `json:"stackSize"`
	MaxStackSize           int       `json:"maxStackSize"`
	X                      int       `json:"x"`
	Y                      int       `json:"y"`
	InventoryId            string    `json:"inventoryId"`
	SocketedItems          []Item    `json:"socketedItems"`
	IsRelic                bool      `json:"isRelic"`
	TalismanTier           int       `json:"talismanTier"`
	ProphecyText           string    `json:"prophecyText"`
	ProphecyDifficultyText string    `json:"prophecyDiffText"`
}
