package fileutil

import (
	"io"
	"io/fs"
	"log"
	"net/http"
	neturl "net/url"
	"os"
	"path"
	"path/filepath"
	"slices"
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

// GetFullName 从给定的文件路径中获取基础名（包含扩展名）。
// filePath 参数是文件路径字符串。
// f 参数是一个可选的函数，用于对文件路径进行自定义处理。如果为 nil，则不进行任何处理。
// ps: 使用时请注意，如果传入的路径中并不包含带后缀名的文件名，函数会返回路径中的最后一个路径元素
func GetFullName(filePath string, f func(filePath string) string) string {
	// 有的文件路径需要一些特殊处理逻辑，可由func进行自定义处理,否则直接传nil即可
	if f != nil {
		processedPath := f(filePath)
		if processedPath != "" {
			filePath = processedPath
		}
	}
	// 添加简单的检查以确保 filePath 不为空
	if filePath == "" {
		return ""
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

// GetBaseName 根据文件路径获取文件的基础名称。
// 此函数主要用于去除文件的扩展名，返回文件的基础名称部分。
// 参数：
//
//	filePath - 文件的路径。可以是绝对路径或相对路径。
//
// 返回值：
//
//	返回文件的基础名称。如果输入的filePath为空，则返回空字符串。
func GetBaseName(filePath string) string {
	// 检查文件路径是否为空
	if filePath == "" {
		return ""
	}
	// 获取文件的全名，不考虑任何参数
	fullName := GetFullName(filePath, nil)
	// 获取文件的扩展名
	extension := GetExtension(fullName)
	// 移除文件的扩展名，返回基础名称
	return strings.TrimSuffix(fullName, extension)
}

// GetExtension 获取文件扩展名
func GetExtension(filePath string) string {
	if filePath == "" {
		return ""
	}
	return path.Ext(filePath)
}

// GetExtNoDot 函数用于获取文件路径的扩展名，不包括点字符。
// 如果文件路径为空，则返回空字符串。
// 参数:
//
//	filePath - 文件的路径。
//
// 返回值:
//
//	文件的扩展名，不包括点字符。如果无法确定扩展名或输入为空，则返回空字符串。
func GetExtNoDot(filePath string) string {
	// 当文件路径为空时，直接返回空字符串
	if filePath == "" {
		return ""
	}
	// 使用path.Ext函数获取文件扩展名，该函数返回的是以点开头的扩展名，这里将其转换为字符串并去除首字符（即点）后返回
	suffix := []rune(path.Ext(filePath))
	return string(suffix[1:])
}

type DirFileListConfig struct {
	fileterSuffixes []string
	filterFileNames []string
	hasPrefix       bool
}

type Option func(cfg *DirFileListConfig)

func WithFileterSuffixes(suffixes ...string) Option {
	return func(cfg *DirFileListConfig) {
		cfg.fileterSuffixes = suffixes
	}
}

func WithFilterFileNames(filterFileNames ...string) Option {
	return func(cfg *DirFileListConfig) {
		cfg.filterFileNames = filterFileNames
	}
}

func WithHasPrefix(hasPrefix bool) Option {
	return func(cfg *DirFileListConfig) {
		cfg.hasPrefix = hasPrefix
	}
}

// GetDirFileList 该函数获取指定目录下的所有文件路径列表。
// ps: 会包括系统自动生成的隐藏文件，如 .DS_Store，注意过滤
func GetDirFileList(dir string, options ...Option) ([]string, error) {
	cfg := &DirFileListConfig{}
	for _, option := range options {
		option(cfg)
	}
	cfg.fileterSuffixes = append(cfg.fileterSuffixes, []string{".DS_Store"}...)
	// 初始化文件路径切片
	var fullFileNames []string
	// 使用 filepath.Walk 遍历目录，err 用于接收 filepath.Walk 的执行结果
	fileNamesLen, fileSuffixesLen := len(cfg.filterFileNames), len(cfg.fileterSuffixes)
	err := filepath.Walk(dir, func(path string, info fs.FileInfo, err error) error {
		// 如果遍历过程中发生错误，直接返回错误
		if err != nil {
			return err
		}
		ext := GetExtension(path)
		isFile := !info.IsDir()
		if fileNamesLen > 0 {
			isFile = isFile && !slices.Contains(cfg.filterFileNames, info.Name())
		}
		if fileSuffixesLen > 0 {
			isFile = isFile && !slices.Contains(cfg.fileterSuffixes, ext)
		}
		// 如果当前路径不是目录，则将其添加到文件路径切片中
		if isFile {
			fullName := info.Name()
			if cfg.hasPrefix {
				fullName = path
			}
			fullFileNames = append(fullFileNames, fullName)
		}
		return nil
	})
	// 如果 filepath.Walk 执行过程中出现错误，记录日志并返回错误
	if err != nil {
		log.Printf("Error walking directory: %v", err)
		return nil, err
	}
	// 返回文件路径切片，表示目录遍历成功
	return fullFileNames, nil
}

func GetDirFileListV2(dir string, hasPrefix bool) ([]string, error) {
	fileList, err := os.ReadDir(dir)
	if err != nil {
		return nil, err
	}
	var fullFileNames []string
	for _, f := range fileList {
		fullName := f.Name()
		if hasPrefix {
			fullName = filepath.Join(dir, f.Name())
		}
		fullFileNames = append(fullFileNames, fullName)
	}
	return fullFileNames, nil
}

// DirOrFileExists 检查指定路径是否存在。
//
// 参数:
//
//	path - 需要检查的文件或目录路径。
//
// 返回值:
//
//	第一个返回值表示指定路径的文件或目录是否存在，true表示存在，false表示不存在。
//	第二个返回值在发生错误时返回错误详情，否则返回nil。
func DirOrFileExists(path string) (bool, error) {
	// 尝试获取指定路径的文件或目录的信息。
	_, err := os.Stat(path)
	// 如果没有错误，说明文件或目录存在。
	if err == nil {
		return true, nil
	}
	// 如果错误是“不存在”，则表明文件或目录不存在，返回false和nil。
	if os.IsNotExist(err) {
		return false, nil
	}
	// 如果出现其他错误，将错误返回。
	return false, err
}
