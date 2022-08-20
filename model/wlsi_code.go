package model

import (
	"context"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// WLSI大会兑换码
type WLSICode struct {
	mgm.IDField `json:",inline" bson:",inline"`
	Code        string `json:"code" bson:"code"`
}

func (r *WLSICode) CollectionName() string {
	return "wlsi_code"
}

func (r *WLSICode) GetByID(ctx context.Context, id primitive.ObjectID) error {
	err := mgm.Coll(r).FindByIDWithCtx(ctx, id, r)
	return err
}

func (r *WLSICode) GetByCode(ctx context.Context, code string) (err error) {
	err = mgm.Coll(r).FindOne(ctx, bson.M{"code": code}).Decode(r)
	return
}
