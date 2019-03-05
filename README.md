go-kmip
=======

[![Build Status](https://travis-ci.org/smira/go-kmip.svg?branch=master)](https://travis-ci.org/smira/go-kmip)
[![codecov](https://codecov.io/gh/smira/go-kmip/branch/master/graph/badge.svg)](https://codecov.io/gh/smira/go-kmip)
[![Documentation](https://godoc.org/github.com/smira/go-kmip?status.svg)](http://godoc.org/github.com/smira/go-kmip)
[![Go Report Card](https://goreportcard.com/badge/github.com/smira/go-kmip)](https://goreportcard.com/report/github.com/smira/go-kmip)

go-kmip implements subset of [KMIP 1.4](http://docs.oasis-open.org/kmip/spec/v1.4/os/kmip-spec-v1.4-os.html) protocol.

Basic encoding/decoding is fully implemented, as well as the basic client/server implementations. Other operations and fields
could be implemented by adding required Go structures with KMIP tags.

Note
----

This is WIP at the moment.

License
-------

This code is licensed under [MPL 2.0](https://www.mozilla.org/en-US/MPL/2.0/).
