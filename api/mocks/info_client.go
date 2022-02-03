// Code generated by mockery v2.9.4. DO NOT EDIT.

package mocks

import (
	context "context"

	info "github.com/ava-labs/avalanchego/api/info"
	ids "github.com/ava-labs/avalanchego/ids"

	mock "github.com/stretchr/testify/mock"

	network "github.com/ava-labs/avalanchego/network"
)

// InfoClient is an autogenerated mock type for the Client type
type InfoClient struct {
	mock.Mock
}

// GetBlockchainID provides a mock function with given fields: _a0, _a1
func (_m *InfoClient) GetBlockchainID(_a0 context.Context, _a1 string) (ids.ID, error) {
	ret := _m.Called(_a0, _a1)

	var r0 ids.ID
	if rf, ok := ret.Get(0).(func(context.Context, string) ids.ID); ok {
		r0 = rf(_a0, _a1)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ids.ID)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNetworkID provides a mock function with given fields: _a0
func (_m *InfoClient) GetNetworkID(_a0 context.Context) (uint32, error) {
	ret := _m.Called(_a0)

	var r0 uint32
	if rf, ok := ret.Get(0).(func(context.Context) uint32); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(uint32)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNetworkName provides a mock function with given fields: _a0
func (_m *InfoClient) GetNetworkName(_a0 context.Context) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNodeID provides a mock function with given fields: _a0
func (_m *InfoClient) GetNodeID(_a0 context.Context) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNodeIP provides a mock function with given fields: _a0
func (_m *InfoClient) GetNodeIP(_a0 context.Context) (string, error) {
	ret := _m.Called(_a0)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context) string); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetNodeVersion provides a mock function with given fields: _a0
func (_m *InfoClient) GetNodeVersion(_a0 context.Context) (*info.GetNodeVersionReply, error) {
	ret := _m.Called(_a0)

	var r0 *info.GetNodeVersionReply
	if rf, ok := ret.Get(0).(func(context.Context) *info.GetNodeVersionReply); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*info.GetNodeVersionReply)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetTxFee provides a mock function with given fields: _a0
func (_m *InfoClient) GetTxFee(_a0 context.Context) (*info.GetTxFeeResponse, error) {
	ret := _m.Called(_a0)

	var r0 *info.GetTxFeeResponse
	if rf, ok := ret.Get(0).(func(context.Context) *info.GetTxFeeResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*info.GetTxFeeResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsBootstrapped provides a mock function with given fields: _a0, _a1
func (_m *InfoClient) IsBootstrapped(_a0 context.Context, _a1 string) (bool, error) {
	ret := _m.Called(_a0, _a1)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(_a0, _a1)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Peers provides a mock function with given fields: _a0
func (_m *InfoClient) Peers(_a0 context.Context) ([]network.PeerInfo, error) {
	ret := _m.Called(_a0)

	var r0 []network.PeerInfo
	if rf, ok := ret.Get(0).(func(context.Context) []network.PeerInfo); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]network.PeerInfo)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Uptime provides a mock function with given fields: _a0
func (_m *InfoClient) Uptime(_a0 context.Context) (*info.UptimeResponse, error) {
	ret := _m.Called(_a0)

	var r0 *info.UptimeResponse
	if rf, ok := ret.Get(0).(func(context.Context) *info.UptimeResponse); ok {
		r0 = rf(_a0)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*info.UptimeResponse)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(_a0)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
