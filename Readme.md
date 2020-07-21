# Slog
### Simple Logging Platform for Go

#### Logging
Slog is a logging package that originated in
[Nan0](github.com/yomiji/nan0). Slog is used to inform the consumer of
the details of operation. The logging mechanism is remotely similar to
some packages in other languages such as Java's slf4j; or at least I'd
like it to be similar to something most people have seen before. That
being said there are some functions that must be discussed, as the
default settings may be more verbose than you need.

* There are four logging levels: ***Info***, ***Warn***,
  ***Fail/Error***, and ***Debug***
* All of the logging levels are enabled by default, to disable them, you
  must set the corresponding logger.
    ```go
      package main
      
      import "github.com/yomiji/slog"
      
      func main() {
      	// set logging levels enabled/disabled
      	// info: true, warn: true, fail: true, debug: false
        slog.ToggleLogging(true, true, true, false)
        
        slog.Info("this is a test") //works, prints
        slog.Debug("not going to print") // won't print
      }
    ```