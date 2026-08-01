/**
 * @Author: developer 500000000@example.com
 * @Date: 2023-09-16 15:38:22
 * @LastEditors: developer 500000000@example.com
 * @LastEditTime: 2025-06-12 17:52:17
 * @FilePath: \go-pbmo-benchmark\models\nested.go
 * @Description: 嵌套模型定义 - 2-3 层嵌套结构和切片嵌套
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
// 场景 1：基础嵌套（2 层）- Address 嵌套
// ══════════════════════════════════════════════════════════════════════════════

// AddressPB 地址信息 Protobuf 结构
type AddressPB struct {
	Country string
	State   string
	City    string
	Street  string
	ZipCode string
}

// AddressModel 地址信息数据库模型
type AddressModel struct {
	Country string `gorm:"column:country;type:varchar(50)"`
	State   string `gorm:"column:state;type:varchar(50)"`
	City    string `gorm:"column:city;type:varchar(50)"`
	Street  string `gorm:"column:street;type:varchar(200)"`
	ZipCode string `gorm:"column:zip_code;type:varchar(20)"`
}

// UserWithAddressPB 用户地址 Protobuf 结构
type UserWithAddressPB struct {
	Id      uint64
	Name    string
	Email   string
	Address *AddressPB
}

// UserWithAddressModel 用户地址数据库模型
type UserWithAddressModel struct {
	ID      uint64        `pbmo:"Id" gorm:"column:id;type:bigint unsigned;primaryKey"`
	Name    string        `gorm:"column:name;type:varchar(100);not null"`
	Email   string        `gorm:"column:email;type:varchar(200);index"`
	Address *AddressModel `gorm:"embedded;embeddedPrefix:addr_"`
}

// TableName 指定表名
func (UserWithAddressModel) TableName() string {
	return "user_with_address"
}

// ══════════════════════════════════════════════════════════════════════════════
// 场景 2：深度嵌套（3 层）
// ══════════════════════════════════════════════════════════════════════════════

// InnerPB 内层 Protobuf 结构
type InnerPB struct {
	Label string
	Count int32
}

// InnerModel 内层数据库模型
type InnerModel struct {
	Label string `gorm:"column:label;type:varchar(50)"`
	Count int32  `gorm:"column:count;type:int"`
}

// MiddlePB 中层 Protobuf 结构
type MiddlePB struct {
	Name  string
	Inner *InnerPB
}

// MiddleModel 中层数据库模型
type MiddleModel struct {
	Name  string      `gorm:"column:name;type:varchar(100)"`
	Inner *InnerModel `gorm:"embedded;embeddedPrefix:inner_"`
}

// OuterPB 外层 Protobuf 结构
type OuterPB struct {
	Title  string
	Middle *MiddlePB
}

// OuterModel 外层数据库模型
type OuterModel struct {
	Title  string       `gorm:"column:title;type:varchar(100);primaryKey"`
	Middle *MiddleModel `gorm:"embedded;embeddedPrefix:middle_"`
}

// TableName 指定表名
func (OuterModel) TableName() string {
	return "outer_info"
}

// ══════════════════════════════════════════════════════════════════════════════
// 场景 3：切片嵌套
// ══════════════════════════════════════════════════════════════════════════════

// ListPB 列表 Protobuf 结构
type ListPB struct {
	Name  string
	Items []*InnerPB
}

// ListModel 列表数据库模型
type ListModel struct {
	Name  string        `gorm:"column:name;type:varchar(100);primaryKey"`
	Items []*InnerModel `gorm:"column:items;type:json;serializer:json"`
}

// TableName 指定表名
func (ListModel) TableName() string {
	return "list_info"
}

// ══════════════════════════════════════════════════════════════════════════════
// 场景 4：Location 嵌套
// ══════════════════════════════════════════════════════════════════════════════

// LocationInfoPB 位置信息 Protobuf 结构
type LocationInfoPB struct {
	Country   string
	Province  string
	City      string
	District  string
	Address   string
	Latitude  float64
	Longitude float64
	ZipCode   string
}

// LocationInfoModel 位置信息数据库模型
type LocationInfoModel struct {
	Country   string  `gorm:"column:country;type:varchar(50)"`
	Province  string  `gorm:"column:province;type:varchar(50)"`
	City      string  `gorm:"column:city;type:varchar(50)"`
	District  string  `gorm:"column:district;type:varchar(50)"`
	Address   string  `gorm:"column:address;type:varchar(500)"`
	Latitude  float64 `gorm:"column:latitude;type:decimal(10,7)"`
	Longitude float64 `gorm:"column:longitude;type:decimal(10,7)"`
	ZipCode   string  `gorm:"column:zip_code;type:varchar(20)"`
}

// StoreInfoPB 店铺信息 Protobuf 结构（包含位置嵌套）
type StoreInfoPB struct {
	StoreId     string
	StoreName   string
	StoreCode   string
	Category    int32
	Status      int32
	Location    *LocationInfoPB
	ManagerId   string
	PhoneNumber string
	Email       string
	CreatedAt   *timestamppb.Timestamp
	UpdatedAt   *timestamppb.Timestamp
}

// StoreInfoModel 店铺信息数据库模型
type StoreInfoModel struct {
	StoreID     string             `pbmo:"StoreId" gorm:"column:store_id;type:varchar(64);primaryKey"`
	StoreName   string             `gorm:"column:store_name;type:varchar(200);index;not null"`
	StoreCode   string             `gorm:"column:store_code;type:varchar(50);uniqueIndex;not null"`
	Category    int                `gorm:"column:category;type:int;index"`
	Status      int                `gorm:"column:status;type:int;default:1;index"`
	Location    *LocationInfoModel `gorm:"embedded;embeddedPrefix:location_"`
	ManagerID   string             `pbmo:"ManagerId" gorm:"column:manager_id;type:varchar(64);index"`
	PhoneNumber string             `gorm:"column:phone_number;type:varchar(20)"`
	Email       string             `gorm:"column:email;type:varchar(200)"`
	CreatedAt   time.Time          `gorm:"column:created_at;type:datetime;not null;index"`
	UpdatedAt   time.Time          `gorm:"column:updated_at;type:datetime;not null"`
}

// TableName 指定表名
func (StoreInfoModel) TableName() string {
	return "store_info"
}

// ══════════════════════════════════════════════════════════════════════════════
// 场景 5：带 Slice[T] 的嵌套切片
// ══════════════════════════════════════════════════════════════════════════════

// CategoryTagPB 分类标签 Protobuf 结构
type CategoryTagPB struct {
	TagId     int32
	TagName   string
	TagType   int32
	Priority  int32
	ExpiredAt *timestamppb.Timestamp
}

// CategoryTagModel 分类标签数据库模型
type CategoryTagModel struct {
	TagID     int32     `pbmo:"TagId" json:"tag_id" gorm:"column:tag_id"`
	TagName   string    `json:"tag_name" gorm:"column:tag_name"`
	TagType   int       `json:"tag_type" gorm:"column:tag_type"`
	Priority  int       `json:"priority" gorm:"column:priority"`
	ExpiredAt time.Time `json:"expired_at" gorm:"column:expired_at"`
}

// ProductCatalogPB 产品目录 Protobuf 结构（包含标签切片）
type ProductCatalogPB struct {
	ProductId   string
	ProductName string
	ProductCode string
	CategoryId  int64
	BrandId     string
	Price       float64
	Stock       int32
	Tags        []*CategoryTagPB
	Enabled     *wrapperspb.BoolValue
	SortOrder   int32
	CreatedAt   *timestamppb.Timestamp
	UpdatedAt   *timestamppb.Timestamp
}

// ProductCatalogModel 产品目录数据库模型
type ProductCatalogModel struct {
	ProductID   string                           `pbmo:"ProductId" gorm:"column:product_id;type:varchar(64);primaryKey"`
	ProductName string                           `gorm:"column:product_name;type:varchar(200);index;not null"`
	ProductCode string                           `gorm:"column:product_code;type:varchar(50);uniqueIndex;not null"`
	CategoryID  int64                            `pbmo:"CategoryId" gorm:"column:category_id;type:bigint;index"`
	BrandID     string                           `pbmo:"BrandId" gorm:"column:brand_id;type:varchar(64);index"`
	Price       float64                          `gorm:"column:price;type:decimal(18,2);not null"`
	Stock       int                              `gorm:"column:stock;type:int;default:0"`
	Tags        sqltypes.Slice[CategoryTagModel] `gorm:"column:tags;type:json"`
	Enabled     *bool                            `gorm:"column:enabled;type:tinyint(1);default:1"`
	SortOrder   int                              `gorm:"column:sort_order;type:int;default:0;index"`
	CreatedAt   time.Time                        `gorm:"column:created_at;type:datetime;not null;index"`
	UpdatedAt   time.Time                        `gorm:"column:updated_at;type:datetime;not null"`
}

// TableName 指定表名
func (ProductCatalogModel) TableName() string {
	return "product_catalog"
}

// ══════════════════════════════════════════════════════════════════════════════
// 场景 6：深度嵌套（3 层）- Contact/Person/Enterprise
// ══════════════════════════════════════════════════════════════════════════════

// ContactDetailPB 联系方式详情 Protobuf 结构
type ContactDetailPB struct {
	Email       string
	PhoneNumber string
	WechatId    string
	QqNumber    string
}

// ContactDetailModel 联系方式详情数据库模型
type ContactDetailModel struct {
	Email       string `gorm:"column:email;type:varchar(200)"`
	PhoneNumber string `gorm:"column:phone_number;type:varchar(20)"`
	WechatID    string `pbmo:"WechatId" gorm:"column:wechat_id;type:varchar(50)"`
	QqNumber    string `gorm:"column:qq_number;type:varchar(20)"`
}

// PersonInfoPB 个人信息 Protobuf 结构（包含联系方式）
type PersonInfoPB struct {
	PersonName string
	PersonAge  int32
	PersonRole string
	Contact    *ContactDetailPB
}

// PersonInfoModel 个人信息数据库模型
type PersonInfoModel struct {
	PersonName string              `gorm:"column:person_name;type:varchar(100)"`
	PersonAge  int                 `gorm:"column:person_age;type:int"`
	PersonRole string              `gorm:"column:person_role;type:varchar(50)"`
	Contact    *ContactDetailModel `gorm:"embedded;embeddedPrefix:contact_"`
}

// EnterpriseInfoPB 企业信息 Protobuf 结构（3 层嵌套）
type EnterpriseInfoPB struct {
	EnterpriseId   string
	EnterpriseName string
	EnterpriseCode string
	Industry       string
	Status         int32
	Location       *LocationInfoPB
	Manager        *PersonInfoPB
	CreatedAt      *timestamppb.Timestamp
	UpdatedAt      *timestamppb.Timestamp
}

// EnterpriseInfoModel 企业信息数据库模型
type EnterpriseInfoModel struct {
	EnterpriseID   string             `pbmo:"EnterpriseId" gorm:"column:enterprise_id;type:varchar(64);primaryKey"`
	EnterpriseName string             `gorm:"column:enterprise_name;type:varchar(200);index;not null"`
	EnterpriseCode string             `gorm:"column:enterprise_code;type:varchar(50);uniqueIndex;not null"`
	Industry       string             `gorm:"column:industry;type:varchar(50);index"`
	Status         int                `gorm:"column:status;type:int;default:1;index"`
	Location       *LocationInfoModel `gorm:"embedded;embeddedPrefix:location_"`
	Manager        *PersonInfoModel   `gorm:"embedded;embeddedPrefix:manager_"`
	CreatedAt      time.Time          `gorm:"column:created_at;type:datetime;not null;index"`
	UpdatedAt      time.Time          `gorm:"column:updated_at;type:datetime;not null"`
}

// TableName 指定表名
func (EnterpriseInfoModel) TableName() string {
	return "enterprise_info"
}
