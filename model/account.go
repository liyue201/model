package model

import (
	"context"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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
	// mandatory: false. only used for DId Card NFT.
	DIdCardEquipped bool  `json:"dIdCardEquipped,omitempty" bson:"dIdCardEquipped"`
	// mandatory: false. The equipped token ability.
	DIdCardEquippedAbility *DIdCardAbility `json:"dIdCardEquippedAbility,omitempty" bson:"dIdCardEquippedAbility"`
}

func (a *TokenAccount) CollectionName() string {
	return "token_accounts"
}

func ListTokenAccounts(ctx context.Context, skip, limit int64, filter interface{}, order interface{}) ([]*TokenAccount, error) {
	var ts []*TokenAccount
	opt := &options.FindOptions{}
	opt.SetSkip(skip)
	opt.SetLimit(limit)
	opt.SetSort(order)
	err := mgm.Coll(&TokenAccount{}).SimpleFindWithCtx(ctx, &ts, filter, opt)
	return ts, err
}

func CountTokenAccounts(filter interface{}) (count int64, err error) {
	return mgm.Coll(&TokenAccount{}).CountDocuments(mgm.Ctx(), filter)
}

func (a *TokenAccount) GetByID(ctx context.Context, id primitive.ObjectID) error {
	err := mgm.Coll(a).FindByIDWithCtx(ctx, id, a)
	return err
}