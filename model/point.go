package model

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type Point struct {
	mgm.IDField `json:",inline" bson:",inline"`
	UserId              primitive.ObjectID `json:"userId,omitempty" bson:"userId"`
	InfluencerId        primitive.ObjectID `json:"influencerId,omitempty" bson:"influencerId"`
	Points              uint64             `json:"points" bson:"points"`
	GuessPoint          uint64             `json:"guessPoint" bson:"guessPoint"`
	ClickActivities     []string           `json:"clickActivities,omitempty" bson:"clickActivities"`
	FacebookActivities  []string           `json:"facebookActivities,omitempty" bson:"facebookActivities"`
	InstagramActivities []string           `json:"instagramActivities,omitempty" bson:"instagramActivities"`
	TwitterActivities   []string           `json:"twitterActivities,omitempty" bson:"twitterActivities"`
	CreatedAt           *time.Time         `json:"createdAt,omitempty" bson:"created_at"`
	UpdatedAt           *time.Time         `json:"lastUpdatedAt,omitempty" bson:"last_updated_at"`
}