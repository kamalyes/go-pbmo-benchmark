/**
 * @Author: developer 500000000@example.com
 * @Date: 2023-09-16 17:08:15
 * @LastEditors: developer 500000000@example.com
 * @LastEditTime: 2025-06-12 17:15:33
 * @FilePath: \go-pbmo-benchmark\models\wrapper.go
 * @Description: Wrapper 字段模型 - 指针和 Wrapper 类型转换
 *
 * Copyright (c) 2025 by developer, All Rights Reserved.
 */

package models

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// ══════════════════════════════════════════════════════════════════════════════
// 场景 1：Wrapper 字段转换
// ══════════════════════════════════════════════════════════════════════════════

// WrapperPB Wrapper 字段 Protobuf 结构
type WrapperPB struct {
	Name     string
	IsActive *wrapperspb.BoolValue
	Count    *wrapperspb.Int32Value
	Score    *wrapperspb.DoubleValue
	Label    *wrapperspb.StringValue
}

// WrapperModel Wrapper 字段数据库模型
type WrapperModel struct {
	Name     string   `gorm:"column:name;type:varchar(100);primaryKey"`
	IsActive *bool    `gorm:"column:is_active;type:tinyint(1)"`
	Count    *int32   `gorm:"column:count;type:int"`
	Score    *float64 `gorm:"column:score;type:decimal(10,2)"`
	Label    *string  `gorm:"column:label;type:varchar(200)"`
}

// TableName 指定表名
func (WrapperModel) TableName() string {
	return "wrapper_info"
}

// ══════════════════════════════════════════════════════════════════════════════
// 场景 2：时间指针转换
// ══════════════════════════════════════════════════════════════════════════════

// TimePtrPB 时间指针 Protobuf 结构
type TimePtrPB struct {
	Name        string
	ScheduledAt *timestamppb.Timestamp
	ReleasedAt  *timestamppb.Timestamp
}

// TimePtrModel 时间指针数据库模型
type TimePtrModel struct {
	Name        string     `gorm:"column:name;type:varchar(100);primaryKey"`
	ScheduledAt *time.Time `gorm:"column:scheduled_at;type:datetime"`
	ReleasedAt  *time.Time `gorm:"column:released_at;type:datetime"`
}

// TableName 指定表名
func (TimePtrModel) TableName() string {
	return "time_ptr_info"
}

// ══════════════════════════════════════════════════════════════════════════════
// 场景 3：零值时间处理
// ══════════════════════════════════════════════════════════════════════════════

// TimeZeroPB 零值时间 Protobuf 结构
type TimeZeroPB struct {
	Name      string
	CreatedAt *timestamppb.Timestamp
	UpdatedAt *timestamppb.Timestamp
}

// TimeZeroModel 零值时间数据库模型
type TimeZeroModel struct {
	Name      string    `gorm:"column:name;type:varchar(100);primaryKey"`
	CreatedAt time.Time `gorm:"column:created_at;type:datetime"`
	UpdatedAt time.Time `gorm:"column:updated_at;type:datetime"`
}

// TableName 指定表名
func (TimeZeroModel) TableName() string {
	return "time_zero_info"
}

// ══════════════════════════════════════════════════════════════════════════════
// 场景 4：命名切片类型
// ══════════════════════════════════════════════════════════════════════════════

// StringSlice 自定义字符串切片类型
type StringSlice []string

// NamedSlicePB 命名切片 Protobuf 结构
type NamedSlicePB struct {
	Name  string
	Tags  []string
	Items []string
}

// NamedSliceModel 命名切片数据库模型
type NamedSliceModel struct {
	Name  string      `gorm:"column:name;type:varchar(100);primaryKey"`
	Tags  StringSlice `gorm:"column:tags;type:json"`
	Items StringSlice `gorm:"column:items;type:json"`
}

// TableName 指定表名
func (NamedSliceModel) TableName() string {
	return "named_slice_info"
}
