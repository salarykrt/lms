package masters

import (
	"gorm.io/gorm"
)

type MasterDocument struct {
	gorm.Model
	Heading         string `gorm:"NOT NULL"`
	DocsType        string `gorm:"NOT NULL;comment:DOCS TYPE"`        // DOCS TYPE
	DocsSubType     string `gorm:"NOT NULL;comment:DOCS SUB TYPE"`    // DOCS SUB TYPE
	DocsRequired    uint8  `gorm:"NOT NULL;comment:1=>DOCS REQUIRED"` // DOCS REQUIRED
	DocumentActive  uint8  `gorm:"default:1;NOT NULL"`
	DocumentDeleted uint8  `gorm:"default:0;NOT NULL"`
}
