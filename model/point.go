package model

import (
	"context"
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Point struct {
	mgm.IDField         `json:",inline" bson:",inline"`
	UserId              primitive.ObjectID `json:"userId,omitempty" bson:"userId"`
	InfluencerId        primitive.ObjectID `json:"influencerId,omitempty" bson:"influencerId"`
	Points              float64            `json:"points" bson:"points"`
	GuessPoint          uint64             `json:"guessPoint" bson:"guessPoint"`
	ClickActivities     []string           `json:"clickActivities,omitempty" bson:"clickActivities"`
	FacebookActivities  []string           `json:"facebookActivities,omitempty" bson:"facebookActivities"`
	InstagramActivities []string           `json:"instagramActivities,omitempty" bson:"instagramActivities"`
	TwitterActivities   []string           `json:"twitterActivities,omitempty" bson:"twitterActivities"`
	CreatedAt           *time.Time         `json:"createdAt,omitempty" bson:"created_at"`
	UpdatedAt           *time.Time         `json:"lastUpdatedAt,omitempty" bson:"last_updated_at"`
}

func (p *Point) CollectionName() string {
	return "points"
}

func ListPoints(ctx context.Context, skip, limit int64, filter interface{}, order interface{}) ([]*Point, error) {
	var ps []*Point
	opt := &options.FindOptions{}
	opt.SetSkip(skip)
	opt.SetLimit(limit)
	opt.SetSort(order)
	err := mgm.Coll(&Point{}).SimpleFindWithCtx(ctx, &ps, filter, opt)
	return ps, err
}

func CountPoints(filter interface{}) (count int64, err error) {
	return mgm.Coll(&Point{}).CountDocuments(mgm.Ctx(), filter)
}

////////////////////////////////////////////////////////////////////////////////

type PointSource uint

const (
	PointSourceInvitation             PointSource = 1
	PointSourceActivityClickSupport   PointSource = 100
	PointSourceActivityShareFacebook  PointSource = 101
	PointSourceActivityShareInstagram PointSource = 102
	PointSourceActivityShareTwitter   PointSource = 103
)

type PointDetail struct {
	mgm.IDField `json:",inline" bson:",inline"`
	UserId      primitive.ObjectID `json:"userId,omitempty" bson:"userId"`
	Source      PointSource        `json:"source,omitempty" bson:"source"`
	// Source为社区社交活动时为InfluencerId，Invitation时为受邀请人的ID
	RelatedId    *primitive.ObjectID `json:"influencerId,omitempty" bson:"influencerId"`
	PointAwarded float64             `json:"pointAwarded" bson:"pointAwarded"`
	BaseFactor   float64             `json:"baseFactor" bson:"baseFactor"`
	NFTFactors   []SmartNftAbility   `json:"nftFactors" bson:"nftFactors"`
	CreatedAt    *time.Time          `json:"createdAt,omitempty" bson:"created_at"`
}

func (p *PointDetail) CollectionName() string {
	return "point_details"
}
