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
	// mandatory: false. Network chain id.
	ChainId int `json:"chainId,omitempty" bson:"chainId"`
	// mandatory: false. Network chain currency symbol.
	CurrencySymbol string  `json:"currencySymbol,omitempty" bson:"currencySymbol"`
	// mandatory: false. Network chain block explorer url.
	BlockExplorerURL  string  `json:"blockExplorerURL,omitempty" bson:"blockExplorerURL"`
	// mandatory: false. Network icon
	IconUri     string              `json:"iconUri,omitempty" bson:"iconUri"`
	// mandatory: false. If the network native token is self-defined ERC20
	NativeTokenId *primitive.ObjectID `json:"nativeTokenId,omitempty" bson:"nativeTokenId"`
	// mandatory: false. If the network is on chain, native currency is needed.
	NativeCurrency *NativeCurrency  `json:"nativeCurrency,omitempty" bson:"nativeCurrency"`
	//mandatory: false.  display order
	Order int `json:"order,omitempty" bson:"order"`
}

type NativeCurrency struct {
	Name string `json:"name"`
	Symbol string `json:"symbol"`
	Decimals uint `json:"decimals"`
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
