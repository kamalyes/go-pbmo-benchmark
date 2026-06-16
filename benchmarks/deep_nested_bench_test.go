/**
 * @Author: developer 500000000@example.com
 * @Date: 2023-09-16 22:05:18
 * @LastEditors: developer 500000000@example.com
 * @LastEditTime: 2025-06-12 22:28:35
 * @FilePath: \go-pbmo-benchmark\benchmarks\deep_nested_bench_test.go
 * @Description: 深度嵌套模型性能测试 - 4-6 层嵌套结构测试
 *
 * Copyright (c) 2025 by developer, All Rights Reserved.
 */

package benchmarks

import (
	"testing"
	"time"

	"github.com/kamalyes/go-pbmo"
	"github.com/kamalyes/go-pbmo-benchmark/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var (
	benchDeepNested4ModelResult *models.DeepNested4Model
	benchDeepNested4PBResult    *models.DeepNested4PB
)

func init() {
	pbmo.Register[models.Level4PB, models.Level4Model]()
	pbmo.Register[models.Level3PB, models.Level3Model]()
	pbmo.Register[models.Level2PB, models.Level2Model]()
	pbmo.Register[models.DeepNested4PB, models.DeepNested4Model]()
}

// ══════════════════════════════════════════════════════════════════════════════
// DeepNested4PB Benchmarks (4 层嵌套) - PBMO vs Native
// ══════════════════════════════════════════════════════════════════════════════

func BenchmarkPBMO_DeepNested4_PBToModel(b *testing.B) {
	now := time.Now()
	pb := models.DeepNested4PB{
		Id:   "dn4_001",
		Name: "deep nested 4",
		Level2: &models.Level2PB{
			Id:     "l2_001",
			Title:  "level 2",
			Status: int32(1),
			Level3: &models.Level3PB{
				Id:    "l3_001",
				Name:  "level 3",
				Count: int32(100),
				Level4: &models.Level4PB{
					Id:    "l4_001",
					Name:  "level 4",
					Value: int32(999),
					Data:  "deep nested data content",
				},
			},
		},
		CreatedAt: timestamppb.New(now),
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.FromPB[models.DeepNested4PB, models.DeepNested4Model](&pb)
	}
}

