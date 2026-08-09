/*
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2026-06-15 00:00:00
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2026-06-15 00:00:00
 * @FilePath: \go-pbmo-benchmark\benchmarks\desensitize_bench_test.go
 * @Description: 脱敏转换基准 - WithDesensitize 开/关对照，验证掩码编织成本与关闭态零装配
 *
 * Copyright (c) 2026 by kamalyes, All Rights Reserved.
 */

package benchmarks

import (
	"testing"

	"github.com/kamalyes/go-pbmo"
	"github.com/kamalyes/go-pbmo-benchmark/models"
)

// ══════════════════════════════════════════════════════════════════════════════
// 脱敏转换基准（掩码编织进拷贝计划，开关关闭时零装配）
// ══════════════════════════════════════════════════════════════════════════════

// newDesensitizePB 构造携带敏感字段的 PB 实例
func newDesensitizePB() models.DesensitizePB {
	return models.DesensitizePB{
		Name:     "zhangsanfeng",
		Phone:    "13800138000",
		Email:    "zhangsan@example.com",
		IdCard:   "11010119900307123X",
		BankCard: "6222021234567890123",
		Token:    "sk-live-abcdefghijklmnop",
		Address:  "No.1 Zhongguancun Street Haidian Beijing",
		Memo:     "keep-me",
	}
}

// BenchmarkPBMO_DesensitizeOff_PBToModel_Internal 开关关闭：计划不装掩码闭包，
// 与引入脱敏前逐字节相同，作为零装配对照
func BenchmarkPBMO_DesensitizeOff_PBToModel_Internal(b *testing.B) {
	converter := pbmo.NewBidiConverter(models.DesensitizePB{}, models.DesensitizeModel{}, pbmo.WithDesensitize(false))
	pb := newDesensitizePB()

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var model models.DesensitizeModel
		if err := converter.ConvertPBToModel(pb, &model); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkPBMO_DesensitizeOn_PBToModel_Internal 开关开启：同一拷贝计划内联掩码，
// 与关闭态的差值即掩码本身的成本
func BenchmarkPBMO_DesensitizeOn_PBToModel_Internal(b *testing.B) {
	converter := pbmo.NewBidiConverter(models.DesensitizePB{}, models.DesensitizeModel{}, pbmo.WithDesensitize(true))
	pb := newDesensitizePB()

	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		var model models.DesensitizeModel
		if err := converter.ConvertPBToModel(pb, &model); err != nil {
			b.Fatal(err)
		}
	}
}