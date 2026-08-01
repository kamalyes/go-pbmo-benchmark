/**
 * @Author: developer 500000000@example.com
 * @Date: 2023-09-16 23:10:25
 * @LastEditors: developer 500000000@example.com
 * @LastEditTime: 2025-06-12 23:35:17
 * @FilePath: \go-pbmo-benchmark\benchmarks\huge_bench_test.go
 * @Description: 超大字段模型性能测试 - 30/50/100 字段测试
 *
 * Copyright (c) 2025 by developer, All Rights Reserved.
 */

package benchmarks

import (
	"fmt"
	"testing"
	"time"

	"github.com/kamalyes/go-pbmo"
	"github.com/kamalyes/go-pbmo-benchmark/models"
	sqltypes "github.com/kamalyes/go-sqlbuilder/types"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var (
	benchHuge30ModelResult  *models.Huge30Model
	benchHuge30PBResult     *models.Huge30PB
	benchHuge50ModelResult  *models.Huge50Model
	benchHuge50PBResult     *models.Huge50PB
	benchHuge100ModelResult *models.Huge100Model
	benchHuge100PBResult    *models.Huge100PB
)

func init() {
	pbmo.Register[models.Huge30PB, models.Huge30Model]()
	pbmo.Register[models.Huge50PB, models.Huge50Model]()
	pbmo.Register[models.Huge100PB, models.Huge100Model]()
}

// ══════════════════════════════════════════════════════════════════════════════
// Huge30PB Benchmarks (30 字段) - PBMO vs Native
// ══════════════════════════════════════════════════════════════════════════════

func BenchmarkPBMO_Huge30_PBToModel(b *testing.B) {
	now := time.Now()
	pb := models.Huge30PB{
		Id:              "ent001",
		Code:            "ENT001",
		Name:            "Test Enterprise",
		ShortName:       "TE",
		EnglishName:     "Test Enterprise Ltd",
		Status:          int32(1),
		Type:            int32(1),
		Category:        int32(1),
		Level:           int32(1),
		Priority:        int32(5),
		ContactPerson:   "Zhang San",
		ContactPhone:    "13800138000",
		ContactEmail:    "contact@example.com",
		Website:         "https://example.com",
		Fax:             "0755-12345678",
		Address:         "High-tech Park, Nanshan",
		Country:         "China",
		Province:        "Guangdong",
		City:            "Shenzhen",
		District:        "Nanshan",
		Industry:        "IT",
		BusinessScope:   "Software Development",
		RegisterDate:    "2020-01-01",
		LegalPerson:     "Zhang San",
		RegisterCapital: 10000000.00,
		EmployeeCount:   int32(500),
		Tags:            []string{"tech", "enterprise", "verified"},
		Settings:        map[string]any{"key1": "val1", "key2": int64(42)},
		CreatedAt:       timestamppb.New(now),
		UpdatedAt:       timestamppb.New(now),
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.FromPB[models.Huge30PB, models.Huge30Model](&pb)
	}
}

func BenchmarkNative_Huge30_PBToModel(b *testing.B) {
	now := time.Now()
	pb := models.Huge30PB{
		Id:              "ent001",
		Code:            "ENT001",
		Name:            "Test Enterprise",
		ShortName:       "TE",
		EnglishName:     "Test Enterprise Ltd",
		Status:          int32(1),
		Type:            int32(1),
		Category:        int32(1),
		Level:           int32(1),
		Priority:        int32(5),
		ContactPerson:   "Zhang San",
		ContactPhone:    "13800138000",
		ContactEmail:    "contact@example.com",
		Website:         "https://example.com",
		Fax:             "0755-12345678",
		Address:         "High-tech Park, Nanshan",
		Country:         "China",
		Province:        "Guangdong",
		City:            "Shenzhen",
		District:        "Nanshan",
		Industry:        "IT",
		BusinessScope:   "Software Development",
		RegisterDate:    "2020-01-01",
		LegalPerson:     "Zhang San",
		RegisterCapital: 10000000.00,
		EmployeeCount:   int32(500),
		Tags:            []string{"tech", "enterprise", "verified"},
		Settings:        map[string]any{"key1": "val1", "key2": int64(42)},
		CreatedAt:       timestamppb.New(now),
		UpdatedAt:       timestamppb.New(now),
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.Huge30Model
	for i := 0; i < b.N; i++ {
		result = &models.Huge30Model{
			ID:              pb.Id,
			Code:            pb.Code,
			Name:            pb.Name,
			ShortName:       pb.ShortName,
			EnglishName:     pb.EnglishName,
			Status:          int(pb.Status),
			Type:            int(pb.Type),
			Category:        int(pb.Category),
			Level:           int(pb.Level),
			Priority:        int(pb.Priority),
			ContactPerson:   pb.ContactPerson,
			ContactPhone:    pb.ContactPhone,
			ContactEmail:    pb.ContactEmail,
			Website:         pb.Website,
			Fax:             pb.Fax,
			Address:         pb.Address,
			Country:         pb.Country,
			Province:        pb.Province,
			City:            pb.City,
			District:        pb.District,
			Industry:        pb.Industry,
			BusinessScope:   pb.BusinessScope,
			RegisterDate:    pb.RegisterDate,
			LegalPerson:     pb.LegalPerson,
			RegisterCapital: pb.RegisterCapital,
			EmployeeCount:   int(pb.EmployeeCount),
			Tags:            pb.Tags,
			Settings:        sqltypes.MapAny(pb.Settings),
			CreatedAt:       pb.CreatedAt.AsTime(),
			UpdatedAt:       pb.UpdatedAt.AsTime(),
		}
	}
	benchHuge30ModelResult = result
}

func BenchmarkPBMO_Huge30_ModelToPB(b *testing.B) {
	now := time.Now()
	model := models.Huge30Model{
		ID:              "ent001",
		Code:            "ENT001",
		Name:            "Test Enterprise",
		ShortName:       "TE",
		EnglishName:     "Test Enterprise Ltd",
		Status:          1,
		Type:            1,
		Category:        1,
		Level:           1,
		Priority:        5,
		ContactPerson:   "Zhang San",
		ContactPhone:    "13800138000",
		ContactEmail:    "contact@example.com",
		Website:         "https://example.com",
		Fax:             "0755-12345678",
		Address:         "High-tech Park, Nanshan",
		Country:         "China",
		Province:        "Guangdong",
		City:            "Shenzhen",
		District:        "Nanshan",
		Industry:        "IT",
		BusinessScope:   "Software Development",
		RegisterDate:    "2020-01-01",
		LegalPerson:     "Zhang San",
		RegisterCapital: 10000000.00,
		EmployeeCount:   500,
		Tags:            []string{"tech", "enterprise", "verified"},
		Settings:        sqltypes.MapAny{"key1": "val1", "key2": int64(42)},
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.ToPB[models.Huge30Model, models.Huge30PB](&model)
	}
}

func BenchmarkNative_Huge30_ModelToPB(b *testing.B) {
	now := time.Now()
	model := models.Huge30Model{
		ID:              "ent001",
		Code:            "ENT001",
		Name:            "Test Enterprise",
		ShortName:       "TE",
		EnglishName:     "Test Enterprise Ltd",
		Status:          1,
		Type:            1,
		Category:        1,
		Level:           1,
		Priority:        5,
		ContactPerson:   "Zhang San",
		ContactPhone:    "13800138000",
		ContactEmail:    "contact@example.com",
		Website:         "https://example.com",
		Fax:             "0755-12345678",
		Address:         "High-tech Park, Nanshan",
		Country:         "China",
		Province:        "Guangdong",
		City:            "Shenzhen",
		District:        "Nanshan",
		Industry:        "IT",
		BusinessScope:   "Software Development",
		RegisterDate:    "2020-01-01",
		LegalPerson:     "Zhang San",
		RegisterCapital: 10000000.00,
		EmployeeCount:   500,
		Tags:            []string{"tech", "enterprise", "verified"},
		Settings:        sqltypes.MapAny{"key1": "val1", "key2": int64(42)},
		CreatedAt:       now,
		UpdatedAt:       now,
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.Huge30PB
	for i := 0; i < b.N; i++ {
		result = &models.Huge30PB{
			Id:              model.ID,
			Code:            model.Code,
			Name:            model.Name,
			ShortName:       model.ShortName,
			EnglishName:     model.EnglishName,
			Status:          int32(model.Status),
			Type:            int32(model.Type),
			Category:        int32(model.Category),
			Level:           int32(model.Level),
			Priority:        int32(model.Priority),
			ContactPerson:   model.ContactPerson,
			ContactPhone:    model.ContactPhone,
			ContactEmail:    model.ContactEmail,
			Website:         model.Website,
			Fax:             model.Fax,
			Address:         model.Address,
			Country:         model.Country,
			Province:        model.Province,
			City:            model.City,
			District:        model.District,
			Industry:        model.Industry,
			BusinessScope:   model.BusinessScope,
			RegisterDate:    model.RegisterDate,
			LegalPerson:     model.LegalPerson,
			RegisterCapital: model.RegisterCapital,
			EmployeeCount:   int32(model.EmployeeCount),
			Tags:            model.Tags,
			Settings:        map[string]any(model.Settings),
			CreatedAt:       timestamppb.New(model.CreatedAt),
			UpdatedAt:       timestamppb.New(model.UpdatedAt),
		}
	}
	benchHuge30PBResult = result
}

// ══════════════════════════════════════════════════════════════════════════════
// Huge50PB Benchmarks (50 字段) - PBMO vs Native
// ══════════════════════════════════════════════════════════════════════════════

func BenchmarkPBMO_Huge50_PBToModel(b *testing.B) {
	now := time.Now()
	enabled := true
	pb := newHuge50PB(now, enabled)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.FromPB[models.Huge50PB, models.Huge50Model](&pb)
	}
}

func BenchmarkNative_Huge50_PBToModel(b *testing.B) {
	now := time.Now()
	enabled := true
	pb := newHuge50PB(now, enabled)

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.Huge50Model
	for i := 0; i < b.N; i++ {
		var enabledPtr *bool
		var verifiedPtr *bool
		if pb.Enabled != nil {
			v := pb.Enabled.Value
			enabledPtr = &v
		}
		if pb.Verified != nil {
			v := pb.Verified.Value
			verifiedPtr = &v
		}
		result = &models.Huge50Model{
			ID:             pb.Id,
			Username:       pb.Username,
			Nickname:       pb.Nickname,
			RealName:       pb.RealName,
			IdCard:         pb.IdCard,
			Email:          pb.Email,
			Phone:          pb.Phone,
			Mobile:         pb.Mobile,
			Qq:             pb.Qq,
			Wechat:         pb.Wechat,
			Avatar:         pb.Avatar,
			Gender:         pb.Gender,
			BirthDate:      pb.BirthDate,
			Age:            int(pb.Age),
			Status:         int(pb.Status),
			Level:          int(pb.Level),
			AccountType:    pb.AccountType,
			VipLevel:       pb.VipLevel,
			Points:         pb.Points,
			Score:          pb.Score,
			Balance:        pb.Balance,
			TotalRecharge:  pb.TotalRecharge,
			TotalConsume:   pb.TotalConsume,
			CreditScore:    int(pb.CreditScore),
			SecurityLevel:  int(pb.SecurityLevel),
			LastLoginTime:  pb.LastLoginTime,
			LastLoginIp:    pb.LastLoginIp,
			RegisterTime:   pb.RegisterTime,
			RegisterIp:     pb.RegisterIp,
			RegisterDevice: pb.RegisterDevice,
			RegisterSource: pb.RegisterSource,
			ReferrerID:     pb.ReferrerId,
			InviteCode:     pb.InviteCode,
			Country:        pb.Country,
			Province:       pb.Province,
			City:           pb.City,
			District:       pb.District,
			Address:        pb.Address,
			ZipCode:        pb.ZipCode,
			Company:        pb.Company,
			Department:     pb.Department,
			Position:       pb.Position,
			Industry:       pb.Industry,
			Education:      pb.Education,
			Hobby:          pb.Hobby,
			Signature:      pb.Signature,
			Introduction:   pb.Introduction,
			Tags:           pb.Tags,
			Settings:       sqltypes.MapAny(pb.Settings),
			Enabled:        enabledPtr,
			Verified:       verifiedPtr,
			CreatedAt:      pb.CreatedAt.AsTime(),
			UpdatedAt:      pb.UpdatedAt.AsTime(),
		}
	}
	benchHuge50ModelResult = result
}

func BenchmarkPBMO_Huge50_ModelToPB(b *testing.B) {
	now := time.Now()
	model := newHuge50Model(now)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.ToPB[models.Huge50Model, models.Huge50PB](&model)
	}
}

func BenchmarkNative_Huge50_ModelToPB(b *testing.B) {
	now := time.Now()
	model := newHuge50Model(now)

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.Huge50PB
	for i := 0; i < b.N; i++ {
		result = &models.Huge50PB{
			Id:             model.ID,
			Username:       model.Username,
			Nickname:       model.Nickname,
			RealName:       model.RealName,
			IdCard:         model.IdCard,
			Email:          model.Email,
			Phone:          model.Phone,
			Mobile:         model.Mobile,
			Qq:             model.Qq,
			Wechat:         model.Wechat,
			Avatar:         model.Avatar,
			Gender:         model.Gender,
			BirthDate:      model.BirthDate,
			Age:            int32(model.Age),
			Status:         int32(model.Status),
			Level:          int32(model.Level),
			AccountType:    model.AccountType,
			VipLevel:       model.VipLevel,
			Points:         model.Points,
			Score:          model.Score,
			Balance:        model.Balance,
			TotalRecharge:  model.TotalRecharge,
			TotalConsume:   model.TotalConsume,
			CreditScore:    int32(model.CreditScore),
			SecurityLevel:  int32(model.SecurityLevel),
			LastLoginTime:  model.LastLoginTime,
			LastLoginIp:    model.LastLoginIp,
			RegisterTime:   model.RegisterTime,
			RegisterIp:     model.RegisterIp,
			RegisterDevice: model.RegisterDevice,
			RegisterSource: model.RegisterSource,
			ReferrerId:     model.ReferrerID,
			InviteCode:     model.InviteCode,
			Country:        model.Country,
			Province:       model.Province,
			City:           model.City,
			District:       model.District,
			Address:        model.Address,
			ZipCode:        model.ZipCode,
			Company:        model.Company,
			Department:     model.Department,
			Position:       model.Position,
			Industry:       model.Industry,
			Education:      model.Education,
			Hobby:          model.Hobby,
			Signature:      model.Signature,
			Introduction:   model.Introduction,
			Tags:           model.Tags,
			Settings:       map[string]any(model.Settings),
			CreatedAt:      timestamppb.New(model.CreatedAt),
			UpdatedAt:      timestamppb.New(model.UpdatedAt),
		}
		if model.Enabled != nil {
			result.Enabled = wrapperspb.Bool(*model.Enabled)
		}
		if model.Verified != nil {
			result.Verified = wrapperspb.Bool(*model.Verified)
		}
	}
	benchHuge50PBResult = result
}

func newHuge50PB(now time.Time, enabled bool) models.Huge50PB {
	return models.Huge50PB{
		Id:             "user50_001",
		Username:       "testuser50",
		Nickname:       "Test50",
		RealName:       "Zhang San",
		IdCard:         "440305199001011234",
		Email:          "test50@example.com",
		Phone:          "13800138000",
		Mobile:         "13900139000",
		Qq:             "123456789",
		Wechat:         "wx_test50",
		Avatar:         "https://example.com/avatar.jpg",
		Gender:         "male",
		BirthDate:      "1990-01-01",
		Age:            int32(35),
		Status:         int32(1),
		Level:          int32(5),
		AccountType:    int64(1),
		VipLevel:       int64(3),
		Points:         int64(10000),
		Score:          int64(8500),
		Balance:        int64(5000),
		TotalRecharge:  int64(100000),
		TotalConsume:   int64(80000),
		CreditScore:    int32(800),
		SecurityLevel:  int32(3),
		LastLoginTime:  fmt.Sprintf("%d", now.Unix()),
		LastLoginIp:    "192.168.1.1",
		RegisterTime:   fmt.Sprintf("%d", now.AddDate(-1, 0, 0).Unix()),
		RegisterIp:     "192.168.1.100",
		RegisterDevice: "iPhone",
		RegisterSource: "web",
		ReferrerId:     "ref001",
		InviteCode:     "INV001",
		Country:        "China",
		Province:       "Guangdong",
		City:           "Shenzhen",
		District:       "Nanshan",
		Address:        "High-tech Park",
		ZipCode:        "518000",
		Company:        "Test Inc",
		Department:     "Engineering",
		Position:       "Developer",
		Industry:       "IT",
		Education:      "Bachelor",
		Hobby:          "coding",
		Signature:      "Hello World",
		Introduction:   "A software engineer",
		Tags:           []string{"vip", "verified"},
		Settings:       map[string]any{"theme": "dark"},
		Enabled:        wrapperspb.Bool(enabled),
		Verified:       wrapperspb.Bool(true),
		CreatedAt:      timestamppb.New(now),
		UpdatedAt:      timestamppb.New(now),
	}
}

func newHuge50Model(now time.Time) models.Huge50Model {
	enabled := true
	verified := true
	return models.Huge50Model{
		ID:             "user50_001",
		Username:       "testuser50",
		Nickname:       "Test50",
		RealName:       "Zhang San",
		IdCard:         "440305199001011234",
		Email:          "test50@example.com",
		Phone:          "13800138000",
		Mobile:         "13900139000",
		Qq:             "123456789",
		Wechat:         "wx_test50",
		Avatar:         "https://example.com/avatar.jpg",
		Gender:         "male",
		BirthDate:      "1990-01-01",
		Age:            35,
		Status:         1,
		Level:          5,
		AccountType:    int64(1),
		VipLevel:       int64(3),
		Points:         int64(10000),
		Score:          int64(8500),
		Balance:        int64(5000),
		TotalRecharge:  int64(100000),
		TotalConsume:   int64(80000),
		CreditScore:    800,
		SecurityLevel:  3,
		LastLoginTime:  fmt.Sprintf("%d", now.Unix()),
		LastLoginIp:    "192.168.1.1",
		RegisterTime:   fmt.Sprintf("%d", now.Unix()),
		RegisterIp:     "192.168.1.100",
		RegisterDevice: "iPhone",
		RegisterSource: "web",
		ReferrerID:     "ref001",
		InviteCode:     "INV001",
		Country:        "China",
		Province:       "Guangdong",
		City:           "Shenzhen",
		District:       "Nanshan",
		Address:        "High-tech Park",
		ZipCode:        "518000",
		Company:        "Test Inc",
		Department:     "Engineering",
		Position:       "Developer",
		Industry:       "IT",
		Education:      "Bachelor",
		Hobby:          "coding",
		Signature:      "Hello World",
		Introduction:   "A software engineer",
		Tags:           []string{"vip", "verified"},
		Settings:       sqltypes.MapAny{"theme": "dark"},
		Enabled:        &enabled,
		Verified:       &verified,
		CreatedAt:      now,
		UpdatedAt:      now,
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// Huge100PB Benchmarks (100 字段) - PBMO vs Native
// ══════════════════════════════════════════════════════════════════════════════

func BenchmarkPBMO_Huge100_PBToModel(b *testing.B) {
	now := time.Now()
	pb := newHuge100PB(now)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.FromPB[models.Huge100PB, models.Huge100Model](&pb)
	}
}

func BenchmarkNative_Huge100_PBToModel(b *testing.B) {
	now := time.Now()
	pb := newHuge100PB(now)

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.Huge100Model
	for i := 0; i < b.N; i++ {
		m := &models.Huge100Model{}
		m.F01 = pb.F01
		m.F02 = pb.F02
		m.F03 = pb.F03
		m.F04 = pb.F04
		m.F05 = pb.F05
		m.F06 = pb.F06
		m.F07 = pb.F07
		m.F08 = pb.F08
		m.F09 = pb.F09
		m.F10 = pb.F10
		m.F11 = int(pb.F11)
		m.F12 = int(pb.F12)
		m.F13 = int(pb.F13)
		m.F14 = int(pb.F14)
		m.F15 = int(pb.F15)
		m.F16 = int(pb.F16)
		m.F17 = int(pb.F17)
		m.F18 = int(pb.F18)
		m.F19 = int(pb.F19)
		m.F20 = int(pb.F20)
		m.F21 = pb.F21
		m.F22 = pb.F22
		m.F23 = pb.F23
		m.F24 = pb.F24
		m.F25 = pb.F25
		m.F26 = pb.F26
		m.F27 = pb.F27
		m.F28 = pb.F28
		m.F29 = pb.F29
		m.F30 = pb.F30
		m.F31 = pb.F31
		m.F32 = pb.F32
		m.F33 = pb.F33
		m.F34 = pb.F34
		m.F35 = pb.F35
		m.F36 = pb.F36
		m.F37 = pb.F37
		m.F38 = pb.F38
		m.F39 = pb.F39
		m.F40 = pb.F40
		m.F41 = pb.F41
		m.F42 = pb.F42
		m.F43 = pb.F43
		m.F44 = pb.F44
		m.F45 = pb.F45
		m.F46 = pb.F46
		m.F47 = pb.F47
		m.F48 = pb.F48
		m.F49 = pb.F49
		m.F50 = pb.F50
		m.F51 = int(pb.F51)
		m.F52 = int(pb.F52)
		m.F53 = int(pb.F53)
		m.F54 = int(pb.F54)
		m.F55 = int(pb.F55)
		m.F56 = int(pb.F56)
		m.F57 = int(pb.F57)
		m.F58 = int(pb.F58)
		m.F59 = int(pb.F59)
		m.F60 = int(pb.F60)
		m.F61 = pb.F61
		m.F62 = pb.F62
		m.F63 = pb.F63
		m.F64 = pb.F64
		m.F65 = pb.F65
		m.F66 = pb.F66
		m.F67 = pb.F67
		m.F68 = pb.F68
		m.F69 = pb.F69
		m.F70 = pb.F70
		m.F71 = pb.F71
		m.F72 = pb.F72
		m.F73 = pb.F73
		m.F74 = pb.F74
		m.F75 = pb.F75
		m.F76 = pb.F76
		m.F77 = pb.F77
		m.F78 = pb.F78
		m.F79 = pb.F79
		m.F80 = pb.F80
		m.F81 = pb.F81
		m.F82 = pb.F82
		m.F83 = pb.F83
		m.F84 = sqltypes.MapAny(pb.F84)
		m.F85 = sqltypes.MapAny(pb.F85)
		m.F86 = sqltypes.MapAny(pb.F86)
		if pb.F87 != nil {
			v := pb.F87.Value
			m.F87 = &v
		}
		if pb.F88 != nil {
			v := pb.F88.Value
			m.F88 = &v
		}
		if pb.F89 != nil {
			v := pb.F89.Value
			m.F89 = &v
		}
		if pb.F90 != nil {
			v := pb.F90.Value
			m.F90 = &v
		}
		if pb.F91 != nil {
			v := pb.F91.Value
			m.F91 = &v
		}
		if pb.F92 != nil {
			v := pb.F92.Value
			m.F92 = &v
		}
		m.F94 = pb.F94.AsTime()
		m.F95 = pb.F95.AsTime()
		if pb.F96 != nil {
			m.F96 = pb.F96.AsTime()
		}
		if pb.F97 != nil {
			m.F97 = pb.F97.AsTime()
		}
		if pb.F98 != nil {
			m.F98 = pb.F98.AsTime()
		}
		m.F99 = pb.F99.AsTime()
		m.F100 = pb.F100.AsTime()
		result = m
	}
	benchHuge100ModelResult = result
}

func BenchmarkPBMO_Huge100_ModelToPB(b *testing.B) {
	now := time.Now()
	model := newHuge100Model(now)

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.ToPB[models.Huge100Model, models.Huge100PB](&model)
	}
}

func BenchmarkNative_Huge100_ModelToPB(b *testing.B) {
	now := time.Now()
	model := newHuge100Model(now)

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.Huge100PB
	for i := 0; i < b.N; i++ {
		pb := &models.Huge100PB{
			F01: model.F01, F02: model.F02, F03: model.F03, F04: model.F04, F05: model.F05,
			F06: model.F06, F07: model.F07, F08: model.F08, F09: model.F09, F10: model.F10,
			F11: int32(model.F11), F12: int32(model.F12), F13: int32(model.F13), F14: int32(model.F14), F15: int32(model.F15),
			F16: int32(model.F16), F17: int32(model.F17), F18: int32(model.F18), F19: int32(model.F19), F20: int32(model.F20),
			F21: model.F21, F22: model.F22, F23: model.F23, F24: model.F24, F25: model.F25,
			F26: model.F26, F27: model.F27, F28: model.F28, F29: model.F29, F30: model.F30,
			F31: model.F31, F32: model.F32, F33: model.F33, F34: model.F34, F35: model.F35,
			F36: model.F36, F37: model.F37, F38: model.F38, F39: model.F39, F40: model.F40,
			F41: model.F41, F42: model.F42, F43: model.F43, F44: model.F44, F45: model.F45,
			F46: model.F46, F47: model.F47, F48: model.F48, F49: model.F49, F50: model.F50,
			F51: int32(model.F51), F52: int32(model.F52), F53: int32(model.F53), F54: int32(model.F54), F55: int32(model.F55),
			F56: int32(model.F56), F57: int32(model.F57), F58: int32(model.F58), F59: int32(model.F59), F60: int32(model.F60),
			F61: model.F61, F62: model.F62, F63: model.F63, F64: model.F64, F65: model.F65,
			F66: model.F66, F67: model.F67, F68: model.F68, F69: model.F69, F70: model.F70,
			F71: model.F71, F72: model.F72, F73: model.F73, F74: model.F74, F75: model.F75,
			F76: model.F76, F77: model.F77, F78: model.F78, F79: model.F79, F80: model.F80,
			F81: model.F81, F82: model.F82, F83: model.F83,
			F84: map[string]any(model.F84), F85: map[string]any(model.F85), F86: map[string]any(model.F86),
			F94: timestamppb.New(model.F94), F95: timestamppb.New(model.F95),
			F99: timestamppb.New(model.F99), F100: timestamppb.New(model.F100),
		}
		if model.F87 != nil {
			pb.F87 = wrapperspb.Int32(*model.F87)
		}
		if model.F88 != nil {
			pb.F88 = wrapperspb.Int64(*model.F88)
		}
		if model.F89 != nil {
			pb.F89 = wrapperspb.Float(*model.F89)
		}
		if model.F90 != nil {
			pb.F90 = wrapperspb.Double(*model.F90)
		}
		if model.F91 != nil {
			pb.F91 = wrapperspb.Bool(*model.F91)
		}
		if model.F92 != nil {
			pb.F92 = wrapperspb.String(*model.F92)
		}
		result = pb
	}
	benchHuge100PBResult = result
}

func newHuge100Model(now time.Time) models.Huge100Model {
	f87 := int32(87)
	f89 := float32(89.1)
	f90 := float64(90.2)
	f91 := true
	f92 := "f92"
	return models.Huge100Model{
		F01: "str01", F02: "str02", F03: "str03", F04: "str04", F05: "str05",
		F06: "str06", F07: "str07", F08: "str08", F09: "str09", F10: "str10",
		F11: 11, F12: 12, F13: 13, F14: 14, F15: 15,
		F16: 16, F17: 17, F18: 18, F19: 19, F20: 20,
		F21: int64(21), F22: int64(22), F23: int64(23), F24: int64(24), F25: int64(25),
		F26: int64(26), F27: int64(27), F28: int64(28), F29: int64(29), F30: int64(30),
		F31: 31.1, F32: 32.2, F33: 33.3, F34: 34.4, F35: 35.5,
		F36: 36.6, F37: 37.7, F38: 38.8, F39: 39.9, F40: 40.0,
		F41: "f41", F42: "f42", F43: "f43", F44: "f44", F45: "f45",
		F46: "f46", F47: "f47", F48: "f48", F49: "f49", F50: "f50",
		F51: 51, F52: 52, F53: 53, F54: 54, F55: 55,
		F56: 56, F57: 57, F58: 58, F59: 59, F60: 60,
		F61: "f61", F62: "f62", F63: "f63", F64: "f64", F65: "f65",
		F66: "f66", F67: "f67", F68: "f68", F69: "f69", F70: "f70",
		F71: int64(71), F72: int64(72), F73: int64(73), F74: int64(74), F75: int64(75),
		F76: true, F77: true, F78: false, F79: false, F80: true,
		F81: []string{"s1", "s2"}, F82: []string{"s3"}, F83: []string{"s4"},
		F84: sqltypes.MapAny{"k1": "v1"}, F85: sqltypes.MapAny{"k2": int64(2)}, F86: sqltypes.MapAny{"k3": true},
		F87: &f87, F89: &f89, F90: &f90, F91: &f91, F92: &f92,
		F99: now, F100: now,
	}
}

func newHuge100PB(now time.Time) models.Huge100PB {
	return models.Huge100PB{
		F01: "str01", F02: "str02", F03: "str03", F04: "str04", F05: "str05",
		F06: "str06", F07: "str07", F08: "str08", F09: "str09", F10: "str10",
		F11: int32(11), F12: int32(12), F13: int32(13), F14: int32(14), F15: int32(15),
		F16: int32(16), F17: int32(17), F18: int32(18), F19: int32(19), F20: int32(20),
		F21: int64(21), F22: int64(22), F23: int64(23), F24: int64(24), F25: int64(25),
		F26: int64(26), F27: int64(27), F28: int64(28), F29: int64(29), F30: int64(30),
		F31: 31.1, F32: 32.2, F33: 33.3, F34: 34.4, F35: 35.5,
		F36: 36.6, F37: 37.7, F38: 38.8, F39: 39.9, F40: 40.0,
		F41: "f41", F42: "f42", F43: "f43", F44: "f44", F45: "f45",
		F46: "f46", F47: "f47", F48: "f48", F49: "f49", F50: "f50",
		F51: int32(51), F52: int32(52), F53: int32(53), F54: int32(54), F55: int32(55),
		F56: int32(56), F57: int32(57), F58: int32(58), F59: int32(59), F60: int32(60),
		F61: "f61", F62: "f62", F63: "f63", F64: "f64", F65: "f65",
		F66: "f66", F67: "f67", F68: "f68", F69: "f69", F70: "f70",
		F71: int64(71), F72: int64(72), F73: int64(73), F74: int64(74), F75: int64(75),
		F76: true, F77: true, F78: false, F79: false, F80: true,
		F81: []string{"s1", "s2"}, F82: []string{"s3"}, F83: []string{"s4"},
		F84: map[string]any{"k1": "v1"}, F85: map[string]any{"k2": int64(2)}, F86: map[string]any{"k3": true},
		F87: wrapperspb.Int32(87), F88: wrapperspb.Int64(88), F89: wrapperspb.Float(89.1),
		F90: wrapperspb.Double(90.2), F91: wrapperspb.Bool(true), F92: wrapperspb.String("f92"),
		F94: timestamppb.New(now), F95: timestamppb.New(now), F99: timestamppb.New(now), F100: timestamppb.New(now),
	}
}
