/**
 * Created by g7tianyi on Apr 13, 2022
 */

package model

import (
	"context"
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Influencer 名人信息
type Influencer struct {
	// ID
	mgm.IDField  `json:",inline" bson:",inline"`
	// 姓名
	Name      string             `json:"name,omitempty" bson:"name"`
	// 头像
	Avatar    string             `json:"avatar,omitempty" bson:"avatar"`
	// 简介
	Bio       string             `json:"bio,omitempty" bson:"bio"`
	// 简介图片
	BioImages []Image            `json:"bioImages,omitempty" bson:"bioImages"`
	//Followers     []primitive.ObjectID `json:"followers,omitempty" bson:"followers"`
	//FollowerCount uint64               `json:"followerCount,omitempty" bson:"followerCount"`
	//Hots          uint64               `json:"hots,omitempty" bson:"hots"`
	// 1: active 2: pending
	Status      uint64       `json:"status,omitempty" bson:"status"`
	GuessImages []GuessImage `json:"guessImages,omitempty" bson:"guessImages"`
	Roadmap     []string     `json:"roadmap,omitempty" bson:"roadmap"`
	CreatedAt   *time.Time   `json:"createdAt,omitempty" bson:"createdAt"`
	UpdatedAt   *time.Time   `json:"lastUpdatedAt,omitempty" bson:"last_updated_at"`
}

func (c *Influencer) Validate() error {
	return nil
}

type GuessImage struct {
	Uri           string     `json:"uri,omitempty" bson:"uri"`
	RevealedAfter *time.Time `json:"revealedAfter,omitempty" bson:"revealedAfter"`
}

func (c *Influencer) CollectionName() string { return "influencers" }

func (c *Influencer) GetByID(ctx context.Context, id primitive.ObjectID) error {
	err := mgm.Coll(c).FindByIDWithCtx(ctx, id, c)
	return err
}