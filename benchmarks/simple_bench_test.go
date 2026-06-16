/**
 * @Author: developer 500000000@example.com
 * @Date: 2023-09-17 00:08:15
 * @LastEditors: developer 500000000@example.com
 * @LastEditTime: 2025-06-15 00:05:18
 * @FilePath: \go-pbmo-benchmark\benchmarks\simple_bench_test.go
 * @Description: 简单模型性能测试 - PBMO vs Native - 4-6 字段
 *
 * Copyright (c) 2025 by developer, All Rights Reserved.
 */

package benchmarks

import (
	"testing"
	"time"

	"github.com/kamalyes/go-pbmo"
	"github.com/kamalyes/go-pbmo-benchmark/models"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// 全局变量，防止编译器优化掉 benchmark 结果
var (
	benchSimplePBResult     *models.SimplePB
	benchSimpleModelResult  *models.SimpleModel
	benchAccountPBResult    *models.AccountInfoPB
	benchAccountModelResult *models.AccountInfoModel
	benchMappedPBResult     *models.MappedPB
	benchMappedModelResult  *models.MappedModel
	benchSimpleModelsResult []*models.SimpleModel
	benchSimplePBsResult    []*models.SimplePB
)

func init() {
	pbmo.Register[models.SimplePB, models.SimpleModel]()
	pbmo.Register[models.AccountInfoPB, models.AccountInfoModel]()
	pbmo.Register[models.MappedPB, models.MappedModel]()
}

// ══════════════════════════════════════════════════════════════════════════════
// SimplePB Benchmarks (4 字段) - PBMO vs Native
// ══════════════════════════════════════════════════════════════════════════════

func BenchmarkPBMO_Simple_PBToModel(b *testing.B) {
	pb := models.SimplePB{
		Id:    uint64(1),
		Name:  "testuser",
		Email: "test@example.com",
		Age:   int32(25),
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.FromPB[models.SimplePB, models.SimpleModel](&pb)
	}
}

func BenchmarkPBMO_Simple_ModelToPB(b *testing.B) {
	model := models.SimpleModel{
		ID:    uint64(1),
		Name:  "testuser",
		Email: "test@example.com",
		Age:   25,
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.ToPB[models.SimpleModel, models.SimplePB](&model)
	}
}

func BenchmarkNative_Simple_PBToModel(b *testing.B) {
	pb := models.SimplePB{
		Id:    uint64(1),
		Name:  "testuser",
		Email: "test@example.com",
		Age:   int32(25),
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.SimpleModel
	for i := 0; i < b.N; i++ {
		result = &models.SimpleModel{
			ID:    pb.Id,
			Name:  pb.Name,
			Email: pb.Email,
			Age:   int(pb.Age),
		}
	}
	benchSimpleModelResult = result
}

func BenchmarkNative_Simple_ModelToPB(b *testing.B) {
	model := models.SimpleModel{
		ID:    uint64(1),
		Name:  "testuser",
		Email: "test@example.com",
		Age:   25,
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.SimplePB
	for i := 0; i < b.N; i++ {
		result = &models.SimplePB{
			Id:    model.ID,
			Name:  model.Name,
			Email: model.Email,
			Age:   int32(model.Age),
		}
	}
	benchSimplePBResult = result
}

func BenchmarkPBMO_Simple_PBToModel_Parallel(b *testing.B) {
	simplePB := models.SimplePB{
		Id:    uint64(1),
		Name:  "testuser",
		Email: "test@example.com",
		Age:   int32(25),
	}

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		var result *models.SimpleModel
		for pb.Next() {
			result, _ = pbmo.FromPB[models.SimplePB, models.SimpleModel](&simplePB)
		}
		benchSimpleModelResult = result
	})
}

func BenchmarkNative_Simple_PBToModel_Parallel(b *testing.B) {
	simplePB := models.SimplePB{
		Id:    uint64(1),
		Name:  "testuser",
		Email: "test@example.com",
		Age:   int32(25),
	}

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(pb *testing.PB) {
		var result *models.SimpleModel
		for pb.Next() {
			result = &models.SimpleModel{
				ID:    simplePB.Id,
				Name:  simplePB.Name,
				Email: simplePB.Email,
				Age:   int(simplePB.Age),
			}
		}
		benchSimpleModelResult = result
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// AccountInfoPB Benchmarks (5 字段 + 时间戳)
// ══════════════════════════════════════════════════════════════════════════════

func BenchmarkPBMO_Account_PBToModel(b *testing.B) {
	now := time.Now()
	pb := models.AccountInfoPB{
		AccountId:  "acc001",
		Username:   "testuser",
		Status:     int32(1),
		CreatedAt:  timestamppb.New(now),
		LastActive: timestamppb.New(now),
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.FromPB[models.AccountInfoPB, models.AccountInfoModel](&pb)
	}
}

func BenchmarkNative_Account_PBToModel(b *testing.B) {
	now := time.Now()
	pb := models.AccountInfoPB{
		AccountId:  "acc001",
		Username:   "testuser",
		Status:     int32(1),
		CreatedAt:  timestamppb.New(now),
		LastActive: timestamppb.New(now),
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.AccountInfoModel
	for i := 0; i < b.N; i++ {
		result = &models.AccountInfoModel{
			AccountID:  pb.AccountId,
			Username:   pb.Username,
			Status:     int(pb.Status),
			CreatedAt:  pb.CreatedAt.AsTime(),
			LastActive: pb.LastActive.AsTime(),
		}
	}
	benchAccountModelResult = result
}

func BenchmarkPBMO_Account_ModelToPB(b *testing.B) {
	now := time.Now()
	model := models.AccountInfoModel{
		AccountID:  "acc001",
		Username:   "testuser",
		Status:     1,
		CreatedAt:  now,
		LastActive: now,
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.ToPB[models.AccountInfoModel, models.AccountInfoPB](&model)
	}
}

func BenchmarkNative_Account_ModelToPB(b *testing.B) {
	now := time.Now()
	model := models.AccountInfoModel{
		AccountID:  "acc001",
		Username:   "testuser",
		Status:     1,
		CreatedAt:  now,
		LastActive: now,
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.AccountInfoPB
	for i := 0; i < b.N; i++ {
		result = &models.AccountInfoPB{
			AccountId:  model.AccountID,
			Username:   model.Username,
			Status:     int32(model.Status),
			CreatedAt:  timestamppb.New(model.CreatedAt),
			LastActive: timestamppb.New(model.LastActive),
		}
	}
	benchAccountPBResult = result
}

// ══════════════════════════════════════════════════════════════════════════════
// MappedPB Benchmarks (3 字段 + 字段映射)
// ══════════════════════════════════════════════════════════════════════════════

func BenchmarkPBMO_Mapped_PBToModel(b *testing.B) {
	pb := models.MappedPB{
		ClientId:  uint64(123),
		UserName:  "testuser",
		UserEmail: "test@example.com",
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.FromPB[models.MappedPB, models.MappedModel](&pb)
	}
}

func BenchmarkNative_Mapped_PBToModel(b *testing.B) {
	pb := models.MappedPB{
		ClientId:  uint64(123),
		UserName:  "testuser",
		UserEmail: "test@example.com",
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.MappedModel
	for i := 0; i < b.N; i++ {
		result = &models.MappedModel{
			ID:    pb.ClientId,
			Name:  pb.UserName,
			Email: pb.UserEmail,
		}
	}
	benchMappedModelResult = result
}

func BenchmarkPBMO_Mapped_ModelToPB(b *testing.B) {
	model := models.MappedModel{
		ID:    uint64(123),
		Name:  "testuser",
		Email: "test@example.com",
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.ToPB[models.MappedModel, models.MappedPB](&model)
	}
}

func BenchmarkNative_Mapped_ModelToPB(b *testing.B) {
	model := models.MappedModel{
		ID:    uint64(123),
		Name:  "testuser",
		Email: "test@example.com",
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.MappedPB
	for i := 0; i < b.N; i++ {
		result = &models.MappedPB{
			ClientId:  model.ID,
			UserName:  model.Name,
			UserEmail: model.Email,
		}
	}
	benchMappedPBResult = result
}

// ══════════════════════════════════════════════════════════════════════════════
// 批量转换测试
// ══════════════════════════════════════════════════════════════════════════════

func BenchmarkPBMO_Simple_Batch100_PBToModel(b *testing.B) {
	pbs := make([]*models.SimplePB, 100)
	for i := range pbs {
		pbs[i] = &models.SimplePB{
			Id:    uint64(i),
			Name:  "user" + string(rune(i)),
			Email: "user" + string(rune(i)) + "@example.com",
			Age:   int32(20 + i%50),
		}
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.FromPBs[models.SimplePB, models.SimpleModel](pbs)
	}
}

func BenchmarkNative_Simple_Batch100_PBToModel(b *testing.B) {
	pbs := make([]*models.SimplePB, 100)
	for i := range pbs {
		pbs[i] = &models.SimplePB{
			Id:    uint64(i),
			Name:  "user" + string(rune(i)),
			Email: "user" + string(rune(i)) + "@example.com",
			Age:   int32(20 + i%50),
		}
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result []*models.SimpleModel
	for i := 0; i < b.N; i++ {
		result = make([]*models.SimpleModel, len(pbs))
		for j, pb := range pbs {
			result[j] = &models.SimpleModel{
				ID:    pb.Id,
				Name:  pb.Name,
				Email: pb.Email,
				Age:   int(pb.Age),
			}
		}
		benchSimpleModelsResult = result
	}
}

func BenchmarkPBMO_Simple_Batch1000_PBToModel(b *testing.B) {
	pbs := make([]*models.SimplePB, 1000)
	for i := range pbs {
		pbs[i] = &models.SimplePB{
			Id:    uint64(i),
			Name:  "user" + string(rune(i)),
			Email: "user" + string(rune(i)) + "@example.com",
			Age:   int32(20 + i%50),
		}
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.FromPBs[models.SimplePB, models.SimpleModel](pbs)
	}
}

func BenchmarkNative_Simple_Batch1000_PBToModel(b *testing.B) {
	pbs := make([]*models.SimplePB, 1000)
	for i := range pbs {
		pbs[i] = &models.SimplePB{
			Id:    uint64(i),
			Name:  "user" + string(rune(i)),
			Email: "user" + string(rune(i)) + "@example.com",
			Age:   int32(20 + i%50),
		}
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result []*models.SimpleModel
	for i := 0; i < b.N; i++ {
		result = make([]*models.SimpleModel, len(pbs))
		for j, pb := range pbs {
			result[j] = &models.SimpleModel{
				ID:    pb.Id,
				Name:  pb.Name,
				Email: pb.Email,
				Age:   int(pb.Age),
			}
		}
		benchSimpleModelsResult = result
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// 验证测试 - 确保 PBMO 和 Native 转换结果一致
// ══════════════════════════════════════════════════════════════════════════════

func TestSimple_PBToModel_Consistency(t *testing.T) {
	pb := models.SimplePB{
		Id:    uint64(1),
		Name:  "testuser",
		Email: "test@example.com",
		Age:   int32(25),
	}

	// PBMO 转换
	pbmoResult, err := pbmo.FromPB[models.SimplePB, models.SimpleModel](&pb)
	assert.Nil(t, err)
	assert.NotNil(t, pbmoResult)

	// Native 转换
	nativeResult := &models.SimpleModel{
		ID:    pb.Id,
		Name:  pb.Name,
		Email: pb.Email,
		Age:   int(pb.Age),
	}

	// 验证一致性
	assert.Equal(t, nativeResult.ID, pbmoResult.ID)
	assert.Equal(t, nativeResult.Name, pbmoResult.Name)
	assert.Equal(t, nativeResult.Email, pbmoResult.Email)
	assert.Equal(t, nativeResult.Age, pbmoResult.Age)
}

func TestSimple_ModelToPB_Consistency(t *testing.T) {
	model := models.SimpleModel{
		ID:    uint64(1),
		Name:  "testuser",
		Email: "test@example.com",
		Age:   25,
	}

	// PBMO 转换
	pbmoResult, err := pbmo.ToPB[models.SimpleModel, models.SimplePB](&model)
	assert.Nil(t, err)
	assert.NotNil(t, pbmoResult)

	// Native 转换
	nativeResult := &models.SimplePB{
		Id:    model.ID,
		Name:  model.Name,
		Email: model.Email,
		Age:   int32(model.Age),
	}

	// 验证一致性
	assert.Equal(t, nativeResult.Id, pbmoResult.Id)
	assert.Equal(t, nativeResult.Name, pbmoResult.Name)
	assert.Equal(t, nativeResult.Email, pbmoResult.Email)
	assert.Equal(t, nativeResult.Age, pbmoResult.Age)
}

func TestAccount_PBToModel_Consistency(t *testing.T) {
	now := time.Now()
	pb := models.AccountInfoPB{
		AccountId:  "acc001",
		Username:   "testuser",
		Status:     int32(1),
		CreatedAt:  timestamppb.New(now),
		LastActive: timestamppb.New(now),
	}

	// PBMO 转换
	pbmoResult, err := pbmo.FromPB[models.AccountInfoPB, models.AccountInfoModel](&pb)
	assert.Nil(t, err)
	assert.NotNil(t, pbmoResult)

	// Native 转换
	nativeResult := &models.AccountInfoModel{
		AccountID:  pb.AccountId,
		Username:   pb.Username,
		Status:     int(pb.Status),
		CreatedAt:  pb.CreatedAt.AsTime(),
		LastActive: pb.LastActive.AsTime(),
	}

	// 验证一致性
	assert.Equal(t, nativeResult.AccountID, pbmoResult.AccountID)
	assert.Equal(t, nativeResult.Username, pbmoResult.Username)
	assert.Equal(t, nativeResult.Status, pbmoResult.Status)
	assert.Equal(t, nativeResult.CreatedAt.Unix(), pbmoResult.CreatedAt.Unix())
	assert.Equal(t, nativeResult.LastActive.Unix(), pbmoResult.LastActive.Unix())
}

func TestMapped_PBToModel_Consistency(t *testing.T) {
	pb := models.MappedPB{
		ClientId:  uint64(123),
		UserName:  "testuser",
		UserEmail: "test@example.com",
	}

	// PBMO 转换
	pbmoResult, err := pbmo.FromPB[models.MappedPB, models.MappedModel](&pb)
	assert.Nil(t, err)
	assert.NotNil(t, pbmoResult)

	// Native 转换
	nativeResult := &models.MappedModel{
		ID:    pb.ClientId,
		Name:  pb.UserName,
		Email: pb.UserEmail,
	}

	// 验证一致性
	assert.Equal(t, nativeResult.ID, pbmoResult.ID)
	assert.Equal(t, nativeResult.Name, pbmoResult.Name)
	assert.Equal(t, nativeResult.Email, pbmoResult.Email)
}
