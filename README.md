# go-logger-facade

Tooling and facade around https://github.com/uber-go/zap. 
I needed to have my applications use a facade to allow swapping out loggers amd provide a backup for when it goes wrong.
After swapping logging libs once I wanted a facade to make it easy in the future that supported both structured & unstructured.
Functions for both unstructured and structured logging, panic handling and special handling for cancelled contexts 
(don't log errors on normal network disconnects).
Unstructured functions are marked Deprecated.

## Context Injection

```go
package main

import (
	"context"
	"github.com/davidwartell/go-logger-facade/logger"
	"os"
)

func main() {
	logI := logger.NewLogger()
	logI.StartTask(logger.WithProductNameShort("example"))
	defer func() {
		// make sure the log buffers are flushed/synced
		_ = logI.StopTask
	}()
	isDevEnv := true
	if isDevEnv {
		logger.Instance().SetConsoleLogging(true)
	}

	ctx, cancel := context.WithCancel(logger.WithLogger(context.Background(), logI))
	defer cancel()

	doSomeWork(ctx)

	os.Exit(0)
}

func doSomeWork(ctx context.Context) {
	// add some fields to the logger that get used on every log after this using the context
	ctx = logger.WithFields(ctx, logger.String("some_field", "some_value"))
	doSomeWorkForReal(ctx)
}

func doSomeWorkForReal(ctx context.Context) {
	if err := someFunctionThatMightFail(); err != nil {
		logger.OfMust(ctx).Error("some message",
			logger.Error(err),
			logger.Int("another_field", 1),
		)
	}
}

func someFunctionThatMightFail() error {
	return nil
}
```

## Contributing

Happy to accept PRs.

# Author

**davidwartell**

* <http://github.com/davidwartell>
* <http://linkedin.com/in/wartell>
