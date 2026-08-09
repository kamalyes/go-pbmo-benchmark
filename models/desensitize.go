/**
 * @Author: developer 500000000@example.com
 * @Date: 2023-09-16 15:18:33
 * @LastEditors: developer 500000000@example.com
 * @LastEditTime: 2025-06-12 17:05:12
 * @FilePath: \go-pbmo-benchmark\models\desensitize.go
 * @Description: 脱敏模型定义 - 携带 desensitize tag 的敏感字段转换场景
 *
 * Copyright (c) 2025 by developer, All Rights Reserved.
 */

package models

// ══════════════════════════════════════════════════════════════════════════════
// 场景：脱敏字段（WithDesensitize 开关的掩码对照）
// ══════════════════════════════════════════════════════════════════════════════

// DesensitizePB 敏感信息 Protobuf 结构
type DesensitizePB struct {
	Name     string
	Phone    string
	Email    string
	IdCard   string
	BankCard string
	Token    string
	Address  string
	Memo     string
}

// DesensitizeModel 敏感信息数据库模型，目标侧字段携带 desensitize tag
type DesensitizeModel struct {
	Name     string `pbmo:"Name" gorm:"column:name;type:varchar(100);not null" desensitize:"name"`
	Phone    string `pbmo:"Phone" gorm:"column:phone;type:varchar(20);index" desensitize:"mobilePhone"`
	Email    string `pbmo:"Email" gorm:"column:email;type:varchar(200)" desensitize:"email"`
	IdCard   string `pbmo:"IdCard" gorm:"column:id_card;type:varchar(20)" desensitize:"identityCard"`
	BankCard string `pbmo:"BankCard" gorm:"column:bank_card;type:varchar(32)" desensitize:"bankCard"`
	Token    string `pbmo:"Token" gorm:"column:token;type:varchar(128)" desensitize:"secret"`
	Address  string `pbmo:"Address" gorm:"column:address;type:varchar(255)" desensitize:"address"`
	Memo     string `pbmo:"Memo" gorm:"column:memo;type:varchar(255)"`
}

// TableName 指定表名
func (DesensitizeModel) TableName() string {
	return "desensitize_info"
}