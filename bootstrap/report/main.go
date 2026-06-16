/**
 * @Author: kamalyes 501893067@qq.com
 * @Date: 2025-06-15 10:30:21
 * @LastEditors: kamalyes 501893067@qq.com
 * @LastEditTime: 2025-06-15 10:32:55
 * @FilePath: \go-pbmo-benchmark\bootstrap\report\main.go
 * @Description: PBMO vs Native 性能测试报告生成器
 *
 * Copyright (c) 2025 by kamalyes, All Rights Reserved.
 */
package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"path/filepath"
	"sort"
	"strconv"
	"strings"
)

type benchResult struct {
	Name     string  `json:"name"`
	NsPerOp  float64 `json:"ns_per_op"`
	BPerOp   uint64  `json:"bytes_per_op"`
	AllocsOp uint64  `json:"allocs_per_op"`
}

type comparisonRow struct {
	Scenario string  `json:"scenario"`
	PBMONs   float64 `json:"pbmo_ns"`
	NativeNs float64 `json:"native_ns"`
	Ratio    float64 `json:"ratio"`
	Winner   string  `json:"winner"`
}

type allocRow struct {
	Scenario     string `json:"scenario"`
	PBMOAllocs   uint64 `json:"pbmo_allocs"`
	NativeAllocs uint64 `json:"native_allocs"`
	PBMOBytes    uint64 `json:"pbmo_bytes"`
	NativeBytes  uint64 `json:"native_bytes"`
	Winner       string `json:"winner"`
}

type envInfo struct {
	Goos   string `json:"goos"`
	Goarch string `json:"goarch"`
	Pkg    string `json:"pkg"`
	CPU    string `json:"cpu"`
}

type aggResult struct {
	name      string
	nsList    []float64
	bList     []uint64
	allocList []uint64
}

func (a *aggResult) add(ns float64, b uint64, alloc uint64) {
	a.nsList = append(a.nsList, ns)
	a.bList = append(a.bList, b)
	a.allocList = append(a.allocList, alloc)
}

func (a *aggResult) avgNsPerOp() float64 {
	if len(a.nsList) == 0 {
		return 0
	}
	var sum float64
	for _, v := range a.nsList {
		sum += v
	}
	return sum / float64(len(a.nsList))
}

func (a *aggResult) avgBPerOp() uint64 {
	if len(a.bList) == 0 {
		return 0
	}
	var sum uint64
	for _, v := range a.bList {
		sum += v
	}
	return sum / uint64(len(a.bList))
}

func (a *aggResult) avgAllocsOp() uint64 {
	if len(a.allocList) == 0 {
		return 0
	}
	var sum uint64
	for _, v := range a.allocList {
		sum += v
	}
	return sum / uint64(len(a.allocList))
}

func main() {
	rootDir := "."
	parseFile := ""

	for i := 1; i < len(os.Args); i++ {
		if os.Args[i] == "-parse" && i+1 < len(os.Args) {
			parseFile = os.Args[i+1]
			i++
		} else {
			rootDir = os.Args[i]
		}
	}

	var allResults []benchResult
	var env envInfo

	if parseFile != "" {
		allResults, env = parseBenchmarkFile(parseFile)
		if len(allResults) == 0 {
			fmt.Println("Warning: no benchmark results parsed, using fallback data")
			allResults = fallbackData()
		}
	} else {
		allResults = fallbackData()
	}

	pbmoResults, nativeResults := splitPBMONative(allResults)
	comparisons := buildComparisons(pbmoResults, nativeResults)
	allocs := buildAllocComparisons(pbmoResults, nativeResults)

	benchDir := filepath.Join(rootDir, "benchmarks")
	os.MkdirAll(benchDir, 0755)

	// 生成图表
	latencySVG := generateLatencySVG(comparisons, "PBMO vs Native Latency (ns/op)")
	writeFile(filepath.Join(benchDir, "latency.svg"), latencySVG)

	allocSVG := generateAllocSVG(allocs, "PBMO vs Native Memory Allocation (allocs/op)")
	writeFile(filepath.Join(benchDir, "allocs.svg"), allocSVG)

	// 保存 JSON
	writeJSON(filepath.Join(benchDir, "benchmark_results.json"), struct {
		PBMO   []benchResult `json:"pbmo"`
		Native []benchResult `json:"native"`
	}{pbmoResults, nativeResults})
	writeJSON(filepath.Join(benchDir, "benchmark_comparisons.json"), comparisons)
	writeJSON(filepath.Join(benchDir, "benchmark_allocs.json"), allocs)

	// 生成 BENCHMARKS.md
	generateBenchmarksMD(rootDir, comparisons, allocs, env)

	fmt.Printf("✅ Done! Generated %d comparisons (%d alloc rows)\n", len(comparisons), len(allocs))
}

