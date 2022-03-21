package model

import (
	"context"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func (c *Community) GetByID(ctx context.Context, id primitive.ObjectID) error {
	err := mgm.Coll(c).FindByIDWithCtx(ctx, id, c)
	return err
}

func (c *Community) CollectionName() string {
	return "communities"
}

func ListCommunities(ctx context.Context, skip, limit int64, filter interface{}, order interface{}) ([]*Community, error) {
	var cs []*Community
	opt := &options.FindOptions{}
	opt.SetSkip(skip)
	opt.SetLimit(limit)
	opt.SetSort(order)
	err := mgm.Coll(&Community{}).SimpleFindWithCtx(ctx, &cs, filter, opt)
	return cs, err
}

func CountCommunities(filter interface{}) (count int64, err error) {
	return mgm.Coll(&Community{}).CountDocuments(mgm.Ctx(), filter)
}