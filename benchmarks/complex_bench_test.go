/**
 * @Author: developer 500000000@example.com
 * @Date: 2023-09-16 13:15:16
 * @LastEditors: developer 500000000@example.com
 * @LastEditTime: 2025-06-12 20:38:21
 * @FilePath: \go-pbmo-benchmark\benchmarks\complex_bench_test.go
 * @Description: 复杂模型性能测试 - 20-27 字段大模型
 *
 * Copyright (c) 2025 by developer, All Rights Reserved.
 */

package benchmarks

import (
	"testing"
	"time"

	"github.com/kamalyes/go-pbmo"
	"github.com/kamalyes/go-pbmo-benchmark/models"
	sqltypes "github.com/kamalyes/go-sqlbuilder/types"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var (
	benchLargeModelResult        *models.LargeModel
	benchLargePBResult           *models.LargePB
	benchOrganizationModelResult *models.OrganizationModel
	benchOrganizationPBResult    *models.OrganizationPB
	benchUserProfileModelResult  *models.UserProfileModel
	benchUserProfilePBResult     *models.UserProfilePB
	benchLargeModelsResult       []*models.LargeModel
)

func init() {
	pbmo.Register[models.LargePB, models.LargeModel]()
	pbmo.Register[models.OrganizationPB, models.OrganizationModel]()
	pbmo.Register[models.UserProfilePB, models.UserProfileModel]()
}

// ══════════════════════════════════════════════════════════════════════════════
// LargePB Benchmarks (20 字段) - PBMO vs Native
// ══════════════════════════════════════════════════════════════════════════════

func BenchmarkPBMO_Large_PBToModel(b *testing.B) {
	now := time.Now()
	minVal := int32(10)
	maxVal := int32(100)
	pb := models.LargePB{
		Id:          uint64(1),
		Name:        "largetest",
		Email:       "large@example.com",
		Phone:       "13800138000",
		Status:      int32(1),
		Priority:    int32(5),
		Score:       95.5,
		Rating:      4.8,
		Active:      true,
		Verified:    true,
		Tags:        []string{"tag1", "tag2", "tag3"},
		Country:     "China",
		Region:      "Guangdong",
		City:        "Shenzhen",
		ZipCode:     "518000",
		Description: "Large model test",
		CreatedAt:   timestamppb.New(now),
		UpdatedAt:   timestamppb.New(now),
		MinVal:      wrapperspb.Int32(minVal),
		MaxVal:      wrapperspb.Int32(maxVal),
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.FromPB[models.LargePB, models.LargeModel](&pb)
	}
}

func BenchmarkNative_Large_PBToModel(b *testing.B) {
	now := time.Now()
	minVal := int32(10)
	maxVal := int32(100)
	pb := models.LargePB{
		Id:          uint64(1),
		Name:        "largetest",
		Email:       "large@example.com",
		Phone:       "13800138000",
		Status:      int32(1),
		Priority:    int32(5),
		Score:       95.5,
		Rating:      4.8,
		Active:      true,
		Verified:    true,
		Tags:        []string{"tag1", "tag2", "tag3"},
		Country:     "China",
		Region:      "Guangdong",
		City:        "Shenzhen",
		ZipCode:     "518000",
		Description: "Large model test",
		CreatedAt:   timestamppb.New(now),
		UpdatedAt:   timestamppb.New(now),
		MinVal:      wrapperspb.Int32(minVal),
		MaxVal:      wrapperspb.Int32(maxVal),
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.LargeModel
	for i := 0; i < b.N; i++ {
		var minValPtr, maxValPtr *int32
		if pb.MinVal != nil {
			v := pb.MinVal.Value
			minValPtr = &v
		}
		if pb.MaxVal != nil {
			v := pb.MaxVal.Value
			maxValPtr = &v
		}
		result = &models.LargeModel{
			ID:          pb.Id,
			Name:        pb.Name,
			Email:       pb.Email,
			Phone:       pb.Phone,
			Status:      int(pb.Status),
			Priority:    int(pb.Priority),
			Score:       pb.Score,
			Rating:      pb.Rating,
			Active:      pb.Active,
			Verified:    pb.Verified,
			Tags:        pb.Tags,
			Country:     pb.Country,
			Region:      pb.Region,
			City:        pb.City,
			ZipCode:     pb.ZipCode,
			Description: pb.Description,
			CreatedAt:   pb.CreatedAt.AsTime(),
			UpdatedAt:   pb.UpdatedAt.AsTime(),
			MinVal:      minValPtr,
			MaxVal:      maxValPtr,
		}
	}
	benchLargeModelResult = result
}

func BenchmarkPBMO_Large_ModelToPB(b *testing.B) {
	now := time.Now()
	minVal := int32(10)
	maxVal := int32(100)
	model := models.LargeModel{
		ID:          uint64(1),
		Name:        "largetest",
		Email:       "large@example.com",
		Phone:       "13800138000",
		Status:      1,
		Priority:    5,
		Score:       95.5,
		Rating:      4.8,
		Active:      true,
		Verified:    true,
		Tags:        []string{"tag1", "tag2", "tag3"},
		Country:     "China",
		Region:      "Guangdong",
		City:        "Shenzhen",
		ZipCode:     "518000",
		Description: "Large model test",
		CreatedAt:   now,
		UpdatedAt:   now,
		MinVal:      &minVal,
		MaxVal:      &maxVal,
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.ToPB[models.LargeModel, models.LargePB](&model)
	}
}

func BenchmarkNative_Large_ModelToPB(b *testing.B) {
	now := time.Now()
	minVal := int32(10)
	maxVal := int32(100)
	model := models.LargeModel{
		ID:          uint64(1),
		Name:        "largetest",
		Email:       "large@example.com",
		Phone:       "13800138000",
		Status:      1,
		Priority:    5,
		Score:       95.5,
		Rating:      4.8,
		Active:      true,
		Verified:    true,
		Tags:        []string{"tag1", "tag2", "tag3"},
		Country:     "China",
		Region:      "Guangdong",
		City:        "Shenzhen",
		ZipCode:     "518000",
		Description: "Large model test",
		CreatedAt:   now,
		UpdatedAt:   now,
		MinVal:      &minVal,
		MaxVal:      &maxVal,
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.LargePB
	for i := 0; i < b.N; i++ {
		pb := &models.LargePB{
			Id:          model.ID,
			Name:        model.Name,
			Email:       model.Email,
			Phone:       model.Phone,
			Status:      int32(model.Status),
			Priority:    int32(model.Priority),
			Score:       model.Score,
			Rating:      model.Rating,
			Active:      model.Active,
			Verified:    model.Verified,
			Tags:        model.Tags,
			Country:     model.Country,
			Region:      model.Region,
			City:        model.City,
			ZipCode:     model.ZipCode,
			Description: model.Description,
			CreatedAt:   timestamppb.New(model.CreatedAt),
			UpdatedAt:   timestamppb.New(model.UpdatedAt),
		}
		if model.MinVal != nil {
			pb.MinVal = wrapperspb.Int32(*model.MinVal)
		}
		if model.MaxVal != nil {
			pb.MaxVal = wrapperspb.Int32(*model.MaxVal)
		}
		result = pb
	}
	benchLargePBResult = result
}

func BenchmarkPBMO_Large_PBToModel_Parallel(b *testing.B) {
	now := time.Now()
	minVal := int32(10)
	maxVal := int32(100)
	pb := models.LargePB{
		Id:          uint64(1),
		Name:        "largetest",
		Email:       "large@example.com",
		Phone:       "13800138000",
		Status:      int32(1),
		Priority:    int32(5),
		Score:       95.5,
		Rating:      4.8,
		Active:      true,
		Verified:    true,
		Tags:        []string{"tag1", "tag2", "tag3"},
		Country:     "China",
		Region:      "Guangdong",
		City:        "Shenzhen",
		ZipCode:     "518000",
		Description: "Large model test",
		CreatedAt:   timestamppb.New(now),
		UpdatedAt:   timestamppb.New(now),
		MinVal:      wrapperspb.Int32(minVal),
		MaxVal:      wrapperspb.Int32(maxVal),
	}

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		var result *models.LargeModel
		for p.Next() {
			result, _ = pbmo.FromPB[models.LargePB, models.LargeModel](&pb)
		}
		benchLargeModelResult = result
	})
}

func BenchmarkNative_Large_PBToModel_Parallel(b *testing.B) {
	now := time.Now()
	minVal := int32(10)
	maxVal := int32(100)
	pb := models.LargePB{
		Id:          uint64(1),
		Name:        "largetest",
		Email:       "large@example.com",
		Phone:       "13800138000",
		Status:      int32(1),
		Priority:    int32(5),
		Score:       95.5,
		Rating:      4.8,
		Active:      true,
		Verified:    true,
		Tags:        []string{"tag1", "tag2", "tag3"},
		Country:     "China",
		Region:      "Guangdong",
		City:        "Shenzhen",
		ZipCode:     "518000",
		Description: "Large model test",
		CreatedAt:   timestamppb.New(now),
		UpdatedAt:   timestamppb.New(now),
		MinVal:      wrapperspb.Int32(minVal),
		MaxVal:      wrapperspb.Int32(maxVal),
	}

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(tpb *testing.PB) {
		var result *models.LargeModel
		for tpb.Next() {
			var minValPtr, maxValPtr *int32
			if pb.MinVal != nil {
				v := pb.MinVal.Value
				minValPtr = &v
			}
			if pb.MaxVal != nil {
				v := pb.MaxVal.Value
				maxValPtr = &v
			}
			result = &models.LargeModel{
				ID:          pb.Id,
				Name:        pb.Name,
				Email:       pb.Email,
				Phone:       pb.Phone,
				Status:      int(pb.Status),
				Priority:    int(pb.Priority),
				Score:       pb.Score,
				Rating:      pb.Rating,
				Active:      pb.Active,
				Verified:    pb.Verified,
				Tags:        pb.Tags,
				Country:     pb.Country,
				Region:      pb.Region,
				City:        pb.City,
				ZipCode:     pb.ZipCode,
				Description: pb.Description,
				CreatedAt:   pb.CreatedAt.AsTime(),
				UpdatedAt:   pb.UpdatedAt.AsTime(),
				MinVal:      minValPtr,
				MaxVal:      maxValPtr,
			}
		}
		benchLargeModelResult = result
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// OrganizationPB Benchmarks (19 字段) - PBMO vs Native
// ══════════════════════════════════════════════════════════════════════════════

func BenchmarkPBMO_Organization_PBToModel(b *testing.B) {
	now := time.Now()
	maxUsers := int32(1000)
	maxStorage := int64(10737418240)
	pb := models.OrganizationPB{
		OrgId:           "org001",
		OrgCode:         "ORG001",
		OrgName:         "Test Organization",
		Status:          int32(1),
		Settings:        map[string]any{"key1": "val1", "key2": int64(42)},
		OwnerUserId:     "user001",
		AdminUserIds:    []string{"admin1", "admin2"},
		DepartmentCount: int32(10),
		EmployeeCount:   int32(100),
		RegionCode:      "GD",
		Industry:        "IT",
		BusinessType:    int32(1),
		ContactEmail:    "contact@example.com",
		ContactPhone:    "400-888-8888",
		CreatedAt:       timestamppb.New(now),
		UpdatedAt:       timestamppb.New(now),
		ExpiredAt:       timestamppb.New(now.AddDate(1, 0, 0)),
		MaxUsers:        wrapperspb.Int32(maxUsers),
		MaxStorage:      wrapperspb.Int64(maxStorage),
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.FromPB[models.OrganizationPB, models.OrganizationModel](&pb)
	}
}

func BenchmarkNative_Organization_PBToModel(b *testing.B) {
	now := time.Now()
	maxUsers := int32(1000)
	maxStorage := int64(10737418240)
	pb := models.OrganizationPB{
		OrgId:           "org001",
		OrgCode:         "ORG001",
		OrgName:         "Test Organization",
		Status:          int32(1),
		Settings:        map[string]any{"key1": "val1", "key2": int64(42)},
		OwnerUserId:     "user001",
		AdminUserIds:    []string{"admin1", "admin2"},
		DepartmentCount: int32(10),
		EmployeeCount:   int32(100),
		RegionCode:      "GD",
		Industry:        "IT",
		BusinessType:    int32(1),
		ContactEmail:    "contact@example.com",
		ContactPhone:    "400-888-8888",
		CreatedAt:       timestamppb.New(now),
		UpdatedAt:       timestamppb.New(now),
		ExpiredAt:       timestamppb.New(now.AddDate(1, 0, 0)),
		MaxUsers:        wrapperspb.Int32(maxUsers),
		MaxStorage:      wrapperspb.Int64(maxStorage),
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.OrganizationModel
	for i := 0; i < b.N; i++ {
		var maxUsersPtr *int32
		var maxStoragePtr *int64
		if pb.MaxUsers != nil {
			v := pb.MaxUsers.Value
			maxUsersPtr = &v
		}
		if pb.MaxStorage != nil {
			v := pb.MaxStorage.Value
			maxStoragePtr = &v
		}
		result = &models.OrganizationModel{
			OrgID:           pb.OrgId,
			OrgCode:         pb.OrgCode,
			OrgName:         pb.OrgName,
			Status:          int(pb.Status),
			Settings:        sqltypes.MapAny(pb.Settings),
			OwnerUserID:     pb.OwnerUserId,
			AdminUserIDs:    sqltypes.StringSlice(pb.AdminUserIds),
			DepartmentCount: int(pb.DepartmentCount),
			EmployeeCount:   int(pb.EmployeeCount),
			RegionCode:      pb.RegionCode,
			Industry:        pb.Industry,
			BusinessType:    int(pb.BusinessType),
			ContactEmail:    pb.ContactEmail,
			ContactPhone:    pb.ContactPhone,
			CreatedAt:       pb.CreatedAt.AsTime(),
			UpdatedAt:       pb.UpdatedAt.AsTime(),
			ExpiredAt:       pb.ExpiredAt.AsTime(),
			MaxUsers:        maxUsersPtr,
			MaxStorage:      maxStoragePtr,
		}
	}
	benchOrganizationModelResult = result
}

func BenchmarkPBMO_Organization_ModelToPB(b *testing.B) {
	now := time.Now()
	maxUsers := int32(1000)
	maxStorage := int64(10737418240)
	model := models.OrganizationModel{
		OrgID:           "org001",
		OrgCode:         "ORG001",
		OrgName:         "Test Organization",
		Status:          1,
		Settings:        sqltypes.MapAny{"key1": "val1", "key2": int64(42)},
		OwnerUserID:     "user001",
		AdminUserIDs:    sqltypes.StringSlice{"admin1", "admin2"},
		DepartmentCount: 10,
		EmployeeCount:   100,
		RegionCode:      "GD",
		Industry:        "IT",
		BusinessType:    1,
		ContactEmail:    "contact@example.com",
		ContactPhone:    "400-888-8888",
		CreatedAt:       now,
		UpdatedAt:       now,
		ExpiredAt:       now.AddDate(1, 0, 0),
		MaxUsers:        &maxUsers,
		MaxStorage:      &maxStorage,
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.ToPB[models.OrganizationModel, models.OrganizationPB](&model)
	}
}

func BenchmarkNative_Organization_ModelToPB(b *testing.B) {
	now := time.Now()
	maxUsers := int32(1000)
	maxStorage := int64(10737418240)
	model := models.OrganizationModel{
		OrgID:           "org001",
		OrgCode:         "ORG001",
		OrgName:         "Test Organization",
		Status:          1,
		Settings:        sqltypes.MapAny{"key1": "val1", "key2": int64(42)},
		OwnerUserID:     "user001",
		AdminUserIDs:    sqltypes.StringSlice{"admin1", "admin2"},
		DepartmentCount: 10,
		EmployeeCount:   100,
		RegionCode:      "GD",
		Industry:        "IT",
		BusinessType:    1,
		ContactEmail:    "contact@example.com",
		ContactPhone:    "400-888-8888",
		CreatedAt:       now,
		UpdatedAt:       now,
		ExpiredAt:       now.AddDate(1, 0, 0),
		MaxUsers:        &maxUsers,
		MaxStorage:      &maxStorage,
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.OrganizationPB
	for i := 0; i < b.N; i++ {
		pb := &models.OrganizationPB{
			OrgId:           model.OrgID,
			OrgCode:         model.OrgCode,
			OrgName:         model.OrgName,
			Status:          int32(model.Status),
			Settings:        map[string]any(model.Settings),
			OwnerUserId:     model.OwnerUserID,
			AdminUserIds:    []string(model.AdminUserIDs),
			DepartmentCount: int32(model.DepartmentCount),
			EmployeeCount:   int32(model.EmployeeCount),
			RegionCode:      model.RegionCode,
			Industry:        model.Industry,
			BusinessType:    int32(model.BusinessType),
			ContactEmail:    model.ContactEmail,
			ContactPhone:    model.ContactPhone,
			CreatedAt:       timestamppb.New(model.CreatedAt),
			UpdatedAt:       timestamppb.New(model.UpdatedAt),
			ExpiredAt:       timestamppb.New(model.ExpiredAt),
		}
		if model.MaxUsers != nil {
			pb.MaxUsers = wrapperspb.Int32(*model.MaxUsers)
		}
		if model.MaxStorage != nil {
			pb.MaxStorage = wrapperspb.Int64(*model.MaxStorage)
		}
		result = pb
	}
	benchOrganizationPBResult = result
}

// ══════════════════════════════════════════════════════════════════════════════
// UserProfilePB Benchmarks (27 字段) - PBMO vs Native
// ══════════════════════════════════════════════════════════════════════════════

func BenchmarkPBMO_UserProfile_PBToModel(b *testing.B) {
	now := time.Now()
	pb := models.UserProfilePB{
		UserId:         "user001",
		Username:       "testuser",
		Nickname:       "Test User",
		AvatarUrl:      "https://example.com/avatar.jpg",
		Email:          "test@example.com",
		PhoneNumber:    "13800138000",
		Gender:         int32(1),
		BirthDate:      "1990-01-01",
		RegionCode:     "GD",
		CountryCode:    "CN",
		CityCode:       "SZ",
		Address:        "Test Address",
		ZipCode:        "518000",
		Status:         int32(1),
		Level:          int32(5),
		Score:          int64(1000),
		Balance:        100.50,
		TotalSpent:     5000.00,
		LastLoginTime:  now.Unix(),
		LastLoginIp:    "192.168.1.1",
		RegisterTime:   now.AddDate(-1, 0, 0).Unix(),
		RegisterIp:     "192.168.1.100",
		RegisterDevice: "iPhone",
		VerifiedEmail:  true,
		VerifiedPhone:  true,
		CreatedAt:      timestamppb.New(now),
		UpdatedAt:      timestamppb.New(now),
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.FromPB[models.UserProfilePB, models.UserProfileModel](&pb)
	}
}

func BenchmarkNative_UserProfile_PBToModel(b *testing.B) {
	now := time.Now()
	pb := models.UserProfilePB{
		UserId:         "user001",
		Username:       "testuser",
		Nickname:       "Test User",
		AvatarUrl:      "https://example.com/avatar.jpg",
		Email:          "test@example.com",
		PhoneNumber:    "13800138000",
		Gender:         int32(1),
		BirthDate:      "1990-01-01",
		RegionCode:     "GD",
		CountryCode:    "CN",
		CityCode:       "SZ",
		Address:        "Test Address",
		ZipCode:        "518000",
		Status:         int32(1),
		Level:          int32(5),
		Score:          int64(1000),
		Balance:        100.50,
		TotalSpent:     5000.00,
		LastLoginTime:  now.Unix(),
		LastLoginIp:    "192.168.1.1",
		RegisterTime:   now.AddDate(-1, 0, 0).Unix(),
		RegisterIp:     "192.168.1.100",
		RegisterDevice: "iPhone",
		VerifiedEmail:  true,
		VerifiedPhone:  true,
		CreatedAt:      timestamppb.New(now),
		UpdatedAt:      timestamppb.New(now),
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.UserProfileModel
	for i := 0; i < b.N; i++ {
		result = &models.UserProfileModel{
			UserID:         pb.UserId,
			Username:       pb.Username,
			Nickname:       pb.Nickname,
			AvatarUrl:      pb.AvatarUrl,
			Email:          pb.Email,
			PhoneNumber:    pb.PhoneNumber,
			Gender:         int(pb.Gender),
			BirthDate:      pb.BirthDate,
			RegionCode:     pb.RegionCode,
			CountryCode:    pb.CountryCode,
			CityCode:       pb.CityCode,
			Address:        pb.Address,
			ZipCode:        pb.ZipCode,
			Status:         int(pb.Status),
			Level:          int(pb.Level),
			Score:          pb.Score,
			Balance:        pb.Balance,
			TotalSpent:     pb.TotalSpent,
			LastLoginTime:  pb.LastLoginTime,
			LastLoginIp:    pb.LastLoginIp,
			RegisterTime:   pb.RegisterTime,
			RegisterIp:     pb.RegisterIp,
			RegisterDevice: pb.RegisterDevice,
			VerifiedEmail:  pb.VerifiedEmail,
			VerifiedPhone:  pb.VerifiedPhone,
			CreatedAt:      pb.CreatedAt.AsTime(),
			UpdatedAt:      pb.UpdatedAt.AsTime(),
		}
	}
	benchUserProfileModelResult = result
}

func BenchmarkPBMO_UserProfile_ModelToPB(b *testing.B) {
	now := time.Now()
	model := models.UserProfileModel{
		UserID:         "user001",
		Username:       "testuser",
		Nickname:       "Test User",
		AvatarUrl:      "https://example.com/avatar.jpg",
		Email:          "test@example.com",
		PhoneNumber:    "13800138000",
		Gender:         1,
		BirthDate:      "1990-01-01",
		RegionCode:     "GD",
		CountryCode:    "CN",
		CityCode:       "SZ",
		Address:        "Test Address",
		ZipCode:        "518000",
		Status:         1,
		Level:          5,
		Score:          1000,
		Balance:        100.50,
		TotalSpent:     5000.00,
		LastLoginTime:  now.Unix(),
		LastLoginIp:    "192.168.1.1",
		RegisterTime:   now.AddDate(-1, 0, 0).Unix(),
		RegisterIp:     "192.168.1.100",
		RegisterDevice: "iPhone",
		VerifiedEmail:  true,
		VerifiedPhone:  true,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.ToPB[models.UserProfileModel, models.UserProfilePB](&model)
	}
}

func BenchmarkNative_UserProfile_ModelToPB(b *testing.B) {
	now := time.Now()
	model := models.UserProfileModel{
		UserID:         "user001",
		Username:       "testuser",
		Nickname:       "Test User",
		AvatarUrl:      "https://example.com/avatar.jpg",
		Email:          "test@example.com",
		PhoneNumber:    "13800138000",
		Gender:         1,
		BirthDate:      "1990-01-01",
		RegionCode:     "GD",
		CountryCode:    "CN",
		CityCode:       "SZ",
		Address:        "Test Address",
		ZipCode:        "518000",
		Status:         1,
		Level:          5,
		Score:          1000,
		Balance:        100.50,
		TotalSpent:     5000.00,
		LastLoginTime:  now.Unix(),
		LastLoginIp:    "192.168.1.1",
		RegisterTime:   now.AddDate(-1, 0, 0).Unix(),
		RegisterIp:     "192.168.1.100",
		RegisterDevice: "iPhone",
		VerifiedEmail:  true,
		VerifiedPhone:  true,
		CreatedAt:      now,
		UpdatedAt:      now,
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.UserProfilePB
	for i := 0; i < b.N; i++ {
		result = &models.UserProfilePB{
			UserId:         model.UserID,
			Username:       model.Username,
			Nickname:       model.Nickname,
			AvatarUrl:      model.AvatarUrl,
			Email:          model.Email,
			PhoneNumber:    model.PhoneNumber,
			Gender:         int32(model.Gender),
			BirthDate:      model.BirthDate,
			RegionCode:     model.RegionCode,
			CountryCode:    model.CountryCode,
			CityCode:       model.CityCode,
			Address:        model.Address,
			ZipCode:        model.ZipCode,
			Status:         int32(model.Status),
			Level:          int32(model.Level),
			Score:          model.Score,
			Balance:        model.Balance,
			TotalSpent:     model.TotalSpent,
			LastLoginTime:  model.LastLoginTime,
			LastLoginIp:    model.LastLoginIp,
			RegisterTime:   model.RegisterTime,
			RegisterIp:     model.RegisterIp,
			RegisterDevice: model.RegisterDevice,
			VerifiedEmail:  model.VerifiedEmail,
			VerifiedPhone:  model.VerifiedPhone,
			CreatedAt:      timestamppb.New(model.CreatedAt),
			UpdatedAt:      timestamppb.New(model.UpdatedAt),
		}
	}
	benchUserProfilePBResult = result
}

func BenchmarkPBMO_UserProfile_PBToModel_Parallel(b *testing.B) {
	now := time.Now()
	pb := models.UserProfilePB{
		UserId:         "user001",
		Username:       "testuser",
		Nickname:       "Test User",
		AvatarUrl:      "https://example.com/avatar.jpg",
		Email:          "test@example.com",
		PhoneNumber:    "13800138000",
		Gender:         int32(1),
		BirthDate:      "1990-01-01",
		RegionCode:     "GD",
		CountryCode:    "CN",
		CityCode:       "SZ",
		Address:        "Test Address",
		ZipCode:        "518000",
		Status:         int32(1),
		Level:          int32(5),
		Score:          int64(1000),
		Balance:        100.50,
		TotalSpent:     5000.00,
		LastLoginTime:  now.Unix(),
		LastLoginIp:    "192.168.1.1",
		RegisterTime:   now.AddDate(-1, 0, 0).Unix(),
		RegisterIp:     "192.168.1.100",
		RegisterDevice: "iPhone",
		VerifiedEmail:  true,
		VerifiedPhone:  true,
		CreatedAt:      timestamppb.New(now),
		UpdatedAt:      timestamppb.New(now),
	}

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		var result *models.UserProfileModel
		for p.Next() {
			result, _ = pbmo.FromPB[models.UserProfilePB, models.UserProfileModel](&pb)
		}
		benchUserProfileModelResult = result
	})
}

func BenchmarkNative_UserProfile_PBToModel_Parallel(b *testing.B) {
	now := time.Now()
	pb := models.UserProfilePB{
		UserId:         "user001",
		Username:       "testuser",
		Nickname:       "Test User",
		AvatarUrl:      "https://example.com/avatar.jpg",
		Email:          "test@example.com",
		PhoneNumber:    "13800138000",
		Gender:         int32(1),
		BirthDate:      "1990-01-01",
		RegionCode:     "GD",
		CountryCode:    "CN",
		CityCode:       "SZ",
		Address:        "Test Address",
		ZipCode:        "518000",
		Status:         int32(1),
		Level:          int32(5),
		Score:          int64(1000),
		Balance:        100.50,
		TotalSpent:     5000.00,
		LastLoginTime:  now.Unix(),
		LastLoginIp:    "192.168.1.1",
		RegisterTime:   now.AddDate(-1, 0, 0).Unix(),
		RegisterIp:     "192.168.1.100",
		RegisterDevice: "iPhone",
		VerifiedEmail:  true,
		VerifiedPhone:  true,
		CreatedAt:      timestamppb.New(now),
		UpdatedAt:      timestamppb.New(now),
	}

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(tpb *testing.PB) {
		var result *models.UserProfileModel
		for tpb.Next() {
			result = &models.UserProfileModel{
				UserID:         pb.UserId,
				Username:       pb.Username,
				Nickname:       pb.Nickname,
				AvatarUrl:      pb.AvatarUrl,
				Email:          pb.Email,
				PhoneNumber:    pb.PhoneNumber,
				Gender:         int(pb.Gender),
				BirthDate:      pb.BirthDate,
				RegionCode:     pb.RegionCode,
				CountryCode:    pb.CountryCode,
				CityCode:       pb.CityCode,
				Address:        pb.Address,
				ZipCode:        pb.ZipCode,
				Status:         int(pb.Status),
				Level:          int(pb.Level),
				Score:          pb.Score,
				Balance:        pb.Balance,
				TotalSpent:     pb.TotalSpent,
				LastLoginTime:  pb.LastLoginTime,
				LastLoginIp:    pb.LastLoginIp,
				RegisterTime:   pb.RegisterTime,
				RegisterIp:     pb.RegisterIp,
				RegisterDevice: pb.RegisterDevice,
				VerifiedEmail:  pb.VerifiedEmail,
				VerifiedPhone:  pb.VerifiedPhone,
				CreatedAt:      pb.CreatedAt.AsTime(),
				UpdatedAt:      pb.UpdatedAt.AsTime(),
			}
		}
		benchUserProfileModelResult = result
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// 批量转换测试 - PBMO vs Native
// ══════════════════════════════════════════════════════════════════════════════

func BenchmarkPBMO_Large_Batch100_PBToModel(b *testing.B) {
	now := time.Now()
	minVal := int32(10)
	maxVal := int32(100)

	pbs := make([]*models.LargePB, 100)
	for i := range pbs {
		pbs[i] = &models.LargePB{
			Id:          uint64(i),
			Name:        "user",
			Email:       "test@example.com",
			Phone:       "13800138000",
			Status:      int32(1),
			Priority:    int32(5),
			Score:       95.5,
			Rating:      4.8,
			Active:      true,
			Verified:    true,
			Tags:        []string{"tag1", "tag2"},
			Country:     "China",
			Region:      "Guangdong",
			City:        "Shenzhen",
			ZipCode:     "518000",
			Description: "test",
			CreatedAt:   timestamppb.New(now),
			UpdatedAt:   timestamppb.New(now),
			MinVal:      wrapperspb.Int32(minVal),
			MaxVal:      wrapperspb.Int32(maxVal),
		}
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.FromPBs[models.LargePB, models.LargeModel](pbs)
	}
}

func BenchmarkNative_Large_Batch100_PBToModel(b *testing.B) {
	now := time.Now()
	minVal := int32(10)
	maxVal := int32(100)

	pbs := make([]*models.LargePB, 100)
	for i := range pbs {
		pbs[i] = &models.LargePB{
			Id:          uint64(i),
			Name:        "user",
			Email:       "test@example.com",
			Phone:       "13800138000",
			Status:      int32(1),
			Priority:    int32(5),
			Score:       95.5,
			Rating:      4.8,
			Active:      true,
			Verified:    true,
			Tags:        []string{"tag1", "tag2"},
			Country:     "China",
			Region:      "Guangdong",
			City:        "Shenzhen",
			ZipCode:     "518000",
			Description: "test",
			CreatedAt:   timestamppb.New(now),
			UpdatedAt:   timestamppb.New(now),
			MinVal:      wrapperspb.Int32(minVal),
			MaxVal:      wrapperspb.Int32(maxVal),
		}
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result []*models.LargeModel
	for i := 0; i < b.N; i++ {
		result = make([]*models.LargeModel, len(pbs))
		for j, pb := range pbs {
			var minValPtr, maxValPtr *int32
			if pb.MinVal != nil {
				v := pb.MinVal.Value
				minValPtr = &v
			}
			if pb.MaxVal != nil {
				v := pb.MaxVal.Value
				maxValPtr = &v
			}
			result[j] = &models.LargeModel{
				ID:          pb.Id,
				Name:        pb.Name,
				Email:       pb.Email,
				Phone:       pb.Phone,
				Status:      int(pb.Status),
				Priority:    int(pb.Priority),
				Score:       pb.Score,
				Rating:      pb.Rating,
				Active:      pb.Active,
				Verified:    pb.Verified,
				Tags:        pb.Tags,
				Country:     pb.Country,
				Region:      pb.Region,
				City:        pb.City,
				ZipCode:     pb.ZipCode,
				Description: pb.Description,
				CreatedAt:   pb.CreatedAt.AsTime(),
				UpdatedAt:   pb.UpdatedAt.AsTime(),
				MinVal:      minValPtr,
				MaxVal:      maxValPtr,
			}
		}
	}
	benchLargeModelsResult = result
}
