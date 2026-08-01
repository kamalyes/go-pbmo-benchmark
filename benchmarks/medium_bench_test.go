/**
 * @Author: developer 500000000@example.com
 * @Date: 2023-09-16 17:25:07
 * @LastEditors: developer 500000000@example.com
 * @LastEditTime: 2025-06-15 00:16:15
 * @FilePath: \go-pbmo-benchmark\benchmarks\medium_bench_test.go
 * @Description: 中等复杂度模型性能测试 - 8-13 字段
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
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// 全局变量，防止编译器优化掉 benchmark 结果
var (
	benchMediumModelResult        *models.MediumModel
	benchMediumPBResult           *models.MediumPB
	benchFullModelResult          *models.FullModel
	benchFullPBResult             *models.FullPB
	benchMemberModelResult        *models.MemberProfileModel
	benchMemberPBResult           *models.MemberProfilePB
	benchServiceConfigModelResult *models.ServiceConfigModel
	benchServiceConfigPBResult    *models.ServiceConfigPB
)

func init() {
	pbmo.Register[models.MediumPB, models.MediumModel]()
	pbmo.Register[models.FullPB, models.FullModel]()
	pbmo.Register[models.MemberProfilePB, models.MemberProfileModel]()
	pbmo.Register[models.ServiceConfigPB, models.ServiceConfigModel]()
}

// ══════════════════════════════════════════════════════════════════════════════
// MediumPB Benchmarks (8 字段) - PBMO vs Native
// ══════════════════════════════════════════════════════════════════════════════

func BenchmarkPBMO_Medium_PBToModel(b *testing.B) {
	pb := models.MediumPB{
		Id:       uint64(1),
		Name:     "testuser",
		Email:    "test@example.com",
		Age:      int32(25),
		Score:    95.5,
		Active:   true,
		Tags:     []string{"tag1", "tag2", "tag3"},
		Priority: int32(5),
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.FromPB[models.MediumPB, models.MediumModel](&pb)
	}
}

func BenchmarkNative_Medium_PBToModel(b *testing.B) {
	pb := models.MediumPB{
		Id:       uint64(1),
		Name:     "testuser",
		Email:    "test@example.com",
		Age:      int32(25),
		Score:    95.5,
		Active:   true,
		Tags:     []string{"tag1", "tag2", "tag3"},
		Priority: int32(5),
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.MediumModel
	for i := 0; i < b.N; i++ {
		result = &models.MediumModel{
			ID:       pb.Id,
			Name:     pb.Name,
			Email:    pb.Email,
			Age:      int(pb.Age),
			Score:    pb.Score,
			Active:   pb.Active,
			Tags:     pb.Tags,
			Priority: int(pb.Priority),
		}
	}
	benchMediumModelResult = result
}

func BenchmarkPBMO_Medium_ModelToPB(b *testing.B) {
	model := models.MediumModel{
		ID:       uint64(1),
		Name:     "testuser",
		Email:    "test@example.com",
		Age:      25,
		Score:    95.5,
		Active:   true,
		Tags:     []string{"tag1", "tag2", "tag3"},
		Priority: 5,
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.ToPB[models.MediumModel, models.MediumPB](&model)
	}
}

func BenchmarkNative_Medium_ModelToPB(b *testing.B) {
	model := models.MediumModel{
		ID:       uint64(1),
		Name:     "testuser",
		Email:    "test@example.com",
		Age:      25,
		Score:    95.5,
		Active:   true,
		Tags:     []string{"tag1", "tag2", "tag3"},
		Priority: 5,
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.MediumPB
	for i := 0; i < b.N; i++ {
		result = &models.MediumPB{
			Id:       model.ID,
			Name:     model.Name,
			Email:    model.Email,
			Age:      int32(model.Age),
			Score:    model.Score,
			Active:   model.Active,
			Tags:     model.Tags,
			Priority: int32(model.Priority),
		}
	}
	benchMediumPBResult = result
}

// ══════════════════════════════════════════════════════════════════════════════
// FullPB Benchmarks (13 字段 + 时间戳 + Wrapper) - PBMO vs Native
// ══════════════════════════════════════════════════════════════════════════════

func BenchmarkPBMO_Full_PBToModel(b *testing.B) {
	now := time.Now()
	minVal := int32(10)
	maxVal := int32(100)
	desc := "test description"
	pb := models.FullPB{
		Id:          uint64(1),
		Name:        "testuser",
		Email:       "test@example.com",
		Age:         int32(25),
		Score:       95.5,
		Active:      true,
		Tags:        []string{"tag1", "tag2"},
		CreatedAt:   timestamppb.New(now),
		UpdatedAt:   timestamppb.New(now),
		MinVal:      wrapperspb.Int32(minVal),
		MaxVal:      wrapperspb.Int32(maxVal),
		Description: wrapperspb.String(desc),
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.FromPB[models.FullPB, models.FullModel](&pb)
	}
}

func BenchmarkNative_Full_PBToModel(b *testing.B) {
	now := time.Now()
	minVal := int32(10)
	maxVal := int32(100)
	desc := "test description"
	pb := models.FullPB{
		Id:          uint64(1),
		Name:        "testuser",
		Email:       "test@example.com",
		Age:         int32(25),
		Score:       95.5,
		Active:      true,
		Tags:        []string{"tag1", "tag2"},
		CreatedAt:   timestamppb.New(now),
		UpdatedAt:   timestamppb.New(now),
		MinVal:      wrapperspb.Int32(minVal),
		MaxVal:      wrapperspb.Int32(maxVal),
		Description: wrapperspb.String(desc),
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.FullModel
	for i := 0; i < b.N; i++ {
		var minValPtr, maxValPtr *int32
		var descPtr *string
		if pb.MinVal != nil {
			v := pb.MinVal.Value
			minValPtr = &v
		}
		if pb.MaxVal != nil {
			v := pb.MaxVal.Value
			maxValPtr = &v
		}
		if pb.Description != nil {
			v := pb.Description.Value
			descPtr = &v
		}
		result = &models.FullModel{
			ID:          pb.Id,
			Name:        pb.Name,
			Email:       pb.Email,
			Age:         int(pb.Age),
			Score:       pb.Score,
			Active:      pb.Active,
			Tags:        pb.Tags,
			CreatedAt:   pb.CreatedAt.AsTime(),
			UpdatedAt:   pb.UpdatedAt.AsTime(),
			MinVal:      minValPtr,
			MaxVal:      maxValPtr,
			Description: descPtr,
		}
	}
	benchFullModelResult = result
}

func BenchmarkPBMO_Full_ModelToPB(b *testing.B) {
	now := time.Now()
	minVal := int32(10)
	maxVal := int32(100)
	desc := "test description"
	model := models.FullModel{
		ID:          uint64(1),
		Name:        "testuser",
		Email:       "test@example.com",
		Age:         25,
		Score:       95.5,
		Active:      true,
		Tags:        []string{"tag1", "tag2"},
		CreatedAt:   now,
		UpdatedAt:   now,
		MinVal:      &minVal,
		MaxVal:      &maxVal,
		Description: &desc,
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.ToPB[models.FullModel, models.FullPB](&model)
	}
}

func BenchmarkNative_Full_ModelToPB(b *testing.B) {
	now := time.Now()
	minVal := int32(10)
	maxVal := int32(100)
	desc := "test description"
	model := models.FullModel{
		ID:          uint64(1),
		Name:        "testuser",
		Email:       "test@example.com",
		Age:         25,
		Score:       95.5,
		Active:      true,
		Tags:        []string{"tag1", "tag2"},
		CreatedAt:   now,
		UpdatedAt:   now,
		MinVal:      &minVal,
		MaxVal:      &maxVal,
		Description: &desc,
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.FullPB
	for i := 0; i < b.N; i++ {
		pb := &models.FullPB{
			Id:        model.ID,
			Name:      model.Name,
			Email:     model.Email,
			Age:       int32(model.Age),
			Score:     model.Score,
			Active:    model.Active,
			Tags:      model.Tags,
			CreatedAt: timestamppb.New(model.CreatedAt),
			UpdatedAt: timestamppb.New(model.UpdatedAt),
		}
		if model.MinVal != nil {
			pb.MinVal = wrapperspb.Int32(*model.MinVal)
		}
		if model.MaxVal != nil {
			pb.MaxVal = wrapperspb.Int32(*model.MaxVal)
		}
		if model.Description != nil {
			pb.Description = wrapperspb.String(*model.Description)
		}
		result = pb
	}
	benchFullPBResult = result
}

// ══════════════════════════════════════════════════════════════════════════════
// MemberProfilePB Benchmarks (11 字段 + MapAny) - PBMO vs Native
// ══════════════════════════════════════════════════════════════════════════════

func BenchmarkPBMO_Member_PBToModel(b *testing.B) {
	now := time.Now()
	balance := 100.50
	pb := models.MemberProfilePB{
		MemberId:    "mem001",
		Username:    "testuser",
		Email:       "test@example.com",
		PhoneNumber: "13800138000",
		Level:       int32(3),
		Score:       95.5,
		IsActive:    true,
		Tags:        []string{"vip", "verified"},
		Balance:     wrapperspb.Double(balance),
		CreatedAt:   timestamppb.New(now),
		UpdatedAt:   timestamppb.New(now),
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.FromPB[models.MemberProfilePB, models.MemberProfileModel](&pb)
	}
}

func BenchmarkNative_Member_PBToModel(b *testing.B) {
	now := time.Now()
	balance := 100.50
	pb := models.MemberProfilePB{
		MemberId:    "mem001",
		Username:    "testuser",
		Email:       "test@example.com",
		PhoneNumber: "13800138000",
		Level:       int32(3),
		Score:       95.5,
		IsActive:    true,
		Tags:        []string{"vip", "verified"},
		Balance:     wrapperspb.Double(balance),
		CreatedAt:   timestamppb.New(now),
		UpdatedAt:   timestamppb.New(now),
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.MemberProfileModel
	for i := 0; i < b.N; i++ {
		var balancePtr *float64
		if pb.Balance != nil {
			v := pb.Balance.Value
			balancePtr = &v
		}
		result = &models.MemberProfileModel{
			MemberID:    pb.MemberId,
			Username:    pb.Username,
			Email:       pb.Email,
			PhoneNumber: pb.PhoneNumber,
			Level:       int(pb.Level),
			Score:       pb.Score,
			IsActive:    pb.IsActive,
			Tags:        sqltypes.StringSlice(pb.Tags),
			Balance:     balancePtr,
			CreatedAt:   pb.CreatedAt.AsTime(),
			UpdatedAt:   pb.UpdatedAt.AsTime(),
		}
	}
	benchMemberModelResult = result
}

func BenchmarkPBMO_Member_ModelToPB(b *testing.B) {
	now := time.Now()
	balance := 100.50
	model := models.MemberProfileModel{
		MemberID:    "mem001",
		Username:    "testuser",
		Email:       "test@example.com",
		PhoneNumber: "13800138000",
		Level:       3,
		Score:       95.5,
		IsActive:    true,
		Tags:        sqltypes.StringSlice{"vip", "verified"},
		Balance:     &balance,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.ToPB[models.MemberProfileModel, models.MemberProfilePB](&model)
	}
}

func BenchmarkNative_Member_ModelToPB(b *testing.B) {
	now := time.Now()
	balance := 100.50
	model := models.MemberProfileModel{
		MemberID:    "mem001",
		Username:    "testuser",
		Email:       "test@example.com",
		PhoneNumber: "13800138000",
		Level:       3,
		Score:       95.5,
		IsActive:    true,
		Tags:        sqltypes.StringSlice{"vip", "verified"},
		Balance:     &balance,
		CreatedAt:   now,
		UpdatedAt:   now,
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.MemberProfilePB
	for i := 0; i < b.N; i++ {
		pb := &models.MemberProfilePB{
			MemberId:    model.MemberID,
			Username:    model.Username,
			Email:       model.Email,
			PhoneNumber: model.PhoneNumber,
			Level:       int32(model.Level),
			Score:       model.Score,
			IsActive:    model.IsActive,
			Tags:        []string(model.Tags),
			CreatedAt:   timestamppb.New(model.CreatedAt),
			UpdatedAt:   timestamppb.New(model.UpdatedAt),
		}
		if model.Balance != nil {
			pb.Balance = wrapperspb.Double(*model.Balance)
		}
		result = pb
	}
	benchMemberPBResult = result
}

// ══════════════════════════════════════════════════════════════════════════════
// ServiceConfigPB Benchmarks (9 字段) - PBMO vs Native
// ══════════════════════════════════════════════════════════════════════════════

func BenchmarkPBMO_ServiceConfig_PBToModel(b *testing.B) {
	now := time.Now()
	enabled := true
	priority := int32(5)
	pb := models.ServiceConfigPB{
		ServiceId:   "svc001",
		ServiceName: "test-service",
		Version:     "1.0.0",
		Enabled:     wrapperspb.Bool(enabled),
		Port:        int32(8080),
		Endpoints:   []string{"/api/v1", "/api/v2"},
		Settings:    map[string]any{"timeout": 30, "retries": 3},
		Priority:    wrapperspb.Int32(priority),
		UpdatedAt:   timestamppb.New(now),
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.FromPB[models.ServiceConfigPB, models.ServiceConfigModel](&pb)
	}
}

func BenchmarkNative_ServiceConfig_PBToModel(b *testing.B) {
	now := time.Now()
	enabled := true
	priority := int32(5)
	pb := models.ServiceConfigPB{
		ServiceId:   "svc001",
		ServiceName: "test-service",
		Version:     "1.0.0",
		Enabled:     wrapperspb.Bool(enabled),
		Port:        int32(8080),
		Endpoints:   []string{"/api/v1", "/api/v2"},
		Settings:    map[string]any{"timeout": 30, "retries": 3},
		Priority:    wrapperspb.Int32(priority),
		UpdatedAt:   timestamppb.New(now),
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.ServiceConfigModel
	for i := 0; i < b.N; i++ {
		var enabledPtr *bool
		if pb.Enabled != nil {
			v := pb.Enabled.Value
			enabledPtr = &v
		}
		var priorityPtr *int32
		if pb.Priority != nil {
			v := pb.Priority.Value
			priorityPtr = &v
		}
		result = &models.ServiceConfigModel{
			ServiceID:   pb.ServiceId,
			ServiceName: pb.ServiceName,
			Version:     pb.Version,
			Enabled:     enabledPtr,
			Port:        int(pb.Port),
			Endpoints:   sqltypes.StringSlice(pb.Endpoints),
			Settings:    sqltypes.MapAny(pb.Settings),
			Priority:    priorityPtr,
			UpdatedAt:   pb.UpdatedAt.AsTime(),
		}
	}
	benchServiceConfigModelResult = result
}

func BenchmarkPBMO_ServiceConfig_ModelToPB(b *testing.B) {
	now := time.Now()
	enabled := true
	priority := int32(5)
	model := models.ServiceConfigModel{
		ServiceID:   "svc001",
		ServiceName: "test-service",
		Version:     "1.0.0",
		Enabled:     &enabled,
		Port:        8080,
		Endpoints:   sqltypes.StringSlice{"/api/v1", "/api/v2"},
		Settings:    sqltypes.MapAny{"timeout": 30, "retries": 3},
		Priority:    &priority,
		UpdatedAt:   now,
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.ToPB[models.ServiceConfigModel, models.ServiceConfigPB](&model)
	}
}

func BenchmarkNative_ServiceConfig_ModelToPB(b *testing.B) {
	now := time.Now()
	enabled := true
	priority := int32(5)
	model := models.ServiceConfigModel{
		ServiceID:   "svc001",
		ServiceName: "test-service",
		Version:     "1.0.0",
		Enabled:     &enabled,
		Port:        8080,
		Endpoints:   sqltypes.StringSlice{"/api/v1", "/api/v2"},
		Settings:    sqltypes.MapAny{"timeout": 30, "retries": 3},
		Priority:    &priority,
		UpdatedAt:   now,
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.ServiceConfigPB
	for i := 0; i < b.N; i++ {
		pb := &models.ServiceConfigPB{
			ServiceId:   model.ServiceID,
			ServiceName: model.ServiceName,
			Version:     model.Version,
			Port:        int32(model.Port),
			Endpoints:   []string(model.Endpoints),
			Settings:    map[string]any(model.Settings),
			UpdatedAt:   timestamppb.New(model.UpdatedAt),
		}
		if model.Enabled != nil {
			pb.Enabled = wrapperspb.Bool(*model.Enabled)
		}
		if model.Priority != nil {
			pb.Priority = wrapperspb.Int32(*model.Priority)
		}
		result = pb
	}
	benchServiceConfigPBResult = result
}

// ══════════════════════════════════════════════════════════════════════════════
// 并行测试 - PBMO vs Native
// ══════════════════════════════════════════════════════════════════════════════

func BenchmarkPBMO_Medium_PBToModel_Parallel(b *testing.B) {
	mediumPB := models.MediumPB{
		Id:       uint64(1),
		Name:     "testuser",
		Email:    "test@example.com",
		Age:      int32(25),
		Score:    95.5,
		Active:   true,
		Tags:     []string{"tag1", "tag2", "tag3"},
		Priority: int32(5),
	}

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		var result *models.MediumModel
		for pb.Next() {
			result, _ = pbmo.FromPB[models.MediumPB, models.MediumModel](&mediumPB)
		}
		benchMediumModelResult = result
	})
}

func BenchmarkPBMO_Full_PBToModel_Parallel(b *testing.B) {
	now := time.Now()
	minVal := int32(10)
	maxVal := int32(100)
	desc := "test description"
	fullPB := models.FullPB{
		Id:          uint64(1),
		Name:        "testuser",
		Email:       "test@example.com",
		Age:         int32(25),
		Score:       95.5,
		Active:      true,
		Tags:        []string{"tag1", "tag2"},
		CreatedAt:   timestamppb.New(now),
		UpdatedAt:   timestamppb.New(now),
		MinVal:      wrapperspb.Int32(minVal),
		MaxVal:      wrapperspb.Int32(maxVal),
		Description: wrapperspb.String(desc),
	}

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		var result *models.FullModel
		for pb.Next() {
			result, _ = pbmo.FromPB[models.FullPB, models.FullModel](&fullPB)
		}
		benchFullModelResult = result
	})
}

func BenchmarkNative_Medium_PBToModel_Parallel(b *testing.B) {
	pb := models.MediumPB{
		Id:       uint64(1),
		Name:     "testuser",
		Email:    "test@example.com",
		Age:      int32(25),
		Score:    95.5,
		Active:   true,
		Tags:     []string{"tag1", "tag2", "tag3"},
		Priority: int32(5),
	}

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(tpb *testing.PB) {
		var result *models.MediumModel
		for tpb.Next() {
			result = &models.MediumModel{
				ID:       pb.Id,
				Name:     pb.Name,
				Email:    pb.Email,
				Age:      int(pb.Age),
				Score:    pb.Score,
				Active:   pb.Active,
				Tags:     pb.Tags,
				Priority: int(pb.Priority),
			}
		}
		benchMediumModelResult = result
	})
}

func BenchmarkNative_Full_PBToModel_Parallel(b *testing.B) {
	now := time.Now()
	minVal := int32(10)
	maxVal := int32(100)
	desc := "test description"
	pb := models.FullPB{
		Id:          uint64(1),
		Name:        "testuser",
		Email:       "test@example.com",
		Age:         int32(25),
		Score:       95.5,
		Active:      true,
		Tags:        []string{"tag1", "tag2"},
		CreatedAt:   timestamppb.New(now),
		UpdatedAt:   timestamppb.New(now),
		MinVal:      wrapperspb.Int32(minVal),
		MaxVal:      wrapperspb.Int32(maxVal),
		Description: wrapperspb.String(desc),
	}

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(tpb *testing.PB) {
		var result *models.FullModel
		for tpb.Next() {
			var minValPtr, maxValPtr *int32
			var descPtr *string
			if pb.MinVal != nil {
				v := pb.MinVal.Value
				minValPtr = &v
			}
			if pb.MaxVal != nil {
				v := pb.MaxVal.Value
				maxValPtr = &v
			}
			if pb.Description != nil {
				v := pb.Description.Value
				descPtr = &v
			}
			result = &models.FullModel{
				ID:          pb.Id,
				Name:        pb.Name,
				Email:       pb.Email,
				Age:         int(pb.Age),
				Score:       pb.Score,
				Active:      pb.Active,
				Tags:        pb.Tags,
				CreatedAt:   pb.CreatedAt.AsTime(),
				UpdatedAt:   pb.UpdatedAt.AsTime(),
				MinVal:      minValPtr,
				MaxVal:      maxValPtr,
				Description: descPtr,
			}
		}
		benchFullModelResult = result
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// 验证测试 - 确保 PBMO 和 Native 转换结果一致
// ══════════════════════════════════════════════════════════════════════════════

func TestMedium_PBToModel_Consistency(t *testing.T) {
	pb := models.MediumPB{
		Id:       uint64(1),
		Name:     "testuser",
		Email:    "test@example.com",
		Age:      int32(25),
		Score:    95.5,
		Active:   true,
		Tags:     []string{"tag1", "tag2", "tag3"},
		Priority: int32(5),
	}

	pbmoResult, err := pbmo.FromPB[models.MediumPB, models.MediumModel](&pb)
	assert.Nil(t, err)
	assert.NotNil(t, pbmoResult)

	nativeResult := &models.MediumModel{
		ID:       pb.Id,
		Name:     pb.Name,
		Email:    pb.Email,
		Age:      int(pb.Age),
		Score:    pb.Score,
		Active:   pb.Active,
		Tags:     pb.Tags,
		Priority: int(pb.Priority),
	}

	assert.Equal(t, nativeResult.ID, pbmoResult.ID)
	assert.Equal(t, nativeResult.Name, pbmoResult.Name)
	assert.Equal(t, nativeResult.Email, pbmoResult.Email)
	assert.Equal(t, nativeResult.Age, pbmoResult.Age)
	assert.Equal(t, nativeResult.Score, pbmoResult.Score)
	assert.Equal(t, nativeResult.Active, pbmoResult.Active)
	assert.Equal(t, nativeResult.Priority, pbmoResult.Priority)
}

func TestFull_PBToModel_Consistency(t *testing.T) {
	now := time.Now()
	minVal := int32(10)
	maxVal := int32(100)
	desc := "test description"
	pb := models.FullPB{
		Id:          uint64(1),
		Name:        "testuser",
		Email:       "test@example.com",
		Age:         int32(25),
		Score:       95.5,
		Active:      true,
		Tags:        []string{"tag1", "tag2"},
		CreatedAt:   timestamppb.New(now),
		UpdatedAt:   timestamppb.New(now),
		MinVal:      wrapperspb.Int32(minVal),
		MaxVal:      wrapperspb.Int32(maxVal),
		Description: wrapperspb.String(desc),
	}

	pbmoResult, err := pbmo.FromPB[models.FullPB, models.FullModel](&pb)
	assert.Nil(t, err)
	assert.NotNil(t, pbmoResult)

	assert.Equal(t, pb.Id, pbmoResult.ID)
	assert.Equal(t, pb.Name, pbmoResult.Name)
	assert.Equal(t, pb.Email, pbmoResult.Email)
	assert.Equal(t, int(pb.Age), pbmoResult.Age)
	assert.Equal(t, pb.Score, pbmoResult.Score)
	assert.Equal(t, pb.Active, pbmoResult.Active)
	assert.Equal(t, minVal, *pbmoResult.MinVal)
	assert.Equal(t, maxVal, *pbmoResult.MaxVal)
	assert.Equal(t, desc, *pbmoResult.Description)
}

func TestMemberProfile_PBToModel_Consistency(t *testing.T) {
	now := time.Now()
	balance := 100.50
	pb := models.MemberProfilePB{
		MemberId:    "mem001",
		Username:    "testuser",
		Email:       "test@example.com",
		PhoneNumber: "13800138000",
		Level:       int32(3),
		Score:       95.5,
		IsActive:    true,
		Tags:        []string{"vip", "verified"},
		Balance:     wrapperspb.Double(balance),
		CreatedAt:   timestamppb.New(now),
		UpdatedAt:   timestamppb.New(now),
	}

	pbmoResult, err := pbmo.FromPB[models.MemberProfilePB, models.MemberProfileModel](&pb)
	assert.Nil(t, err)
	assert.NotNil(t, pbmoResult)

	assert.Equal(t, pb.MemberId, pbmoResult.MemberID)
	assert.Equal(t, pb.Username, pbmoResult.Username)
	assert.Equal(t, pb.Email, pbmoResult.Email)
	assert.Equal(t, int(pb.Level), pbmoResult.Level)
	assert.Equal(t, pb.Score, pbmoResult.Score)
	assert.Equal(t, pb.IsActive, pbmoResult.IsActive)
	assert.Equal(t, balance, *pbmoResult.Balance)
}

func TestServiceConfig_PBToModel_Consistency(t *testing.T) {
	now := time.Now()
	enabled := true
	priority := int32(5)
	pb := models.ServiceConfigPB{
		ServiceId:   "svc001",
		ServiceName: "test-service",
		Version:     "1.0.0",
		Enabled:     wrapperspb.Bool(enabled),
		Port:        int32(8080),
		Endpoints:   []string{"/api/v1", "/api/v2"},
		Settings:    map[string]any{"timeout": 30, "retries": 3},
		Priority:    wrapperspb.Int32(priority),
		UpdatedAt:   timestamppb.New(now),
	}

	pbmoResult, err := pbmo.FromPB[models.ServiceConfigPB, models.ServiceConfigModel](&pb)
	assert.Nil(t, err)
	assert.NotNil(t, pbmoResult)

	assert.Equal(t, pb.ServiceId, pbmoResult.ServiceID)
	assert.Equal(t, pb.ServiceName, pbmoResult.ServiceName)
	assert.Equal(t, pb.Version, pbmoResult.Version)
	assert.Equal(t, int(pb.Port), pbmoResult.Port)
	assert.Equal(t, enabled, *pbmoResult.Enabled)
	assert.Equal(t, priority, *pbmoResult.Priority)
}
