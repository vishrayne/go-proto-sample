# go-proto-sample
[![baby-gopher](https://raw.githubusercontent.com/drnic/babygopher-site/gh-pages/images/babygopher-logo-small.png)](http://www.babygopher.org)

Just curious on how protocol buffers work based on [this](https://developers.google.com/protocol-buffers/docs/gotutorial) tutorial

# Prerequisite

1. Install protoc compiler (I've used 3.5.1) from [here](https://developers.google.com/protocol-buffers/docs/downloads).
2. Get Protobuf runtime for golang. \
`go get -u github.com/golang/protobuf/protoc-gen-go`

# Setup

1. Clone this repo
2. A basic CLI version, but not complete \
`$ go install github.com/vishrayne/go-proto-sample/cmd/go-proto-cli`
3. A mock client and server available by running \
`$ go install github.com/vishrayne/go-proto-sample/cmd/go-proto-server` \
`$ go install github.com/vishrayne/go-proto-sample/cmd/go-proto-mock-client`

