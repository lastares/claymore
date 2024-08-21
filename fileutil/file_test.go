package fileutil

import (
	"strings"
	"testing"
)

func TestDownloadFile(t *testing.T) {
	testUrl := "http://iph.href.lu/579x200"
	testFilePath := "./testfile"
	// Execute
	err := DownloadFile(testUrl, testFilePath)
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
			got := GetBaseFullName(tt.filePath, tt.f)
			if got != tt.want {
				t.Errorf("GetBaseFullName() = %v, want %v", got, tt.want)
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
