/**
 * @Author: developer 500000000@example.com
 * @Date: 2026-10-11 00:00:00
 * @LastEditors: developer 500000000@example.com
 * @LastEditTime: 2026-10-11 00:00:00
 * @FilePath: \go-pbmo-benchmark\benchmarks\image_bench_test.go
 * @Description: go-argus 图片校验真实样本测试与基准 - 样本由测试程序化生成，全程内存，测试结束自动销毁
 *
 * Copyright (c) 2026 by developer, All Rights Reserved.
 */

package benchmarks

import (
	"bytes"
	"encoding/binary"
	"image"
	"image/color"
	"image/color/palette"
	"image/draw"
	"image/gif"
	"image/jpeg"
	"image/png"
	"testing"

	validator "github.com/kamalyes/go-argus"
	"github.com/stretchr/testify/assert"
)

// ══════════════════════════════════════════════════════════════════════════════
// 真实图片样本生成（全程内存，测试结束随变量作用域销毁）
// ══════════════════════════════════════════════════════════════════════════════

// drawScene 绘制 256x256 生产风格场景图：对角渐变 + 圆形 + 矩形 + 描边
func drawScene() *image.RGBA {
	img := image.NewRGBA(image.Rect(0, 0, 256, 256))
	for y := 0; y < 256; y++ {
		for x := 0; x < 256; x++ {
			img.SetRGBA(x, y, color.RGBA{uint8(x), uint8(y), uint8((x + y) / 2), 255})
		}
	}
	for y := 48; y <= 208; y++ {
		for x := 48; x <= 208; x++ {
			dx, dy := x-128, y-128
			if dx*dx+dy*dy <= 80*80 {
				img.SetRGBA(x, y, color.RGBA{220, 40, 40, 255})
			}
		}
	}
	for y := 30; y < 60; y++ {
		for x := 30; x < 90; x++ {
			img.SetRGBA(x, y, color.RGBA{30, 90, 220, 255})
		}
	}
	for x := 180; x < 240; x++ {
		img.SetRGBA(x, 190, color.RGBA{255, 230, 40, 255})
		img.SetRGBA(x, 230, color.RGBA{255, 230, 40, 255})
	}
	for y := 190; y <= 230; y++ {
		img.SetRGBA(180, y, color.RGBA{255, 230, 40, 255})
		img.SetRGBA(239, y, color.RGBA{255, 230, 40, 255})
	}
	return img
}

// genRealPNG 标准库真实编码的 PNG
func genRealPNG() []byte {
	var buf bytes.Buffer
	if err := png.Encode(&buf, drawScene()); err != nil {
		panic(err)
	}
	return buf.Bytes()
}

// genRealJPEG 标准库真实编码的 JPEG（quality 85）
func genRealJPEG() []byte {
	var buf bytes.Buffer
	if err := jpeg.Encode(&buf, drawScene(), &jpeg.Options{Quality: 85}); err != nil {
		panic(err)
	}
	return buf.Bytes()
}

// genRealGIF 标准库真实编码的 GIF（Plan9 调色板）
func genRealGIF() []byte {
	img := drawScene()
	palImg := image.NewPaletted(img.Bounds(), palette.Plan9)
	draw.Draw(palImg, palImg.Bounds(), img, image.Point{}, draw.Src)
	var buf bytes.Buffer
	if err := gif.Encode(&buf, palImg, nil); err != nil {
		panic(err)
	}
	return buf.Bytes()
}

// genRealSVG 真实矢量文档
func genRealSVG() []byte {
	return []byte(`<?xml version="1.0" encoding="UTF-8"?>
<svg xmlns="http://www.w3.org/2000/svg" width="256" height="256" viewBox="0 0 256 256">
  <defs>
    <linearGradient id="bg" x1="0" y1="0" x2="1" y2="1">
      <stop offset="0" stop-color="#3366ff"/>
      <stop offset="1" stop-color="#ff6633"/>
    </linearGradient>
  </defs>
  <rect width="256" height="256" fill="url(#bg)"/>
  <circle cx="128" cy="128" r="64" fill="#ffffff" fill-opacity="0.8"/>
  <rect x="30" y="30" width="60" height="30" fill="#1e5adc" opacity="0.6"/>
</svg>
`)
}

// vp8lBitWriter 按 LSB-first 打包位流（VP8L 无损 WebP 要求的位序）
type vp8lBitWriter struct {
	data []byte
	cur  byte
	bit  uint
}

