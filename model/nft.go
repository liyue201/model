package model

import (
	"context"
	"github.com/Overealityio/overeality-server-model/util"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)


type MintReq struct {
	Community   string             `json:"community,omitempty" bson:"-"`
	Title       string             `json:"title,omitempty" bson:"title"`
	SubTitle    string             `json:"subTitle,omitempty" bson:"subTitle"`
	Description string             `json:"description,omitempty" bson:"description"`
	Uri         string             `json:"uri,omitempty" bson:"uri" validate:"required"`
}

type Nft struct {
	mgm.IDField `json:",inline" bson:",inline"`
	MintReq `json:",inline" bson:",inline"`
	CommunityId *primitive.ObjectID `json:"communityId,omitempty" bson:"communityId"`
	Owner       primitive.ObjectID `json:"owner,omitempty" bson:"owner"`
	NumId       int                `json:"numId,omitempty" bson:"numId"`
	MintBy      primitive.ObjectID `json:"mintBy,omitempty" bson:"mintBy"`
	CreatedAt   *time.Time         `json:"createdAt,omitempty" bson:"createdAt"`
	UpdatedAt   *time.Time         `json:"updatedAt,omitempty" bson:"updatedAt"`
}


func (n *Nft) CollectionName() string {
	return "nfts"
}

func (n *Nft) Creating() error {
	n.CreatedAt = util.TimeNow()
	return nil
}

func (n *Nft) Saving() error {
	n.UpdatedAt = util.TimeNow()
	return nil
}

func (n *Nft) Create(ctx context.Context) (err error) {
	err = mgm.Coll(n).CreateWithCtx(ctx, n)
	return err
}

func (n *Nft) UpdateWithCtx(ctx context.Context) error {
	opt := &options.UpdateOptions{}
	opt.SetUpsert(false)
	return mgm.Coll(n).UpdateWithCtx(ctx, n, opt)
}

func (n *Nft) GetByIDWithCtx(ctx context.Context, id primitive.ObjectID) error {
	err := mgm.Coll(n).FindByIDWithCtx(ctx, id, n)
	return err
}

func (n *Nft) FirstWithCtx(ctx context.Context, filter interface{}) error {
	err := mgm.Coll(n).FirstWithCtx(ctx, filter, n)
	return err
}

func ListNftWithCtx(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (ns []*Nft, err error) {
	ns = []*Nft{}
	err = mgm.Coll(&Nft{}).SimpleFindWithCtx(ctx, &ns, filter, opts...)
	return
}

func CountNftWithCtx(ctx context.Context, filter interface{}) (count int64, err error) {
	return mgm.Coll(&Nft{}).CountDocuments(ctx, filter)
}
