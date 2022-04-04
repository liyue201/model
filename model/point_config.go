/**
 * Created by g7tianyi on Mar 22, 2022
 */

package model

import (
	"time"

	"github.com/Overealityio/overeality-server-model/util"
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

const PointConfigTableName = "point_configs"

// PointConfig 积分计算的配置信息
type PointConfig struct {
	mgm.IDField `json:",inline" bson:",inline"`
	// 社交分享的基础分数
	SocialShareFactor int `json:"socialShareFactor" bson:"socialShareFactor"`
	// 邀请一个用户加入Overeality时获得的积分
	InvitationFactor int `json:"invitationFactor" bson:"invitationFactor"`
	// 2022年4月4号，我们决定把用户总积分信息在有邀请用户或社交分享等活动发生时候，做实时更新，结果保存至User表的TotalPoint字段
	// 但在这个之前，需要把从前的积分计算好放入User表中，相当于对User表中TotalPoints等字段进行初始化
	// 这个Collection里存入的这个字段，是表征这个初始化过程是否已执行
	UserPointCalculated bool `json:"userPointCalculated" bson:"userPointCalculated"`
	// 是否正在重新计算用户积分的排名
	IsCalculatingUserPointRank bool `json:"isCalculatingUserPointRank" bson:"isCalculatingUserPointRank"`
	// 用户积分的排名最近一次更新的时间
	UserPointRankLastUpdatedAt int64 `json:"userPointRankLastUpdatedAt" bson:"userPointRankLastUpdatedAt"`
	// 该配置最后一次是由谁更新的
	LastUpdatedBy primitive.ObjectID `json:"lastUpdatedBy,omitempty" bson:"lastUpdatedBy"`
	// 记录创建时间
	CreatedAt *time.Time `json:"createdAt,omitempty" bson:"createdAt"`
	// 记录更新时间
	UpdatedAt *time.Time `json:"updatedAt,omitempty" bson:"updatedAt"`
}

func (m *PointConfig) CollectionName() string { return PointConfigTableName }

func (m *PointConfig) Validate() error {
	return nil
}

func (m *PointConfig) WithDefaults() *PointConfig {
	if m.SocialShareFactor == 0 {
		m.SocialShareFactor = 3
	}
	if m.InvitationFactor == 0 {
		m.InvitationFactor = 5
	}
	if m.CreatedAt == nil {
		m.CreatedAt = util.TimeNow()
	}
	if m.UpdatedAt == nil {
		m.UpdatedAt = util.TimeNow()
	}
	return m
}

func (m *PointConfig) CopyProperties(o *PointConfig) {
	if m.SocialShareFactor != o.SocialShareFactor {
		m.SocialShareFactor = o.SocialShareFactor
	}
	if m.InvitationFactor != o.InvitationFactor {
		m.InvitationFactor = o.InvitationFactor
	}
	if m.LastUpdatedBy != o.LastUpdatedBy {
		m.LastUpdatedBy = o.LastUpdatedBy
	}
}