func (w *vp8lBitWriter) put(v uint32, n uint) {
	for i := uint(0); i < n; i++ {
		if v&(1<<i) != 0 {
			w.cur |= 1 << w.bit
		}
		w.bit++
		if w.bit == 8 {
			w.data = append(w.data, w.cur)
			w.cur = 0
			w.bit = 0
		}
	}
}

func (w *vp8lBitWriter) bytes() []byte {
	if w.bit > 0 {
		w.data = append(w.data, w.cur)
	}
	return w.data
}

// genRealWebP 手工编码 1x1 无损 VP8L 真实码流（可被标准 WebP 解码器解码）
func genRealWebP() []byte {
	w := &vp8lBitWriter{}
	w.put(0x2F, 8) // VP8L 签名
	w.put(0, 14)   // width-1 = 0
	w.put(0, 14)   // height-1 = 0
	w.put(0, 1)    // alpha hint
	w.put(0, 3)    // version
	w.put(0, 1)    // 无 transform
	w.put(0, 1)    // 无 color cache
	w.put(0, 1)    // 无 meta huffman
	// 五个前缀码均为双符号 simple code：green={0,0}, red={0,0x20}, blue={0,0x40}, alpha={0,0xFF}, distance={0,0}
	// 每码格式：is_simple=1, 2 符号, 首符号长度选择位=0（1 bit）, 首符号值 1bit, 次符号值 8bit
	for _, pair := range [][2]uint32{{0, 0}, {0, 0x20}, {0, 0x40}, {0, 0xFF}, {0, 0}} {
		w.put(1, 1)       // simple code
		w.put(1, 1)       // 2 个符号
		w.put(0, 1)       // 首符号字段长度选择位：0 表示 1 bit
		w.put(pair[0], 1) // 首符号值
		w.put(pair[1], 8) // 次符号值
	}
	// 像素数据：1x1 literal，各通道取次符号（编码 bit=1）
	w.put(0, 1) // green = 0
	w.put(1, 1) // red = 0x20
	w.put(1, 1) // blue = 0x40
	w.put(1, 1) // alpha = 0xFF
	vp8l := w.bytes()

	webp := bytes.NewBuffer(nil)
	webp.WriteString("RIFF")
	binary.Write(webp, binary.LittleEndian, uint32(4+8+len(vp8l)))
	webp.WriteString("WEBP")
	webp.WriteString("VP8L")
	binary.Write(webp, binary.LittleEndian, uint32(len(vp8l)))
	webp.Write(vp8l)
	return webp.Bytes()
}

// genRealBMP 2x2 真实 24 位像素 BMP（自下而上、行 4 字节对齐）
func genRealBMP() []byte {
	var bmp bytes.Buffer
	bmp.WriteString("BM")
	binary.Write(&bmp, binary.LittleEndian, uint32(54+16)) // 文件大小
	bmp.Write([]byte{0, 0, 0, 0})                          // 保留
	binary.Write(&bmp, binary.LittleEndian, uint32(54))    // 像素数据偏移
	binary.Write(&bmp, binary.LittleEndian, uint32(40))    // BITMAPINFOHEADER 大小
	binary.Write(&bmp, binary.LittleEndian, int32(2))      // 宽
	binary.Write(&bmp, binary.LittleEndian, int32(2))      // 高
	binary.Write(&bmp, binary.LittleEndian, uint16(1))     // 位平面
	binary.Write(&bmp, binary.LittleEndian, uint16(24))    // 每像素位数
	binary.Write(&bmp, binary.LittleEndian, uint32(0))     // BI_RGB
	binary.Write(&bmp, binary.LittleEndian, uint32(16))    // 图像字节数
	binary.Write(&bmp, binary.LittleEndian, int32(2835))   // 水平分辨率
	binary.Write(&bmp, binary.LittleEndian, int32(2835))   // 垂直分辨率
	binary.Write(&bmp, binary.LittleEndian, uint32(0))     // 调色板
	binary.Write(&bmp, binary.LittleEndian, uint32(0))     // 重要色
	// 每行 2px×3B=6B + 2B 填充 = 8B（BGR 字节序）
	bmp.Write([]byte{0xFF, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}) // 第 2 行：蓝
	bmp.Write([]byte{0x00, 0x00, 0xFF, 0x00, 0x00, 0x00, 0x00, 0x00}) // 第 1 行：红
	return bmp.Bytes()
}

