package model

import (
	"fmt"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Community struct {
	mgm.IDField  `json:",inline" bson:",inline"`
	Name         string               `json:"name,omitempty" bson:"name"`
	IntroContent string               `json:"introContent,omitempty" bson:"introContent"`
	IntroImages  []Image              `json:"introImages,omitempty" bson:"introImages"`
	Influencers  []primitive.ObjectID `json:"influencers,omitempty" bson:"influencers"`
	CreatedAt    *time.Time           `json:"createdAt,omitempty" bson:"createdAt"`
	UpdatedAt    *time.Time           `json:"lastUpdatedAt,omitempty" bson:"last_updated_at"`
}

type Image struct {
	Uri string `json:"uri,omitempty" bson:"uri"`
}

func (c *Community) GetByID(id primitive.ObjectID) error {
	err := mgm.Coll(c).FindByID(id, c)
	if err != nil {
		return fmt.Errorf("get err while query Community %s", err.Error())
	}
	return nil
}

func (c *Community) CollectionName() string {
	return "communities"
}