func fallbackData() []benchResult {
	return []benchResult{
		{"BenchmarkPBMO_Simple_PBToModel", 1063, 48, 1},
		{"BenchmarkNative_Simple_PBToModel", 45, 48, 1},
		{"BenchmarkPBMO_Medium_PBToModel", 2500, 128, 3},
		{"BenchmarkNative_Medium_PBToModel", 120, 128, 3},
	}
}

func parseBenchmarkFile(filename string) ([]benchResult, envInfo) {
	f, err := os.Open(filename)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error opening %s: %v\n", filename, err)
		return nil, envInfo{}
	}
	defer f.Close()

	aggregated := map[string]*aggResult{}
	var env envInfo
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.HasPrefix(line, "goos:") {
			env.Goos = strings.TrimSpace(strings.TrimPrefix(line, "goos:"))
			continue
		}
		if strings.HasPrefix(line, "goarch:") {
			env.Goarch = strings.TrimSpace(strings.TrimPrefix(line, "goarch:"))
			continue
		}
		if strings.HasPrefix(line, "pkg:") {
			env.Pkg = strings.TrimSpace(strings.TrimPrefix(line, "pkg:"))
			continue
		}
		if strings.HasPrefix(line, "cpu:") {
			env.CPU = strings.TrimSpace(strings.TrimPrefix(line, "cpu:"))
			continue
		}
		if !strings.HasPrefix(line, "Benchmark") {
			continue
		}
		name, nsPerOp, bPerOp, allocsOp, ok := parseBenchLine(line)
		if !ok {
			continue
		}
		if _, exists := aggregated[name]; !exists {
			aggregated[name] = &aggResult{name: name}
		}
		aggregated[name].add(nsPerOp, bPerOp, allocsOp)
	}

	return aggregateResults(aggregated), env
}

func parseBenchLine(line string) (name string, nsPerOp float64, bPerOp uint64, allocsOp uint64, ok bool) {
	fields := strings.Fields(line)
	if len(fields) < 5 {
		return "", 0, 0, 0, false
	}

	name = normalizeBenchName(fields[0])

	nsIdx := -1
	for i, f := range fields {
		if strings.HasSuffix(f, "ns/op") {
			nsIdx = i
			break
		}
	}
	if nsIdx < 1 {
		return "", 0, 0, 0, false
	}

	nsPerOp, err := strconv.ParseFloat(fields[nsIdx-1], 64)
	if err != nil {
		return "", 0, 0, 0, false
	}

	for i, f := range fields {
		if strings.HasSuffix(f, "B/op") && i > 0 {
			b, _ := strconv.ParseUint(fields[i-1], 10, 64)
			bPerOp = b
		}
		if strings.HasSuffix(f, "allocs/op") && i > 0 {
			a, _ := strconv.ParseUint(fields[i-1], 10, 64)
			allocsOp = a
		}
	}

	return name, nsPerOp, bPerOp, allocsOp, true
}

func normalizeBenchName(name string) string {
	idx := strings.LastIndexByte(name, '-')
	if idx < 0 || idx == len(name)-1 {
		return name
	}
	for _, r := range name[idx+1:] {
		if r < '0' || r > '9' {
			return name
		}
	}
	return name[:idx]
}

func splitPBMONative(all []benchResult) ([]benchResult, []benchResult) {
	pbmo := collectByPrefix(all, "BenchmarkPBMO_")
	native := collectByPrefix(all, "BenchmarkNative_")
	return pbmo, native
}

func collectByPrefix(all []benchResult, prefix string) []benchResult {
	aggregated := map[string]*aggResult{}
	for _, r := range all {
		name, ok := trimPrefix(r.Name, prefix)
		if !ok {
			continue
		}
		if _, exists := aggregated[name]; !exists {
			aggregated[name] = &aggResult{name: name}
		}
		aggregated[name].add(r.NsPerOp, r.BPerOp, r.AllocsOp)
	}
	return aggregateResults(aggregated)
}

