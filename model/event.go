/**
 * Created by g7tianyi on Apr 18, 2022
 */

package model

import (
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Event 活动
type Event struct {
	// ID
	mgm.IDField  `json:",inline" bson:",inline"`
	CommunityId  primitive.ObjectID   `json:"communityId,omitempty" bson:"communityId"`
	Title        string               `json:"title,omitempty" bson:"title"`
	TextContents []string             `json:"textContents,omitempty" bson:"textContents"`
	Images       []Image              `json:"images,omitempty" bson:"images"`
	Uri          string               `json:"uri,omitempty" bson:"uri"`
	PublishedAt  *time.Time           `json:"publishedAt,omitempty" bson:"publishedAt"`
	StartAt      *time.Time           `json:"startAt,omitempty" bson:"startAt"`
	EndAt        *time.Time           `json:"endAt,omitempty" bson:"endAt"`
	Status       uint64               `json:"status,omitempty" bson:"status"`
	OrganizedBy  []primitive.ObjectID `json:"organizedBy,omitempty" bson:"organizedBy"`
	CreatedAt    *time.Time           `json:"createdAt,omitempty" bson:"created_at"`
	UpdatedAt    *time.Time           `json:"lastUpdatedAt,omitempty" bson:"last_updated_at"`
}

func (c *Event) CollectionName() string {
	return "events"
}

func (c *Event) Validate() error {
	return nil
}
