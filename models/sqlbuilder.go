/**
 * @Author: developer 500000000@example.com
 * @Date: 2023-09-16 17:18:27
 * @LastEditors: developer 500000000@example.com
 * @LastEditTime: 2025-06-12 17:35:52
 * @FilePath: \go-pbmo-benchmark\models\sqlbuilder.go
 * @Description: SQLBuilder types 集成模型 - 高级类型转换
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
// 场景 1：JSON[T] 字段
// ══════════════════════════════════════════════════════════════════════════════

// ConfigData 配置数据结构
type ConfigData struct {
	Theme    string `json:"theme"`
	Language string `json:"language"`
	Timezone string `json:"timezone"`
}

// ConfigPB 配置 Protobuf 结构
type ConfigPB struct {
	Name   string
	Config map[string]any
}

// ConfigModel 配置数据库模型
type ConfigModel struct {
	Name   string                    `gorm:"column:name;type:varchar(100);primaryKey"`
	Config sqltypes.JSON[ConfigData] `gorm:"column:config;type:json"`
}

// TableName 指定表名
func (ConfigModel) TableName() string {
	return "config_info"
}

// ══════════════════════════════════════════════════════════════════════════════
// 场景 2：Slice[T] 字段
// ══════════════════════════════════════════════════════════════════════════════

// ScoreEntry 评分条目
type ScoreEntry struct {
	Label string `json:"label"`
	Value int64  `json:"value"`
}

// ScoreEntryPB 评分条目 Protobuf 结构
type ScoreEntryPB struct {
	Label string
	Value int64
}

// ScoreBoardPB 评分榜 Protobuf 结构
type ScoreBoardPB struct {
	Name   string
	Scores []*ScoreEntryPB
}

// ScoreBoardModel 评分榜数据库模型
type ScoreBoardModel struct {
	Name   string                     `gorm:"column:name;type:varchar(100);primaryKey"`
	Scores sqltypes.Slice[ScoreEntry] `gorm:"column:scores;type:json"`
}

// TableName 指定表名
func (ScoreBoardModel) TableName() string {
	return "score_board"
}

// ══════════════════════════════════════════════════════════════════════════════
// 场景 3：MapAny 字段
// ══════════════════════════════════════════════════════════════════════════════

// MetadataPB 元数据 Protobuf 结构
type MetadataPB struct {
	Name     string
	Metadata map[string]any
}

// MetadataModel 元数据数据库模型
type MetadataModel struct {
	Name     string          `gorm:"column:name;type:varchar(100);primaryKey"`
	Metadata sqltypes.MapAny `gorm:"column:metadata;type:json"`
}

// TableName 指定表名
func (MetadataModel) TableName() string {
	return "metadata_info"
}

// ══════════════════════════════════════════════════════════════════════════════
// 场景 4：综合模型（全 sqlbuilder types 集成）
// ══════════════════════════════════════════════════════════════════════════════

// ComprehensivePB 综合模型 Protobuf 结构
type ComprehensivePB struct {
	Id        string
	Name      string
	Status    int32
	Tags      []string
	Config    map[string]any
	Scores    []*ScoreEntryPB
	CreatedAt *timestamppb.Timestamp
	UpdatedAt *timestamppb.Timestamp
	Enabled   *wrapperspb.BoolValue
	Priority  *wrapperspb.Int32Value
}

// ComprehensiveModel 综合模型数据库模型
type ComprehensiveModel struct {
	ID        string                     `pbmo:"Id" gorm:"column:id;type:varchar(64);primaryKey"`
	Name      string                     `gorm:"column:name;type:varchar(100);index;not null"`
	Status    int                        `gorm:"column:status;type:int;default:1;index"`
	Tags      sqltypes.StringSlice       `gorm:"column:tags;type:json"`
	Config    sqltypes.JSON[ConfigData]  `gorm:"column:config;type:json"`
	Scores    sqltypes.Slice[ScoreEntry] `gorm:"column:scores;type:json"`
	CreatedAt time.Time                  `gorm:"column:created_at;type:datetime;not null;index"`
	UpdatedAt time.Time                  `gorm:"column:updated_at;type:datetime;not null"`
	Enabled   *bool                      `gorm:"column:enabled;type:tinyint(1)"`
	Priority  *int32                     `gorm:"column:priority;type:int"`
}

// TableName 指定表名
func (ComprehensiveModel) TableName() string {
	return "comprehensive_info"
}