func BenchmarkNative_DeepNested4_PBToModel(b *testing.B) {
	now := time.Now()
	pb := models.DeepNested4PB{
		Id:   "dn4_001",
		Name: "deep nested 4",
		Level2: &models.Level2PB{
			Id:     "l2_001",
			Title:  "level 2",
			Status: int32(1),
			Level3: &models.Level3PB{
				Id:    "l3_001",
				Name:  "level 3",
				Count: int32(100),
				Level4: &models.Level4PB{
					Id:    "l4_001",
					Name:  "level 4",
					Value: int32(999),
					Data:  "deep nested data content",
				},
			},
		},
		CreatedAt: timestamppb.New(now),
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.DeepNested4Model
	for i := 0; i < b.N; i++ {
		var level2 *models.Level2Model
		if pb.Level2 != nil {
			var level3 *models.Level3Model
			if pb.Level2.Level3 != nil {
				var level4 *models.Level4Model
				if pb.Level2.Level3.Level4 != nil {
					level4 = &models.Level4Model{
						ID:    pb.Level2.Level3.Level4.Id,
						Name:  pb.Level2.Level3.Level4.Name,
						Value: int(pb.Level2.Level3.Level4.Value),
						Data:  pb.Level2.Level3.Level4.Data,
					}
				}
				level3 = &models.Level3Model{
					ID:     pb.Level2.Level3.Id,
					Name:   pb.Level2.Level3.Name,
					Count:  int(pb.Level2.Level3.Count),
					Level4: level4,
				}
			}
			level2 = &models.Level2Model{
				ID:     pb.Level2.Id,
				Title:  pb.Level2.Title,
				Status: int(pb.Level2.Status),
				Level3: level3,
			}
		}
		result = &models.DeepNested4Model{
			ID:        pb.Id,
			Name:      pb.Name,
			Level2:    level2,
			CreatedAt: pb.CreatedAt.AsTime(),
		}
	}
	benchDeepNested4ModelResult = result
}

func BenchmarkPBMO_DeepNested4_ModelToPB(b *testing.B) {
	now := time.Now()
	model := models.DeepNested4Model{
		ID:   "dn4_001",
		Name: "deep nested 4",
		Level2: &models.Level2Model{
			ID:     "l2_001",
			Title:  "level 2",
			Status: 1,
			Level3: &models.Level3Model{
				ID:    "l3_001",
				Name:  "level 3",
				Count: 100,
				Level4: &models.Level4Model{
					ID:    "l4_001",
					Name:  "level 4",
					Value: 999,
					Data:  "deep nested data content",
				},
			},
		},
		CreatedAt: now,
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.ToPB[models.DeepNested4Model, models.DeepNested4PB](&model)
	}
}

func BenchmarkNative_DeepNested4_ModelToPB(b *testing.B) {
	now := time.Now()
	model := models.DeepNested4Model{
		ID:   "dn4_001",
		Name: "deep nested 4",
		Level2: &models.Level2Model{
			ID:     "l2_001",
			Title:  "level 2",
			Status: 1,
			Level3: &models.Level3Model{
				ID:    "l3_001",
				Name:  "level 3",
				Count: 100,
				Level4: &models.Level4Model{
					ID:    "l4_001",
					Name:  "level 4",
					Value: 999,
					Data:  "deep nested data content",
				},
			},
		},
		CreatedAt: now,
	}

	b.ReportAllocs()
	b.ResetTimer()
	var result *models.DeepNested4PB
	for i := 0; i < b.N; i++ {
		var level2 *models.Level2PB
		if model.Level2 != nil {
			var level3 *models.Level3PB
			if model.Level2.Level3 != nil {
				var level4 *models.Level4PB
				if model.Level2.Level3.Level4 != nil {
					level4 = &models.Level4PB{
						Id:    model.Level2.Level3.Level4.ID,
						Name:  model.Level2.Level3.Level4.Name,
						Value: int32(model.Level2.Level3.Level4.Value),
						Data:  model.Level2.Level3.Level4.Data,
					}
				}
				level3 = &models.Level3PB{
					Id:     model.Level2.Level3.ID,
					Name:   model.Level2.Level3.Name,
					Count:  int32(model.Level2.Level3.Count),
					Level4: level4,
				}
			}
			level2 = &models.Level2PB{
				Id:     model.Level2.ID,
				Title:  model.Level2.Title,
				Status: int32(model.Level2.Status),
				Level3: level3,
			}
		}
		result = &models.DeepNested4PB{
			Id:        model.ID,
			Name:      model.Name,
			Level2:    level2,
			CreatedAt: timestamppb.New(model.CreatedAt),
		}
	}
	benchDeepNested4PBResult = result
}

// ══════════════════════════════════════════════════════════════════════════════
// 并行测试 - PBMO vs Native
// ══════════════════════════════════════════════════════════════════════════════

func BenchmarkPBMO_DeepNested4_PBToModel_Parallel(b *testing.B) {
	now := time.Now()
	pb := models.DeepNested4PB{
		Id:   "dn4_001",
		Name: "deep nested 4",
		Level2: &models.Level2PB{
			Id:     "l2_001",
			Title:  "level 2",
			Status: int32(1),
			Level3: &models.Level3PB{
				Id:    "l3_001",
				Name:  "level 3",
				Count: int32(100),
				Level4: &models.Level4PB{
					Id:    "l4_001",
					Name:  "level 4",
					Value: int32(999),
					Data:  "deep nested data content",
				},
			},
		},
		CreatedAt: timestamppb.New(now),
	}

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		var result *models.DeepNested4Model
		for p.Next() {
			result, _ = pbmo.FromPB[models.DeepNested4PB, models.DeepNested4Model](&pb)
		}
		benchDeepNested4ModelResult = result
	})
}

