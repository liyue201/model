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
	mgm.IDField       `json:",inline" bson:",inline"`
	SocialShareFactor int                `json:"socialShareFactor" bson:"socialShareFactor"`
	InvitationFactor  int                `json:"invitationFactor" bson:"invitationFactor"`
	LastUpdatedBy     primitive.ObjectID `json:"lastUpdatedBy,omitempty" bson:"lastUpdatedBy"`
	CreatedAt         *time.Time         `json:"createdAt,omitempty" bson:"createdAt"`
	UpdatedAt         *time.Time         `json:"updatedAt,omitempty" bson:"updatedAt"`
}

func (m *PointConfig) TableName() string { return PointConfigTableName }

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
	m.CreatedAt = util.TimeNow()
	m.UpdatedAt = util.TimeNow()
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
