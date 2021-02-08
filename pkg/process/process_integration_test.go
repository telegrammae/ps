package process_test

import (
	"task-manager/pkg/process"
	"testing"
	"time"
)

func TestUnixProcess(t *testing.T) {
	p := process.NewProcess("sleep", "3s")

	p.Start()

	if !p.IsRunning() {
		t.Errorf("process is supposed to be running")
	}

	time.Sleep(time.Second)

	err := p.Stop()
	if err != nil {
		t.Fatal(err)
	}

	if p.IsRunning() {
		t.Errorf("process is not supposed to be running")
	}
}
