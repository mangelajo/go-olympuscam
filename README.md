# go-olympuscam

go-olympuscam is a library designed to interface with the olympus
O-MD cameras, including the E-M10 model, using the olympus image share
protocol.

We provide a command line utility `olycam`, but the camera directory
contains a library which can be imported from any go software.


## Install

```bash
go get github.com/mangelajo/go-olympuscam/olycam
```


## Use as go module in your software

```
export GO111MODULE=on
go get github.com/mangelajo/go-olympuscam
```

Then see the [examples](examples) directory.
