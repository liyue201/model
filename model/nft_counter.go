package model

import (
	"time"

	"github.com/kamva/mgm/v3"
)

// NftCounter only works for ERC721
type NftCounter struct {
	mgm.IDField `json:",inline" bson:",inline"`
	// mandatory: true. Uri of ERC721 token nft metadata
	Uri string `json:"uri,omitempty" bson:"uri"`
	// mandatory: true. counter always starts from 1
	Count int `json:"count,omitempty" bson:"count"`
	// mandatory: true.
	CreatedAt *time.Time `json:"createdAt,omitempty" bson:"createdAt"`
	// mandatory: true.
	UpdatedAt *time.Time `json:"updatedAt,omitempty" bson:"updatedAt"`
}
