/**
 * @Author: developer 500000000@example.com
 * @Date: 2023-09-17 00:08:15
 * @LastEditors: developer 500000000@example.com
 * @LastEditTime: 2025-06-15 00:05:18
 * @FilePath: \go-pbmo-benchmark\benchmarks\wrapper_bench_test.go
 * @Description: Wrapper性能测试 - PBMO vs Native
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
	"google.golang.org/protobuf/types/known/wrapperspb"
)

var (
	benchWrapperModelResult    *models.WrapperModel
	benchWrapperPBResult       *models.WrapperPB
	benchTimePtrModelResult    *models.TimePtrModel
	benchTimePtrPBResult       *models.TimePtrPB
	benchTimeZeroModelResult   *models.TimeZeroModel
	benchTimeZeroPBResult      *models.TimeZeroPB
	benchNamedSliceModelResult *models.NamedSliceModel
	benchNamedSlicePBResult    *models.NamedSlicePB
)

func init() {
	pbmo.Register[models.WrapperPB, models.WrapperModel]()
	pbmo.Register[models.TimePtrPB, models.TimePtrModel]()
	pbmo.Register[models.TimeZeroPB, models.TimeZeroModel]()
	pbmo.Register[models.NamedSlicePB, models.NamedSliceModel]()
}

func BenchmarkPBMO_Wrapper_PBToModel(b *testing.B) {
	isActive := true
	count := int32(10)
	score := 95.5
	label := "test"
	pb := models.WrapperPB{
		Name:     "wrapper_test",
		IsActive: wrapperspb.Bool(isActive),
		Count:    wrapperspb.Int32(count),
		Score:    wrapperspb.Double(score),
		Label:    wrapperspb.String(label),
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.FromPB[models.WrapperPB, models.WrapperModel](&pb)
	}
}

func BenchmarkNative_Wrapper_PBToModel(b *testing.B) {
	isActive := true
	count := int32(10)
	score := 95.5
	label := "test"
	pb := models.WrapperPB{
		Name:     "wrapper_test",
		IsActive: wrapperspb.Bool(isActive),
		Count:    wrapperspb.Int32(count),
		Score:    wrapperspb.Double(score),
		Label:    wrapperspb.String(label),
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.WrapperModel
	for i := 0; i < b.N; i++ {
		var isActivePtr *bool
		var countPtr *int32
		var scorePtr *float64
		var labelPtr *string
		if pb.IsActive != nil {
			v := pb.IsActive.Value
			isActivePtr = &v
		}
		if pb.Count != nil {
			v := pb.Count.Value
			countPtr = &v
		}
		if pb.Score != nil {
			v := pb.Score.Value
			scorePtr = &v
		}
		if pb.Label != nil {
			v := pb.Label.Value
			labelPtr = &v
		}
		result = &models.WrapperModel{
			Name:     pb.Name,
			IsActive: isActivePtr,
			Count:    countPtr,
			Score:    scorePtr,
			Label:    labelPtr,
		}
	}
	benchWrapperModelResult = result
}

func BenchmarkPBMO_Wrapper_ModelToPB(b *testing.B) {
	isActive := true
	count := int32(10)
	score := 95.5
	label := "test"
	model := models.WrapperModel{
		Name:     "wrapper_test",
		IsActive: &isActive,
		Count:    &count,
		Score:    &score,
		Label:    &label,
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.ToPB[models.WrapperModel, models.WrapperPB](&model)
	}
}

func BenchmarkNative_Wrapper_ModelToPB(b *testing.B) {
	isActive := true
	count := int32(10)
	score := 95.5
	label := "test"
	model := models.WrapperModel{
		Name:     "wrapper_test",
		IsActive: &isActive,
		Count:    &count,
		Score:    &score,
		Label:    &label,
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.WrapperPB
	for i := 0; i < b.N; i++ {
		pb := &models.WrapperPB{
			Name: model.Name,
		}
		if model.IsActive != nil {
			pb.IsActive = wrapperspb.Bool(*model.IsActive)
		}
		if model.Count != nil {
			pb.Count = wrapperspb.Int32(*model.Count)
		}
		if model.Score != nil {
			pb.Score = wrapperspb.Double(*model.Score)
		}
		if model.Label != nil {
			pb.Label = wrapperspb.String(*model.Label)
		}
		result = pb
	}
	benchWrapperPBResult = result
}

func BenchmarkPBMO_TimePtr_PBToModel(b *testing.B) {
	now := time.Now()
	pb := models.TimePtrPB{
		Name:        "timeptr_test",
		ScheduledAt: timestamppb.New(now),
		ReleasedAt:  timestamppb.New(now.AddDate(0, 0, 1)),
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.FromPB[models.TimePtrPB, models.TimePtrModel](&pb)
	}
}

func BenchmarkNative_TimePtr_PBToModel(b *testing.B) {
	now := time.Now()
	pb := models.TimePtrPB{
		Name:        "timeptr_test",
		ScheduledAt: timestamppb.New(now),
		ReleasedAt:  timestamppb.New(now.AddDate(0, 0, 1)),
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.TimePtrModel
	for i := 0; i < b.N; i++ {
		var scheduledAt, releasedAt *time.Time
		if pb.ScheduledAt != nil {
			t := pb.ScheduledAt.AsTime()
			scheduledAt = &t
		}
		if pb.ReleasedAt != nil {
			t := pb.ReleasedAt.AsTime()
			releasedAt = &t
		}
		result = &models.TimePtrModel{
			Name:        pb.Name,
			ScheduledAt: scheduledAt,
			ReleasedAt:  releasedAt,
		}
	}
	benchTimePtrModelResult = result
}

func BenchmarkPBMO_TimePtr_ModelToPB(b *testing.B) {
	now := time.Now()
	model := models.TimePtrModel{
		Name:        "timeptr_test",
		ScheduledAt: &now,
		ReleasedAt:  nil,
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.ToPB[models.TimePtrModel, models.TimePtrPB](&model)
	}
}

func BenchmarkNative_TimePtr_ModelToPB(b *testing.B) {
	now := time.Now()
	model := models.TimePtrModel{
		Name:        "timeptr_test",
		ScheduledAt: &now,
		ReleasedAt:  nil,
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.TimePtrPB
	for i := 0; i < b.N; i++ {
		pb := &models.TimePtrPB{
			Name: model.Name,
		}
		if model.ScheduledAt != nil {
			pb.ScheduledAt = timestamppb.New(*model.ScheduledAt)
		}
		if model.ReleasedAt != nil {
			pb.ReleasedAt = timestamppb.New(*model.ReleasedAt)
		}
		result = pb
	}
	benchTimePtrPBResult = result
}

func BenchmarkPBMO_TimeZero_PBToModel(b *testing.B) {
	now := time.Now()
	pb := models.TimeZeroPB{
		Name:      "timezero_test",
		CreatedAt: timestamppb.New(now),
		UpdatedAt: timestamppb.New(now),
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.FromPB[models.TimeZeroPB, models.TimeZeroModel](&pb)
	}
}

func BenchmarkNative_TimeZero_PBToModel(b *testing.B) {
	now := time.Now()
	pb := models.TimeZeroPB{
		Name:      "timezero_test",
		CreatedAt: timestamppb.New(now),
		UpdatedAt: timestamppb.New(now),
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.TimeZeroModel
	for i := 0; i < b.N; i++ {
		result = &models.TimeZeroModel{
			Name:      pb.Name,
			CreatedAt: pb.CreatedAt.AsTime(),
			UpdatedAt: pb.UpdatedAt.AsTime(),
		}
	}
	benchTimeZeroModelResult = result
}

func BenchmarkPBMO_TimeZero_ModelToPB(b *testing.B) {
	now := time.Now()
	model := models.TimeZeroModel{
		Name:      "timezero_test",
		CreatedAt: now,
		UpdatedAt: now,
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.ToPB[models.TimeZeroModel, models.TimeZeroPB](&model)
	}
}

func BenchmarkNative_TimeZero_ModelToPB(b *testing.B) {
	now := time.Now()
	model := models.TimeZeroModel{
		Name:      "timezero_test",
		CreatedAt: now,
		UpdatedAt: now,
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.TimeZeroPB
	for i := 0; i < b.N; i++ {
		result = &models.TimeZeroPB{
			Name:      model.Name,
			CreatedAt: timestamppb.New(model.CreatedAt),
			UpdatedAt: timestamppb.New(model.UpdatedAt),
		}
	}
	benchTimeZeroPBResult = result
}

func BenchmarkPBMO_NamedSlice_PBToModel(b *testing.B) {
	pb := models.NamedSlicePB{
		Name:  "namedslice_test",
		Tags:  []string{"tag1", "tag2", "tag3"},
		Items: []string{"item1", "item2", "item3"},
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.FromPB[models.NamedSlicePB, models.NamedSliceModel](&pb)
	}
}

func BenchmarkNative_NamedSlice_PBToModel(b *testing.B) {
	pb := models.NamedSlicePB{
		Name:  "namedslice_test",
		Tags:  []string{"tag1", "tag2", "tag3"},
		Items: []string{"item1", "item2", "item3"},
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.NamedSliceModel
	for i := 0; i < b.N; i++ {
		result = &models.NamedSliceModel{
			Name:  pb.Name,
			Tags:  models.StringSlice(pb.Tags),
			Items: models.StringSlice(pb.Items),
		}
	}
	benchNamedSliceModelResult = result
}

func BenchmarkPBMO_NamedSlice_ModelToPB(b *testing.B) {
	model := models.NamedSliceModel{
		Name:  "namedslice_test",
		Tags:  models.StringSlice{"tag1", "tag2", "tag3"},
		Items: models.StringSlice{"item1", "item2", "item3"},
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.ToPB[models.NamedSliceModel, models.NamedSlicePB](&model)
	}
}

func BenchmarkNative_NamedSlice_ModelToPB(b *testing.B) {
	model := models.NamedSliceModel{
		Name:  "namedslice_test",
		Tags:  models.StringSlice{"tag1", "tag2", "tag3"},
		Items: models.StringSlice{"item1", "item2", "item3"},
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.NamedSlicePB
	for i := 0; i < b.N; i++ {
		result = &models.NamedSlicePB{
			Name:  model.Name,
			Tags:  []string(model.Tags),
			Items: []string(model.Items),
		}
	}
	benchNamedSlicePBResult = result
}

func BenchmarkPBMO_Wrapper_PBToModel_Parallel(b *testing.B) {
	isActive := true
	count := int32(10)
	score := 95.5
	label := "test"
	wrapperPB := models.WrapperPB{
		Name:     "wrapper_test",
		IsActive: wrapperspb.Bool(isActive),
		Count:    wrapperspb.Int32(count),
		Score:    wrapperspb.Double(score),
		Label:    wrapperspb.String(label),
	}

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		var result *models.WrapperModel
		for p.Next() {
			result, _ = pbmo.FromPB[models.WrapperPB, models.WrapperModel](&wrapperPB)
		}
		benchWrapperModelResult = result
	})
}

func BenchmarkNative_Wrapper_PBToModel_Parallel(b *testing.B) {
	isActive := true
	count := int32(10)
	score := 95.5
	label := "test"
	pb := models.WrapperPB{
		Name:     "wrapper_test",
		IsActive: wrapperspb.Bool(isActive),
		Count:    wrapperspb.Int32(count),
		Score:    wrapperspb.Double(score),
		Label:    wrapperspb.String(label),
	}

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(tpb *testing.PB) {
		var result *models.WrapperModel
		for tpb.Next() {
			var isActivePtr *bool
			var countPtr *int32
			var scorePtr *float64
			var labelPtr *string
			if pb.IsActive != nil {
				v := pb.IsActive.Value
				isActivePtr = &v
			}
			if pb.Count != nil {
				v := pb.Count.Value
				countPtr = &v
			}
			if pb.Score != nil {
				v := pb.Score.Value
				scorePtr = &v
			}
			if pb.Label != nil {
				v := pb.Label.Value
				labelPtr = &v
			}
			result = &models.WrapperModel{
				Name:     pb.Name,
				IsActive: isActivePtr,
				Count:    countPtr,
				Score:    scorePtr,
				Label:    labelPtr,
			}
		}
		benchWrapperModelResult = result
	})
}

func TestWrapper_PBToModel_Consistency(t *testing.T) {
	isActive := true
	count := int32(10)
	score := 95.5
	label := "test"
	pb := models.WrapperPB{
		Name:     "wrapper_test",
		IsActive: wrapperspb.Bool(isActive),
		Count:    wrapperspb.Int32(count),
		Score:    wrapperspb.Double(score),
		Label:    wrapperspb.String(label),
	}

	pbmoResult, err := pbmo.FromPB[models.WrapperPB, models.WrapperModel](&pb)
	assert.Nil(t, err)
	assert.NotNil(t, pbmoResult)

	assert.Equal(t, pb.Name, pbmoResult.Name)
	assert.Equal(t, isActive, *pbmoResult.IsActive)
	assert.Equal(t, count, *pbmoResult.Count)
	assert.Equal(t, score, *pbmoResult.Score)
	assert.Equal(t, label, *pbmoResult.Label)
}

func TestTimePtr_PBToModel_Consistency(t *testing.T) {
	now := time.Now()
	pb := models.TimePtrPB{
		Name:        "timeptr_test",
		ScheduledAt: timestamppb.New(now),
		ReleasedAt:  nil,
	}

	pbmoResult, err := pbmo.FromPB[models.TimePtrPB, models.TimePtrModel](&pb)
	assert.Nil(t, err)
	assert.NotNil(t, pbmoResult)

	assert.Equal(t, pb.Name, pbmoResult.Name)
	assert.NotNil(t, pbmoResult.ScheduledAt)
	assert.Nil(t, pbmoResult.ReleasedAt)
}

func TestNamedSlice_PBToModel_Consistency(t *testing.T) {
	pb := models.NamedSlicePB{
		Name:  "namedslice_test",
		Tags:  []string{"tag1", "tag2", "tag3"},
		Items: []string{"item1", "item2", "item3"},
	}

	pbmoResult, err := pbmo.FromPB[models.NamedSlicePB, models.NamedSliceModel](&pb)
	assert.Nil(t, err)
	assert.NotNil(t, pbmoResult)

	assert.Equal(t, pb.Name, pbmoResult.Name)
	assert.Equal(t, []string(pb.Tags), []string(pbmoResult.Tags))
	assert.Equal(t, []string(pb.Items), []string(pbmoResult.Items))
}
