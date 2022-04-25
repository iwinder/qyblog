package shutdown

import "sync"

type ShutdownCallback interface {
	OnShutdown(string) error
}
type ShutdownFunc func(string) error

func (f ShutdownFunc) OnShutdown(shutdownManager string) error {
	return f(shutdownManager)
}

type GSInterface interface {
	StartShutdown(sm ShutdownManager)
	ReportError(er error)
	AddShutdownCallback(callback ShutdownCallback)
}

type ShutdownManager interface {
	GetName() string
	Start(gs GSInterface) error
	ShutdownStart() error
	ShutdownFinish() error
}
type ErrorHandler interface {
	OnError(err error)
}

type ErrorFunc func(err error)

func (f ErrorFunc) OnError(err error) {
	f(err)
}

type GracefulShutdown struct {
	callBacks    []ShutdownCallback
	managers     []ShutdownManager
	errorHandler ErrorHandler
}

func (gs *GracefulShutdown) StartShutdown(sm ShutdownManager) {
	gs.ReportError(sm.ShutdownStart())

	var wg sync.WaitGroup
	for _, shutdownCallback := range gs.callBacks {
		wg.Add(1)
		go func(shutdownCallback ShutdownCallback) {
			defer wg.Done()
			gs.ReportError(shutdownCallback.OnShutdown(sm.GetName()))
		}(shutdownCallback)
	}

	wg.Wait()
	gs.ReportError(sm.ShutdownFinish())
}

func (gs *GracefulShutdown) ReportError(err error) {
	if err != nil && gs.errorHandler != nil {
		gs.errorHandler.OnError(err)
	}
}

func (gs *GracefulShutdown) AddShutdownCallback(callback ShutdownCallback) {
	gs.callBacks = append(gs.callBacks, callback)
}
func (gs *GracefulShutdown) AddShutdownManagee(manager ShutdownManager) {
	gs.managers = append(gs.managers, manager)
}

func (gs *GracefulShutdown) SetErrorHandler(errorHandler ErrorHandler) {
	gs.errorHandler = errorHandler
}

func New() *GracefulShutdown {
	return &GracefulShutdown{
		callBacks: make([]ShutdownCallback, 0, 10),
		managers:  make([]ShutdownManager, 0, 3),
	}
}

func (gs *GracefulShutdown) Start() error {
	for _, manager := range gs.managers {
		if err := manager.Start(gs); err != nil {
			return err
		}
	}
	return nil
}
