/**
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-06-15 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-06-15 00:00:00
 * @FilePath: \go-pbmo-benchmark\benchmarks\internal_bench_test.go
 * @Description: 内部功能基准测试 - 转换器创建/注册/Transformer/枚举/校验/SafeConverter
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */

package benchmarks

import (
	"fmt"
	"reflect"
	"testing"
	"time"

	"github.com/kamalyes/go-pbmo"
	"github.com/kamalyes/go-pbmo-benchmark/models"
	"github.com/stretchr/testify/assert"
	"google.golang.org/protobuf/types/known/timestamppb"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

// ══════════════════════════════════════════════════════════════════════════════
// 全局变量，防止编译器优化
// ══════════════════════════════════════════════════════════════════════════════

var (
	benchConverterResult     *pbmo.BidiConverter
	benchTransformerResult   reflect.Value
	benchSafeConverterResult *models.SimpleModel
	benchInternalTypesResult nativeBidiMetadata
	benchInternalEnumResult  int32
	benchInternalBoolResult  bool
	benchInternalIntResult   int
	benchInternalMappedModel *models.MappedModel
	benchInternalMappedPB    *models.MappedPB
	benchInternalModels      []models.SimpleModel
	benchInternalPBs         []models.SimplePB
	benchInternalModelPtrs   []*models.SimpleModel
	benchInternalWrapperPB   *models.WrapperPB
	benchInternalTimeZeroPB  *models.TimeZeroPB
)

type nativeBidiMetadata struct {
	pbType    reflect.Type
	modelType reflect.Type
	autoTime  bool
	validate  bool
	fieldMap  map[string]string
}

// ══════════════════════════════════════════════════════════════════════════════
// 转换器创建基准
// ══════════════════════════════════════════════════════════════════════════════

func BenchmarkPBMO_NewBidiConverter_Internal(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		pbmo.NewBidiConverter(models.SimplePB{}, models.SimpleModel{})
	}
}

