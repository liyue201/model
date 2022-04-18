/**
 * Created by g7tianyi on Apr 18, 2022
 */

package model

import (
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// News 新闻
type News struct {
	mgm.IDField  `json:",inline" bson:",inline"`
	Contributors []string           `json:"contributors,omitempty" bson:"contributors"`
	CommunityId  primitive.ObjectID `json:"communityId,omitempty" bson:"communityId"`
	Title        string             `json:"title,omitempty" bson:"title"`
	TextContents []string           `json:"textContents,omitempty" bson:"textContents"`
	Images       []Image            `json:"images,omitempty" bson:"images"`
	Uri          string             `json:"uri,omitempty" bson:"uri"`
	Type         uint64             `json:"type,omitempty" bson:"type"`
	PublishedAt  *time.Time         `json:"publishedAt,omitempty" bson:"publishedAt"`
	CreatedAt    *time.Time         `json:"createdAt,omitempty" bson:"created_at"`
	UpdatedAt    *time.Time         `json:"lastUpdatedAt,omitempty" bson:"last_updated_at"`
}

func (m *News) CollectionName() string {
	return "news"
}

func (m *News) Validate() error {
	return nil
}
