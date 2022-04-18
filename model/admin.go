/**
 * Created by g7tianyi on Apr 05, 2022
 */

package model

import (
	"context"

	"github.com/kamva/mgm/v3"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Admin @Deprecated
type Admin struct {
	mgm.IDField `json:",inline" bson:",inline"`
	*User       `json:",inline" bson:",inline"`
	UserId      primitive.ObjectID   `json:"userId" bson:"userId"`
	AdminRoles  []primitive.ObjectID `json:"adminRoles" bson:"adminRoles"`
}

func (a *Admin) Validate() error {
	return nil
}

func (a *Admin) CollectionName() string {
	return "admins"
}

func (a *Admin) GetByEmail(email string) error {
	coll := mgm.Coll(a)
	if err := coll.First(bson.M{"email": email}, a); err != nil {
		return errors.Errorf("GetByEmail %s failed: %v", email, err.Error())
	}
	return nil
}

func (a *Admin) GetByUserId(ctx context.Context, userId primitive.ObjectID) (err error) {
	err = mgm.Coll(a).FindOne(ctx, bson.M{"userId": userId}).Decode(a)
	return
}

func (a *Admin) ImportFromUser(user *User) (err error) {
	a.IDField = mgm.IDField{ID: primitive.NewObjectID()}
	a.User = user
	a.UserId = user.ID
	a.AdminRoles = []primitive.ObjectID{}
	err = mgm.Coll(a).Create(a)
	return
}
