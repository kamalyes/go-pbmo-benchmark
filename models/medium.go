/**
 * @Author: developer 500000000@example.com
 * @Date: 2023-09-16 15:25:11
 * @LastEditors: developer 500000000@example.com
 * @LastEditTime: 2025-06-12 18:05:28
 * @FilePath: \go-pbmo-benchmark\models\medium.go
 * @Description: 中等复杂度模型 - 8-12 字段多类型组合
 *
 * Copyright (c) 2025 by developer, All Rights Reserved.
 */

package models

import (
	"time"

	sqltypes "github.com/kamalyes/go-sqlbuilder/types"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// ══════════════════════════════════════════════════════════════════════════════
// 场景 1：会员资料（11 字段）
// ══════════════════════════════════════════════════════════════════════════════
type MemberProfilePB struct {
	MemberId    string
	Username    string
	Email       string
	PhoneNumber string
	Level       int32
	Score       float64
	IsActive    bool
	Tags        []string
	Balance     *wrapperspb.DoubleValue
	CreatedAt   *timestamppb.Timestamp
	UpdatedAt   *timestamppb.Timestamp
}

// MemberProfileModel 会员资料数据库模型
type MemberProfileModel struct {
	MemberID    string               `pbmo:"MemberId" gorm:"column:member_id;type:varchar(64);primaryKey"`
	Username    string               `gorm:"column:username;type:varchar(100);uniqueIndex;not null"`
	Email       string               `gorm:"column:email;type:varchar(200);index"`
	PhoneNumber string               `gorm:"column:phone_number;type:varchar(20);index"`
	Level       int                  `gorm:"column:level;type:int;default:1;index"`
	Score       float64              `gorm:"column:score;type:decimal(10,2);default:0"`
	IsActive    bool                 `gorm:"column:is_active;type:tinyint(1);default:1;index"`
	Tags        sqltypes.StringSlice `gorm:"column:tags;type:json"`
	Balance     *float64             `gorm:"column:balance;type:decimal(18,2)"`
	CreatedAt   time.Time            `gorm:"column:created_at;type:datetime;not null;index"`
	UpdatedAt   time.Time            `gorm:"column:updated_at;type:datetime;not null"`
}

// TableName 指定表名
func (MemberProfileModel) TableName() string {
	return "member_profile"
}

// ServiceConfigPB 服务配置 Protobuf 结构
type ServiceConfigPB struct {
	ServiceId   string
	ServiceName string
	Version     string
	Enabled     *wrapperspb.BoolValue
	Port        int32
	Endpoints   []string
	Settings    map[string]any
	Priority    *wrapperspb.Int32Value
	UpdatedAt   *timestamppb.Timestamp
}

// ServiceConfigModel 服务配置数据库模型
type ServiceConfigModel struct {
	ServiceID   string               `pbmo:"ServiceId" gorm:"column:service_id;type:varchar(64);primaryKey"`
	ServiceName string               `gorm:"column:service_name;type:varchar(100);index;not null"`
	Version     string               `gorm:"column:version;type:varchar(20);not null"`
	Enabled     *bool                `gorm:"column:enabled;type:tinyint(1);default:1"`
	Port        int                  `gorm:"column:port;type:int;not null"`
	Endpoints   sqltypes.StringSlice `gorm:"column:endpoints;type:json"`
	Settings    sqltypes.MapAny      `gorm:"column:settings;type:json"`
	Priority    *int32               `gorm:"column:priority;type:int;default:5"`
	UpdatedAt   time.Time            `gorm:"column:updated_at;type:datetime;not null"`
}

// TableName 指定表名
func (ServiceConfigModel) TableName() string {
	return "service_config"
}

// ══════════════════════════════════════════════════════════════════════════════
// 场景 3：标准中等模型（8 字段）
// ══════════════════════════════════════════════════════════════════════════════

// MediumPB 中等复杂度 Protobuf 结构
type MediumPB struct {
	Id       uint64
	Name     string
	Email    string
	Age      int32
	Score    float64
	Active   bool
	Tags     []string
	Priority int32
}

// MediumModel 中等复杂度数据库模型
type MediumModel struct {
	ID       uint64   `pbmo:"Id" gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement"`
	Name     string   `gorm:"column:name;type:varchar(100);not null;index"`
	Email    string   `gorm:"column:email;type:varchar(200);index"`
	Age      int      `gorm:"column:age;type:int"`
	Score    float64  `gorm:"column:score;type:decimal(10,2);default:0"`
	Active   bool     `gorm:"column:active;type:tinyint(1);default:1;index"`
	Tags     []string `gorm:"column:tags;type:json"`
	Priority int      `gorm:"column:priority;type:int;default:0"`
}

// TableName 指定表名
func (MediumModel) TableName() string {
	return "medium_info"
}

// ══════════════════════════════════════════════════════════════════════════════
// 场景 4：完整功能模型（13 字段）
// ══════════════════════════════════════════════════════════════════════════════

// FullPB 完整功能 Protobuf 结构
type FullPB struct {
	Id          uint64
	Name        string
	Email       string
	Age         int32
	Score       float64
	Active      bool
	Tags        []string
	CreatedAt   *timestamppb.Timestamp
	UpdatedAt   *timestamppb.Timestamp
	MinVal      *wrapperspb.Int32Value
	MaxVal      *wrapperspb.Int32Value
	Description *wrapperspb.StringValue
}

// FullModel 完整功能数据库模型
type FullModel struct {
	ID          uint64    `pbmo:"Id" gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement"`
	Name        string    `gorm:"column:name;type:varchar(100);not null;index"`
	Email       string    `gorm:"column:email;type:varchar(200);index"`
	Age         int       `gorm:"column:age;type:int"`
	Score       float64   `gorm:"column:score;type:decimal(10,2);default:0"`
	Active      bool      `gorm:"column:active;type:tinyint(1);default:1;index"`
	Tags        []string  `gorm:"column:tags;type:json"`
	CreatedAt   time.Time `gorm:"column:created_at;type:datetime;not null;index"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:datetime;not null"`
	MinVal      *int32    `gorm:"column:min_val;type:int"`
	MaxVal      *int32    `gorm:"column:max_val;type:int"`
	Description *string   `gorm:"column:description;type:text"`
}

// TableName 指定表名
func (FullModel) TableName() string {
	return "full_info"
}
