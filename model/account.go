package model

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type TokenAccount struct {
	mgm.IDField  `json:",inline" bson:",inline"`
	// mandatory: true. User for this account
	User  primitive.ObjectID `json:"user,omitempty" bson:"user"`
	// mandatory: true. network of this account
	NetworkId  primitive.ObjectID `json:"networkId,omitempty" bson:"networkId"`
	// mandatory: true. corresponding token id stored in mongo.
	Token primitive.ObjectID `json:"token,omitempty" bson:"token"`
	// mandatory: true. Token name
	Name         string               `json:"name,omitempty" bson:"name"`
	// mandatory: true. Token Standard
	Standard     TokenStandardEnum    `json:"standard,omitempty" bson:"standard"`
	// mandatory: true. balance for this account
	Balance primitive.Decimal128 `json:"balance,omitempty" bson:"balance"`
	// mandatory: true. create time of this token.
	CreatedAt   *time.Time         `json:"createdAt,omitempty" bson:"createdAt"`
	// mandatory: true. update time of this token.
	UpdatedAt   *time.Time         `json:"updatedAt,omitempty" bson:"updatedAt"`
}
