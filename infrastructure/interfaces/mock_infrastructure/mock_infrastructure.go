// Code generated by MockGen. DO NOT EDIT.
// Source: github.com/int128/amefuriso/infrastructure/interfaces (interfaces: WeatherClient,SlackClient)

// Package mock_infrastructure is a generated GoMock package.
package mock_infrastructure

import (
	context "context"
	gomock "github.com/golang/mock/gomock"
	weather "github.com/int128/go-yahoo-weather/weather"
	slack "github.com/int128/slack"
	reflect "reflect"
)

// MockWeatherClient is a mock of WeatherClient interface
type MockWeatherClient struct {
	ctrl     *gomock.Controller
	recorder *MockWeatherClientMockRecorder
}

// MockWeatherClientMockRecorder is the mock recorder for MockWeatherClient
type MockWeatherClientMockRecorder struct {
	mock *MockWeatherClient
}

// NewMockWeatherClient creates a new mock instance
func NewMockWeatherClient(ctrl *gomock.Controller) *MockWeatherClient {
	mock := &MockWeatherClient{ctrl: ctrl}
	mock.recorder = &MockWeatherClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockWeatherClient) EXPECT() *MockWeatherClientMockRecorder {
	return m.recorder
}

// Get mocks base method
func (m *MockWeatherClient) Get(arg0 context.Context, arg1 string, arg2 weather.Request) ([]weather.Weather, error) {
	ret := m.ctrl.Call(m, "Get", arg0, arg1, arg2)
	ret0, _ := ret[0].([]weather.Weather)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Get indicates an expected call of Get
func (mr *MockWeatherClientMockRecorder) Get(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Get", reflect.TypeOf((*MockWeatherClient)(nil).Get), arg0, arg1, arg2)
}

// MockSlackClient is a mock of SlackClient interface
type MockSlackClient struct {
	ctrl     *gomock.Controller
	recorder *MockSlackClientMockRecorder
}

// MockSlackClientMockRecorder is the mock recorder for MockSlackClient
type MockSlackClientMockRecorder struct {
	mock *MockSlackClient
}

// NewMockSlackClient creates a new mock instance
func NewMockSlackClient(ctrl *gomock.Controller) *MockSlackClient {
	mock := &MockSlackClient{ctrl: ctrl}
	mock.recorder = &MockSlackClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockSlackClient) EXPECT() *MockSlackClientMockRecorder {
	return m.recorder
}

// Send mocks base method
func (m *MockSlackClient) Send(arg0 context.Context, arg1 string, arg2 slack.Message) error {
	ret := m.ctrl.Call(m, "Send", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Send indicates an expected call of Send
func (mr *MockSlackClientMockRecorder) Send(arg0, arg1, arg2 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Send", reflect.TypeOf((*MockSlackClient)(nil).Send), arg0, arg1, arg2)
}
