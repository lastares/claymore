package fileutil

import (
	"fmt"
	"os"
	"path/filepath"
	"slices"
	"strings"
	"testing"
)

var dir = "/Users/ares/GolandProjects/claymore/ziputil"

func TestDownload(t *testing.T) {
	testUrl := "http://iph.href.lu/579x200"
	testFilePath := "./testfile"
	// Execute
	err := Download(testUrl, testFilePath)
	// Verify
	if err != nil {
		t.Errorf("Failed to download file: %s", err)
	}
}

func TestGetBaseFullName(t *testing.T) {
	tests := []struct {
		filePath string
		want     string
		f        func(filePath string) string
	}{
		{filePath: "C:\\Windows\\System32", f: nil, want: "System32"},
		{filePath: "C:\\Windows\\System32\\", f: nil, want: "System32"},
		{filePath: "C:\\Windows\\System32\\notepad.exe", f: nil, want: "notepad.exe"},
		{filePath: "/usr/bin", f: nil, want: "bin"},
		{filePath: "/usr/bin/哈哈aaa.txt", f: nil, want: "哈哈aaa.txt"},
		{filePath: "%2Fusr%2Fbin%2F%E5%93%88%E5%93%88aaa.txt", f: nil, want: "哈哈aaa.txt"},
		{filePath: "/usr/bin/", f: nil, want: "bin"},
		{filePath: "/usr/bin/vim", f: nil, want: "vim"},
		{filePath: "/usr/bin/vim/aa.txt", f: nil, want: "aa.txt"},
		{filePath: "C:\\Program%20Files\\", f: nil, want: "Program%20Files"},
		{filePath: "/home/user/documents/report.txt", f: nil, want: "report.txt"},
		{filePath: "C:\\Windows\\System32\\notepad.exe", f: nil, want: "notepad.exe"},
		{
			filePath: "http://www.baidu.com/xxx/232dd/232lskdjfskdfl323.jpg?st=xxxxxxxzsdsdfsdfsdf",
			f: func(filePath string) string {
				if strings.Contains(filePath, "?") {
					return strings.Split(filePath, "?")[0]
				}
				return filePath
			},
			want: "232lskdjfskdfl323.jpg"},
		{filePath: "/usr/bin/", f: nil, want: "bin"},
		{filePath: "", f: nil, want: ""},
		{filePath: "http://www.baidu.com/sdsdf/232s/aaaaa.txt", f: nil, want: "aaaaa.txt"},
	}

	for _, tt := range tests {
		t.Run(tt.filePath, func(t *testing.T) {
			got := GetFullName(tt.filePath, tt.f)
			if got != tt.want {
				t.Errorf("GetFullName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestGetExtension(t *testing.T) {
	tests := []struct {
		filepath string
		want     string
	}{
		{"https://www.baidu.com/xxx/test.txt", ".txt"},
		{"test.txt", ".txt"},
		{"test", ""},
		{"", ""},
	}

	for _, tt := range tests {
		t.Run(tt.filepath, func(t *testing.T) {
			if got := GetExtension(tt.filepath); got != tt.want {
				t.Errorf("GetExtension() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestGetBaseName 测试 GetBaseName 函数
func TestGetBaseName(t *testing.T) {
	testCases := []struct {
		filePath string
		want     string
	}{
		{"http://www.test.com/223/33/哈哈哈.jpg", "哈哈哈"},
		{"/Users/ares/test.txt", "test"},
		{"", ""},
	}

	for _, tc := range testCases {
		t.Run(tc.filePath, func(t *testing.T) {
			got := GetBaseName(tc.filePath)
			if got != tc.want {
				t.Errorf("GetBaseName(%q) = %q; want %q", tc.filePath, got, tc.want)
			}
		})
	}
}

// TestGetExtNoDot 测试 GetExtNoDot 函数
func TestGetExtNoDot(t *testing.T) {
	tests := []struct {
		filePath string
		want     string
	}{
		{"", ""},
		{"file.txt", "txt"},
		{"file.txt.gz", "gz"},
	}
	for _, tt := range tests {
		t.Run(tt.filePath, func(t *testing.T) {
			got := GetExtNoDot(tt.filePath)
			if got != tt.want {
				t.Errorf("GetExtNoDot(%q) = %q, want %q", tt.filePath, got, tt.want)
			}
		})
	}
}

func TestGetDirFileListV2(t *testing.T) {
	dir := "/Users/ares/GolandProjects/claymore/ziputil"
	tests := []struct {
		dir       string
		hasPrefix bool
		want      []string
	}{
		// 添加测试用例：目录路径，是否有前缀
		{dir, false, []string{"zip.go", "zip_test.go", "课程专栏.zip"}},
		{dir, true, []string{
			filepath.Join(dir, "zip.go"),
			filepath.Join(dir, "zip_test.go"),
			filepath.Join(dir, "课程专栏.zip"),
		}},
	}
	for _, test := range tests {
		t.Run(test.dir, func(t *testing.T) {
			got, err := GetDirFileListV2(test.dir, test.hasPrefix)
			if err != nil {
				t.Errorf("GetDirFileList(%q) error: %v", test.dir, err)
			}
			for _, file := range got {
				if !slices.Contains(test.want, file) {
					t.Errorf("GetDirFileList(%q) got unexpected file: %q", test.dir, file)
				}
			}
		})
	}
}

// TestGetDirFileList 测试 GetDirFileList 函数
func TestGetDirFileList(t *testing.T) {
	// 定义测试用例
	cases := []struct {
		dir             string
		hasPrefix       bool
		fileterSuffixes []string
		fileterFiles    []string
	}{
		// 添加测试用例：目录路径，是否有前缀
		{dir, true, []string{}, []string{"zip.go"}},
		//{"/Users/ares/GolandProjects/claymore/ziputil", true},
	}

	for _, c := range cases {
		// 调用待测试的函数
		files, err := GetDirFileList(c.dir, WithHasPrefix(c.hasPrefix), WithFileterSuffixes(c.fileterSuffixes...), WithFilterFileNames(c.fileterFiles...))
		for _, file := range files {
			fmt.Println(file)
		}
		// 验证结果是否正确
		if err != nil {
			t.Errorf("GetDirFileList(%q, %v) error: %v", c.dir, c.hasPrefix, err)
		}

		// 验证返回的文件列表是否不为空
		if len(files) == 0 {
			t.Errorf("GetDirFileList(%q, %v) returned empty slice", c.dir, c.hasPrefix)
		}

		// 验证返回的文件列表是否包含预期的文件
		expectedFiles := []string{"课程专栏.zip", "zip_test.go"}
		if c.hasPrefix {
			expectedFiles = []string{
				filepath.Join(c.dir, "课程专栏.zip"),
				filepath.Join(c.dir, "zip_test.go"),
				//filepath.Join(c.dir, "zip.go"),
			}
		}
		for _, expectedFile := range expectedFiles {
			found := false
			for _, file := range files {
				if file == expectedFile {
					found = true
					break
				}
			}
			if !found {
				t.Errorf("GetDirFileList(%q, %v) did not find expected file: %q", c.dir, c.hasPrefix, expectedFile)
			}
		}
	}
}

// TestDirOrFileExists 测试 DirOrFileExists 函数
func TestDirOrFileExists(t *testing.T) {
	// 定义测试用例
	tests := []struct {
		path     string
		expected bool
	}{
		// 添加测试用例：存在的情况
		{dir, true},             // 当前目录
		{dir + "/zip.go", true}, // main.go 文件，假定在同一目录下
		// 添加测试用例：不存在的情况
		{dir + "/zip1.go", false},
	}

	for _, test := range tests {
		t.Run(test.path, func(t *testing.T) {
			// 调用待测函数
			exists, err := DirOrFileExists(test.path)

			// 验证结果是否正确
			if exists != test.expected {
				t.Errorf("DirOrFileExists(%s) = %v; want %v", test.path, exists, test.expected)
			}

			// 验证是否有错误发生，除了不存在的情况
			if err != nil && !os.IsNotExist(err) {
				t.Errorf("DirOrFileExists(%s) error = %v; want no error", test.path, err)
			}
		})
	}
}
