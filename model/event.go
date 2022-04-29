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
	mgm.IDField `json:",inline" bson:",inline"`
	// 社区ID
	CommunityId primitive.ObjectID `json:"communityId,omitempty" bson:"communityId"`
	// 标题
	Title string `json:"title,omitempty" bson:"title"`
	// 文字内容
	TextContents []string `json:"textContents,omitempty" bson:"textContents"`
	// 展示时的图片
	Images []Image `json:"images,omitempty" bson:"images"`
	// 文章的URL
	Uri string `json:"uri,omitempty" bson:"uri"`
	// 发布时间
	PublishedAt *time.Time `json:"publishedAt,omitempty" bson:"publishedAt"`
	// 活动开始时间
	StartAt *time.Time `json:"startAt,omitempty" bson:"startAt"`
	// 活动结束时间
	EndAt *time.Time `json:"endAt,omitempty" bson:"endAt"`
	// 活动状态. 1: Active; 2: Pending; 3: Inactive
	Status uint64 `json:"status,omitempty" bson:"status"`
	// 组织人
	OrganizedBy []primitive.ObjectID `json:"organizedBy,omitempty" bson:"organizedBy"`
	// 创建时间
	CreatedAt *time.Time `json:"createdAt,omitempty" bson:"created_at"`
	// 最近一次更新时间
	UpdatedAt *time.Time `json:"lastUpdatedAt,omitempty" bson:"last_updated_at"`
}

func (c *Event) CollectionName() string {
	return "events"
}

func (c *Event) Validate() error {
	return nil
}