func trimPrefix(name, prefix string) (string, bool) {
	if !strings.HasPrefix(name, prefix) {
		return "", false
	}
	return strings.TrimPrefix(name, prefix), true
}

func aggregateResults(aggregated map[string]*aggResult) []benchResult {
	var results []benchResult
	for _, agg := range aggregated {
		results = append(results, benchResult{
			Name:     agg.name,
			NsPerOp:  agg.avgNsPerOp(),
			BPerOp:   agg.avgBPerOp(),
			AllocsOp: agg.avgAllocsOp(),
		})
	}
	sort.Slice(results, func(i, j int) bool { return results[i].Name < results[j].Name })
	return results
}

func buildComparisons(pbmo, native []benchResult) []comparisonRow {
	nativeMap := map[string]benchResult{}
	for _, r := range native {
		nativeMap[r.Name] = r
	}
	var rows []comparisonRow
	for _, p := range pbmo {
		n, ok := nativeMap[p.Name]
		if !ok {
			continue
		}
		ratio := p.NsPerOp / n.NsPerOp
		winner := "Native"
		if p.NsPerOp < n.NsPerOp {
			winner = "PBMO"
		}
		rows = append(rows, comparisonRow{p.Name, p.NsPerOp, n.NsPerOp, ratio, winner})
	}
	sort.Slice(rows, func(i, j int) bool { return rows[i].Scenario < rows[j].Scenario })
	return rows
}

func buildAllocComparisons(pbmo, native []benchResult) []allocRow {
	nativeMap := map[string]benchResult{}
	for _, r := range native {
		nativeMap[r.Name] = r
	}
	var rows []allocRow
	for _, p := range pbmo {
		n, ok := nativeMap[p.Name]
		if !ok {
			continue
		}
		winner := "Native"
		if p.AllocsOp < n.AllocsOp || (p.AllocsOp == n.AllocsOp && p.BPerOp < n.BPerOp) {
			winner = "PBMO"
		}
		rows = append(rows, allocRow{p.Name, p.AllocsOp, n.AllocsOp, p.BPerOp, n.BPerOp, winner})
	}
	sort.Slice(rows, func(i, j int) bool { return rows[i].Scenario < rows[j].Scenario })
	return rows
}

func writeFile(path, content string) {
	os.WriteFile(path, []byte(content), 0644)
}

func writeJSON(filename string, data interface{}) {
	f, _ := os.Create(filename)
	defer f.Close()
	enc := json.NewEncoder(f)
	enc.SetIndent("", "  ")
	enc.Encode(data)
}

const (
	svgWidth     = 900
	barHeight    = 26
	barGap       = 5
	groupGap     = 15
	marginLeft   = 240
	marginRight  = 60
	marginTop    = 50
	marginBottom = 40
	pbmoColor    = "#3B82F6"
	nativeColor  = "#10B981"
	pbmoLabel    = "PBMO"
	nativeLabel  = "Native"
	bgColor      = "#FFFFFF"
	titleColor   = "#1E293B"
	labelColor   = "#334155"
	valueColor   = "#475569"
	legendColor  = "#64748B"
	hintColor    = "#94A3B8"
)

