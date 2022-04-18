package model

import (
	"context"
	
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Network struct {
	mgm.IDField `json:",inline" bson:",inline"`
	// mandatory: true. Network name.
	Name string `json:"name,omitempty" bson:"name"`
	// mandatory: false. Network rpc url.
	RpcUrl string `json:"rpcUrl,omitempty" bson:"rpcUrl"`
}

func (n *Network) CollectionName() string {
	return "networks"
}

func (n *Network) Validate() error {
	return nil
}

func (n *Network) GetById(ctx context.Context, id primitive.ObjectID) error {
	err := mgm.Coll(n).FindByIDWithCtx(ctx, id, n)
	return err
}

func ListNetworks(ctx context.Context, skip, limit int64, filter interface{}, order interface{}) ([]*Network, error) {
	var ns []*Network
	opt := &options.FindOptions{}
	opt.SetSkip(skip)
	opt.SetLimit(limit)
	opt.SetSort(order)
	err := mgm.Coll(&Network{}).SimpleFindWithCtx(ctx, &ns, filter, opt)
	return ns, err
}

func CountNetworks(filter interface{}) (count int64, err error) {
	return mgm.Coll(&Network{}).CountDocuments(mgm.Ctx(), filter)
}
