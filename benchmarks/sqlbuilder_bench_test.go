/**
 * @Author: developer 500000000@example.com
 * @Date: 2023-09-17 00:08:15
 * @LastEditors: developer 500000000@example.com
 * @LastEditTime: 2025-06-15 00:05:18
 * @FilePath: \go-pbmo-benchmark\benchmarks\sqlbuilder_bench_test.go
 * @Description: SQLBuilder性能测试 - PBMO vs Native
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

var (
	benchConfigModelResult        *models.ConfigModel
	benchConfigPBResult           *models.ConfigPB
	benchScoreBoardModelResult    *models.ScoreBoardModel
	benchScoreBoardPBResult       *models.ScoreBoardPB
	benchMetadataModelResult      *models.MetadataModel
	benchMetadataPBResult         *models.MetadataPB
	benchComprehensiveModelResult *models.ComprehensiveModel
	benchComprehensivePBResult    *models.ComprehensivePB
)

func init() {
	pbmo.Register[models.ConfigPB, models.ConfigModel]()
	pbmo.Register[models.ScoreEntryPB, models.ScoreEntry]()
	pbmo.Register[models.ScoreBoardPB, models.ScoreBoardModel]()
	pbmo.Register[models.MetadataPB, models.MetadataModel]()
	pbmo.Register[models.ComprehensivePB, models.ComprehensiveModel]()
}

func BenchmarkPBMO_Config_PBToModel(b *testing.B) {
	pb := models.ConfigPB{
		Name:   "config_test",
		Config: map[string]any{"theme": "dark", "language": "zh", "timeout": 30},
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.FromPB[models.ConfigPB, models.ConfigModel](&pb)
	}
}

func BenchmarkNative_Config_PBToModel(b *testing.B) {
	pb := models.ConfigPB{
		Name:   "config_test",
		Config: map[string]any{"theme": "dark", "language": "zh", "timezone": "Asia/Shanghai"},
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.ConfigModel
	for i := 0; i < b.N; i++ {
		var theme, language, timezone string
		if v, ok := pb.Config["theme"].(string); ok {
			theme = v
		}
		if v, ok := pb.Config["language"].(string); ok {
			language = v
		}
		if v, ok := pb.Config["timezone"].(string); ok {
			timezone = v
		}
		result = &models.ConfigModel{
			Name: pb.Name,
			Config: sqltypes.JSON[models.ConfigData]{Data: models.ConfigData{
				Theme:    theme,
				Language: language,
				Timezone: timezone,
			}},
		}
	}
	benchConfigModelResult = result
}

func BenchmarkPBMO_Config_ModelToPB(b *testing.B) {
	model := models.ConfigModel{
		Name: "config_test",
		Config: sqltypes.JSON[models.ConfigData]{Data: models.ConfigData{
			Theme:    "dark",
			Language: "zh",
			Timezone: "Asia/Shanghai",
		}},
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.ToPB[models.ConfigModel, models.ConfigPB](&model)
	}
}

func BenchmarkNative_Config_ModelToPB(b *testing.B) {
	model := models.ConfigModel{
		Name: "config_test",
		Config: sqltypes.JSON[models.ConfigData]{Data: models.ConfigData{
			Theme:    "dark",
			Language: "zh",
			Timezone: "Asia/Shanghai",
		}},
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.ConfigPB
	for i := 0; i < b.N; i++ {
		result = &models.ConfigPB{
			Name:   model.Name,
			Config: map[string]any{"theme": model.Config.Data.Theme, "language": model.Config.Data.Language, "timezone": model.Config.Data.Timezone},
		}
	}
	benchConfigPBResult = result
}

func BenchmarkPBMO_ScoreBoard_PBToModel(b *testing.B) {
	pb := models.ScoreBoardPB{
		Name: "scoreboard_test",
		Scores: []*models.ScoreEntryPB{
			{Label: "math", Value: 95},
			{Label: "english", Value: 88},
			{Label: "science", Value: 92},
		},
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.FromPB[models.ScoreBoardPB, models.ScoreBoardModel](&pb)
	}
}

func BenchmarkNative_ScoreBoard_PBToModel(b *testing.B) {
	pb := models.ScoreBoardPB{
		Name: "scoreboard_test",
		Scores: []*models.ScoreEntryPB{
			{Label: "math", Value: 95},
			{Label: "english", Value: 88},
			{Label: "science", Value: 92},
		},
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.ScoreBoardModel
	for i := 0; i < b.N; i++ {
		scores := make(sqltypes.Slice[models.ScoreEntry], len(pb.Scores))
		for j, s := range pb.Scores {
			scores[j] = models.ScoreEntry{Label: s.Label, Value: s.Value}
		}
		result = &models.ScoreBoardModel{
			Name:   pb.Name,
			Scores: scores,
		}
	}
	benchScoreBoardModelResult = result
}

func BenchmarkPBMO_ScoreBoard_ModelToPB(b *testing.B) {
	model := models.ScoreBoardModel{
		Name: "scoreboard_test",
		Scores: sqltypes.Slice[models.ScoreEntry]{
			{Label: "math", Value: 95},
			{Label: "english", Value: 88},
			{Label: "science", Value: 92},
		},
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.ToPB[models.ScoreBoardModel, models.ScoreBoardPB](&model)
	}
}

func BenchmarkNative_ScoreBoard_ModelToPB(b *testing.B) {
	model := models.ScoreBoardModel{
		Name: "scoreboard_test",
		Scores: sqltypes.Slice[models.ScoreEntry]{
			{Label: "math", Value: 95},
			{Label: "english", Value: 88},
			{Label: "science", Value: 92},
		},
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.ScoreBoardPB
	for i := 0; i < b.N; i++ {
		scores := make([]*models.ScoreEntryPB, len(model.Scores))
		for j, s := range model.Scores {
			scores[j] = &models.ScoreEntryPB{Label: s.Label, Value: s.Value}
		}
		result = &models.ScoreBoardPB{
			Name:   model.Name,
			Scores: scores,
		}
	}
	benchScoreBoardPBResult = result
}

func BenchmarkPBMO_Metadata_PBToModel(b *testing.B) {
	pb := models.MetadataPB{
		Name:     "metadata_test",
		Metadata: map[string]any{"key1": "val1", "key2": int64(42), "key3": true},
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.FromPB[models.MetadataPB, models.MetadataModel](&pb)
	}
}

func BenchmarkNative_Metadata_PBToModel(b *testing.B) {
	pb := models.MetadataPB{
		Name:     "metadata_test",
		Metadata: map[string]any{"key1": "val1", "key2": int64(42), "key3": true},
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.MetadataModel
	for i := 0; i < b.N; i++ {
		result = &models.MetadataModel{
			Name:     pb.Name,
			Metadata: sqltypes.MapAny(pb.Metadata),
		}
	}
	benchMetadataModelResult = result
}

func BenchmarkPBMO_Metadata_ModelToPB(b *testing.B) {
	model := models.MetadataModel{
		Name:     "metadata_test",
		Metadata: sqltypes.MapAny{"key1": "val1", "key2": int64(42), "key3": true},
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.ToPB[models.MetadataModel, models.MetadataPB](&model)
	}
}

func BenchmarkNative_Metadata_ModelToPB(b *testing.B) {
	model := models.MetadataModel{
		Name:     "metadata_test",
		Metadata: sqltypes.MapAny{"key1": "val1", "key2": int64(42), "key3": true},
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.MetadataPB
	for i := 0; i < b.N; i++ {
		result = &models.MetadataPB{
			Name:     model.Name,
			Metadata: map[string]any(model.Metadata),
		}
	}
	benchMetadataPBResult = result
}

func BenchmarkPBMO_Comprehensive_PBToModel(b *testing.B) {
	now := time.Now()
	enabled := true
	priority := int32(5)
	pb := models.ComprehensivePB{
		Id:        "comp001",
		Name:      "comprehensive_test",
		Status:    int32(1),
		Tags:      []string{"tag1", "tag2"},
		Config:    map[string]any{"timeout": 30},
		Scores:    []*models.ScoreEntryPB{{Label: "math", Value: 95}},
		CreatedAt: timestamppb.New(now),
		UpdatedAt: timestamppb.New(now),
		Enabled:   wrapperspb.Bool(enabled),
		Priority:  wrapperspb.Int32(priority),
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.FromPB[models.ComprehensivePB, models.ComprehensiveModel](&pb)
	}
}

func BenchmarkNative_Comprehensive_PBToModel(b *testing.B) {
	now := time.Now()
	enabled := true
	priority := int32(5)
	pb := models.ComprehensivePB{
		Id:        "comp001",
		Name:      "comprehensive_test",
		Status:    int32(1),
		Tags:      []string{"tag1", "tag2"},
		Config:    map[string]any{"timeout": 30},
		Scores:    []*models.ScoreEntryPB{{Label: "math", Value: 95}},
		CreatedAt: timestamppb.New(now),
		UpdatedAt: timestamppb.New(now),
		Enabled:   wrapperspb.Bool(enabled),
		Priority:  wrapperspb.Int32(priority),
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.ComprehensiveModel
	for i := 0; i < b.N; i++ {
		var enabledPtr *bool
		var priorityPtr *int32
		if pb.Enabled != nil {
			v := pb.Enabled.Value
			enabledPtr = &v
		}
		if pb.Priority != nil {
			v := pb.Priority.Value
			priorityPtr = &v
		}
		scores := make(sqltypes.Slice[models.ScoreEntry], len(pb.Scores))
		for j, s := range pb.Scores {
			scores[j] = models.ScoreEntry{Label: s.Label, Value: s.Value}
		}
		result = &models.ComprehensiveModel{
			ID:        pb.Id,
			Name:      pb.Name,
			Status:    int(pb.Status),
			Tags:      sqltypes.StringSlice(pb.Tags),
			Config:    sqltypes.JSON[models.ConfigData]{Data: models.ConfigData{Theme: "", Language: "", Timezone: ""}},
			Scores:    scores,
			CreatedAt: pb.CreatedAt.AsTime(),
			UpdatedAt: pb.UpdatedAt.AsTime(),
			Enabled:   enabledPtr,
			Priority:  priorityPtr,
		}
	}
	benchComprehensiveModelResult = result
}

func BenchmarkPBMO_Comprehensive_ModelToPB(b *testing.B) {
	now := time.Now()
	enabled := true
	priority := int32(5)
	model := models.ComprehensiveModel{
		ID:        "comp001",
		Name:      "comprehensive_test",
		Status:    1,
		Tags:      sqltypes.StringSlice{"tag1", "tag2"},
		Config:    sqltypes.JSON[models.ConfigData]{Data: models.ConfigData{Theme: "dark", Language: "zh", Timezone: "Asia/Shanghai"}},
		Scores:    sqltypes.Slice[models.ScoreEntry]{{Label: "math", Value: 95}},
		CreatedAt: now,
		UpdatedAt: now,
		Enabled:   &enabled,
		Priority:  &priority,
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.ToPB[models.ComprehensiveModel, models.ComprehensivePB](&model)
	}
}

func BenchmarkNative_Comprehensive_ModelToPB(b *testing.B) {
	now := time.Now()
	enabled := true
	priority := int32(5)
	model := models.ComprehensiveModel{
		ID:        "comp001",
		Name:      "comprehensive_test",
		Status:    1,
		Tags:      sqltypes.StringSlice{"tag1", "tag2"},
		Config:    sqltypes.JSON[models.ConfigData]{Data: models.ConfigData{Theme: "dark", Language: "zh", Timezone: "Asia/Shanghai"}},
		Scores:    sqltypes.Slice[models.ScoreEntry]{{Label: "math", Value: 95}},
		CreatedAt: now,
		UpdatedAt: now,
		Enabled:   &enabled,
		Priority:  &priority,
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.ComprehensivePB
	for i := 0; i < b.N; i++ {
		scores := make([]*models.ScoreEntryPB, len(model.Scores))
		for j, s := range model.Scores {
			scores[j] = &models.ScoreEntryPB{Label: s.Label, Value: s.Value}
		}
		pb := &models.ComprehensivePB{
			Id:        model.ID,
			Name:      model.Name,
			Status:    int32(model.Status),
			Tags:      []string(model.Tags),
			Config:    map[string]any{"theme": model.Config.Data.Theme, "language": model.Config.Data.Language},
			Scores:    scores,
			CreatedAt: timestamppb.New(model.CreatedAt),
			UpdatedAt: timestamppb.New(model.UpdatedAt),
		}
		if model.Enabled != nil {
			pb.Enabled = wrapperspb.Bool(*model.Enabled)
		}
		if model.Priority != nil {
			pb.Priority = wrapperspb.Int32(*model.Priority)
		}
		result = pb
	}
	benchComprehensivePBResult = result
}

func TestComprehensive_PBToModel_Consistency(t *testing.T) {
	now := time.Now()
	enabled := true
	priority := int32(5)
	pb := models.ComprehensivePB{
		Id:        "comp001",
		Name:      "comprehensive_test",
		Status:    int32(1),
		Tags:      []string{"tag1", "tag2"},
		Config:    map[string]any{"timeout": 30},
		Scores:    []*models.ScoreEntryPB{{Label: "math", Value: 95}},
		CreatedAt: timestamppb.New(now),
		UpdatedAt: timestamppb.New(now),
		Enabled:   wrapperspb.Bool(enabled),
		Priority:  wrapperspb.Int32(priority),
	}

	pbmoResult, err := pbmo.FromPB[models.ComprehensivePB, models.ComprehensiveModel](&pb)
	assert.Nil(t, err)
	assert.NotNil(t, pbmoResult)

	assert.Equal(t, pb.Id, pbmoResult.ID)
	assert.Equal(t, pb.Name, pbmoResult.Name)
	assert.Equal(t, int(pb.Status), pbmoResult.Status)
	assert.Equal(t, enabled, *pbmoResult.Enabled)
	assert.Equal(t, priority, *pbmoResult.Priority)
}

func TestMetadata_PBToModel_Consistency(t *testing.T) {
	pb := models.MetadataPB{
		Name:     "metadata_test",
		Metadata: map[string]any{"key1": "val1"},
	}

	pbmoResult, err := pbmo.FromPB[models.MetadataPB, models.MetadataModel](&pb)
	assert.Nil(t, err)
	assert.NotNil(t, pbmoResult)

	assert.Equal(t, pb.Name, pbmoResult.Name)
}
