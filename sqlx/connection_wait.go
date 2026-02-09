package sqlx

import (
	"context"
	"fmt"
	"time"

	"github.com/sethvargo/go-retry"
	"github.com/go-logr/logr"
)

// DBWaitOpts are options used to configure the behavior of WaitForDBConnectable
type DBWaitOpts struct {
	// Timeout is the amount of time to wait for the DB to connect before erroring, nil represents an indefinite wait
	Timeout *time.Duration
	// Logger is the logger to which to write connection errors, nil will result in no errors being logged
	Logger *logr.Logger
}

// ContextPinger is an abstraction for unit testing. *sql.DB satisfies this interface and should be used when calling WaitForDBConnectable
type ContextPinger interface {
	PingContext(context.Context) error
}

// WaitForDBConnectable waits for the given ContextPinger (satisfied by *sql.DB) to respond without error. See DBWaitOpts for configuration options
func WaitForDBConnectable(db ContextPinger, opts DBWaitOpts) error {
	ctx := context.Background()
	cancel := func() {}
	if opts.Timeout != nil {
		ctx, cancel = context.WithTimeout(context.Background(), *opts.Timeout)
	}
	defer cancel()

	if err := retry.Fibonacci(ctx, 1*time.Second, func(ctx context.Context) error {
		err := db.PingContext(ctx)
		if err != nil {
			if opts.Logger != nil {
				opts.Logger.Error(err, "error connecting to database")
			}
			return retry.RetryableError(err)
		}
		return nil
	}); err != nil {
		return fmt.Errorf("unable to wait for db connectable: %w", err)
	}

	return nil
}
