/**
 * @Author: developer 500000000@example.com
 * @Date: 2023-09-16 15:18:33
 * @LastEditors: developer 500000000@example.com
 * @LastEditTime: 2025-06-12 17:05:12
 * @FilePath: \go-pbmo-benchmark\models\simple.go
 * @Description: 简单模型定义 - 基础字段类型转换（4-6 字段）
 *
 * Copyright (c) 2025 by developer, All Rights Reserved.
 */

package models

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// ══════════════════════════════════════════════════════════════════════════════
// 场景 1：简单结构体（少量字段，类型相同）
// ══════════════════════════════════════════════════════════════════════════════

// SimplePB 简单信息 Protobuf 结构
type SimplePB struct {
	Id    uint64
	Name  string
	Email string
	Age   int32
}

// SimpleModel 简单信息数据库模型
type SimpleModel struct {
	ID    uint64 `pbmo:"Id" gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement"`
	Name  string `gorm:"column:name;type:varchar(100);not null;index"`
	Email string `gorm:"column:email;type:varchar(200);uniqueIndex"`
	Age   int    `gorm:"column:age;type:int"`
}

// TableName 指定表名
func (SimpleModel) TableName() string {
	return "simple_info"
}

// ══════════════════════════════════════════════════════════════════════════════
// 场景 2：带时间戳的简单模型
// ══════════════════════════════════════════════════════════════════════════════

// AccountInfoPB 账户信息 Protobuf 结构
type AccountInfoPB struct {
	AccountId  string
	Username   string
	Status     int32
	CreatedAt  *timestamppb.Timestamp
	LastActive *timestamppb.Timestamp
}

// AccountInfoModel 账户信息数据库模型
type AccountInfoModel struct {
	AccountID  string    `pbmo:"AccountId" gorm:"column:account_id;type:varchar(64);primaryKey"`
	Username   string    `gorm:"column:username;type:varchar(100);uniqueIndex;not null"`
	Status     int       `gorm:"column:status;type:int;default:1;index"`
	CreatedAt  time.Time `gorm:"column:created_at;type:datetime;not null;index"`
	LastActive time.Time `gorm:"column:last_active;type:datetime"`
}

// TableName 指定表名
func (AccountInfoModel) TableName() string {
	return "account_info"
}

// ══════════════════════════════════════════════════════════════════════════════
// 场景 3：字段映射测试
// ══════════════════════════════════════════════════════════════════════════════

// MappedPB 字段映射 Protobuf 结构
type MappedPB struct {
	ClientId  uint64
	UserName  string
	UserEmail string
}

// MappedModel 字段映射数据库模型
type MappedModel struct {
	ID    uint64 `pbmo:"ClientId" gorm:"column:id;type:bigint unsigned;primaryKey"`
	Name  string `pbmo:"UserName" gorm:"column:name;type:varchar(100);not null"`
	Email string `pbmo:"UserEmail" gorm:"column:email;type:varchar(200);index"`
}

// TableName 指定表名
func (MappedModel) TableName() string {
	return "mapped_info"
}
