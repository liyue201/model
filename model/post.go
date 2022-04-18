package model

import (
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Post struct {
	// ID
	mgm.IDField `json:",inline" bson:",inline"`
	Author      primitive.ObjectID   `json:"author,omitempty" bson:"author"`
	TextContent string               `json:"textContent,omitempty" bson:"textContent"`
	Images      []Image              `json:"images,omitempty" bson:"images"`
	Likes       []primitive.ObjectID `json:"likes,omitempty" bson:"likes"`
	LikesCount  uint64               `json:"likesCount,omitempty" bson:"likesCount"`
	CreatedAt   *time.Time           `json:"createdAt,omitempty" bson:"created_at"`
	UpdatedAt   *time.Time           `json:"lastUpdatedAt,omitempty" bson:"last_updated_at"`
}

func (m *Post) CollectionName() string {
	return "posts"
}

func (m *Post) Validate() error {
	return nil
}

type ListPostResponse struct {
	Data  []*Post `json:"data"`
	Total int64   `json:"total"`
}

func NewListPostResponse(data []*Post, total int64) *ListPostResponse {
	return &ListPostResponse{Data: data, Total: total}
}
