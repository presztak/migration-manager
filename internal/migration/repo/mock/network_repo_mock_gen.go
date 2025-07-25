// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package mock

import (
	"context"
	"sync"

	"github.com/FuturFusion/migration-manager/internal/migration"
)

// Ensure, that NetworkRepoMock does implement migration.NetworkRepo.
// If this is not the case, regenerate this file with moq.
var _ migration.NetworkRepo = &NetworkRepoMock{}

// NetworkRepoMock is a mock implementation of migration.NetworkRepo.
//
//	func TestSomethingThatUsesNetworkRepo(t *testing.T) {
//
//		// make and configure a mocked migration.NetworkRepo
//		mockedNetworkRepo := &NetworkRepoMock{
//			CreateFunc: func(ctx context.Context, network migration.Network) (int64, error) {
//				panic("mock out the Create method")
//			},
//			DeleteByNameAndSourceFunc: func(ctx context.Context, name string, src string) error {
//				panic("mock out the DeleteByNameAndSource method")
//			},
//			GetAllFunc: func(ctx context.Context) (migration.Networks, error) {
//				panic("mock out the GetAll method")
//			},
//			GetAllBySourceFunc: func(ctx context.Context, src string) (migration.Networks, error) {
//				panic("mock out the GetAllBySource method")
//			},
//			GetByNameAndSourceFunc: func(ctx context.Context, name string, src string) (*migration.Network, error) {
//				panic("mock out the GetByNameAndSource method")
//			},
//			UpdateFunc: func(ctx context.Context, network migration.Network) error {
//				panic("mock out the Update method")
//			},
//		}
//
//		// use mockedNetworkRepo in code that requires migration.NetworkRepo
//		// and then make assertions.
//
//	}
type NetworkRepoMock struct {
	// CreateFunc mocks the Create method.
	CreateFunc func(ctx context.Context, network migration.Network) (int64, error)

	// DeleteByNameAndSourceFunc mocks the DeleteByNameAndSource method.
	DeleteByNameAndSourceFunc func(ctx context.Context, name string, src string) error

	// GetAllFunc mocks the GetAll method.
	GetAllFunc func(ctx context.Context) (migration.Networks, error)

	// GetAllBySourceFunc mocks the GetAllBySource method.
	GetAllBySourceFunc func(ctx context.Context, src string) (migration.Networks, error)

	// GetByNameAndSourceFunc mocks the GetByNameAndSource method.
	GetByNameAndSourceFunc func(ctx context.Context, name string, src string) (*migration.Network, error)

	// UpdateFunc mocks the Update method.
	UpdateFunc func(ctx context.Context, network migration.Network) error

	// calls tracks calls to the methods.
	calls struct {
		// Create holds details about calls to the Create method.
		Create []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Network is the network argument value.
			Network migration.Network
		}
		// DeleteByNameAndSource holds details about calls to the DeleteByNameAndSource method.
		DeleteByNameAndSource []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Name is the name argument value.
			Name string
			// Src is the src argument value.
			Src string
		}
		// GetAll holds details about calls to the GetAll method.
		GetAll []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
		}
		// GetAllBySource holds details about calls to the GetAllBySource method.
		GetAllBySource []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Src is the src argument value.
			Src string
		}
		// GetByNameAndSource holds details about calls to the GetByNameAndSource method.
		GetByNameAndSource []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Name is the name argument value.
			Name string
			// Src is the src argument value.
			Src string
		}
		// Update holds details about calls to the Update method.
		Update []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// Network is the network argument value.
			Network migration.Network
		}
	}
	lockCreate                sync.RWMutex
	lockDeleteByNameAndSource sync.RWMutex
	lockGetAll                sync.RWMutex
	lockGetAllBySource        sync.RWMutex
	lockGetByNameAndSource    sync.RWMutex
	lockUpdate                sync.RWMutex
}

