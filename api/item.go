package api

type SocketGroup struct {
	Id        int    `json:"group"`
	Attribute string `json:"attr"`
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
	QuestItemFrameType
)

type Item struct {
	IsVerified          bool           `json:"verified"`
	Width               int            `json:"w"`
	Height              int            `json:"h"`
	ItemLevel           int            `json:"ilvl"`
	Icon                string         `json:"icon"`
	League              string         `json:"league"`
	Id                  string         `json:"id"`
	Sockets             []SocketGroup  `json:"sockets"`
	Name                string         `json:"name"`
	Type                string         `json:"typeLine"`
	IsIdentified        bool           `json:"identified"`
	IsCorrupted         bool           `json:"corrupted"`
	IsLockedToCharacter bool           `json:"lockedToCharacter"`
	Note                string         `json:"note"`
	Properties          []ItemProperty `json:"properties"`
	Requirements        []ItemProperty `json:"requirements"`
	ExplicitMods        []string       `json:"explicitMods"`
	FlavorText          []string       `json:"flavourText"`
	FrameType           FrameType      `json:"frameType"`
	StackSize           int            `json:"stackSize"`
	MaxStackSize        int            `json:"maxStackSize"`
	X                   int            `json:"x"`
	Y                   int            `json:"y"`
	InventoryId         string         `json:"inventoryId"`
	SocketedItems       []Item         `json:"socketedItems"`
}
