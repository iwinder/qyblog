package shutdownmanagers

import (
	"gitee.com/windcoder/qingyucms/internal/pkg/qy-common/shutdown"
	"os"
	"os/signal"
	"syscall"
)

const Name = "PosixSignalManager"

type PosixSignalManager struct {
	signals []os.Signal
}

func (posixSignalManager *PosixSignalManager) ShutdownStart() error {
	return nil
}

func (posixSignalManager *PosixSignalManager) ShutdownFinish() error {
	return nil
}

func NewPosixSignalManager(sig ...os.Signal) *PosixSignalManager {
	if len(sig) == 0 {
		sig = make([]os.Signal, 2)
		sig[0] = os.Interrupt
		sig[1] = syscall.SIGTERM
	}
	return &PosixSignalManager{signals: sig}
}

func (posixSignalManager *PosixSignalManager) GetName() string {
	return Name
}

func (psm *PosixSignalManager) Start(gs shutdown.GSInterface) error {
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, psm.signals...)
		<-c
		gs.StartShutdown(psm)
	}()
	return nil
}