// genRealTIFF 2x2 真实 1 位黑白 TIFF（小端、单 IFD）
func genRealTIFF() []byte {
	var tif bytes.Buffer
	tif.WriteString("II\x2a\x00")
	binary.Write(&tif, binary.LittleEndian, uint32(8)) // IFD 偏移
	dataOffset := 8 + 2 + 6*12 + 4
	binary.Write(&tif, binary.LittleEndian, uint16(6)) // 目录项数
	entry := func(tag uint16, typ uint16, count uint32, value uint32) {
		binary.Write(&tif, binary.LittleEndian, tag)
		binary.Write(&tif, binary.LittleEndian, typ)
		binary.Write(&tif, binary.LittleEndian, count)
		binary.Write(&tif, binary.LittleEndian, value)
	}
	entry(0x0100, 3, 1, 2)                             // ImageWidth SHORT = 2
	entry(0x0101, 3, 1, 2)                             // ImageLength SHORT = 2
	entry(0x0102, 3, 1, 1)                             // BitsPerSample SHORT = 1
	entry(0x0106, 3, 1, 1)                             // Photometric SHORT = BlackIsZero
	entry(0x0111, 4, 1, uint32(dataOffset))            // StripOffsets LONG
	entry(0x0117, 4, 1, 2)                             // StripByteCounts LONG = 2
	binary.Write(&tif, binary.LittleEndian, uint32(0)) // 无下一个 IFD
	tif.Write([]byte{0x40, 0x80})                      // 2 行、每行 1 字节像素
	return tif.Bytes()
}

// genRealICO 1x1 真实 32 位图标（ICONDIR + 目录项 + DIB + AND 掩码）
func genRealICO() []byte {
	var ico bytes.Buffer
	binary.Write(&ico, binary.LittleEndian, uint16(0)) // 保留字
	binary.Write(&ico, binary.LittleEndian, uint16(1)) // 类型：图标
	binary.Write(&ico, binary.LittleEndian, uint16(1)) // 图标数
	binary.Write(&ico, binary.LittleEndian, []byte{1, 1, 0, 0, 1, 0, 32, 0})
	binary.Write(&ico, binary.LittleEndian, uint32(40+4+4)) // 图像字节数
	binary.Write(&ico, binary.LittleEndian, uint32(22))     // 数据偏移（6+16）
	// BITMAPINFOHEADER（高度为图标位图的 2 倍：XOR+AND 掩码）
	binary.Write(&ico, binary.LittleEndian, uint32(40))
	binary.Write(&ico, binary.LittleEndian, int32(1))
	binary.Write(&ico, binary.LittleEndian, int32(2))
	binary.Write(&ico, binary.LittleEndian, uint16(1))
	binary.Write(&ico, binary.LittleEndian, uint16(32))
	binary.Write(&ico, binary.LittleEndian, uint32(0))
	binary.Write(&ico, binary.LittleEndian, uint32(8))
	binary.Write(&ico, binary.LittleEndian, int32(0))
	binary.Write(&ico, binary.LittleEndian, int32(0))
	binary.Write(&ico, binary.LittleEndian, uint32(0))
	binary.Write(&ico, binary.LittleEndian, uint32(0))
	ico.Write([]byte{0xFF, 0x00, 0x00, 0xFF}) // BGRA 红色像素
	ico.Write([]byte{0x00, 0x00, 0x00, 0x00}) // AND 掩码（不透明）
	return ico.Bytes()
}

// genRealBoxBrand 构造真实 ISOBMFF 结构（ftyp 容器 + mdat 载荷）
func genRealBoxBrand(brand, compat string) []byte {
	b := bytes.NewBuffer(nil)
	binary.Write(b, binary.BigEndian, uint32(28)) // ftyp box 大小
	b.WriteString("ftyp")
	b.WriteString(brand)
	b.Write([]byte{0, 0, 0, 0}) // minor version
	b.WriteString(compat)
	binary.Write(b, binary.BigEndian, uint32(8)) // mdat box 大小
	b.WriteString("mdat")
	return b.Bytes()
}

// bankIconRequest 复刻业务场景：银行图标非空、不超 2MiB、仅允许 PNG/JPEG
type bankIconRequest struct {
	Icon []byte `json:"icon" validate:"required,max=2097152,image=png jpg"`
}

