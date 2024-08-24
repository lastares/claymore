package stringutil

import (
	"crypto/md5"
	"fmt"
	"io"
)

// Substr 返回字符串 str 从索引 start 开始，长度为 length 的子字符串。
// 如果 start 为负，则从字符串末尾向前数。如果 length 为负，则返回从 start 开始到字符串末尾的子字符串。
// 参数:
//
//	str - 要处理的字符串。
//	start - 子字符串的起始位置。
//	length - 子字符串的长度。
//
// 返回值: 字符串 str 的子字符串。
func Substr(str string, start, length int) string {
	// 将字符串转换为 rune 切片，以正确处理多字节字符
	rs := []rune(str)
	rl := len(rs)
	// 调整 start 和 length，处理负数的情况
	if start < 0 {
		start = rl + start
	}
	if length < 0 {
		length = start + length
		start = rl + length
		length = -length
	}
	// 计算子字符串的结束位置。
	end := start + length
	// 确保 start 和 end 在有效范围内。
	if start < 0 {
		start = 0
	}
	if start > rl {
		start = rl
	}
	if end < 0 {
		end = 0
	}
	if end > rl {
		end = rl
	}
	// 返回从 start 到 end 的子字符串。
	return string(rs[start:end])
}

// Md5 计算给定字符串的 MD5 哈希值。
// 这个函数接受一个字符串输入 s，并返回该字符串的 MD5 哈希值。
// MD5 哈希是一种广泛使用的散列算法，可以产生一个128位的哈希值，
// 通常用于数据完整性检查和密码存储等场景。
func Md5(s string) string {
	// 初始化一个新的 MD5 哈希计算对象。
	h := md5.New()
	// 使用给定的字符串更新哈希计算对象，即对字符串进行哈希处理。
	io.WriteString(h, s)
	// 格式化并返回计算得到的 MD5 哈希值。
	// "%x" 是用于将二进制数据转换为十六进制字符串的格式说明符。
	return fmt.Sprintf("%x", h.Sum(nil))
}