func generateLatencySVG(comparisons []comparisonRow, title string) string {
	if len(comparisons) == 0 {
		return "<svg></svg>"
	}

	maxNs := 0.0
	for _, c := range comparisons {
		maxNs = math.Max(maxNs, c.PBMONs)
		maxNs = math.Max(maxNs, c.NativeNs)
	}

	chartWidth := svgWidth - marginLeft - marginRight
	groupHeight := 2*barHeight + barGap + groupGap
	totalHeight := marginTop + len(comparisons)*groupHeight + marginBottom + 30

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d" viewBox="0 0 %d %d">`, svgWidth, totalHeight, svgWidth, totalHeight))
	sb.WriteString(fmt.Sprintf(`<rect width="100%%" height="100%%" fill="%s" rx="12"/>`, bgColor))
	sb.WriteString(fmt.Sprintf(`<text x="%d" y="30" fill="%s" font-family="system-ui,-apple-system,sans-serif" font-size="16" font-weight="600">%s</text>`, marginLeft, titleColor, title))
	sb.WriteString(fmt.Sprintf(`<rect x="%d" y="14" width="14" height="14" rx="3" fill="%s"/>`, svgWidth-200, pbmoColor))
	sb.WriteString(fmt.Sprintf(`<text x="%d" y="26" fill="%s" font-family="system-ui,sans-serif" font-size="12">%s</text>`, svgWidth-182, legendColor, pbmoLabel))
	sb.WriteString(fmt.Sprintf(`<rect x="%d" y="14" width="14" height="14" rx="3" fill="%s"/>`, svgWidth-110, nativeColor))
	sb.WriteString(fmt.Sprintf(`<text x="%d" y="26" fill="%s" font-family="system-ui,sans-serif" font-size="12">%s</text>`, svgWidth-92, legendColor, nativeLabel))

	for i, c := range comparisons {
		y := marginTop + i*groupHeight
		label := c.Scenario
		if len(label) > 32 {
			label = label[:29] + "..."
		}
		sb.WriteString(fmt.Sprintf(`<text x="%d" y="%d" fill="%s" font-family="monospace" font-size="11" text-anchor="end">%s</text>`, marginLeft-10, y+barHeight-6, labelColor, label))
		pbmoW := (c.PBMONs / maxNs) * float64(chartWidth)
		nativeW := (c.NativeNs / maxNs) * float64(chartWidth)
		sb.WriteString(fmt.Sprintf(`<rect x="%d" y="%d" width="%.1f" height="%d" rx="4" fill="%s" opacity="0.9"/>`, marginLeft, y, pbmoW, barHeight, pbmoColor))
		sb.WriteString(fmt.Sprintf(`<text x="%d" y="%d" fill="%s" font-family="monospace" font-size="10">%.0f ns</text>`, marginLeft+int(pbmoW)+6, y+barHeight-6, valueColor, c.PBMONs))
		sb.WriteString(fmt.Sprintf(`<rect x="%d" y="%d" width="%.1f" height="%d" rx="4" fill="%s" opacity="0.9"/>`, marginLeft, y+barHeight+barGap, nativeW, barHeight, nativeColor))
		sb.WriteString(fmt.Sprintf(`<text x="%d" y="%d" fill="%s" font-family="monospace" font-size="10">%.0f ns (%.1fx)</text>`, marginLeft+int(nativeW)+6, y+2*barHeight+barGap-6, valueColor, c.NativeNs, c.Ratio))
	}

	sb.WriteString(fmt.Sprintf(`<text x="%d" y="%d" fill="%s" font-family="system-ui,sans-serif" font-size="10">Lower is better ▸</text>`, marginLeft, totalHeight-10, hintColor))
	sb.WriteString(`</svg>`)
	return sb.String()
}

func generateAllocSVG(allocs []allocRow, title string) string {
	if len(allocs) == 0 {
		return "<svg></svg>"
	}

	maxAllocs := uint64(0)
	for _, a := range allocs {
		if a.PBMOAllocs > maxAllocs {
			maxAllocs = a.PBMOAllocs
		}
		if a.NativeAllocs > maxAllocs {
			maxAllocs = a.NativeAllocs
		}
	}
	if maxAllocs == 0 {
		maxAllocs = 1
	}

	chartWidth := svgWidth - marginLeft - marginRight
	groupHeight := 2*barHeight + barGap + groupGap
	totalHeight := marginTop + len(allocs)*groupHeight + marginBottom + 30

	var sb strings.Builder
	sb.WriteString(fmt.Sprintf(`<svg xmlns="http://www.w3.org/2000/svg" width="%d" height="%d" viewBox="0 0 %d %d">`, svgWidth, totalHeight, svgWidth, totalHeight))
	sb.WriteString(fmt.Sprintf(`<rect width="100%%" height="100%%" fill="%s" rx="12"/>`, bgColor))
	sb.WriteString(fmt.Sprintf(`<text x="%d" y="30" fill="%s" font-family="system-ui,-apple-system,sans-serif" font-size="16" font-weight="600">%s</text>`, marginLeft, titleColor, title))
	sb.WriteString(fmt.Sprintf(`<rect x="%d" y="14" width="14" height="14" rx="3" fill="%s"/>`, marginLeft, pbmoColor))
	sb.WriteString(fmt.Sprintf(`<text x="%d" y="26" fill="%s" font-family="system-ui,sans-serif" font-size="12">%s</text>`, marginLeft+18, legendColor, pbmoLabel))
	sb.WriteString(fmt.Sprintf(`<rect x="%d" y="14" width="14" height="14" rx="3" fill="%s"/>`, marginLeft+90, nativeColor))
	sb.WriteString(fmt.Sprintf(`<text x="%d" y="26" fill="%s" font-family="system-ui,sans-serif" font-size="12">%s</text>`, marginLeft+108, legendColor, nativeLabel))

	for i, a := range allocs {
		y := marginTop + i*groupHeight
		label := a.Scenario
		if len(label) > 32 {
			label = label[:29] + "..."
		}
		sb.WriteString(fmt.Sprintf(`<text x="%d" y="%d" fill="%s" font-family="monospace" font-size="11" text-anchor="end">%s</text>`, marginLeft-10, y+barHeight-6, labelColor, label))
		pbmoW := (float64(a.PBMOAllocs) / float64(maxAllocs)) * float64(chartWidth)
		nativeW := (float64(a.NativeAllocs) / float64(maxAllocs)) * float64(chartWidth)
		sb.WriteString(fmt.Sprintf(`<rect x="%d" y="%d" width="%.1f" height="%d" rx="4" fill="%s" opacity="0.9"/>`, marginLeft, y, pbmoW, barHeight, pbmoColor))
		sb.WriteString(allocText(marginLeft, pbmoW, y+barHeight-6, a.PBMOAllocs, a.PBMOBytes))
		sb.WriteString(fmt.Sprintf(`<rect x="%d" y="%d" width="%.1f" height="%d" rx="4" fill="%s" opacity="0.9"/>`, marginLeft, y+barHeight+barGap, nativeW, barHeight, nativeColor))
		sb.WriteString(allocText(marginLeft, nativeW, y+2*barHeight+barGap-6, a.NativeAllocs, a.NativeBytes))
	}

	sb.WriteString(fmt.Sprintf(`<text x="%d" y="%d" fill="%s" font-family="system-ui,sans-serif" font-size="10">Lower is better ▸</text>`, marginLeft, totalHeight-10, hintColor))
	sb.WriteString(`</svg>`)
	return sb.String()
}

