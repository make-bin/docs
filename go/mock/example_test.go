package mock

import (
	"testing"

	applemock "github.com/make-bin/docs/go/mock/mocks"
	"github.com/stretchr/testify/assert"
)

func TestGetAppleSize(t *testing.T) {

	// 创建mock 实现
	mocker := applemock.NewApple(t)
	// 设置mock 预期 , 模拟根据不同入惨,方法返回的值
	mocker.On("AppleSize", 10).Return(10)
	mocker.On("AppleSize", 20).Return(2)
	// 任何入参
	//mocker.On(mock.Anything).Return(1)

	// 测试代码
	s := &Size{
		app: mocker,
	}

	testCases := []struct {
		input int
		want  string
	}{
		{input: 10, want: "aaaaa"},
		{input: 20, want: "bbbbb"},
	}

	for _, tc := range testCases {
		s := s.GetSize(tc.input)
		// 断言输入和 每个 输出结果一致,则通过测试,
		assert.Equal(t, tc.want, s)
	}

	mocker.AssertExpectations(t)
	//mock.Mock 还有一些很实用的常用断言，如：
	//AssertCalled 断言被调用
	//AssertNotCalled 断言没被调用
	//AssertExpectations 断言 On 和 Return 设置的参数和返回值的方法，有被调用
	//AssertNumberOfCalls 断言调用次数
	/*
	   // 创建mock 实现
	   mocker := applemock.NewApple(t)
	   // 设置mock 预期 , 模拟根据不同入惨,方法返回的值
	   mocker.On("AppleSize", 10).Return(1)
	   // 任何入参
	   //mocker.On(mock.Anything).Return(1)

	   // 测试代码

	   	s := &Size{
	   		app: mocker,
	   	}

	   size := s.GetSize()
	   assert.Equal(t, 10, size)

	   mocker.AssertExpectations(t)
	   // 多个mocker
	   //mock.AssertExpectations(t, mocker1, mokcer2)
	*/
}
