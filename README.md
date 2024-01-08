# go-logger-facade

Tooling and facade around https://github.com/uber-go/zap. 
I needed to have my applications use a facade to allow swapping out loggers amd provide a backup for when it goes wrong.
After swapping logging libs once I wanted a facade to make it easy in the future that supported both structured & unstructured.
Functions for both unstructured and structured logging, panic handling and special handling for cancelled contexts 
(don't log errors on normal network disconnects).
Unstructured functions are marked Deprecated.

## Usage

Example
```go
// start the logging service and configure name of your app used in log file names if you enable logging to files
logger.Instance().StartTask(logger.WithProductNameShort("example"))
defer func() {
    // make sure the log buffers are flushed/synced
    _ = logger.Instance().StopTask
}()
isDevEnv := true
if isDevEnv {
    logger.Instance().SetConsoleLogging(true)
}
```

## Contributing

Happy to accept PRs.

# Author

**davidwartell**

* <http://github.com/davidwartell>
* <http://linkedin.com/in/wartell>