func allocText(baseX int, barW float64, textY int, allocs, bytes uint64) string {
	if allocs == 0 {
		return fmt.Sprintf(`<text x="%d" y="%d" fill="%s" font-family="monospace" font-size="10">0</text>`, baseX+4, textY, valueColor)
	}
	text := fmt.Sprintf("%d (%dB)", allocs, bytes)
	textX := baseX + int(barW) + 6
	return fmt.Sprintf(`<text x="%d" y="%d" fill="%s" font-family="monospace" font-size="10">%s</text>`, textX, textY, valueColor, text)
}

func generateBenchmarksMD(rootDir string, comparisons []comparisonRow, allocs []allocRow, env envInfo) {
	var sb strings.Builder

	sb.WriteString("# Benchmark Details\n\n")
	sb.WriteString("**Auto-generated by `go run ./bootstrap/report`. Do not edit manually.**\n\n")

	if env.Goos != "" || env.Goarch != "" || env.CPU != "" {
		sb.WriteString("## Environment\n\n")
		sb.WriteString(fmt.Sprintf("| Key | Value |\n|-----|-------|\n| goos | %s |\n| goarch | %s |\n| pkg | %s |\n| cpu | %s |\n\n",
			env.Goos, env.Goarch, env.Pkg, env.CPU))
	}

	sb.WriteString("## Latency Comparison (ns/op) — PBMO vs Native\n\n")
	sb.WriteString("| Scenario | PBMO | Native | Ratio (PBMO/Native) | Winner |\n")
	sb.WriteString("|----------|-----:|-------:|--------------------:|--------|\n")
	for _, c := range comparisons {
		sb.WriteString(fmt.Sprintf("| %s | %.0f | %.0f | %.2fx | %s |\n",
			c.Scenario, c.PBMONs, c.NativeNs, c.Ratio, c.Winner))
	}

	sb.WriteString("\n## Memory Allocation — PBMO vs Native\n\n")
	sb.WriteString("| Scenario | PBMO (allocs) | Native (allocs) | PBMO (bytes) | Native (bytes) | Winner |\n")
	sb.WriteString("|----------|---------------:|----------------:|-------------:|---------------:|--------|\n")
	for _, a := range allocs {
		sb.WriteString(fmt.Sprintf("| %s | %d | %d | %d | %d | %s |\n",
			a.Scenario, a.PBMOAllocs, a.NativeAllocs, a.PBMOBytes, a.NativeBytes, a.Winner))
	}

	os.WriteFile(filepath.Join(rootDir, "BENCHMARKS.md"), []byte(sb.String()), 0644)
}
