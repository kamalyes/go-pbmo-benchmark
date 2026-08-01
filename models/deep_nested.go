/**
 * @Author: developer 500000000@example.com
 * @Date: 2023-09-16 18:22:11
 * @LastEditors: developer 500000000@example.com
 * @LastEditTime: 2025-06-12 19:38:25
 * @FilePath: \go-pbmo-benchmark\models\deep_nested.go
 * @Description: 深度嵌套模型 - 4-6 层嵌套结构测试
 *
 * Copyright (c) 2025 by developer, All Rights Reserved.
 */

package models

import (
	"time"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// ══════════════════════════════════════════════════════════════════════════════
// 场景 1：4 层嵌套 - Level1 > Level2 > Level3 > Level4
// ══════════════════════════════════════════════════════════════════════════════

// Level4PB 第4层 Protobuf 结构
type Level4PB struct {
	Id    string
	Name  string
	Value int32
	Data  string
}

// Level4Model 第4层数据库模型
type Level4Model struct {
	ID    string `pbmo:"Id" gorm:"column:id;type:varchar(64)"`
	Name  string `gorm:"column:name;type:varchar(100)"`
	Value int    `gorm:"column:value;type:int"`
	Data  string `gorm:"column:data;type:text"`
}

// Level3PB 第3层 Protobuf 结构
type Level3PB struct {
	Id     string
	Name   string
	Count  int32
	Level4 *Level4PB
}

// Level3Model 第3层数据库模型
type Level3Model struct {
	ID     string       `pbmo:"Id" gorm:"column:id;type:varchar(64)"`
	Name   string       `gorm:"column:name;type:varchar(100)"`
	Count  int          `gorm:"column:count;type:int"`
	Level4 *Level4Model `gorm:"embedded;embeddedPrefix:l4_"`
}

// Level2PB 第2层 Protobuf 结构
type Level2PB struct {
	Id     string
	Title  string
	Status int32
	Level3 *Level3PB
}

// Level2Model 第2层数据库模型
type Level2Model struct {
	ID     string       `pbmo:"Id" gorm:"column:id;type:varchar(64)"`
	Title  string       `gorm:"column:title;type:varchar(200)"`
	Status int          `gorm:"column:status;type:int"`
	Level3 *Level3Model `gorm:"embedded;embeddedPrefix:l3_"`
}

// DeepNested4PB 4层嵌套 Protobuf 结构
type DeepNested4PB struct {
	Id        string
	Name      string
	Level2    *Level2PB
	CreatedAt *timestamppb.Timestamp
}

// DeepNested4Model 4层嵌套数据库模型
type DeepNested4Model struct {
	ID        string       `pbmo:"Id" gorm:"column:id;type:varchar(64);primaryKey"`
	Name      string       `gorm:"column:name;type:varchar(100);not null"`
	Level2    *Level2Model `gorm:"embedded;embeddedPrefix:l2_"`
	CreatedAt time.Time    `gorm:"column:created_at;type:datetime;not null"`
}

func (DeepNested4Model) TableName() string { return "deep_nested_4" }