// Create calls CreateFunc.
func (mock *NetworkRepoMock) Create(ctx context.Context, network migration.Network) (int64, error) {
	if mock.CreateFunc == nil {
		panic("NetworkRepoMock.CreateFunc: method is nil but NetworkRepo.Create was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Network migration.Network
	}{
		Ctx:     ctx,
		Network: network,
	}
	mock.lockCreate.Lock()
	mock.calls.Create = append(mock.calls.Create, callInfo)
	mock.lockCreate.Unlock()
	return mock.CreateFunc(ctx, network)
}

// CreateCalls gets all the calls that were made to Create.
// Check the length with:
//
//	len(mockedNetworkRepo.CreateCalls())
func (mock *NetworkRepoMock) CreateCalls() []struct {
	Ctx     context.Context
	Network migration.Network
} {
	var calls []struct {
		Ctx     context.Context
		Network migration.Network
	}
	mock.lockCreate.RLock()
	calls = mock.calls.Create
	mock.lockCreate.RUnlock()
	return calls
}

// DeleteByNameAndSource calls DeleteByNameAndSourceFunc.
func (mock *NetworkRepoMock) DeleteByNameAndSource(ctx context.Context, name string, src string) error {
	if mock.DeleteByNameAndSourceFunc == nil {
		panic("NetworkRepoMock.DeleteByNameAndSourceFunc: method is nil but NetworkRepo.DeleteByNameAndSource was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Name string
		Src  string
	}{
		Ctx:  ctx,
		Name: name,
		Src:  src,
	}
	mock.lockDeleteByNameAndSource.Lock()
	mock.calls.DeleteByNameAndSource = append(mock.calls.DeleteByNameAndSource, callInfo)
	mock.lockDeleteByNameAndSource.Unlock()
	return mock.DeleteByNameAndSourceFunc(ctx, name, src)
}

// DeleteByNameAndSourceCalls gets all the calls that were made to DeleteByNameAndSource.
// Check the length with:
//
//	len(mockedNetworkRepo.DeleteByNameAndSourceCalls())
func (mock *NetworkRepoMock) DeleteByNameAndSourceCalls() []struct {
	Ctx  context.Context
	Name string
	Src  string
} {
	var calls []struct {
		Ctx  context.Context
		Name string
		Src  string
	}
	mock.lockDeleteByNameAndSource.RLock()
	calls = mock.calls.DeleteByNameAndSource
	mock.lockDeleteByNameAndSource.RUnlock()
	return calls
}

// GetAll calls GetAllFunc.
func (mock *NetworkRepoMock) GetAll(ctx context.Context) (migration.Networks, error) {
	if mock.GetAllFunc == nil {
		panic("NetworkRepoMock.GetAllFunc: method is nil but NetworkRepo.GetAll was just called")
	}
	callInfo := struct {
		Ctx context.Context
	}{
		Ctx: ctx,
	}
	mock.lockGetAll.Lock()
	mock.calls.GetAll = append(mock.calls.GetAll, callInfo)
	mock.lockGetAll.Unlock()
	return mock.GetAllFunc(ctx)
}

// GetAllCalls gets all the calls that were made to GetAll.
// Check the length with:
//
//	len(mockedNetworkRepo.GetAllCalls())
func (mock *NetworkRepoMock) GetAllCalls() []struct {
	Ctx context.Context
} {
	var calls []struct {
		Ctx context.Context
	}
	mock.lockGetAll.RLock()
	calls = mock.calls.GetAll
	mock.lockGetAll.RUnlock()
	return calls
}

// GetAllBySource calls GetAllBySourceFunc.
func (mock *NetworkRepoMock) GetAllBySource(ctx context.Context, src string) (migration.Networks, error) {
	if mock.GetAllBySourceFunc == nil {
		panic("NetworkRepoMock.GetAllBySourceFunc: method is nil but NetworkRepo.GetAllBySource was just called")
	}
	callInfo := struct {
		Ctx context.Context
		Src string
	}{
		Ctx: ctx,
		Src: src,
	}
	mock.lockGetAllBySource.Lock()
	mock.calls.GetAllBySource = append(mock.calls.GetAllBySource, callInfo)
	mock.lockGetAllBySource.Unlock()
	return mock.GetAllBySourceFunc(ctx, src)
}

// GetAllBySourceCalls gets all the calls that were made to GetAllBySource.
// Check the length with:
//
//	len(mockedNetworkRepo.GetAllBySourceCalls())
func (mock *NetworkRepoMock) GetAllBySourceCalls() []struct {
	Ctx context.Context
	Src string
} {
	var calls []struct {
		Ctx context.Context
		Src string
	}
	mock.lockGetAllBySource.RLock()
	calls = mock.calls.GetAllBySource
	mock.lockGetAllBySource.RUnlock()
	return calls
}

// GetByNameAndSource calls GetByNameAndSourceFunc.
func (mock *NetworkRepoMock) GetByNameAndSource(ctx context.Context, name string, src string) (*migration.Network, error) {
	if mock.GetByNameAndSourceFunc == nil {
		panic("NetworkRepoMock.GetByNameAndSourceFunc: method is nil but NetworkRepo.GetByNameAndSource was just called")
	}
	callInfo := struct {
		Ctx  context.Context
		Name string
		Src  string
	}{
		Ctx:  ctx,
		Name: name,
		Src:  src,
	}
	mock.lockGetByNameAndSource.Lock()
	mock.calls.GetByNameAndSource = append(mock.calls.GetByNameAndSource, callInfo)
	mock.lockGetByNameAndSource.Unlock()
	return mock.GetByNameAndSourceFunc(ctx, name, src)
}

// GetByNameAndSourceCalls gets all the calls that were made to GetByNameAndSource.
// Check the length with:
//
//	len(mockedNetworkRepo.GetByNameAndSourceCalls())
func (mock *NetworkRepoMock) GetByNameAndSourceCalls() []struct {
	Ctx  context.Context
	Name string
	Src  string
} {
	var calls []struct {
		Ctx  context.Context
		Name string
		Src  string
	}
	mock.lockGetByNameAndSource.RLock()
	calls = mock.calls.GetByNameAndSource
	mock.lockGetByNameAndSource.RUnlock()
	return calls
}

// Update calls UpdateFunc.
func (mock *NetworkRepoMock) Update(ctx context.Context, network migration.Network) error {
	if mock.UpdateFunc == nil {
		panic("NetworkRepoMock.UpdateFunc: method is nil but NetworkRepo.Update was just called")
	}
	callInfo := struct {
		Ctx     context.Context
		Network migration.Network
	}{
		Ctx:     ctx,
		Network: network,
	}
	mock.lockUpdate.Lock()
	mock.calls.Update = append(mock.calls.Update, callInfo)
	mock.lockUpdate.Unlock()
	return mock.UpdateFunc(ctx, network)
}

// UpdateCalls gets all the calls that were made to Update.
// Check the length with:
//
//	len(mockedNetworkRepo.UpdateCalls())
func (mock *NetworkRepoMock) UpdateCalls() []struct {
	Ctx     context.Context
	Network migration.Network
} {
	var calls []struct {
		Ctx     context.Context
		Network migration.Network
	}
	mock.lockUpdate.RLock()
	calls = mock.calls.Update
	mock.lockUpdate.RUnlock()
	return calls
}
