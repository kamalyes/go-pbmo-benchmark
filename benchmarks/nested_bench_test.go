/**
 * @Author: developer 500000000@example.com
 * @Date: 2023-09-16 21:15:37
 * @LastEditors: developer 500000000@example.com
 * @LastEditTime: 2025-06-15 21:55:28
 * @FilePath: \go-pbmo-benchmark\benchmarks\nested_bench_test.go
 * @Description: 嵌套模型性能测试 - 2-3 层嵌套结构和切片嵌套
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
	benchUserWithAddressModelResult *models.UserWithAddressModel
	benchUserWithAddressPBResult    *models.UserWithAddressPB
	benchOuterModelResult           *models.OuterModel
	benchOuterPBResult              *models.OuterPB
	benchListModelResult            *models.ListModel
	benchListPBResult               *models.ListPB
	benchStoreModelResult           *models.StoreInfoModel
	benchStorePBResult              *models.StoreInfoPB
	benchProductModelResult         *models.ProductCatalogModel
	benchProductPBResult            *models.ProductCatalogPB
	benchEnterpriseModelResult      *models.EnterpriseInfoModel
	benchEnterprisePBResult         *models.EnterpriseInfoPB
)

func init() {
	pbmo.Register[models.AddressPB, models.AddressModel]()
	pbmo.Register[models.UserWithAddressPB, models.UserWithAddressModel]()
	pbmo.Register[models.InnerPB, models.InnerModel]()
	pbmo.Register[models.MiddlePB, models.MiddleModel]()
	pbmo.Register[models.OuterPB, models.OuterModel]()
	pbmo.Register[models.ListPB, models.ListModel]()
	pbmo.Register[models.LocationInfoPB, models.LocationInfoModel]()
	pbmo.Register[models.StoreInfoPB, models.StoreInfoModel]()
	pbmo.Register[models.CategoryTagPB, models.CategoryTagModel]()
	pbmo.Register[models.ProductCatalogPB, models.ProductCatalogModel]()
	pbmo.Register[models.ContactDetailPB, models.ContactDetailModel]()
	pbmo.Register[models.PersonInfoPB, models.PersonInfoModel]()
	pbmo.Register[models.EnterpriseInfoPB, models.EnterpriseInfoModel]()
}

// ══════════════════════════════════════════════════════════════════════════════
// UserWithAddressPB Benchmarks (2 层嵌套 - Address) - PBMO vs Native
// ══════════════════════════════════════════════════════════════════════════════

func BenchmarkPBMO_UserWithAddress_PBToModel(b *testing.B) {
	pb := models.UserWithAddressPB{
		Id:    uint64(1),
		Name:  "testuser",
		Email: "test@example.com",
		Address: &models.AddressPB{
			Country: "China",
			State:   "Guangdong",
			City:    "Shenzhen",
			Street:  "Nanshan Street",
			ZipCode: "518000",
		},
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.FromPB[models.UserWithAddressPB, models.UserWithAddressModel](&pb)
	}
}

func BenchmarkNative_UserWithAddress_PBToModel(b *testing.B) {
	pb := models.UserWithAddressPB{
		Id:    uint64(1),
		Name:  "testuser",
		Email: "test@example.com",
		Address: &models.AddressPB{
			Country: "China",
			State:   "Guangdong",
			City:    "Shenzhen",
			Street:  "Nanshan Street",
			ZipCode: "518000",
		},
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.UserWithAddressModel
	for i := 0; i < b.N; i++ {
		var address *models.AddressModel
		if pb.Address != nil {
			address = &models.AddressModel{
				Country: pb.Address.Country,
				State:   pb.Address.State,
				City:    pb.Address.City,
				Street:  pb.Address.Street,
				ZipCode: pb.Address.ZipCode,
			}
		}
		result = &models.UserWithAddressModel{
			ID:      pb.Id,
			Name:    pb.Name,
			Email:   pb.Email,
			Address: address,
		}
	}
	benchUserWithAddressModelResult = result
}

func BenchmarkPBMO_UserWithAddress_ModelToPB(b *testing.B) {
	model := models.UserWithAddressModel{
		ID:    uint64(1),
		Name:  "testuser",
		Email: "test@example.com",
		Address: &models.AddressModel{
			Country: "China",
			State:   "Guangdong",
			City:    "Shenzhen",
			Street:  "Nanshan Street",
			ZipCode: "518000",
		},
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.ToPB[models.UserWithAddressModel, models.UserWithAddressPB](&model)
	}
}

func BenchmarkNative_UserWithAddress_ModelToPB(b *testing.B) {
	model := models.UserWithAddressModel{
		ID:    uint64(1),
		Name:  "testuser",
		Email: "test@example.com",
		Address: &models.AddressModel{
			Country: "China",
			State:   "Guangdong",
			City:    "Shenzhen",
			Street:  "Nanshan Street",
			ZipCode: "518000",
		},
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.UserWithAddressPB
	for i := 0; i < b.N; i++ {
		var address *models.AddressPB
		if model.Address != nil {
			address = &models.AddressPB{
				Country: model.Address.Country,
				State:   model.Address.State,
				City:    model.Address.City,
				Street:  model.Address.Street,
				ZipCode: model.Address.ZipCode,
			}
		}
		result = &models.UserWithAddressPB{
			Id:      model.ID,
			Name:    model.Name,
			Email:   model.Email,
			Address: address,
		}
	}
	benchUserWithAddressPBResult = result
}

// ══════════════════════════════════════════════════════════════════════════════
// OuterPB Benchmarks (3 层嵌套 - Outer > Middle > Inner) - PBMO vs Native
// ══════════════════════════════════════════════════════════════════════════════

func BenchmarkPBMO_Outer_PBToModel(b *testing.B) {
	pb := models.OuterPB{
		Title: "outer",
		Middle: &models.MiddlePB{
			Name: "middle",
			Inner: &models.InnerPB{
				Label: "inner",
				Count: int32(100),
			},
		},
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.FromPB[models.OuterPB, models.OuterModel](&pb)
	}
}

func BenchmarkNative_Outer_PBToModel(b *testing.B) {
	pb := models.OuterPB{
		Title: "outer",
		Middle: &models.MiddlePB{
			Name: "middle",
			Inner: &models.InnerPB{
				Label: "inner",
				Count: int32(100),
			},
		},
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.OuterModel
	for i := 0; i < b.N; i++ {
		var middle *models.MiddleModel
		if pb.Middle != nil {
			var inner *models.InnerModel
			if pb.Middle.Inner != nil {
				inner = &models.InnerModel{
					Label: pb.Middle.Inner.Label,
					Count: pb.Middle.Inner.Count,
				}
			}
			middle = &models.MiddleModel{
				Name:  pb.Middle.Name,
				Inner: inner,
			}
		}
		result = &models.OuterModel{
			Title:  pb.Title,
			Middle: middle,
		}
	}
	benchOuterModelResult = result
}

func BenchmarkPBMO_Outer_ModelToPB(b *testing.B) {
	model := models.OuterModel{
		Title: "outer",
		Middle: &models.MiddleModel{
			Name: "middle",
			Inner: &models.InnerModel{
				Label: "inner",
				Count: 100,
			},
		},
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.ToPB[models.OuterModel, models.OuterPB](&model)
	}
}

func BenchmarkNative_Outer_ModelToPB(b *testing.B) {
	model := models.OuterModel{
		Title: "outer",
		Middle: &models.MiddleModel{
			Name: "middle",
			Inner: &models.InnerModel{
				Label: "inner",
				Count: 100,
			},
		},
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.OuterPB
	for i := 0; i < b.N; i++ {
		var middle *models.MiddlePB
		if model.Middle != nil {
			var inner *models.InnerPB
			if model.Middle.Inner != nil {
				inner = &models.InnerPB{
					Label: model.Middle.Inner.Label,
					Count: int32(model.Middle.Inner.Count),
				}
			}
			middle = &models.MiddlePB{
				Name:  model.Middle.Name,
				Inner: inner,
			}
		}
		result = &models.OuterPB{
			Title:  model.Title,
			Middle: middle,
		}
	}
	benchOuterPBResult = result
}

// ══════════════════════════════════════════════════════════════════════════════
// ListPB Benchmarks (切片嵌套) - PBMO vs Native
// ══════════════════════════════════════════════════════════════════════════════

func BenchmarkPBMO_List_PBToModel(b *testing.B) {
	pb := models.ListPB{
		Name: "testlist",
		Items: []*models.InnerPB{
			{Label: "item1", Count: int32(10)},
			{Label: "item2", Count: int32(20)},
			{Label: "item3", Count: int32(30)},
		},
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.FromPB[models.ListPB, models.ListModel](&pb)
	}
}

func BenchmarkNative_List_PBToModel(b *testing.B) {
	pb := models.ListPB{
		Name: "testlist",
		Items: []*models.InnerPB{
			{Label: "item1", Count: int32(10)},
			{Label: "item2", Count: int32(20)},
			{Label: "item3", Count: int32(30)},
		},
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.ListModel
	for i := 0; i < b.N; i++ {
		items := make([]*models.InnerModel, len(pb.Items))
		for j, item := range pb.Items {
			items[j] = &models.InnerModel{
				Label: item.Label,
				Count: item.Count,
			}
		}
		result = &models.ListModel{
			Name:  pb.Name,
			Items: items,
		}
	}
	benchListModelResult = result
}

func BenchmarkPBMO_List_ModelToPB(b *testing.B) {
	model := models.ListModel{
		Name: "testlist",
		Items: []*models.InnerModel{
			{Label: "item1", Count: 10},
			{Label: "item2", Count: 20},
			{Label: "item3", Count: 30},
		},
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.ToPB[models.ListModel, models.ListPB](&model)
	}
}

func BenchmarkNative_List_ModelToPB(b *testing.B) {
	model := models.ListModel{
		Name: "testlist",
		Items: []*models.InnerModel{
			{Label: "item1", Count: 10},
			{Label: "item2", Count: 20},
			{Label: "item3", Count: 30},
		},
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.ListPB
	for i := 0; i < b.N; i++ {
		items := make([]*models.InnerPB, len(model.Items))
		for j, item := range model.Items {
			items[j] = &models.InnerPB{
				Label: item.Label,
				Count: int32(item.Count),
			}
		}
		result = &models.ListPB{
			Name:  model.Name,
			Items: items,
		}
	}
	benchListPBResult = result
}

// ══════════════════════════════════════════════════════════════════════════════
// StoreInfoPB Benchmarks (包含 Location 嵌套，11 字段) - PBMO vs Native
// ══════════════════════════════════════════════════════════════════════════════

func BenchmarkPBMO_Store_PBToModel(b *testing.B) {
	now := time.Now()
	pb := models.StoreInfoPB{
		StoreId:     "store001",
		StoreName:   "Test Store",
		StoreCode:   "ST001",
		Category:    int32(1),
		Status:      int32(1),
		ManagerId:   "manager001",
		PhoneNumber: "400-888-8888",
		Email:       "store@example.com",
		Location: &models.LocationInfoPB{
			Country:   "China",
			Province:  "Guangdong",
			City:      "Shenzhen",
			District:  "Nanshan",
			Address:   "Nanshan Street 123",
			Latitude:  22.5431,
			Longitude: 114.0579,
			ZipCode:   "518000",
		},
		CreatedAt: timestamppb.New(now),
		UpdatedAt: timestamppb.New(now),
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.FromPB[models.StoreInfoPB, models.StoreInfoModel](&pb)
	}
}

func BenchmarkNative_Store_PBToModel(b *testing.B) {
	now := time.Now()
	pb := models.StoreInfoPB{
		StoreId:     "store001",
		StoreName:   "Test Store",
		StoreCode:   "ST001",
		Category:    int32(1),
		Status:      int32(1),
		ManagerId:   "manager001",
		PhoneNumber: "400-888-8888",
		Email:       "store@example.com",
		Location: &models.LocationInfoPB{
			Country:   "China",
			Province:  "Guangdong",
			City:      "Shenzhen",
			District:  "Nanshan",
			Address:   "Nanshan Street 123",
			Latitude:  22.5431,
			Longitude: 114.0579,
			ZipCode:   "518000",
		},
		CreatedAt: timestamppb.New(now),
		UpdatedAt: timestamppb.New(now),
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.StoreInfoModel
	for i := 0; i < b.N; i++ {
		var location *models.LocationInfoModel
		if pb.Location != nil {
			location = &models.LocationInfoModel{
				Country:   pb.Location.Country,
				Province:  pb.Location.Province,
				City:      pb.Location.City,
				District:  pb.Location.District,
				Address:   pb.Location.Address,
				Latitude:  pb.Location.Latitude,
				Longitude: pb.Location.Longitude,
				ZipCode:   pb.Location.ZipCode,
			}
		}
		result = &models.StoreInfoModel{
			StoreID:     pb.StoreId,
			StoreName:   pb.StoreName,
			StoreCode:   pb.StoreCode,
			Category:    int(pb.Category),
			Status:      int(pb.Status),
			ManagerID:   pb.ManagerId,
			PhoneNumber: pb.PhoneNumber,
			Email:       pb.Email,
			Location:    location,
			CreatedAt:   pb.CreatedAt.AsTime(),
			UpdatedAt:   pb.UpdatedAt.AsTime(),
		}
	}
	benchStoreModelResult = result
}

func BenchmarkPBMO_Store_ModelToPB(b *testing.B) {
	now := time.Now()
	model := models.StoreInfoModel{
		StoreID:     "store001",
		StoreName:   "Test Store",
		StoreCode:   "ST001",
		Category:    1,
		Status:      1,
		ManagerID:   "manager001",
		PhoneNumber: "400-888-8888",
		Email:       "store@example.com",
		Location: &models.LocationInfoModel{
			Country:   "China",
			Province:  "Guangdong",
			City:      "Shenzhen",
			District:  "Nanshan",
			Address:   "Nanshan Street 123",
			Latitude:  22.5431,
			Longitude: 114.0579,
			ZipCode:   "518000",
		},
		CreatedAt: now,
		UpdatedAt: now,
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.ToPB[models.StoreInfoModel, models.StoreInfoPB](&model)
	}
}

func BenchmarkNative_Store_ModelToPB(b *testing.B) {
	now := time.Now()
	model := models.StoreInfoModel{
		StoreID:     "store001",
		StoreName:   "Test Store",
		StoreCode:   "ST001",
		Category:    1,
		Status:      1,
		ManagerID:   "manager001",
		PhoneNumber: "400-888-8888",
		Email:       "store@example.com",
		Location: &models.LocationInfoModel{
			Country:   "China",
			Province:  "Guangdong",
			City:      "Shenzhen",
			District:  "Nanshan",
			Address:   "Nanshan Street 123",
			Latitude:  22.5431,
			Longitude: 114.0579,
			ZipCode:   "518000",
		},
		CreatedAt: now,
		UpdatedAt: now,
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.StoreInfoPB
	for i := 0; i < b.N; i++ {
		var location *models.LocationInfoPB
		if model.Location != nil {
			location = &models.LocationInfoPB{
				Country:   model.Location.Country,
				Province:  model.Location.Province,
				City:      model.Location.City,
				District:  model.Location.District,
				Address:   model.Location.Address,
				Latitude:  model.Location.Latitude,
				Longitude: model.Location.Longitude,
				ZipCode:   model.Location.ZipCode,
			}
		}
		result = &models.StoreInfoPB{
			StoreId:     model.StoreID,
			StoreName:   model.StoreName,
			StoreCode:   model.StoreCode,
			Category:    int32(model.Category),
			Status:      int32(model.Status),
			ManagerId:   model.ManagerID,
			PhoneNumber: model.PhoneNumber,
			Email:       model.Email,
			Location:    location,
			CreatedAt:   timestamppb.New(model.CreatedAt),
			UpdatedAt:   timestamppb.New(model.UpdatedAt),
		}
	}
	benchStorePBResult = result
}

// ══════════════════════════════════════════════════════════════════════════════
// ProductCatalogPB Benchmarks (包含 Slice[CategoryTag]，12 字段)
// ══════════════════════════════════════════════════════════════════════════════

func BenchmarkPBMO_Product_PBToModel(b *testing.B) {
	now := time.Now()
	enabled := true
	pb := models.ProductCatalogPB{
		ProductId:   "prod001",
		ProductName: "Test Product",
		ProductCode: "PROD001",
		CategoryId:  int64(100),
		BrandId:     "brand001",
		Price:       99.99,
		Stock:       int32(1000),
		Tags: []*models.CategoryTagPB{
			{TagId: int32(1), TagName: "Hot", TagType: int32(1), Priority: int32(10), ExpiredAt: timestamppb.New(now)},
			{TagId: int32(2), TagName: "New", TagType: int32(2), Priority: int32(5), ExpiredAt: timestamppb.New(now)},
		},
		Enabled:   wrapperspb.Bool(enabled),
		SortOrder: int32(1),
		CreatedAt: timestamppb.New(now),
		UpdatedAt: timestamppb.New(now),
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.FromPB[models.ProductCatalogPB, models.ProductCatalogModel](&pb)
	}
}

func BenchmarkNative_Product_PBToModel(b *testing.B) {
	now := time.Now()
	enabled := true
	pb := models.ProductCatalogPB{
		ProductId:   "prod001",
		ProductName: "Test Product",
		ProductCode: "PROD001",
		CategoryId:  int64(100),
		BrandId:     "brand001",
		Price:       99.99,
		Stock:       int32(1000),
		Tags: []*models.CategoryTagPB{
			{TagId: int32(1), TagName: "Hot", TagType: int32(1), Priority: int32(10), ExpiredAt: timestamppb.New(now)},
			{TagId: int32(2), TagName: "New", TagType: int32(2), Priority: int32(5), ExpiredAt: timestamppb.New(now)},
		},
		Enabled:   wrapperspb.Bool(enabled),
		SortOrder: int32(1),
		CreatedAt: timestamppb.New(now),
		UpdatedAt: timestamppb.New(now),
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.ProductCatalogModel
	for i := 0; i < b.N; i++ {
		tags := make(sqltypes.Slice[models.CategoryTagModel], len(pb.Tags))
		for j, tag := range pb.Tags {
			tags[j] = models.CategoryTagModel{
				TagID:     tag.TagId,
				TagName:   tag.TagName,
				TagType:   int(tag.TagType),
				Priority:  int(tag.Priority),
				ExpiredAt: tag.ExpiredAt.AsTime(),
			}
		}
		var enabledPtr *bool
		if pb.Enabled != nil {
			v := pb.Enabled.Value
			enabledPtr = &v
		}
		result = &models.ProductCatalogModel{
			ProductID:   pb.ProductId,
			ProductName: pb.ProductName,
			ProductCode: pb.ProductCode,
			CategoryID:  pb.CategoryId,
			BrandID:     pb.BrandId,
			Price:       pb.Price,
			Stock:       int(pb.Stock),
			Tags:        tags,
			Enabled:     enabledPtr,
			SortOrder:   int(pb.SortOrder),
			CreatedAt:   pb.CreatedAt.AsTime(),
			UpdatedAt:   pb.UpdatedAt.AsTime(),
		}
	}
	benchProductModelResult = result
}

func BenchmarkPBMO_Product_ModelToPB(b *testing.B) {
	now := time.Now()
	enabled := true
	model := models.ProductCatalogModel{
		ProductID:   "prod001",
		ProductName: "Test Product",
		ProductCode: "PROD001",
		CategoryID:  int64(100),
		BrandID:     "brand001",
		Price:       99.99,
		Stock:       1000,
		Tags: sqltypes.Slice[models.CategoryTagModel]{
			{TagID: 1, TagName: "Hot", TagType: 1, Priority: 10, ExpiredAt: now},
			{TagID: 2, TagName: "New", TagType: 2, Priority: 5, ExpiredAt: now},
		},
		Enabled:   &enabled,
		SortOrder: 1,
		CreatedAt: now,
		UpdatedAt: now,
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.ToPB[models.ProductCatalogModel, models.ProductCatalogPB](&model)
	}
}

func BenchmarkNative_Product_ModelToPB(b *testing.B) {
	now := time.Now()
	enabled := true
	model := models.ProductCatalogModel{
		ProductID:   "prod001",
		ProductName: "Test Product",
		ProductCode: "PROD001",
		CategoryID:  int64(100),
		BrandID:     "brand001",
		Price:       99.99,
		Stock:       1000,
		Tags: sqltypes.Slice[models.CategoryTagModel]{
			{TagID: 1, TagName: "Hot", TagType: 1, Priority: 10, ExpiredAt: now},
			{TagID: 2, TagName: "New", TagType: 2, Priority: 5, ExpiredAt: now},
		},
		Enabled:   &enabled,
		SortOrder: 1,
		CreatedAt: now,
		UpdatedAt: now,
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.ProductCatalogPB
	for i := 0; i < b.N; i++ {
		tags := make([]*models.CategoryTagPB, len(model.Tags))
		for j, tag := range model.Tags {
			tags[j] = &models.CategoryTagPB{
				TagId:     tag.TagID,
				TagName:   tag.TagName,
				TagType:   int32(tag.TagType),
				Priority:  int32(tag.Priority),
				ExpiredAt: timestamppb.New(tag.ExpiredAt),
			}
		}
		pb := &models.ProductCatalogPB{
			ProductId:   model.ProductID,
			ProductName: model.ProductName,
			ProductCode: model.ProductCode,
			CategoryId:  model.CategoryID,
			BrandId:     model.BrandID,
			Price:       model.Price,
			Stock:       int32(model.Stock),
			Tags:        tags,
			SortOrder:   int32(model.SortOrder),
			CreatedAt:   timestamppb.New(model.CreatedAt),
			UpdatedAt:   timestamppb.New(model.UpdatedAt),
		}
		if model.Enabled != nil {
			pb.Enabled = wrapperspb.Bool(*model.Enabled)
		}
		result = pb
	}
	benchProductPBResult = result
}

// ══════════════════════════════════════════════════════════════════════════════
// EnterpriseInfoPB Benchmarks (3 层嵌套 - Enterprise > Location/Person > Contact，9 字段)
// ══════════════════════════════════════════════════════════════════════════════

func BenchmarkPBMO_Enterprise_PBToModel(b *testing.B) {
	now := time.Now()
	pb := models.EnterpriseInfoPB{
		EnterpriseId:   "ent001",
		EnterpriseName: "Test Enterprise",
		EnterpriseCode: "ENT001",
		Industry:       "IT",
		Status:         int32(1),
		Location: &models.LocationInfoPB{
			Country:   "China",
			Province:  "Guangdong",
			City:      "Shenzhen",
			District:  "Nanshan",
			Address:   "High-tech Park",
			Latitude:  22.5431,
			Longitude: 114.0579,
			ZipCode:   "518000",
		},
		Manager: &models.PersonInfoPB{
			PersonName: "Zhang San",
			PersonAge:  int32(35),
			PersonRole: "CEO",
			Contact: &models.ContactDetailPB{
				Email:       "zhangsan@example.com",
				PhoneNumber: "13800138000",
				WechatId:    "zhangsan_wx",
				QqNumber:    "123456789",
			},
		},
		CreatedAt: timestamppb.New(now),
		UpdatedAt: timestamppb.New(now),
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.FromPB[models.EnterpriseInfoPB, models.EnterpriseInfoModel](&pb)
	}
}

func BenchmarkNative_Enterprise_PBToModel(b *testing.B) {
	now := time.Now()
	pb := models.EnterpriseInfoPB{
		EnterpriseId:   "ent001",
		EnterpriseName: "Test Enterprise",
		EnterpriseCode: "ENT001",
		Industry:       "IT",
		Status:         int32(1),
		Location: &models.LocationInfoPB{
			Country:   "China",
			Province:  "Guangdong",
			City:      "Shenzhen",
			District:  "Nanshan",
			Address:   "High-tech Park",
			Latitude:  22.5431,
			Longitude: 114.0579,
			ZipCode:   "518000",
		},
		Manager: &models.PersonInfoPB{
			PersonName: "Zhang San",
			PersonAge:  int32(35),
			PersonRole: "CEO",
			Contact: &models.ContactDetailPB{
				Email:       "zhangsan@example.com",
				PhoneNumber: "13800138000",
				WechatId:    "zhangsan_wx",
				QqNumber:    "123456789",
			},
		},
		CreatedAt: timestamppb.New(now),
		UpdatedAt: timestamppb.New(now),
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.EnterpriseInfoModel
	for i := 0; i < b.N; i++ {
		var location *models.LocationInfoModel
		if pb.Location != nil {
			location = &models.LocationInfoModel{
				Country:   pb.Location.Country,
				Province:  pb.Location.Province,
				City:      pb.Location.City,
				District:  pb.Location.District,
				Address:   pb.Location.Address,
				Latitude:  pb.Location.Latitude,
				Longitude: pb.Location.Longitude,
				ZipCode:   pb.Location.ZipCode,
			}
		}
		var manager *models.PersonInfoModel
		if pb.Manager != nil {
			var contact *models.ContactDetailModel
			if pb.Manager.Contact != nil {
				contact = &models.ContactDetailModel{
					Email:       pb.Manager.Contact.Email,
					PhoneNumber: pb.Manager.Contact.PhoneNumber,
					WechatID:    pb.Manager.Contact.WechatId,
					QqNumber:    pb.Manager.Contact.QqNumber,
				}
			}
			manager = &models.PersonInfoModel{
				PersonName: pb.Manager.PersonName,
				PersonAge:  int(pb.Manager.PersonAge),
				PersonRole: pb.Manager.PersonRole,
				Contact:    contact,
			}
		}
		result = &models.EnterpriseInfoModel{
			EnterpriseID:   pb.EnterpriseId,
			EnterpriseName: pb.EnterpriseName,
			EnterpriseCode: pb.EnterpriseCode,
			Industry:       pb.Industry,
			Status:         int(pb.Status),
			Location:       location,
			Manager:        manager,
			CreatedAt:      pb.CreatedAt.AsTime(),
			UpdatedAt:      pb.UpdatedAt.AsTime(),
		}
	}
	benchEnterpriseModelResult = result
}

func BenchmarkPBMO_Enterprise_ModelToPB(b *testing.B) {
	now := time.Now()
	model := models.EnterpriseInfoModel{
		EnterpriseID:   "ent001",
		EnterpriseName: "Test Enterprise",
		EnterpriseCode: "ENT001",
		Industry:       "IT",
		Status:         1,
		Location: &models.LocationInfoModel{
			Country:   "China",
			Province:  "Guangdong",
			City:      "Shenzhen",
			District:  "Nanshan",
			Address:   "High-tech Park",
			Latitude:  22.5431,
			Longitude: 114.0579,
			ZipCode:   "518000",
		},
		Manager: &models.PersonInfoModel{
			PersonName: "Zhang San",
			PersonAge:  35,
			PersonRole: "CEO",
			Contact: &models.ContactDetailModel{
				Email:       "zhangsan@example.com",
				PhoneNumber: "13800138000",
				WechatID:    "zhangsan_wx",
				QqNumber:    "123456789",
			},
		},
		CreatedAt: now,
		UpdatedAt: now,
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.ToPB[models.EnterpriseInfoModel, models.EnterpriseInfoPB](&model)
	}
}

func BenchmarkNative_Enterprise_ModelToPB(b *testing.B) {
	now := time.Now()
	model := models.EnterpriseInfoModel{
		EnterpriseID:   "ent001",
		EnterpriseName: "Test Enterprise",
		EnterpriseCode: "ENT001",
		Industry:       "IT",
		Status:         1,
		Location: &models.LocationInfoModel{
			Country:   "China",
			Province:  "Guangdong",
			City:      "Shenzhen",
			District:  "Nanshan",
			Address:   "High-tech Park",
			Latitude:  22.5431,
			Longitude: 114.0579,
			ZipCode:   "518000",
		},
		Manager: &models.PersonInfoModel{
			PersonName: "Zhang San",
			PersonAge:  35,
			PersonRole: "CEO",
			Contact: &models.ContactDetailModel{
				Email:       "zhangsan@example.com",
				PhoneNumber: "13800138000",
				WechatID:    "zhangsan_wx",
				QqNumber:    "123456789",
			},
		},
		CreatedAt: now,
		UpdatedAt: now,
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.EnterpriseInfoPB
	for i := 0; i < b.N; i++ {
		var location *models.LocationInfoPB
		if model.Location != nil {
			location = &models.LocationInfoPB{
				Country:   model.Location.Country,
				Province:  model.Location.Province,
				City:      model.Location.City,
				District:  model.Location.District,
				Address:   model.Location.Address,
				Latitude:  model.Location.Latitude,
				Longitude: model.Location.Longitude,
				ZipCode:   model.Location.ZipCode,
			}
		}
		var manager *models.PersonInfoPB
		if model.Manager != nil {
			var contact *models.ContactDetailPB
			if model.Manager.Contact != nil {
				contact = &models.ContactDetailPB{
					Email:       model.Manager.Contact.Email,
					PhoneNumber: model.Manager.Contact.PhoneNumber,
					WechatId:    model.Manager.Contact.WechatID,
					QqNumber:    model.Manager.Contact.QqNumber,
				}
			}
			manager = &models.PersonInfoPB{
				PersonName: model.Manager.PersonName,
				PersonAge:  int32(model.Manager.PersonAge),
				PersonRole: model.Manager.PersonRole,
				Contact:    contact,
			}
		}
		result = &models.EnterpriseInfoPB{
			EnterpriseId:   model.EnterpriseID,
			EnterpriseName: model.EnterpriseName,
			EnterpriseCode: model.EnterpriseCode,
			Industry:       model.Industry,
			Status:         int32(model.Status),
			Location:       location,
			Manager:        manager,
			CreatedAt:      timestamppb.New(model.CreatedAt),
			UpdatedAt:      timestamppb.New(model.UpdatedAt),
		}
	}
	benchEnterprisePBResult = result
}

// ══════════════════════════════════════════════════════════════════════════════
// 并行测试
// ══════════════════════════════════════════════════════════════════════════════

func BenchmarkPBMO_Outer_PBToModel_Parallel(b *testing.B) {
	pb := models.OuterPB{
		Title: "outer",
		Middle: &models.MiddlePB{
			Name: "middle",
			Inner: &models.InnerPB{
				Label: "inner",
				Count: int32(100),
			},
		},
	}

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		var result *models.OuterModel
		for p.Next() {
			result, _ = pbmo.FromPB[models.OuterPB, models.OuterModel](&pb)
		}
		benchOuterModelResult = result
	})
}

func BenchmarkNative_Outer_PBToModel_Parallel(b *testing.B) {
	pb := models.OuterPB{
		Title: "outer",
		Middle: &models.MiddlePB{
			Name: "middle",
			Inner: &models.InnerPB{
				Label: "inner",
				Count: int32(100),
			},
		},
	}

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(tpb *testing.PB) {
		var result *models.OuterModel
		for tpb.Next() {
			var middle *models.MiddleModel
			if pb.Middle != nil {
				var inner *models.InnerModel
				if pb.Middle.Inner != nil {
					inner = &models.InnerModel{
						Label: pb.Middle.Inner.Label,
						Count: pb.Middle.Inner.Count,
					}
				}
				middle = &models.MiddleModel{
					Name:  pb.Middle.Name,
					Inner: inner,
				}
			}
			result = &models.OuterModel{
				Title:  pb.Title,
				Middle: middle,
			}
		}
		benchOuterModelResult = result
	})
}

func BenchmarkPBMO_Enterprise_PBToModel_Parallel(b *testing.B) {
	now := time.Now()
	pb := models.EnterpriseInfoPB{
		EnterpriseId:   "ent001",
		EnterpriseName: "Test Enterprise",
		EnterpriseCode: "ENT001",
		Industry:       "IT",
		Status:         int32(1),
		Location: &models.LocationInfoPB{
			Country:   "China",
			Province:  "Guangdong",
			City:      "Shenzhen",
			District:  "Nanshan",
			Address:   "High-tech Park",
			Latitude:  22.5431,
			Longitude: 114.0579,
			ZipCode:   "518000",
		},
		Manager: &models.PersonInfoPB{
			PersonName: "Zhang San",
			PersonAge:  int32(35),
			PersonRole: "CEO",
			Contact: &models.ContactDetailPB{
				Email:       "zhangsan@example.com",
				PhoneNumber: "13800138000",
				WechatId:    "zhangsan_wx",
				QqNumber:    "123456789",
			},
		},
		CreatedAt: timestamppb.New(now),
		UpdatedAt: timestamppb.New(now),
	}

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		var result *models.EnterpriseInfoModel
		for p.Next() {
			result, _ = pbmo.FromPB[models.EnterpriseInfoPB, models.EnterpriseInfoModel](&pb)
		}
		benchEnterpriseModelResult = result
	})
}

func BenchmarkNative_Enterprise_PBToModel_Parallel(b *testing.B) {
	now := time.Now()
	pb := models.EnterpriseInfoPB{
		EnterpriseId:   "ent001",
		EnterpriseName: "Test Enterprise",
		EnterpriseCode: "ENT001",
		Industry:       "IT",
		Status:         int32(1),
		Location: &models.LocationInfoPB{
			Country:   "China",
			Province:  "Guangdong",
			City:      "Shenzhen",
			District:  "Nanshan",
			Address:   "High-tech Park",
			Latitude:  22.5431,
			Longitude: 114.0579,
			ZipCode:   "518000",
		},
		Manager: &models.PersonInfoPB{
			PersonName: "Zhang San",
			PersonAge:  int32(35),
			PersonRole: "CEO",
			Contact: &models.ContactDetailPB{
				Email:       "zhangsan@example.com",
				PhoneNumber: "13800138000",
				WechatId:    "zhangsan_wx",
				QqNumber:    "123456789",
			},
		},
		CreatedAt: timestamppb.New(now),
		UpdatedAt: timestamppb.New(now),
	}

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(tpb *testing.PB) {
		var result *models.EnterpriseInfoModel
		for tpb.Next() {
			var location *models.LocationInfoModel
			if pb.Location != nil {
				location = &models.LocationInfoModel{
					Country:   pb.Location.Country,
					Province:  pb.Location.Province,
					City:      pb.Location.City,
					District:  pb.Location.District,
					Address:   pb.Location.Address,
					Latitude:  pb.Location.Latitude,
					Longitude: pb.Location.Longitude,
					ZipCode:   pb.Location.ZipCode,
				}
			}
			var manager *models.PersonInfoModel
			if pb.Manager != nil {
				var contact *models.ContactDetailModel
				if pb.Manager.Contact != nil {
					contact = &models.ContactDetailModel{
						Email:       pb.Manager.Contact.Email,
						PhoneNumber: pb.Manager.Contact.PhoneNumber,
						WechatID:    pb.Manager.Contact.WechatId,
						QqNumber:    pb.Manager.Contact.QqNumber,
					}
				}
				manager = &models.PersonInfoModel{
					PersonName: pb.Manager.PersonName,
					PersonAge:  int(pb.Manager.PersonAge),
					PersonRole: pb.Manager.PersonRole,
					Contact:    contact,
				}
			}
			result = &models.EnterpriseInfoModel{
				EnterpriseID:   pb.EnterpriseId,
				EnterpriseName: pb.EnterpriseName,
				EnterpriseCode: pb.EnterpriseCode,
				Industry:       pb.Industry,
				Status:         int(pb.Status),
				Location:       location,
				Manager:        manager,
				CreatedAt:      pb.CreatedAt.AsTime(),
				UpdatedAt:      pb.UpdatedAt.AsTime(),
			}
		}
		benchEnterpriseModelResult = result
	})
}