func BenchmarkNative_DeepNested4_PBToModel_Parallel(b *testing.B) {
	now := time.Now()
	pb := models.DeepNested4PB{
		Id:   "dn4_001",
		Name: "deep nested 4",
		Level2: &models.Level2PB{
			Id:     "l2_001",
			Title:  "level 2",
			Status: int32(1),
			Level3: &models.Level3PB{
				Id:    "l3_001",
				Name:  "level 3",
				Count: int32(100),
				Level4: &models.Level4PB{
					Id:    "l4_001",
					Name:  "level 4",
					Value: int32(999),
					Data:  "deep nested data content",
				},
			},
		},
		CreatedAt: timestamppb.New(now),
	}

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(tpb *testing.PB) {
		var result *models.DeepNested4Model
		for tpb.Next() {
			var level2 *models.Level2Model
			if pb.Level2 != nil {
				var level3 *models.Level3Model
				if pb.Level2.Level3 != nil {
					var level4 *models.Level4Model
					if pb.Level2.Level3.Level4 != nil {
						level4 = &models.Level4Model{
							ID:    pb.Level2.Level3.Level4.Id,
							Name:  pb.Level2.Level3.Level4.Name,
							Value: int(pb.Level2.Level3.Level4.Value),
							Data:  pb.Level2.Level3.Level4.Data,
						}
					}
					level3 = &models.Level3Model{
						ID:     pb.Level2.Level3.Id,
						Name:   pb.Level2.Level3.Name,
						Count:  int(pb.Level2.Level3.Count),
						Level4: level4,
					}
				}
				level2 = &models.Level2Model{
					ID:     pb.Level2.Id,
					Title:  pb.Level2.Title,
					Status: int(pb.Level2.Status),
					Level3: level3,
				}
			}
			result = &models.DeepNested4Model{
				ID:        pb.Id,
				Name:      pb.Name,
				Level2:    level2,
				CreatedAt: pb.CreatedAt.AsTime(),
			}
		}
		benchDeepNested4ModelResult = result
	})
}

func BenchmarkPBMO_DeepNested4_ModelToPB_Parallel(b *testing.B) {
	now := time.Now()
	model := models.DeepNested4Model{
		ID:   "dn4_001",
		Name: "deep nested 4",
		Level2: &models.Level2Model{
			ID:     "l2_001",
			Title:  "level 2",
			Status: 1,
			Level3: &models.Level3Model{
				ID:    "l3_001",
				Name:  "level 3",
				Count: 100,
				Level4: &models.Level4Model{
					ID:    "l4_001",
					Name:  "level 4",
					Value: 999,
					Data:  "deep nested data content",
				},
			},
		},
		CreatedAt: now,
	}

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(p *testing.PB) {
		var result *models.DeepNested4PB
		for p.Next() {
			result, _ = pbmo.ToPB[models.DeepNested4Model, models.DeepNested4PB](&model)
		}
		benchDeepNested4PBResult = result
	})
}

func BenchmarkNative_DeepNested4_ModelToPB_Parallel(b *testing.B) {
	now := time.Now()
	model := models.DeepNested4Model{
		ID:   "dn4_001",
		Name: "deep nested 4",
		Level2: &models.Level2Model{
			ID:     "l2_001",
			Title:  "level 2",
			Status: 1,
			Level3: &models.Level3Model{
				ID:    "l3_001",
				Name:  "level 3",
				Count: 100,
				Level4: &models.Level4Model{
					ID:    "l4_001",
					Name:  "level 4",
					Value: 999,
					Data:  "deep nested data content",
				},
			},
		},
		CreatedAt: now,
	}

	b.ReportAllocs()
	b.ResetTimer()
	b.RunParallel(func(tpb *testing.PB) {
		var result *models.DeepNested4PB
		for tpb.Next() {
			var level2 *models.Level2PB
			if model.Level2 != nil {
				var level3 *models.Level3PB
				if model.Level2.Level3 != nil {
					var level4 *models.Level4PB
					if model.Level2.Level3.Level4 != nil {
						level4 = &models.Level4PB{
							Id:    model.Level2.Level3.Level4.ID,
							Name:  model.Level2.Level3.Level4.Name,
							Value: int32(model.Level2.Level3.Level4.Value),
							Data:  model.Level2.Level3.Level4.Data,
						}
					}
					level3 = &models.Level3PB{
						Id:     model.Level2.Level3.ID,
						Name:   model.Level2.Level3.Name,
						Count:  int32(model.Level2.Level3.Count),
						Level4: level4,
					}
				}
				level2 = &models.Level2PB{
					Id:     model.Level2.ID,
					Title:  model.Level2.Title,
					Status: int32(model.Level2.Status),
					Level3: level3,
				}
			}
			result = &models.DeepNested4PB{
				Id:        model.ID,
				Name:      model.Name,
				Level2:    level2,
				CreatedAt: timestamppb.New(model.CreatedAt),
			}
		}
		benchDeepNested4PBResult = result
	})
}

