package ziputil

import (
	"archive/zip"
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/text/encoding/simplifiedchinese"
	"golang.org/x/text/encoding/unicode"
	"golang.org/x/text/transform"
)

const (
	DirPerm         os.FileMode = 0755
	FilePerm        os.FileMode = 0644
	GBKEncodingFlag             = 2048
	NoEncodingFlag              = 0
)

// Zip 函数将源目录下的所有文件和子目录压缩到目标目录中。
// 参数 srcDir 是待压缩的源目录路径。
// 参数 destDir 是压缩后的目标目录路径。
// 返回error，表示压缩过程中可能发生的错误。
func Zip(srcDir string, destDir string) error {
	// 删除目标目录，以确保创建新的压缩文件。
	_ = os.RemoveAll(destDir)
	// 创建压缩文件。
	zf, err := os.Create(destDir)
	if err != nil {
		return err
	}
	// 确保在函数结束时关闭压缩文件。
	defer func() { _ = zf.Close() }()
	// 初始化压缩文件的写入器。
	archive := zip.NewWriter(zf)
	// 确保在函数结束时关闭写入器。
	defer func() { _ = archive.Close() }()
	// 遍历源目录，将每个文件和子目录添加到压缩文件中。
	return filepath.Walk(srcDir, func(path string, info os.FileInfo, _ error) error {
		// 跳过源目录本身，只压缩其内容。
		if path == srcDir {
			return nil
		}
		// 根据文件信息创建压缩条目。
		header, err := zip.FileInfoHeader(info)
		if err != nil {
			return err
		}
		// 调整文件路径，以确保在压缩文件中的相对路径是正确的。
		header.Name = strings.TrimPrefix(path, srcDir+string(os.PathSeparator))
		// 如果是目录，确保文件名以斜杠结尾。
		if info.IsDir() {
			header.Name += `/`
		} else {
			// 设置文件的压缩方法为Deflate。
			header.Method = zip.Deflate
		}
		// 创建压缩条目的实际写入器。
		writer, err := archive.CreateHeader(header)
		if err != nil {
			return err
		}
		// 如果不是目录，则打开文件并将其内容写入压缩文件。
		if !info.IsDir() {
			file, err := os.Open(path)
			if err != nil {
				return err
			}
			// 确保在函数结束时关闭打开的文件。
			defer func() { _ = file.Close() }()
			_, err = io.Copy(writer, file)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

// Unzip 解压缩指定的zip文件到目标目录
// zipFileName 是要解压缩的zip文件路径
// targetDirectory 是解压缩后文件存放的目标目录,如果目标目录参数为空，则将其默认设置为ZIP文件名所在目录。
func Unzip(zipFileName string, targetDirectory string) error {
	// url.QueryUnescape 用于解码URL编码的字符串，防止出现错误
	zipFileName, _ = url.QueryUnescape(zipFileName)
	// makeDefaultDirectory 根据zip文件名生成默认的目标目录
	targetDirectory = makeDefaultDirectory(zipFileName, targetDirectory)
	// 打开zip文件以进行读取
	zipReader, err := zip.OpenReader(zipFileName)
	if err != nil {
		// 如果无法打开zip文件，返回错误
		return fmt.Errorf("failed to open zip file: %w", err)
	}
	// 确保在函数结束时关闭zip文件
	defer func(zipReader *zip.ReadCloser) {
		_ = zipReader.Close()
	}(zipReader)
	// 遍历zip文件中的每个文件
	for _, f := range zipReader.File {
		// GetDecodeName 用于获取文件的解码后名称
		decodeName, err := GetDecodeName(f)
		if err != nil {
			// 如果获取解码名称失败，返回错误
			return err
		}
		// 过滤掉与 _MACOSX 相关的文件
		if strings.Contains(decodeName, "_MACOSX") {
			continue
		}
		// 计算文件的完整路径
		fpath := filepath.Clean(filepath.Join(targetDirectory, decodeName))
		// 检查文件路径是否安全
		err = checkDirectorySafe(fpath, targetDirectory)
		if err != nil {
			return err
		}
		// 如果文件不是目录，则处理文件
		if !f.FileInfo().IsDir() {
			// 确保文件所在的目录存在
			dir := filepath.Dir(fpath)
			if err = ensureDirectoryExists(dir, DirPerm); err != nil {
				return err
			}
			// 处理文件内容
			if err := processFile(f, fpath); err != nil {
				return err
			}
		}
	}
	// 所有文件处理完毕，返回nil表示成功
	return nil
}

// ensureDirectoryExists 确保指定路径的目录存在，如果不存在则创建该目录。
// 这个函数主要用于在程序运行前准备好必要的目录资源，避免因为目录缺失导致的错误。
// 参数:
//
//	path: 字符串类型，指定需要检查或创建的目录路径。
//	perm: os.FileMode类型，定义了目录的权限设置。
//
// 返回值:
//
//	如果目录创建失败，返回错误信息；如果成功，返回nil。
func ensureDirectoryExists(path string, perm os.FileMode) error {
	// 使用os.MkdirAll尝试创建目录，如果遇到错误则进行处理。
	// os.MkdirAll函数会创建path路径中的所有不存在的目录，并设置相应的权限。
	if err := os.MkdirAll(path, perm); err != nil {
		// 如果创建目录失败，返回自定义的错误信息，包含原始错误。
		// 这里使用fmt.Errorf进行错误格式化，%s占位符用于插入path字符串，%w占位符用于插入原始错误对象。
		return fmt.Errorf("failed to create directory %s: %w", path, err)
	}
	// 如果目录创建成功或者目标目录已经存在，则返回nil，表示没有遇到错误。
	return nil
}

// makeDefaultDirectory 根据ZIP文件名和目标目录参数来确定默认的目标目录。
// 如果目标目录参数为空，则将其默认设置为ZIP文件名所在目录。
func makeDefaultDirectory(zipFileName, targetDirectory string) string {
	// 检查目标目录是否为空
	if targetDirectory == "" {
		// 如果为空，则将目标目录设置为ZIP文件所在的目录
		targetDirectory = filepath.Dir(zipFileName)
	}
	// 返回最终的目标目录
	return targetDirectory
}

// checkDirectorySafe 检查文件路径是否安全。
// 它确保文件路径位于目标目录内，以防止路径遍历攻击。
// 参数:
//
//	fpath: 需要检查的文件路径。
//	targetDirectory: 目标目录路径。
//
// 返回值:
//
//	如果文件路径安全，则返回nil；否则返回相应的错误。
func checkDirectorySafe(fpath, targetDirectory string) error {
	// 获取目标目录的绝对路径，用于后续比较。
	absTargetDirectory, err := filepath.Abs(targetDirectory)
	if err != nil {
		return fmt.Errorf("failed to get absolute path of target directory: %w", err)
	}
	// 获取文件的绝对路径，用于后续比较。
	absFPath, err := filepath.Abs(fpath)
	if err != nil {
		return fmt.Errorf("failed to get absolute path of file: %w", err)
	}
	// 检查文件的绝对路径是否以目标目录的绝对路径前缀，以确保文件位于目标目录内。
	if !strings.HasPrefix(absFPath, absTargetDirectory) {
		return errors.New("invalid file path")
	}
	// 文件路径安全，无需操作。
	return nil
}

// GetDecodeName 根据 zip 文件的标志位获取解码后的文件名。
// 这个函数处理不同编码的文件名，确保在不同场景下都能正确解读。
//
// 参数:
// f *zip.File: 一个指向 zip 文件的指针，包含文件信息和属性。
//
// 返回值:
// string: 解码后的文件名。
// error: 如果解码过程中发生错误，返回错误信息；否则返回 nil。
func GetDecodeName(f *zip.File) (string, error) {
	// 初始化解码后的文件名变量
	var decodeName string

	// 根据文件标志位选择合适的编码方式
	switch f.Flags {
	case NoEncodingFlag:
		// 默认编码 (UTF-8)
		decodeName = f.Name
	case GBKEncodingFlag:
		// GBK 编码
		b := []byte(f.Name)
		decoder := transform.Chain(simplifiedchinese.GB18030.NewDecoder(), unicode.UTF8.NewDecoder())
		content, _, err := transform.Bytes(decoder, b)
		if err != nil {
			return "", fmt.Errorf("failed to decode filename: %w", err)
		}
		decodeName = string(content)
	default:
		// 如果标志位不是已知的值，则按默认编码处理
		decodeName = f.Name
	}

	// 返回解码后的文件名，如果没有解码过程中的错误，返回 nil
	return decodeName, nil
}

// processFile 处理指定的压缩文件成员，将其内容复制到文件系统中的指定位置。
// 参数:
//
//	file: *zip.File - 压缩文件中的文件成员。
//	fpath: string - 文件系统中的目标路径。
//
// 返回值:
//
//	error - 表示函数执行过程中是否遇到了错误，如果没有错误，返回nil。
func processFile(file *zip.File, fpath string) error {
	// 打开压缩文件成员以读取其内容。
	inFile, err := file.Open()
	if err != nil {
		// 如果无法打开文件，则返回错误。
		return fmt.Errorf("failed to open file %s: %w", file.Name, err)
	}
	// 确保在函数返回后关闭文件。
	defer func(inFile io.ReadCloser) {
		_ = inFile.Close()
	}(inFile)

	// 创建或打开目标文件以写入压缩文件成员的内容。
	outFile, err := os.OpenFile(fpath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, FilePerm)
	if err != nil {
		// 如果无法创建或打开目标文件，则返回错误。
		return fmt.Errorf("failed to create file %s: %w", file.Name, err)
	}
	// 确保在函数返回后关闭目标文件。
	defer func(outFile *os.File) {
		_ = outFile.Close()
	}(outFile)
	// 复制文件内容时发生错误时的处理逻辑。
	_, err = io.Copy(outFile, inFile)
	if err != nil {
		// 尝试回滚操作，即删除部分写入的文件。
		if rollbackErr := os.Remove(fpath); rollbackErr != nil {
			// 如果回滚操作失败，返回错误。
			return fmt.Errorf("failed to remove partially written file %s: %w", file.Name, rollbackErr)
		}
		// 返回复制数据过程中遇到的错误。
		return fmt.Errorf("failed to copy data for file %s: %w", file.Name, err)
	}
	// 如果没有遇到错误，返回nil。
	return nil
}
