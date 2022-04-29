package model

import (
	"time"

	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Post 帖子信息
type Post struct {
	// ID
	mgm.IDField `json:",inline" bson:",inline"`
	// 作者ID
	Author primitive.ObjectID `json:"author,omitempty" bson:"author"`
	// 文本内容
	TextContent string `json:"textContent,omitempty" bson:"textContent"`
	// 图片列表
	Images []Image `json:"images,omitempty" bson:"images"`
	// 点赞的人
	Likes []primitive.ObjectID `json:"likes,omitempty" bson:"likes"`
	// 点赞总数
	LikesCount uint64 `json:"likesCount,omitempty" bson:"likesCount"`
	// 创建时间
	CreatedAt *time.Time `json:"createdAt,omitempty" bson:"created_at"`
	// 最近一次更新时间
	UpdatedAt *time.Time `json:"lastUpdatedAt,omitempty" bson:"last_updated_at"`
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
