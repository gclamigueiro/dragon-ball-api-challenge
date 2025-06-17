package character

type Character struct {
	ID   int    `gorm:"primaryKey;not null" json:"id"`         // Required by DB
	Name string `gorm:"not null;check:name <> ''" json:"name"` // Required and non-empty string
	Ki   string `json:"ki"`
	Race string `json:"race"`
}

func (c *Character) IsValid() bool {
	return c.ID != 0 && c.Name != ""
}
