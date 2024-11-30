package mock

import (
	"testing"

	applemock "github.com/make-bin/docs/go/mock/mocks"
)

func TestGetAppleSize(t *testing.T) {

	// 创建mock 实现
	mocker := applemock.NewApple(t)
	// 设置mock 预期
	mocker.On("AppleSize").Return(10)
	// 测试代码
	s := &Size{
		app: mocker,
	}
	s.GetSize()
}
