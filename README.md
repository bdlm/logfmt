# logfmt

<p align="center">
	<a href="https://github.com/bdlm/logfmt/blob/master/LICENSE"><img src="https://img.shields.io/github/license/bdlm/logfmt.svg" alt="MIT License"></a>
	<a href="https://github.com/mkenney/software-guides/blob/master/STABILITY-BADGES.md#mature"><img src="https://img.shields.io/badge/stability-mature-008000.svg" alt="Mature"></a>
	<a href="https://travis-ci.org/bdlm/logfmt"><img src="https://travis-ci.org/bdlm/logfmt.svg?branch=master" alt="Build status"></a>
	<a href="https://codecov.io/gh/bdlm/logfmt"><img src="https://img.shields.io/codecov/c/github/bdlm/logfmt/master.svg" alt="Coverage status"></a>
	<a href="https://goreportcard.com/report/github.com/bdlm/logfmt"><img src="https://goreportcard.com/badge/github.com/bdlm/logfmt" alt="Go Report Card"></a>
	<a href="https://github.com/bdlm/logfmt/issues"><img src="https://img.shields.io/github/issues-raw/bdlm/logfmt.svg" alt="Github issues"></a>
	<a href="https://github.com/bdlm/logfmt/pulls"><img src="https://img.shields.io/github/issues-pr/bdlm/logfmt.svg" alt="Github pull requests"></a>
	<a href="https://godoc.org/github.com/bdlm/logfmt"><img src="https://godoc.org/github.com/bdlm/logfmt?status.svg" alt="GoDoc"></a>
</p>

A simple https://github.com/sirupsen/logrus log formatter. Includes

* [iso 8601](https://en.wikipedia.org/wiki/ISO_8601) formatted date string - `"2006-01-02T15:04:05.000Z07:00"`
* OS hostname
* Log level
* Log message
* Additional fields
* Caller - "file:line func"

## Usage

```go
import (
	log "github.com/sirupsen/logrus"
	logfmt "github.com/bdlm/logfmt"
)
```

#### Text format

Set the formatter:
```go
log.SetFormatter(&logfmt.TextFormat{})
```

Produces:
```js
time="2018-04-16T05:14:07.559Z" host="k8s-proxy-688fb8b57d-4rzt4" level="info" msg="starting kubernetes proxy" port="80" caller="proxy.go:252 github.com/bdlm/logfmt/pkg/proxy.(*Proxy).Start"
```

#### JSON format

Set the formatter
```go
log.SetFormatter(&logfmt.JSONFormat{})
```

Produces:
```json
{"time":"2018-04-16T06:23:37.133Z","level":"info","host":"k8s-proxy-7b77bfd8bd-7xcvn","msg":"starting kubernetes proxy","data":{"port":"80"},"caller":"proxy.go:258 github.com/bdlm/logfmt/pkg/proxy.(*Proxy).Start"}
```
