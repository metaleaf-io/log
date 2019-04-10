[![Build Status](https://travis-ci.org/metaleaf-io/log.svg)](https://travis-ci.org/metaleaf-io/log)
[![GoDoc](https://godoc.org/github.com/metaleaf-io/log/github?status.svg)](https://godoc.org/github.com/metaleaf-io/log)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue.svg)](https://www.apache.org/licenses/LICENSE-2.0)
[![Go version](https://img.shields.io/badge/go-~%3E1.12.0-green.svg)](https://golang.org/doc/devel/release.html#go1.12)

# log

This extremely simple log library is intended for use with micro service
applications, which have only two logging targets: the console while
developing the application, and a log aggregation service in production.

This library was inspired by:

* [Logging Packages in Golang](https://www.client9.com/logging-packages-in-golang/)
* [logrus](https://github.com/sirupsen/logrus)
* [Benchmarking Logging Libraries for Go](https://github.com/imkira/go-loggers-bench)

## Features

* The `fmt` library is not used to minimize memory allocations.
* Logs to `stdout` by default.
* Can be configured to send messages to a UDP log aggregator.

## Usage

By default, `log` will emit messages on `os.Stderr` which is supposed to be
an unbuffered stream. The destination can be changed with: 

```go
log.SetOutput(os.Stdout)
``` 

To prevent console logging from being visible at all, change the destination
to `/dev/null`:

```go
devnull, err := os.OpenFile(os.DevNull, O_WRONLY, 0666)
if err != null {
	panic(err)
}
defer devnull.Close()
log.SetOutput(devnull)
```

To send output to a UDP log aggregator, just set the address and port of the
service as follows:

```go
log.SetServer("10.10.10.10:8080")
```

Now every log message is formatted in JSON and will be sent to the aggregator:

```go
log.Info("The quick brown fox")
// Output: {"time":1554370662469959000,"name":"main","level":"INFO","message":"The quick brown fox"} 
```

## Contributing

 1.  Fork it
 2.  Create a feature branch (`git checkout -b new-feature`)
 3.  Commit changes (`git commit -am "Added new feature xyz"`)
 4.  Push the branch (`git push origin new-feature`)
 5.  Create a new pull request.

## Maintainers

* [Metaleaf.io](http://github.com/metaleaf-io/)

## License

Copyright 2019 Metaleaf.io

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
