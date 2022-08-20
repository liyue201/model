package model

import (
	"context"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// WLSI大会邀请用户
type WLSIUser struct {
	mgm.IDField   `json:",inline" bson:",inline"`
	Email         string `json:"email,omitempty" bson:"email"`
	Password      string `json:"password,omitempty" bson:"password"`
	WalletAddress string `json:"walletAddress,omitempty" bson:"wallet_address"`
	Code          string `json:"code" bson:"code"`
}

func (r *WLSIUser) CollectionName() string {
	return "wlsi_user"
}

func (r *WLSIUser) Validate() error {
	return nil
}

func (r *WLSIUser) GetByID(ctx context.Context, id primitive.ObjectID) error {
	err := mgm.Coll(r).FindByIDWithCtx(ctx, id, r)
	return err
}

func (r *WLSIUser) GetByEmail(ctx context.Context, email string) error {
	coll := mgm.Coll(r)
	err := coll.FirstWithCtx(ctx, bson.M{"email": email}, r)
	return err
}

func CountWLSIUsers(filter interface{}) (count int64, err error) {
	return mgm.Coll(&WLSIUser{}).CountDocuments(mgm.Ctx(), filter)
}