// ══════════════════════════════════════════════════════════════════════════════
// 批量转换测试 - PBMO vs Native
// ══════════════════════════════════════════════════════════════════════════════

func BenchmarkPBMO_DeepNested4_Batch100_PBToModel(b *testing.B) {
	now := time.Now()

	pbs := make([]*models.DeepNested4PB, 100)
	for i := range pbs {
		pbs[i] = &models.DeepNested4PB{
			Id:   "dn4_001",
			Name: "deep nested 4",
			Level2: &models.Level2PB{
				Id:     "l2_001",
				Title:  "level 2",
				Status: int32(1),
				Level3: &models.Level3PB{
					Id:    "l3_001",
					Name:  "level 3",
					Count: int32(100),
					Level4: &models.Level4PB{
						Id:    "l4_001",
						Name:  "level 4",
						Value: int32(999),
						Data:  "deep nested data",
					},
				},
			},
			CreatedAt: timestamppb.New(now),
		}
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, _ = pbmo.FromPBs[models.DeepNested4PB, models.DeepNested4Model](pbs)
	}
}

func BenchmarkNative_DeepNested4_Batch100_PBToModel(b *testing.B) {
	now := time.Now()

	pbs := make([]*models.DeepNested4PB, 100)
	for i := range pbs {
		pbs[i] = &models.DeepNested4PB{
			Id:   "dn4_001",
			Name: "deep nested 4",
			Level2: &models.Level2PB{
				Id:     "l2_001",
				Title:  "level 2",
				Status: int32(1),
				Level3: &models.Level3PB{
					Id:    "l3_001",
					Name:  "level 3",
					Count: int32(100),
					Level4: &models.Level4PB{
						Id:    "l4_001",
						Name:  "level 4",
						Value: int32(999),
						Data:  "deep nested data",
					},
				},
			},
			CreatedAt: timestamppb.New(now),
		}
	}

	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		result := make([]*models.DeepNested4Model, len(pbs))
		for j, pb := range pbs {
			var level2 *models.Level2Model
			if pb.Level2 != nil {
				var level3 *models.Level3Model
				if pb.Level2.Level3 != nil {
					var level4 *models.Level4Model
					if pb.Level2.Level3.Level4 != nil {
						level4 = &models.Level4Model{
							ID:    pb.Level2.Level3.Level4.Id,
							Name:  pb.Level2.Level3.Level4.Name,
							Value: int(pb.Level2.Level3.Level4.Value),
							Data:  pb.Level2.Level3.Level4.Data,
						}
					}
					level3 = &models.Level3Model{
						ID:     pb.Level2.Level3.Id,
						Name:   pb.Level2.Level3.Name,
						Count:  int(pb.Level2.Level3.Count),
						Level4: level4,
					}
				}
				level2 = &models.Level2Model{
					ID:     pb.Level2.Id,
					Title:  pb.Level2.Title,
					Status: int(pb.Level2.Status),
					Level3: level3,
				}
			}
			result[j] = &models.DeepNested4Model{
				ID:        pb.Id,
				Name:      pb.Name,
				Level2:    level2,
				CreatedAt: pb.CreatedAt.AsTime(),
			}
		}
		_ = result
	}
}
