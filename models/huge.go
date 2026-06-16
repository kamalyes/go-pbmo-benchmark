/**
 * @Author: developer 500000000@example.com
 * @Date: 2023-09-16 18:28:17
 * @LastEditors: developer 500000000@example.com
 * @LastEditTime: 2025-06-12 19:55:32
 * @FilePath: \go-pbmo-benchmark\models\huge.go
 * @Description: 超大字段模型 - 30/50/60/70/80/90/100 字段测试
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
// 场景 1：30 字段模型 - 企业完整档案
// ══════════════════════════════════════════════════════════════════════════════

// Huge30PB 30字段 Protobuf 结构
type Huge30PB struct {
	// 基础信息 (1-10)
	Id          string
	Code        string
	Name        string
	ShortName   string
	EnglishName string
	Status      int32
	Type        int32
	Category    int32
	Level       int32
	Priority    int32
	// 联系信息 (11-20)
	ContactPerson string
	ContactPhone  string
	ContactEmail  string
	Website       string
	Fax           string
	Address       string
	Country       string
	Province      string
	City          string
	District      string
	// 业务信息 (21-30)
	Industry        string
	BusinessScope   string
	RegisterDate    string
	LegalPerson     string
	RegisterCapital float64
	EmployeeCount   int32
	Tags            []string
	Settings        map[string]any
	CreatedAt       *timestamppb.Timestamp
	UpdatedAt       *timestamppb.Timestamp
}

// Huge30Model 30字段数据库模型
type Huge30Model struct {
	ID              string          `pbmo:"Id" gorm:"column:id;type:varchar(64);primaryKey"`
	Code            string          `gorm:"column:code;type:varchar(50);uniqueIndex;not null"`
	Name            string          `gorm:"column:name;type:varchar(200);index;not null"`
	ShortName       string          `gorm:"column:short_name;type:varchar(100);index"`
	EnglishName     string          `gorm:"column:english_name;type:varchar(200)"`
	Status          int             `gorm:"column:status;type:int;default:1;index"`
	Type            int             `gorm:"column:type;type:int;index"`
	Category        int             `gorm:"column:category;type:int;index"`
	Level           int             `gorm:"column:level;type:int;default:1"`
	Priority        int             `gorm:"column:priority;type:int;default:0;index"`
	ContactPerson   string          `gorm:"column:contact_person;type:varchar(100)"`
	ContactPhone    string          `gorm:"column:contact_phone;type:varchar(20);index"`
	ContactEmail    string          `gorm:"column:contact_email;type:varchar(200)"`
	Website         string          `gorm:"column:website;type:varchar(500)"`
	Fax             string          `gorm:"column:fax;type:varchar(20)"`
	Address         string          `gorm:"column:address;type:varchar(500)"`
	Country         string          `gorm:"column:country;type:varchar(50);index"`
	Province        string          `gorm:"column:province;type:varchar(50);index"`
	City            string          `gorm:"column:city;type:varchar(50);index"`
	District        string          `gorm:"column:district;type:varchar(50)"`
	Industry        string          `gorm:"column:industry;type:varchar(100);index"`
	BusinessScope   string          `gorm:"column:business_scope;type:text"`
	RegisterDate    string          `gorm:"column:register_date;type:varchar(10)"`
	LegalPerson     string          `gorm:"column:legal_person;type:varchar(100)"`
	RegisterCapital float64         `gorm:"column:register_capital;type:decimal(18,2)"`
	EmployeeCount   int             `gorm:"column:employee_count;type:int;default:0"`
	Tags            []string        `gorm:"column:tags;type:json"`
	Settings        sqltypes.MapAny `gorm:"column:settings;type:json"`
	CreatedAt       time.Time       `gorm:"column:created_at;type:datetime;not null;index"`
	UpdatedAt       time.Time       `gorm:"column:updated_at;type:datetime;not null"`
}

func (Huge30Model) TableName() string { return "huge_30" }

// ══════════════════════════════════════════════════════════════════════════════
// 场景 2：50 字段模型 - 超级用户档案
// ══════════════════════════════════════════════════════════════════════════════

// Huge50PB 50字段 Protobuf 结构
type Huge50PB struct {
	// 基础信息 (1-15)
	Id, Username, Nickname, RealName, IdCard, Email, Phone, Mobile, Qq, Wechat, Avatar, Gender string
	BirthDate                                                                                  string
	Age, Status, Level                                                                         int32
	// 账户信息 (16-30)
	AccountType, VipLevel, Points, Score, Balance, TotalRecharge, TotalConsume                                   int64
	CreditScore, SecurityLevel                                                                                   int32
	LastLoginTime, LastLoginIp, RegisterTime, RegisterIp, RegisterDevice, RegisterSource, ReferrerId, InviteCode string
	// 业务信息 (31-45)
	Country, Province, City, District, Address, ZipCode, Company, Department, Position, Industry, Education string
	Hobby, Signature, Introduction                                                                          string
	Tags                                                                                                    []string
	// 系统信息 (46-50)
	Settings  map[string]any
	Enabled   *wrapperspb.BoolValue
	Verified  *wrapperspb.BoolValue
	CreatedAt *timestamppb.Timestamp
	UpdatedAt *timestamppb.Timestamp
}

// Huge50Model 50字段数据库模型
type Huge50Model struct {
	ID             string          `pbmo:"Id" gorm:"column:id;type:varchar(64);primaryKey"`
	Username       string          `gorm:"column:username;type:varchar(100);uniqueIndex;not null"`
	Nickname       string          `gorm:"column:nickname;type:varchar(100);index"`
	RealName       string          `gorm:"column:real_name;type:varchar(100)"`
	IdCard         string          `gorm:"column:id_card;type:varchar(20);index"`
	Email          string          `gorm:"column:email;type:varchar(200);index"`
	Phone          string          `gorm:"column:phone;type:varchar(20);index"`
	Mobile         string          `gorm:"column:mobile;type:varchar(20);index"`
	Qq             string          `gorm:"column:qq;type:varchar(20)"`
	Wechat         string          `gorm:"column:wechat;type:varchar(50)"`
	Avatar         string          `gorm:"column:avatar;type:varchar(500)"`
	Gender         string          `gorm:"column:gender;type:varchar(10)"`
	BirthDate      string          `gorm:"column:birth_date;type:varchar(10)"`
	Age            int             `gorm:"column:age;type:int"`
	Status         int             `gorm:"column:status;type:int;default:1;index"`
	Level          int             `gorm:"column:level;type:int;default:1;index"`
	AccountType    int64           `gorm:"column:account_type;type:bigint;index"`
	VipLevel       int64           `gorm:"column:vip_level;type:bigint;default:0;index"`
	Points         int64           `gorm:"column:points;type:bigint;default:0"`
	Score          int64           `gorm:"column:score;type:bigint;default:0"`
	Balance        int64           `gorm:"column:balance;type:bigint;default:0"`
	TotalRecharge  int64           `gorm:"column:total_recharge;type:bigint;default:0"`
	TotalConsume   int64           `gorm:"column:total_consume;type:bigint;default:0"`
	CreditScore    int             `gorm:"column:credit_score;type:int;default:0"`
	SecurityLevel  int             `gorm:"column:security_level;type:int;default:1"`
	LastLoginTime  string          `gorm:"column:last_login_time;type:varchar(20);index"`
	LastLoginIp    string          `gorm:"column:last_login_ip;type:varchar(50)"`
	RegisterTime   string          `gorm:"column:register_time;type:varchar(20);not null;index"`
	RegisterIp     string          `gorm:"column:register_ip;type:varchar(50)"`
	RegisterDevice string          `gorm:"column:register_device;type:varchar(100)"`
	RegisterSource string          `gorm:"column:register_source;type:varchar(50);index"`
	ReferrerID     string          `pbmo:"ReferrerId" gorm:"column:referrer_id;type:varchar(64);index"`
	InviteCode     string          `gorm:"column:invite_code;type:varchar(20);uniqueIndex"`
	Country        string          `gorm:"column:country;type:varchar(50);index"`
	Province       string          `gorm:"column:province;type:varchar(50);index"`
	City           string          `gorm:"column:city;type:varchar(50);index"`
	District       string          `gorm:"column:district;type:varchar(50)"`
	Address        string          `gorm:"column:address;type:varchar(500)"`
	ZipCode        string          `gorm:"column:zip_code;type:varchar(20)"`
	Company        string          `gorm:"column:company;type:varchar(200)"`
	Department     string          `gorm:"column:department;type:varchar(100)"`
	Position       string          `gorm:"column:position;type:varchar(100)"`
	Industry       string          `gorm:"column:industry;type:varchar(100);index"`
	Education      string          `gorm:"column:education;type:varchar(50)"`
	Hobby          string          `gorm:"column:hobby;type:varchar(500)"`
	Signature      string          `gorm:"column:signature;type:varchar(500)"`
	Introduction   string          `gorm:"column:introduction;type:text"`
	Tags           []string        `gorm:"column:tags;type:json"`
	Settings       sqltypes.MapAny `gorm:"column:settings;type:json"`
	Enabled        *bool           `gorm:"column:enabled;type:tinyint(1)"`
	Verified       *bool           `gorm:"column:verified;type:tinyint(1)"`
	CreatedAt      time.Time       `gorm:"column:created_at;type:datetime;not null;index"`
	UpdatedAt      time.Time       `gorm:"column:updated_at;type:datetime;not null"`
}

func (Huge50Model) TableName() string { return "huge_50" }

// ══════════════════════════════════════════════════════════════════════════════
// 场景 3：100 字段模型 - 终极完整档案
// ══════════════════════════════════════════════════════════════════════════════

// Huge100PB 100字段 Protobuf 结构（包含所有可能的业务字段）
type Huge100PB struct {
	// 基础信息区 (1-20)
	F01, F02, F03, F04, F05, F06, F07, F08, F09, F10 string
	F11, F12, F13, F14, F15, F16, F17, F18, F19, F20 int32
	// 账户信息区 (21-40)
	F21, F22, F23, F24, F25, F26, F27, F28, F29, F30 int64
	F31, F32, F33, F34, F35, F36, F37, F38, F39, F40 float64
	// 业务信息区 (41-60)
	F41, F42, F43, F44, F45, F46, F47, F48, F49, F50 string
	F51, F52, F53, F54, F55, F56, F57, F58, F59, F60 int32
	// 扩展信息区 (61-80)
	F61, F62, F63, F64, F65, F66, F67, F68, F69, F70 string
	F71, F72, F73, F74, F75                          int64
	F76, F77, F78, F79, F80                          bool
	// 系统信息区 (81-100)
	F81, F82, F83 []string
	F84, F85, F86 map[string]any
	F87           *wrapperspb.Int32Value
	F88           *wrapperspb.Int64Value
	F89           *wrapperspb.FloatValue
	F90           *wrapperspb.DoubleValue
	F91           *wrapperspb.BoolValue
	F92           *wrapperspb.StringValue
	F93           *wrapperspb.BytesValue
	F94           *timestamppb.Timestamp
	F95           *timestamppb.Timestamp
	F96           *timestamppb.Timestamp
	F97           *timestamppb.Timestamp
	F98           *timestamppb.Timestamp
	F99           *timestamppb.Timestamp
	F100          *timestamppb.Timestamp
}

// Huge100Model 100字段数据库模型
type Huge100Model struct {
	F01  string          `gorm:"column:f01;type:varchar(64);primaryKey"`
	F02  string          `gorm:"column:f02;type:varchar(100);uniqueIndex"`
	F03  string          `gorm:"column:f03;type:varchar(100);index"`
	F04  string          `gorm:"column:f04;type:varchar(100)"`
	F05  string          `gorm:"column:f05;type:varchar(100)"`
	F06  string          `gorm:"column:f06;type:varchar(200)"`
	F07  string          `gorm:"column:f07;type:varchar(200)"`
	F08  string          `gorm:"column:f08;type:varchar(50)"`
	F09  string          `gorm:"column:f09;type:varchar(50)"`
	F10  string          `gorm:"column:f10;type:varchar(50)"`
	F11  int             `gorm:"column:f11;type:int;default:0;index"`
	F12  int             `gorm:"column:f12;type:int;default:0;index"`
	F13  int             `gorm:"column:f13;type:int;default:0"`
	F14  int             `gorm:"column:f14;type:int;default:0"`
	F15  int             `gorm:"column:f15;type:int;default:0"`
	F16  int             `gorm:"column:f16;type:int;default:0"`
	F17  int             `gorm:"column:f17;type:int;default:0"`
	F18  int             `gorm:"column:f18;type:int;default:0"`
	F19  int             `gorm:"column:f19;type:int;default:0"`
	F20  int             `gorm:"column:f20;type:int;default:0"`
	F21  int64           `gorm:"column:f21;type:bigint;default:0"`
	F22  int64           `gorm:"column:f22;type:bigint;default:0"`
	F23  int64           `gorm:"column:f23;type:bigint;default:0"`
	F24  int64           `gorm:"column:f24;type:bigint;default:0"`
	F25  int64           `gorm:"column:f25;type:bigint;default:0"`
	F26  int64           `gorm:"column:f26;type:bigint;default:0"`
	F27  int64           `gorm:"column:f27;type:bigint;default:0"`
	F28  int64           `gorm:"column:f28;type:bigint;default:0"`
	F29  int64           `gorm:"column:f29;type:bigint;default:0"`
	F30  int64           `gorm:"column:f30;type:bigint;default:0"`
	F31  float64         `gorm:"column:f31;type:decimal(18,6);default:0"`
	F32  float64         `gorm:"column:f32;type:decimal(18,6);default:0"`
	F33  float64         `gorm:"column:f33;type:decimal(18,6);default:0"`
	F34  float64         `gorm:"column:f34;type:decimal(18,6);default:0"`
	F35  float64         `gorm:"column:f35;type:decimal(18,6);default:0"`
	F36  float64         `gorm:"column:f36;type:decimal(18,6);default:0"`
	F37  float64         `gorm:"column:f37;type:decimal(18,6);default:0"`
	F38  float64         `gorm:"column:f38;type:decimal(18,6);default:0"`
	F39  float64         `gorm:"column:f39;type:decimal(18,6);default:0"`
	F40  float64         `gorm:"column:f40;type:decimal(18,6);default:0"`
	F41  string          `gorm:"column:f41;type:varchar(200)"`
	F42  string          `gorm:"column:f42;type:varchar(200)"`
	F43  string          `gorm:"column:f43;type:varchar(200)"`
	F44  string          `gorm:"column:f44;type:varchar(200)"`
	F45  string          `gorm:"column:f45;type:varchar(200)"`
	F46  string          `gorm:"column:f46;type:varchar(200)"`
	F47  string          `gorm:"column:f47;type:varchar(200)"`
	F48  string          `gorm:"column:f48;type:varchar(200)"`
	F49  string          `gorm:"column:f49;type:varchar(200)"`
	F50  string          `gorm:"column:f50;type:varchar(200)"`
	F51  int             `gorm:"column:f51;type:int"`
	F52  int             `gorm:"column:f52;type:int"`
	F53  int             `gorm:"column:f53;type:int"`
	F54  int             `gorm:"column:f54;type:int"`
	F55  int             `gorm:"column:f55;type:int"`
	F56  int             `gorm:"column:f56;type:int"`
	F57  int             `gorm:"column:f57;type:int"`
	F58  int             `gorm:"column:f58;type:int"`
	F59  int             `gorm:"column:f59;type:int"`
	F60  int             `gorm:"column:f60;type:int"`
	F61  string          `gorm:"column:f61;type:text"`
	F62  string          `gorm:"column:f62;type:text"`
	F63  string          `gorm:"column:f63;type:text"`
	F64  string          `gorm:"column:f64;type:text"`
	F65  string          `gorm:"column:f65;type:text"`
	F66  string          `gorm:"column:f66;type:varchar(500)"`
	F67  string          `gorm:"column:f67;type:varchar(500)"`
	F68  string          `gorm:"column:f68;type:varchar(500)"`
	F69  string          `gorm:"column:f69;type:varchar(500)"`
	F70  string          `gorm:"column:f70;type:varchar(500)"`
	F71  int64           `gorm:"column:f71;type:bigint"`
	F72  int64           `gorm:"column:f72;type:bigint"`
	F73  int64           `gorm:"column:f73;type:bigint"`
	F74  int64           `gorm:"column:f74;type:bigint"`
	F75  int64           `gorm:"column:f75;type:bigint"`
	F76  bool            `gorm:"column:f76;type:tinyint(1)"`
	F77  bool            `gorm:"column:f77;type:tinyint(1)"`
	F78  bool            `gorm:"column:f78;type:tinyint(1)"`
	F79  bool            `gorm:"column:f79;type:tinyint(1)"`
	F80  bool            `gorm:"column:f80;type:tinyint(1)"`
	F81  []string        `gorm:"column:f81;type:json"`
	F82  []string        `gorm:"column:f82;type:json"`
	F83  []string        `gorm:"column:f83;type:json"`
	F84  sqltypes.MapAny `gorm:"column:f84;type:json"`
	F85  sqltypes.MapAny `gorm:"column:f85;type:json"`
	F86  sqltypes.MapAny `gorm:"column:f86;type:json"`
	F87  *int32          `gorm:"column:f87;type:int"`
	F88  *int64          `gorm:"column:f88;type:bigint"`
	F89  *float32        `gorm:"column:f89;type:float"`
	F90  *float64        `gorm:"column:f90;type:double"`
	F91  *bool           `gorm:"column:f91;type:tinyint(1)"`
	F92  *string         `gorm:"column:f92;type:varchar(500)"`
	F93  *[]byte         `gorm:"column:f93;type:blob"`
	F94  time.Time       `gorm:"column:f94;type:datetime;index"`
	F95  time.Time       `gorm:"column:f95;type:datetime;index"`
	F96  time.Time       `gorm:"column:f96;type:datetime"`
	F97  time.Time       `gorm:"column:f97;type:datetime"`
	F98  time.Time       `gorm:"column:f98;type:datetime"`
	F99  time.Time       `gorm:"column:f99;type:datetime;not null"`
	F100 time.Time       `gorm:"column:f100;type:datetime;not null"`
}

func (Huge100Model) TableName() string { return "huge_100" }