func BenchmarkPBMO_NewBidiConverter_WithOptions_Internal(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		pbmo.NewBidiConverter(
			models.SimplePB{}, models.SimpleModel{},
			pbmo.WithAutoTimeConversion(true),
			pbmo.WithValidation(false),
			pbmo.WithFieldMapping("ID", "Id"),
		)
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// 枚举映射基准
// ══════════════════════════════════════════════════════════════════════════════

func BenchmarkPBMO_EnumMapper_Map_Internal(b *testing.B) {
	mapper := pbmo.NewEnumMapper()
	for i := int32(0); i < 100; i++ {
		mapper.AddMapping(i, i*10)
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		mapper.Map(int32(i%100), 0)
	}
}

func BenchmarkPBMO_GenericEnumMapper_Map_Internal(b *testing.B) {
	mapper := pbmo.NewGenericEnumMapper[int32, int32](0)
	for i := int32(0); i < 100; i++ {
		mapper.Register(i, i*10)
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		mapper.Map(int32(i % 100))
	}
}

func BenchmarkPBMO_AutoEnumConverter_Convert_Internal(b *testing.B) {
	type ProtoStatus int32
	type WsStatus int32

	converter := pbmo.NewAutoEnumConverter[ProtoStatus, WsStatus](0)
	mappings := make(map[ProtoStatus]WsStatus, 100)
	for i := ProtoStatus(0); i < 100; i++ {
		mappings[i] = WsStatus(i) * 10
	}
	converter.AutoRegister(mappings)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		converter.Convert(ProtoStatus(i % 100))
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// 校验器基准
// ══════════════════════════════════════════════════════════════════════════════

func BenchmarkPBMO_Validator_Validate_Internal(b *testing.B) {
	validator := pbmo.NewValidator()
	validator.RegisterRules("SimpleModel",
		pbmo.FieldRule{Name: "Name", Required: true, MinLen: 1, MaxLen: 100},
		pbmo.FieldRule{Name: "Age", Min: 0, Max: 1000},
	)

	model := models.SimpleModel{Name: "testuser", Age: 25}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if err := validator.Validate(&model); err != nil {
			b.Fatal(err)
		}
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// 注册中心基准
// ══════════════════════════════════════════════════════════════════════════════

func BenchmarkPBMO_Registry_Register_Internal(b *testing.B) {
	registry := pbmo.NewRegistry()
	converter := pbmo.NewBidiConverter(models.SimplePB{}, models.SimpleModel{})

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		registry.Unregister(converter.GetPBType(), converter.GetModelType())
		registry.MustRegister(converter)
	}
}

func BenchmarkPBMO_Registry_Lookup_Internal(b *testing.B) {
	registry := pbmo.NewRegistry()
	converter := pbmo.NewBidiConverter(models.SimplePB{}, models.SimpleModel{})
	registry.MustRegister(converter)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		if _, err := registry.LookupByInstance(models.SimplePB{}, models.SimpleModel{}); err != nil {
			b.Fatal(err)
		}
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// Transformer 基准
// ══════════════════════════════════════════════════════════════════════════════

func BenchmarkPBMO_TransformerRegistry_Apply_Internal(b *testing.B) {
	tr := pbmo.NewTransformerRegistry()
	tr.Register("Name", func(v interface{}) interface{} {
		return fmt.Sprintf("prefix_%s", v.(string))
	})

	val := reflect.ValueOf("hello")

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		benchTransformerResult = tr.Apply("Name", val)
	}
}

func BenchmarkPBMO_TransformerRegistry_Count_Internal(b *testing.B) {
	tr := pbmo.NewTransformerRegistry()
	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("field_%d", i)
		tr.Register(name, func(v interface{}) interface{} { return v })
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_ = tr.Count()
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// SafeConverter 基准
// ══════════════════════════════════════════════════════════════════════════════

func BenchmarkPBMO_SafeConverter_SafeConvertPBToModel_Internal(b *testing.B) {
	converter := pbmo.NewSafeConverter(models.SimplePB{}, models.SimpleModel{})
	pb := models.SimplePB{Id: 1, Name: "hello", Email: "test@example.com", Age: 25}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var model models.SimpleModel
		if err := converter.SafeConvertPBToModel(pb, &model); err != nil {
			b.Fatal(err)
		}
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// 带字段映射的转换器基准（不同于 simple_bench 中的 Mapped）
// ══════════════════════════════════════════════════════════════════════════════

func BenchmarkPBMO_Mapped_WithFieldMapping_PBToModel_Internal(b *testing.B) {
	converter := pbmo.NewBidiConverter(
		models.MappedPB{}, models.MappedModel{},
		pbmo.WithFieldMapping("ID", "ClientId"),
		pbmo.WithFieldMapping("Name", "UserName"),
		pbmo.WithFieldMapping("Email", "UserEmail"),
	)
	pb := models.MappedPB{ClientId: 1, UserName: "test", UserEmail: "test@example.com"}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var model models.MappedModel
		if err := converter.ConvertPBToModel(pb, &model); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkPBMO_Mapped_WithFieldMapping_ModelToPB_Internal(b *testing.B) {
	converter := pbmo.NewBidiConverter(
		models.MappedPB{}, models.MappedModel{},
		pbmo.WithFieldMapping("ID", "ClientId"),
		pbmo.WithFieldMapping("Name", "UserName"),
		pbmo.WithFieldMapping("Email", "UserEmail"),
	)
	modelData := models.MappedModel{ID: 1, Name: "test", Email: "test@example.com"}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var pb models.MappedPB
		if err := converter.ConvertModelToPB(modelData, &pb); err != nil {
			b.Fatal(err)
		}
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// 带 Transformer 的转换基准
// ══════════════════════════════════════════════════════════════════════════════

func BenchmarkPBMO_Simple_WithTransformer_PBToModel_Internal(b *testing.B) {
	converter := pbmo.NewBidiConverter(models.SimplePB{}, models.SimpleModel{})
	converter.RegisterTransformer("Name", func(v interface{}) interface{} {
		return fmt.Sprintf("prefix_%s", v.(string))
	})
	pb := models.SimplePB{Id: 1, Name: "hello", Email: "test@example.com", Age: 25}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var model models.SimpleModel
		if err := converter.ConvertPBToModel(pb, &model); err != nil {
			b.Fatal(err)
		}
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// 批量转换基准
// ══════════════════════════════════════════════════════════════════════════════

func BenchmarkPBMO_Simple_BatchConvert100_PBToModel_Internal(b *testing.B) {
	converter := pbmo.NewBidiConverter(models.SimplePB{}, models.SimpleModel{})
	pbs := make([]models.SimplePB, 100)
	for i := range pbs {
		pbs[i] = models.SimplePB{Id: uint64(i), Name: "user", Email: "u@t.com", Age: int32(20 + i%50)}
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var ms []models.SimpleModel
		if err := converter.BatchConvertPBToModel(pbs, &ms); err != nil {
			b.Fatal(err)
		}
	}
}

func BenchmarkPBMO_Simple_BatchConvert100_ModelToPB_Internal(b *testing.B) {
	converter := pbmo.NewBidiConverter(models.SimplePB{}, models.SimpleModel{})
	ms := make([]models.SimpleModel, 100)
	for i := range ms {
		ms[i] = models.SimpleModel{ID: uint64(i), Name: "user", Email: "u@t.com", Age: 20 + i%50}
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var pbs []models.SimplePB
		if err := converter.BatchConvertModelToPB(ms, &pbs); err != nil {
			b.Fatal(err)
		}
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// SafeConverter 批量基准
// ══════════════════════════════════════════════════════════════════════════════

func BenchmarkPBMO_SafeBatchConvert100_PBToModel_Internal(b *testing.B) {
	pbs := make([]*models.SimplePB, 100)
	for i := range pbs {
		pbs[i] = &models.SimplePB{Id: uint64(i), Name: "user", Email: "u@t.com", Age: int32(20 + i%50)}
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.SafeFromPBs[models.SimplePB, models.SimpleModel](pbs)
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// Wrapper 字段转换基准（不与 wrapper_bench 重复的部分）
// ══════════════════════════════════════════════════════════════════════════════

func BenchmarkPBMO_WrapperField_AllTypes_ModelToPB_Internal(b *testing.B) {
	minVal := int32(10)
	scoreVal := 99.5
	labelVal := "hello"
	activeVal := true
	model := &models.WrapperModel{
		Name:     "bench-wrapper",
		IsActive: &activeVal,
		Count:    &minVal,
		Score:    &scoreVal,
		Label:    &labelVal,
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.ToPB[models.WrapperModel, models.WrapperPB](model)
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// 零值时间转换基准
// ══════════════════════════════════════════════════════════════════════════════

func BenchmarkPBMO_TimeZeroZeroValue_ModelToPB_Internal(b *testing.B) {
	pbmo.Register[models.TimeZeroModel, models.TimeZeroPB]()
	modelData := &models.TimeZeroModel{
		Name:      "zero",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.ToPB[models.TimeZeroModel, models.TimeZeroPB](modelData)
	}
}

func BenchmarkPBMO_TimeZeroValidValue_ModelToPB_Internal(b *testing.B) {
	pbmo.Register[models.TimeZeroModel, models.TimeZeroPB]()
	now := time.Now()
	modelData := &models.TimeZeroModel{
		Name:      "valid",
		CreatedAt: now,
		UpdatedAt: now,
	}

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.ToPB[models.TimeZeroModel, models.TimeZeroPB](modelData)
	}
}

// ══════════════════════════════════════════════════════════════════════════════
// 并发基准
// ══════════════════════════════════════════════════════════════════════════════

func BenchmarkPBMO_Simple_PBToModel_Parallel_Internal(b *testing.B) {
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
		benchSafeConverterResult = result
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// 一致性验证测试
// ══════════════════════════════════════════════════════════════════════════════

func BenchmarkNative_NewBidiConverter_Internal(b *testing.B) {
	b.ReportAllocs()
	var result nativeBidiMetadata
	for i := 0; i < b.N; i++ {
		result = nativeBidiMetadata{
			pbType:    reflect.TypeOf(models.SimplePB{}),
			modelType: reflect.TypeOf(models.SimpleModel{}),
			autoTime:  true,
			validate:  false,
		}
	}
	benchInternalTypesResult = result
}

func BenchmarkNative_NewBidiConverter_WithOptions_Internal(b *testing.B) {
	b.ReportAllocs()
	var result nativeBidiMetadata
	for i := 0; i < b.N; i++ {
		result = nativeBidiMetadata{
			pbType:    reflect.TypeOf(models.SimplePB{}),
			modelType: reflect.TypeOf(models.SimpleModel{}),
			autoTime:  true,
			validate:  false,
			fieldMap: map[string]string{
				"ID": "Id",
			},
		}
	}
	benchInternalTypesResult = result
}

func BenchmarkNative_EnumMapper_Map_Internal(b *testing.B) {
	mappings := make(map[int32]int32, 100)
	for i := int32(0); i < 100; i++ {
		mappings[i] = i * 10
	}

	b.ResetTimer()
	b.ReportAllocs()
	var result int32
	for i := 0; i < b.N; i++ {
		v, ok := mappings[int32(i%100)]
		if !ok {
			v = 0
		}
		result = v
	}
	benchInternalEnumResult = result
}

func BenchmarkNative_GenericEnumMapper_Map_Internal(b *testing.B) {
	mappings := make(map[int32]int32, 100)
	for i := int32(0); i < 100; i++ {
		mappings[i] = i * 10
	}

	b.ResetTimer()
	b.ReportAllocs()
	var result int32
	for i := 0; i < b.N; i++ {
		v, ok := mappings[int32(i%100)]
		if !ok {
			v = 0
		}
		result = v
	}
	benchInternalEnumResult = result
}

func BenchmarkNative_AutoEnumConverter_Convert_Internal(b *testing.B) {
	type ProtoStatus int32
	type WsStatus int32

	mappings := make(map[ProtoStatus]WsStatus, 100)
	for i := ProtoStatus(0); i < 100; i++ {
		mappings[i] = WsStatus(i) * 10
	}

	b.ResetTimer()
	b.ReportAllocs()
	var result WsStatus
	for i := 0; i < b.N; i++ {
		v, ok := mappings[ProtoStatus(i%100)]
		if !ok {
			v = 0
		}
		result = v
	}
	benchInternalEnumResult = int32(result)
}

func BenchmarkNative_Validator_Validate_Internal(b *testing.B) {
	model := models.SimpleModel{Name: "testuser", Age: 25}

	b.ResetTimer()
	b.ReportAllocs()
	var result bool
	for i := 0; i < b.N; i++ {
		result = model.Name != "" &&
			len(model.Name) >= 1 &&
			len(model.Name) <= 100 &&
			model.Age >= 0 &&
			model.Age <= 1000
		if !result {
			b.Fatal("validation failed")
		}
	}
	benchInternalBoolResult = result
}

func BenchmarkNative_Registry_Register_Internal(b *testing.B) {
	converter := pbmo.NewBidiConverter(models.SimplePB{}, models.SimpleModel{})
	key := [2]reflect.Type{reflect.TypeOf(models.SimplePB{}), reflect.TypeOf(models.SimpleModel{})}
	registry := make(map[[2]reflect.Type]*pbmo.BidiConverter, 1)

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		delete(registry, key)
		registry[key] = converter
	}
	benchConverterResult = registry[key]
}

func BenchmarkNative_Registry_Lookup_Internal(b *testing.B) {
	converter := pbmo.NewBidiConverter(models.SimplePB{}, models.SimpleModel{})
	key := [2]reflect.Type{reflect.TypeOf(models.SimplePB{}), reflect.TypeOf(models.SimpleModel{})}
	registry := map[[2]reflect.Type]*pbmo.BidiConverter{key: converter}

	b.ResetTimer()
	b.ReportAllocs()
	var result *pbmo.BidiConverter
	for i := 0; i < b.N; i++ {
		result = registry[key]
		if result == nil {
			b.Fatal("converter not found")
		}
	}
	benchConverterResult = result
}

func BenchmarkNative_TransformerRegistry_Apply_Internal(b *testing.B) {
	transformer := func(v interface{}) interface{} {
		return fmt.Sprintf("prefix_%s", v.(string))
	}
	val := reflect.ValueOf("hello")

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		benchTransformerResult = reflect.ValueOf(transformer(val.Interface()))
	}
}

func BenchmarkNative_TransformerRegistry_Count_Internal(b *testing.B) {
	transformers := make(map[string]func(interface{}) interface{}, 10)
	for i := 0; i < 10; i++ {
		name := fmt.Sprintf("field_%d", i)
		transformers[name] = func(v interface{}) interface{} { return v }
	}

	b.ResetTimer()
	b.ReportAllocs()
	var result int
	for i := 0; i < b.N; i++ {
		result = len(transformers)
	}
	benchInternalIntResult = result
}

func BenchmarkNative_SafeConverter_SafeConvertPBToModel_Internal(b *testing.B) {
	pb := models.SimplePB{Id: 1, Name: "hello", Email: "test@example.com", Age: 25}

	b.ResetTimer()
	b.ReportAllocs()
	var result *models.SimpleModel
	for i := 0; i < b.N; i++ {
		result = &models.SimpleModel{
			ID:    pb.Id,
			Name:  pb.Name,
			Email: pb.Email,
			Age:   int(pb.Age),
		}
	}
	benchSafeConverterResult = result
}

func BenchmarkNative_Mapped_WithFieldMapping_PBToModel_Internal(b *testing.B) {
	pb := models.MappedPB{ClientId: 1, UserName: "test", UserEmail: "test@example.com"}

	b.ResetTimer()
	b.ReportAllocs()
	var result *models.MappedModel
	for i := 0; i < b.N; i++ {
		result = &models.MappedModel{
			ID:    pb.ClientId,
			Name:  pb.UserName,
			Email: pb.UserEmail,
		}
	}
	benchInternalMappedModel = result
}

func BenchmarkNative_Mapped_WithFieldMapping_ModelToPB_Internal(b *testing.B) {
	modelData := models.MappedModel{ID: 1, Name: "test", Email: "test@example.com"}

	b.ResetTimer()
	b.ReportAllocs()
	var result *models.MappedPB
	for i := 0; i < b.N; i++ {
		result = &models.MappedPB{
			ClientId:  modelData.ID,
			UserName:  modelData.Name,
			UserEmail: modelData.Email,
		}
	}
	benchInternalMappedPB = result
}

func BenchmarkNative_Simple_WithTransformer_PBToModel_Internal(b *testing.B) {
	pb := models.SimplePB{Id: 1, Name: "hello", Email: "test@example.com", Age: 25}

	b.ResetTimer()
	b.ReportAllocs()
	var result *models.SimpleModel
	for i := 0; i < b.N; i++ {
		result = &models.SimpleModel{
			ID:    pb.Id,
			Name:  fmt.Sprintf("prefix_%s", pb.Name),
			Email: pb.Email,
			Age:   int(pb.Age),
		}
	}
	benchSafeConverterResult = result
}

func BenchmarkNative_Simple_BatchConvert100_PBToModel_Internal(b *testing.B) {
	pbs := make([]models.SimplePB, 100)
	for i := range pbs {
		pbs[i] = models.SimplePB{Id: uint64(i), Name: "user", Email: "u@t.com", Age: int32(20 + i%50)}
	}

	b.ResetTimer()
	b.ReportAllocs()
	var result []models.SimpleModel
	for i := 0; i < b.N; i++ {
		result = make([]models.SimpleModel, len(pbs))
		for j := range pbs {
			result[j] = models.SimpleModel{
				ID:    pbs[j].Id,
				Name:  pbs[j].Name,
				Email: pbs[j].Email,
				Age:   int(pbs[j].Age),
			}
		}
	}
	benchInternalModels = result
}

func BenchmarkNative_Simple_BatchConvert100_ModelToPB_Internal(b *testing.B) {
	ms := make([]models.SimpleModel, 100)
	for i := range ms {
		ms[i] = models.SimpleModel{ID: uint64(i), Name: "user", Email: "u@t.com", Age: 20 + i%50}
	}

	b.ResetTimer()
	b.ReportAllocs()
	var result []models.SimplePB
	for i := 0; i < b.N; i++ {
		result = make([]models.SimplePB, len(ms))
		for j := range ms {
			result[j] = models.SimplePB{
				Id:    ms[j].ID,
				Name:  ms[j].Name,
				Email: ms[j].Email,
				Age:   int32(ms[j].Age),
			}
		}
	}
	benchInternalPBs = result
}

func BenchmarkNative_SafeBatchConvert100_PBToModel_Internal(b *testing.B) {
	pbs := make([]*models.SimplePB, 100)
	for i := range pbs {
		pbs[i] = &models.SimplePB{Id: uint64(i), Name: "user", Email: "u@t.com", Age: int32(20 + i%50)}
	}

	b.ResetTimer()
	b.ReportAllocs()
	var result []*models.SimpleModel
	var successCount int
	for i := 0; i < b.N; i++ {
		result = make([]*models.SimpleModel, len(pbs))
		successCount = 0
		for j, pb := range pbs {
			if pb == nil {
				continue
			}
			result[j] = &models.SimpleModel{
				ID:    pb.Id,
				Name:  pb.Name,
				Email: pb.Email,
				Age:   int(pb.Age),
			}
			successCount++
		}
	}
	benchInternalModelPtrs = result
	benchInternalIntResult = successCount
}

func BenchmarkNative_WrapperField_AllTypes_ModelToPB_Internal(b *testing.B) {
	minVal := int32(10)
	scoreVal := 99.5
	labelVal := "hello"
	activeVal := true
	model := &models.WrapperModel{
		Name:     "bench-wrapper",
		IsActive: &activeVal,
		Count:    &minVal,
		Score:    &scoreVal,
		Label:    &labelVal,
	}

	b.ResetTimer()
	b.ReportAllocs()
	var result *models.WrapperPB
	for i := 0; i < b.N; i++ {
		pb := &models.WrapperPB{Name: model.Name}
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
	benchInternalWrapperPB = result
}

func BenchmarkNative_TimeZeroZeroValue_ModelToPB_Internal(b *testing.B) {
	modelData := &models.TimeZeroModel{
		Name:      "zero",
		CreatedAt: time.Time{},
		UpdatedAt: time.Time{},
	}

	b.ResetTimer()
	b.ReportAllocs()
	var result *models.TimeZeroPB
	for i := 0; i < b.N; i++ {
		pb := &models.TimeZeroPB{Name: modelData.Name}
		if !modelData.CreatedAt.IsZero() {
			pb.CreatedAt = timestamppb.New(modelData.CreatedAt)
		}
		if !modelData.UpdatedAt.IsZero() {
			pb.UpdatedAt = timestamppb.New(modelData.UpdatedAt)
		}
		result = pb
	}
	benchInternalTimeZeroPB = result
}

func BenchmarkNative_TimeZeroValidValue_ModelToPB_Internal(b *testing.B) {
	now := time.Now()
	modelData := &models.TimeZeroModel{
		Name:      "valid",
		CreatedAt: now,
		UpdatedAt: now,
	}

	b.ResetTimer()
	b.ReportAllocs()
	var result *models.TimeZeroPB
	for i := 0; i < b.N; i++ {
		result = &models.TimeZeroPB{
			Name:      modelData.Name,
			CreatedAt: timestamppb.New(modelData.CreatedAt),
			UpdatedAt: timestamppb.New(modelData.UpdatedAt),
		}
	}
	benchInternalTimeZeroPB = result
}

func BenchmarkNative_Simple_PBToModel_Parallel_Internal(b *testing.B) {
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
		benchSafeConverterResult = result
	})
}

func TestPBMO_Internal_Simple_PBToModel_Consistency(t *testing.T) {
	pb := models.SimplePB{
		Id:    uint64(1),
		Name:  "testuser",
		Email: "test@example.com",
		Age:   int32(25),
	}

	pbmoResult, err := pbmo.FromPB[models.SimplePB, models.SimpleModel](&pb)
	assert.Nil(t, err)
	assert.NotNil(t, pbmoResult)

	nativeResult := &models.SimpleModel{
		ID:    pb.Id,
		Name:  pb.Name,
		Email: pb.Email,
		Age:   int(pb.Age),
	}

	assert.Equal(t, nativeResult.ID, pbmoResult.ID)
	assert.Equal(t, nativeResult.Name, pbmoResult.Name)
	assert.Equal(t, nativeResult.Email, pbmoResult.Email)
	assert.Equal(t, nativeResult.Age, pbmoResult.Age)
}

func TestPBMO_Internal_Account_PBToModel_Consistency(t *testing.T) {
	now := time.Now()
	pb := models.AccountInfoPB{
		AccountId:  "acc001",
		Username:   "testuser",
		Status:     int32(1),
		CreatedAt:  timestamppb.New(now),
		LastActive: timestamppb.New(now),
	}

	pbmoResult, err := pbmo.FromPB[models.AccountInfoPB, models.AccountInfoModel](&pb)
	assert.Nil(t, err)
	assert.NotNil(t, pbmoResult)

	nativeResult := &models.AccountInfoModel{
		AccountID:  pb.AccountId,
		Username:   pb.Username,
		Status:     int(pb.Status),
		CreatedAt:  pb.CreatedAt.AsTime(),
		LastActive: pb.LastActive.AsTime(),
	}

	assert.Equal(t, nativeResult.AccountID, pbmoResult.AccountID)
	assert.Equal(t, nativeResult.Username, pbmoResult.Username)
	assert.Equal(t, nativeResult.Status, pbmoResult.Status)
	assert.Equal(t, nativeResult.CreatedAt.Unix(), pbmoResult.CreatedAt.Unix())
}

func TestPBMO_Internal_Wrapper_PBToModel_Consistency(t *testing.T) {
	pb := models.WrapperPB{
		Name:     "test",
		IsActive: wrapperspb.Bool(true),
		Count:    wrapperspb.Int32(42),
		Score:    wrapperspb.Double(99.5),
		Label:    wrapperspb.String("hello"),
	}

	pbmoResult, err := pbmo.FromPB[models.WrapperPB, models.WrapperModel](&pb)
	assert.Nil(t, err)
	assert.NotNil(t, pbmoResult)

	assert.Equal(t, "test", pbmoResult.Name)
	assert.True(t, *pbmoResult.IsActive)
	assert.Equal(t, int32(42), *pbmoResult.Count)
	assert.Equal(t, 99.5, *pbmoResult.Score)
	assert.Equal(t, "hello", *pbmoResult.Label)
}

func TestPBMO_Internal_Mapped_WithFieldMapping_Consistency(t *testing.T) {
	pb := models.MappedPB{
		ClientId:  uint64(123),
		UserName:  "testuser",
		UserEmail: "test@example.com",
	}

	converter := pbmo.NewBidiConverter(
		models.MappedPB{}, models.MappedModel{},
		pbmo.WithFieldMapping("ID", "ClientId"),
		pbmo.WithFieldMapping("Name", "UserName"),
		pbmo.WithFieldMapping("Email", "UserEmail"),
	)

	var model models.MappedModel
	err := converter.ConvertPBToModel(pb, &model)
	assert.Nil(t, err)

	assert.Equal(t, uint64(123), model.ID)
	assert.Equal(t, "testuser", model.Name)
	assert.Equal(t, "test@example.com", model.Email)
}
