# Simple Fix
[![Generic badge](https://img.shields.io/github/v/release/b2broker/simplefix-go.svg?style=for-the-badge)](https://github.com/b2broker/simplefix-go/releases/latest)
[![Generic badge](https://goreportcard.com/badge/github.com/b2broker/simplefix-go?style=for-the-badge)](https://goreportcard.com/report/github.com/b2broker/simplefix-go)
[![Generic badge](https://img.shields.io/github/stars/b2broker/simplefix-go?style=for-the-badge&logo=GitHub)](https://github.com/b2broker/simplefix-go/stargazers)
[![Generic badge](https://img.shields.io/badge/Go->=1.16-blue.svg?style=for-the-badge&logo=go)](https://golang.org/doc/go1.16)
[![Generic badge](https://img.shields.io/badge/semver-semantic_release-blue.svg?style=for-the-badge&logo=semantic-release)](https://github.com/go-semantic-release/semantic-release)

:warning: This is beta-version. It may contain errors.

<details>
<summary>Table of content</summary>

## Table of contents
- [What is FIX](#what-is-fix-api)
- [Installation](#installation)
- [Getting started](#getting-started)
- [Generator](#generator)
- [Client (Initiator)](#starting-as-client)
- [Server (Acceptor)](#starting-as-server)
- [Features](#features)

</details>

## What is FIX API?

[The exhaustive material of FIX API](https://www.onixs.biz/fix-dictionary.html) 

## Installation

Download SimpleFix-Go library
```sh
$ go get -u github.com/b2broker/simplefix-go
```

Install *generator* if you want to use your own scheme.
```sh
$ cd $GOPATH/src/github.com/b2broker/simplefix-go && go install ./...
```

## Generator

Generator create message structures, tag and msg type constants and methods for work with any FIX API.

You can see basic result of Generator working at [tests directory](https://github.com/b2broker/simplefix-go/tree/master/tests/fix44). It is short version of FIX 4.4 generated from scheme 
located in ./source directory.

### Simple generating example

```sh
fixgen -o=./fix44 -s=./source/fix44.xml -t=./source/types.xml
```

So, now you have your library code in ./fix44 directory.

### Parameters

`-o` – output directory

`-s` – path to main XML-scheme, you can see examples in ./source directory

`-t` – path to XML file with types mapping. This file provides generator information about type casting.
FIX protocol has a lot of types, but in Go we should have small set of types for use FIX API.

According to parameters we must have two XML files for make our own library. You can use existing files or modify them 
as you want.

## Getting started

### Session Options

This is message builder, field and message tags for session pipelines. This structure will be generated by fixgen 
command very soon.

```go
// fixgen is your generated fix package

var sessionOpts = session.Opts{
	MessageBuilders: session.MessageBuilders{
		HeaderBuilder:        fixgen.Header{}.New(),
		TrailerBuilder:       fixgen.Trailer{}.New(),
		LogonBuilder:         fixgen.Logon{}.New(),
		LogoutBuilder:        fixgen.Logout{}.New(),
		RejectBuilder:        fixgen.Reject{}.New(),
		HeartbeatBuilder:     fixgen.Heartbeat{}.New(),
		TestRequestBuilder:   fixgen.TestRequest{}.New(),
		ResendRequestBuilder: fixgen.ResendRequest{}.New(),
	},
	Tags: &messages.Tags{
		MsgType:         mustConvToInt(fixgen.FieldMsgType),
		MsgSeqNum:       mustConvToInt(fixgen.FieldMsgSeqNum),
		HeartBtInt:      mustConvToInt(fixgen.FieldHeartBtInt),
		EncryptedMethod: mustConvToInt(fixgen.FieldEncryptMethod),
	},
	AllowedEncryptedMethods: map[string]struct{}{
		fixgen.EnumEncryptMethodNoneother: {},
	},
	SessionErrorCodes: &messages.SessionErrorCodes{
		InvalidTagNumber:            mustConvToInt(fixgen.EnumSessionRejectReasonInvalidtagnumber),
		RequiredTagMissing:          mustConvToInt(fixgen.EnumSessionRejectReasonRequiredtagmissing),
		TagNotDefinedForMessageType: mustConvToInt(fixgen.EnumSessionRejectReasonTagNotDefinedForThisMessageType),
		UndefinedTag:                mustConvToInt(fixgen.EnumSessionRejectReasonUndefinedtag),
		TagSpecialWithoutValue:      mustConvToInt(fixgen.EnumSessionRejectReasonTagspecifiedwithoutavalue),
		IncorrectValue:              mustConvToInt(fixgen.EnumSessionRejectReasonValueisincorrectoutofrangeforthistag),
		IncorrectDataFormatValue:    mustConvToInt(fixgen.EnumSessionRejectReasonIncorrectdataformatforvalue),
		DecryptionProblem:           mustConvToInt(fixgen.EnumSessionRejectReasonDecryptionproblem),
		SignatureProblem:            mustConvToInt(fixgen.EnumSessionRejectReasonSignatureproblem),
		CompIDProblem:               mustConvToInt(fixgen.EnumSessionRejectReasonCompidproblem),
		Other:                       mustConvToInt(fixgen.EnumSessionRejectReasonOther),
	},
}
```

### Starting as client

*Initiator* is a FIX API client, which connect to an existing server by.

You can see [example here](https://github.com/b2broker/simplefix-go/blob/master/examples/initiator/main.go).

### Starting as server

*Acceptor* is a listener. It accepts and handles client connections. 
According to the FIX protocol acceptor could be provider or receiver of data.
It means acceptor can send requests to the clients and read data streams of them. 

You can see [example here](https://github.com/b2broker/simplefix-go/blob/master/examples/acceptor/main.go).

## Features

- [x] Add custom fields to protocol
- [x] Add custom messages to protocol
- [ ] Add custom types to protocol
- [x] Features for session pipelines
- [x] Acceptor (server)
- [x] Initiator (client)
- [ ] Validation of incoming messages
- [x] Validation of outgoing messages
- [ ] Mock server with testing data 
- [ ] Anything you want Just with submit a new issue
