package ziputil

import (
	"testing"
)

func TestUnzip(t *testing.T) {
	// 此处假设 test.zip 是一个有效的zip文件路径
	zipFilePath := "/Users/ares/GolandProjects/claymore/ziputil/课程专栏.zip"
	// 调用待测试的函数
	err := Unzip(zipFilePath, "")
	if err != nil {
		t.Errorf("Unzip() error = %v", err)
	}
}
