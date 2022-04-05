/**
 * Created by g7tianyi on Apr 05, 2022
 */

package model

import (
	"github.com/kamva/mgm/v3"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Admin struct {
	mgm.IDField `json:",inline" bson:",inline"`
	*User       `json:",inline" bson:",inline"`
	AdminRoles  []primitive.ObjectID `json:"adminRoles" bson:"adminRoles"`
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