// ══════════════════════════════════════════════════════════════════════════════
// 真实样本正确性测试
// ══════════════════════════════════════════════════════════════════════════════

func TestArgus_Image_FormatDetection_Consistency(t *testing.T) {
	// 样本在用例内生成，随作用域销毁
	samples := map[string][]byte{
		"png":  genRealPNG(),
		"jpeg": genRealJPEG(),
		"gif":  genRealGIF(),
		"webp": genRealWebP(),
		"bmp":  genRealBMP(),
		"tiff": genRealTIFF(),
		"ico":  genRealICO(),
		"avif": genRealBoxBrand("avif", "avifmif1"),
		"heic": genRealBoxBrand("heic", "mif1heic"),
		"svg":  genRealSVG(),
	}

	for format, data := range samples {
		assert.Equal(t, format, validator.DetectImageFormat(data), "格式嗅探应命中 %s", format)
		assert.True(t, validator.IsImage(data), "真实 %s 样本应通过 IsImage", format)
	}

	// 格式专属规则
	assert.True(t, validator.IsPNG(samples["png"]))
	assert.True(t, validator.IsJPEG(samples["jpeg"]))
	assert.True(t, validator.IsGIF(samples["gif"]))
	assert.True(t, validator.IsWebP(samples["webp"]))
	assert.True(t, validator.IsBMP(samples["bmp"]))
	assert.True(t, validator.IsTIFF(samples["tiff"]))
	assert.True(t, validator.IsICO(samples["ico"]))
	assert.True(t, validator.IsAVIF(samples["avif"]))
	assert.True(t, validator.IsHEIC(samples["heic"]))
	assert.True(t, validator.IsSVG(samples["svg"]))

	// 白名单：PNG 通过、GIF 被拒、未知格式名 fail-closed
	assert.True(t, validator.IsImageInFormats(samples["png"], "png jpg"))
	assert.False(t, validator.IsImageInFormats(samples["gif"], "png jpg"))
	assert.False(t, validator.IsImageInFormats(samples["png"], "png tiffx"))

	// 截断样本：去掉 IEND 尾块后必须被完整性检查拒绝
	truncated := samples["png"][:len(samples["png"])-6]
	assert.False(t, validator.IsPNG(truncated))
	assert.False(t, validator.IsImage(truncated))
}

func TestArgus_Image_StructValidation_Consistency(t *testing.T) {
	v := validator.New()
	png, jpeg, gif := genRealPNG(), genRealJPEG(), genRealGIF()

	// 合法图标通过
	assert.Nil(t, v.Struct(bankIconRequest{Icon: png}))
	assert.Nil(t, v.Struct(bankIconRequest{Icon: jpeg}))

	// 空、白名单外格式、截断、超限均失败
	assert.NotNil(t, v.Struct(bankIconRequest{Icon: nil}))
	assert.NotNil(t, v.Struct(bankIconRequest{Icon: gif}))
	assert.NotNil(t, v.Struct(bankIconRequest{Icon: png[:len(png)-6]}))
	assert.NotNil(t, v.Struct(bankIconRequest{Icon: make([]byte, 2097153)}))

	// 失败错误经 i18n 渲染为可读消息
	err := v.Struct(bankIconRequest{Icon: gif})
	messages := validator.TranslateValidationErrors(err, "zh-CN")
	assert.Len(t, messages, 1)
	assert.Equal(t, "image", messages[0].Tag)
	assert.NotEmpty(t, messages[0].Message)
}

// ══════════════════════════════════════════════════════════════════════════════
// 图片校验性能基准
// ══════════════════════════════════════════════════════════════════════════════

func BenchmarkArgus_Image_DetectFormat(b *testing.B) {
	png := genRealPNG()
	jpeg := genRealJPEG()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = validator.DetectImageFormat(png)
		_ = validator.DetectImageFormat(jpeg)
	}
}

func BenchmarkArgus_Image_IsImage(b *testing.B) {
	png := genRealPNG()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = validator.IsImage(png)
	}
}

func BenchmarkArgus_Image_IsImageInFormats(b *testing.B) {
	png := genRealPNG()
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = validator.IsImageInFormats(png, "png jpg")
	}
}

func BenchmarkArgus_Image_StructTag(b *testing.B) {
	v := validator.New()
	req := bankIconRequest{Icon: genRealPNG()}
	b.ReportAllocs()
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_ = v.Struct(&req)
	}
}
