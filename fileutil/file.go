package fileutil

import (
	"io"
	"net/http"
	neturl "net/url"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// Download 从指定URL下载文件到本地路径。
// url: 文件的URL地址。
// filepath: 本地保存文件的路径。
// 返回值: 如果下载过程中发生错误，返回相应的错误；如果成功，返回nil。
func Download(url, filepath string) error {
	// 定义局部变量
	var (
		out     *os.File
		err     error
		resp    *http.Response
		retries = 3 // 最多支持重试3次
	)
	url, _ = neturl.QueryUnescape(url)
	// 创建文件
	out, err = os.Create(filepath)
	if err != nil {
		return err
	}
	// 确保文件关闭
	defer func() {
		_ = out.Close()
	}()

	// 获取数据
	for retries > 0 {
		resp, err = http.Get(url)
		if err != nil {
			retries--
		} else {
			break
		}
	}
	if err != nil {
		return err
	}
	// 确保响应体关闭
	defer func() {
		_ = resp.Body.Close()
	}()
	// 将响应体写入文件
	_, err = io.Copy(out, resp.Body)
	if err != nil {
		return err
	}
	return nil
}

// GetBaseFullName 从给定的文件路径中获取基础名（包含扩展名）。
// filePath 参数是文件路径字符串。
// f 参数是一个可选的函数，用于对文件路径进行自定义处理。如果为 nil，则不进行任何处理。
// ps: 使用时请注意，如果传入的路径中并不包含带后缀名的文件名，函数会返回路径中的最后一个路径元素
func GetBaseFullName(filePath string, f func(filePath string) string) string {
	// 添加简单的检查以确保 filePath 不为空
	if filePath == "" {
		return ""
	}
	// 有的文件路径需要一些特殊处理逻辑，可由func进行自定义处理,否则直接传nil即可
	if f != nil {
		processedPath := f(filePath)
		if processedPath != "" {
			filePath = processedPath
		}
	}
	// 处理url中中文被编码的问题
	filePath, _ = neturl.QueryUnescape(filePath)
	// 标准化路径中的斜杠，确保与当前操作系统兼容
	normalizedPath := filepath.FromSlash(filePath)
	separatorStr := string(filepath.Separator)
	// 替换路径中的斜杠，以匹配当前操作系统的路径分隔符
	if strings.Contains(normalizedPath, "\\") {
		normalizedPath = strings.ReplaceAll(normalizedPath, "\\", separatorStr)
	}
	// 去除路径两端的路径分隔符
	normalizedPath = strings.Trim(normalizedPath, separatorStr)
	// 获取文件的扩展名
	ext := filepath.Ext(normalizedPath)
	// 如果没有扩展名，则尝试通过分割路径来获取文件名
	if ext == "" {
		pathSplits := strings.Split(normalizedPath, separatorStr)
		return pathSplits[len(pathSplits)-1]
	}
	// 使用 filepath.Base 函数获取文件的基础名称
	baseName := filepath.Base(normalizedPath)
	// 检查结果是否为空或仅为路径分隔符
	if baseName == "" || strings.HasPrefix(baseName, string(filepath.Separator)) {
		return ""
	}
	// 返回文件的基础名称
	return baseName
}

// GetExtension 获取文件扩展名
func GetExtension(filepath string) string {
	return path.Ext(filepath)
}
