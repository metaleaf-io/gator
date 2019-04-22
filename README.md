[![Build Status](https://travis-ci.org/metaleaf-io/gator.svg)](https://travis-ci.org/metaleaf-io/gator)
[![GoDoc](https://godoc.org/github.com/metaleaf-io/gator/github?status.svg)](https://godoc.org/github.com/metaleaf-io/gator)
[![License](https://img.shields.io/badge/license-Apache--2.0-blue.svg)](https://www.apache.org/licenses/LICENSE-2.0)
[![Go version](https://img.shields.io/badge/go-~%3E1.12.0-green.svg)](https://golang.org/doc/devel/release.html#go1.12)

# gator

Gator is a log aggregator that accepts log entries via UDP. These messages
are stored in a PostgreSQL database ordered by time, and the messages are
indexed for full-text searching.

## Features

* Very small and lightweight
* Designed to run as a container

## Usage

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
