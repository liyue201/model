package model

import (
	"context"
	"github.com/Overealityio/overeality-server-model/util"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

type Counter struct {
	mgm.IDField `json:",inline" bson:",inline"`
	Uri         string             `json:"uri,omitempty" bson:"uri"`
	Count       int                `json:"count,omitempty" bson:"count"`
	CreatedAt   *time.Time         `json:"createdAt,omitempty" bson:"createdAt"`
	UpdatedAt   *time.Time         `json:"updatedAt,omitempty" bson:"updatedAt"`
}


func (c *Counter) CollectionName() string {
	return "counters"
}

func (c *Counter) Creating() error {
	c.CreatedAt = util.TimeNow()
	return nil
}

func (c *Counter) Saving() error {
	c.UpdatedAt = util.TimeNow()
	return nil
}

func (c *Counter) Create() (err error) {
	err = mgm.Coll(c).Create(c)
	return err
}

func (c *Counter) Update() error {
	opt := &options.UpdateOptions{}
	opt.SetUpsert(false)
	return mgm.Coll(c).Update(c, opt)
}

func (c *Counter) GetByUri(uri string) error {
	err := mgm.Coll(c).First(bson.M{"uri":uri}, c)
	return err
}

func (c *Counter) GetByIDWithCtx(ctx context.Context, id primitive.ObjectID) error {
	err := mgm.Coll(c).FindByIDWithCtx(ctx, id, c)
	return err
}

func ListCounterWithCtx(ctx context.Context,filter interface{}, opts ...*options.FindOptions) (cs []*Counter, err error) {
	cs = []*Counter{}
	err = mgm.Coll(&Counter{}).SimpleFindWithCtx(ctx, &cs, filter, opts...)
	return
}

func CountCounterWithCtx(ctx context.Context, filter interface{}) (count int64, err error) {
	return mgm.Coll(&Counter{}).CountDocuments(ctx, filter)
}



