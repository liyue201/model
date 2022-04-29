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
	mgm.IDField `json:",inline" bson:",inline"`
	// 暂时不知道这个字段的意义
	Contributors []string `json:"contributors,omitempty" bson:"contributors"`
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
	// 类型. 1: Main; 2: Feature; 3: Basic
	Type uint64 `json:"type,omitempty" bson:"type"`
	// 发布时间
	PublishedAt *time.Time `json:"publishedAt,omitempty" bson:"publishedAt"`
	// 创建时间
	CreatedAt *time.Time `json:"createdAt,omitempty" bson:"created_at"`
	// 最近一次更新时间
	UpdatedAt *time.Time `json:"lastUpdatedAt,omitempty" bson:"last_updated_at"`
}

func (m *News) CollectionName() string {
	return "news"
}

func (m *News) Validate() error {
	return nil
}
