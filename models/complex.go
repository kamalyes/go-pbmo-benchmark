/**
 * @Author: developer 500000000@example.com
 * @Date: 2023-09-16 15:32:55
 * @LastEditors: developer 500000000@example.com
 * @LastEditTime: 2025-06-12 18:15:05
 * @FilePath: \go-pbmo-benchmark\models\complex.go
 * @Description: 复杂模型定义 - 20+ 字段大模型
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
// 场景 1：大型模型（20+ 字段）- LargePB
// ══════════════════════════════════════════════════════════════════════════════

// LargePB 大型模型 Protobuf 结构
type LargePB struct {
	Id          uint64
	Name        string
	Email       string
	Phone       string
	Status      int32
	Priority    int32
	Score       float64
	Rating      float64
	Active      bool
	Verified    bool
	Tags        []string
	Country     string
	Region      string
	City        string
	ZipCode     string
	Description string
	CreatedAt   *timestamppb.Timestamp
	UpdatedAt   *timestamppb.Timestamp
	MinVal      *wrapperspb.Int32Value
	MaxVal      *wrapperspb.Int32Value
}

// LargeModel 大型模型数据库模型
type LargeModel struct {
	ID          uint64    `pbmo:"Id" gorm:"column:id;type:bigint unsigned;primaryKey;autoIncrement"`
	Name        string    `gorm:"column:name;type:varchar(100);not null;index"`
	Email       string    `gorm:"column:email;type:varchar(200);index"`
	Phone       string    `gorm:"column:phone;type:varchar(20);index"`
	Status      int       `gorm:"column:status;type:int;default:1;index"`
	Priority    int       `gorm:"column:priority;type:int;default:0;index"`
	Score       float64   `gorm:"column:score;type:decimal(10,2);default:0"`
	Rating      float64   `gorm:"column:rating;type:decimal(3,2);default:0"`
	Active      bool      `gorm:"column:active;type:tinyint(1);default:1;index"`
	Verified    bool      `gorm:"column:verified;type:tinyint(1);default:0"`
	Tags        []string  `gorm:"column:tags;type:json"`
	Country     string    `gorm:"column:country;type:varchar(50);index"`
	Region      string    `gorm:"column:region;type:varchar(50);index"`
	City        string    `gorm:"column:city;type:varchar(50);index"`
	ZipCode     string    `gorm:"column:zip_code;type:varchar(20)"`
	Description string    `gorm:"column:description;type:text"`
	CreatedAt   time.Time `gorm:"column:created_at;type:datetime;not null;index"`
	UpdatedAt   time.Time `gorm:"column:updated_at;type:datetime;not null"`
	MinVal      *int32    `gorm:"column:min_val;type:int"`
	MaxVal      *int32    `gorm:"column:max_val;type:int"`
}

// TableName 指定表名
func (LargeModel) TableName() string {
	return "large_info"
}

// ══════════════════════════════════════════════════════════════════════════════
// 场景 2：组织信息模型（19 字段）
// ══════════════════════════════════════════════════════════════════════════════

// OrganizationPB 组织信息 Protobuf 结构
type OrganizationPB struct {
	OrgId           string
	OrgCode         string
	OrgName         string
	Status          int32
	Settings        map[string]any
	OwnerUserId     string
	AdminUserIds    []string
	DepartmentCount int32
	EmployeeCount   int32
	RegionCode      string
	Industry        string
	BusinessType    int32
	ContactEmail    string
	ContactPhone    string
	CreatedAt       *timestamppb.Timestamp
	UpdatedAt       *timestamppb.Timestamp
	ExpiredAt       *timestamppb.Timestamp
	MaxUsers        *wrapperspb.Int32Value
	MaxStorage      *wrapperspb.Int64Value
}

// OrganizationModel 组织信息数据库模型
type OrganizationModel struct {
	OrgID           string               `pbmo:"OrgId" gorm:"column:org_id;type:varchar(64);primaryKey"`
	OrgCode         string               `gorm:"column:org_code;type:varchar(50);uniqueIndex;not null"`
	OrgName         string               `gorm:"column:org_name;type:varchar(200);index;not null"`
	Status          int                  `gorm:"column:status;type:int;default:1;index"`
	Settings        sqltypes.MapAny      `gorm:"column:settings;type:json"`
	OwnerUserID     string               `pbmo:"OwnerUserId" gorm:"column:owner_user_id;type:varchar(64);index"`
	AdminUserIDs    sqltypes.StringSlice `gorm:"column:admin_user_ids;type:json"`
	DepartmentCount int                  `gorm:"column:department_count;type:int;default:0"`
	EmployeeCount   int                  `gorm:"column:employee_count;type:int;default:0"`
	RegionCode      string               `gorm:"column:region_code;type:varchar(20);index"`
	Industry        string               `gorm:"column:industry;type:varchar(50);index"`
	BusinessType    int                  `gorm:"column:business_type;type:int;index"`
	ContactEmail    string               `gorm:"column:contact_email;type:varchar(200)"`
	ContactPhone    string               `gorm:"column:contact_phone;type:varchar(20)"`
	CreatedAt       time.Time            `gorm:"column:created_at;type:datetime;not null;index"`
	UpdatedAt       time.Time            `gorm:"column:updated_at;type:datetime;not null"`
	ExpiredAt       time.Time            `gorm:"column:expired_at;type:datetime;index"`
	MaxUsers        *int32               `gorm:"column:max_users;type:int"`
	MaxStorage      *int64               `gorm:"column:max_storage;type:bigint"`
}

// TableName 指定表名
func (OrganizationModel) TableName() string {
	return "organization"
}

// UserProfilePB 用户资料 Protobuf 结构（26+ 字段）
type UserProfilePB struct {
	UserId         string
	Username       string
	Nickname       string
	AvatarUrl      string
	Email          string
	PhoneNumber    string
	Gender         int32
	BirthDate      string
	RegionCode     string
	CountryCode    string
	CityCode       string
	Address        string
	ZipCode        string
	Status         int32
	Level          int32
	Score          int64
	Balance        float64
	TotalSpent     float64
	LastLoginTime  int64
	LastLoginIp    string
	RegisterTime   int64
	RegisterIp     string
	RegisterDevice string
	VerifiedEmail  bool
	VerifiedPhone  bool
	CreatedAt      *timestamppb.Timestamp
	UpdatedAt      *timestamppb.Timestamp
}

// UserProfileModel 用户资料数据库模型
type UserProfileModel struct {
	UserID         string    `pbmo:"UserId" gorm:"column:user_id;type:varchar(64);primaryKey"`
	Username       string    `gorm:"column:username;type:varchar(100);uniqueIndex;not null"`
	Nickname       string    `gorm:"column:nickname;type:varchar(100);index"`
	AvatarUrl      string    `gorm:"column:avatar_url;type:varchar(500)"`
	Email          string    `gorm:"column:email;type:varchar(200);index"`
	PhoneNumber    string    `gorm:"column:phone_number;type:varchar(20);index"`
	Gender         int       `gorm:"column:gender;type:tinyint;default:0"`
	BirthDate      string    `gorm:"column:birth_date;type:varchar(10)"`
	RegionCode     string    `gorm:"column:region_code;type:varchar(20);index"`
	CountryCode    string    `gorm:"column:country_code;type:varchar(10);index"`
	CityCode       string    `gorm:"column:city_code;type:varchar(20);index"`
	Address        string    `gorm:"column:address;type:varchar(500)"`
	ZipCode        string    `gorm:"column:zip_code;type:varchar(20)"`
	Status         int       `gorm:"column:status;type:int;default:1;index"`
	Level          int       `gorm:"column:level;type:int;default:1;index"`
	Score          int64     `gorm:"column:score;type:bigint;default:0"`
	Balance        float64   `gorm:"column:balance;type:decimal(18,2);default:0"`
	TotalSpent     float64   `gorm:"column:total_spent;type:decimal(18,2);default:0"`
	LastLoginTime  int64     `gorm:"column:last_login_time;type:bigint;index"`
	LastLoginIp    string    `gorm:"column:last_login_ip;type:varchar(50)"`
	RegisterTime   int64     `gorm:"column:register_time;type:bigint;not null;index"`
	RegisterIp     string    `gorm:"column:register_ip;type:varchar(50)"`
	RegisterDevice string    `gorm:"column:register_device;type:varchar(100)"`
	VerifiedEmail  bool      `gorm:"column:verified_email;type:tinyint(1);default:0"`
	VerifiedPhone  bool      `gorm:"column:verified_phone;type:tinyint(1);default:0"`
	CreatedAt      time.Time `gorm:"column:created_at;type:datetime;not null;index"`
	UpdatedAt      time.Time `gorm:"column:updated_at;type:datetime;not null"`
}

// TableName 指定表名
func (UserProfileModel) TableName() string {
	return "user_profile"
}
