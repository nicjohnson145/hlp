package sqlx

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

type mockPinger struct {
	ErrorTimes int
	count      int
}

func (m *mockPinger) PingContext(_ context.Context) error {
	m.count += 1

	if m.ErrorTimes == -1 || m.count <= m.ErrorTimes {
		return fmt.Errorf("some error")
	}

	return nil
}

func TestWaitForDBConnectable(t *testing.T) {
	t.Run("1_error_then_connect", func(t *testing.T) {
		pinger := &mockPinger{ErrorTimes: 1}

		timer := time.NewTimer(3 * time.Second)
		defer timer.Stop()

		waitTimeout := 2 * time.Second

		run := func(c chan bool, e chan error) {
			err := WaitForDBConnectable(pinger, DBWaitOpts{Timeout: &waitTimeout})
			e <- err
			c <- true
		}

		doneChan := make(chan bool, 1)
		errChan := make(chan error, 1)
		go run(doneChan, errChan)

		select {
		case <-timer.C:
			t.Fatal("Timeout exceeded")
		case <-doneChan:
			// Neat, test passes
		}

		require.NoError(t, <-errChan)
	})
}
