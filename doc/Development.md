# Development and Testing Guide <br/> Quotes Microservice Client SDK for Golang

This document provides high-level instructions on how to build and test the client SDK.

* [Environment Setup](#setup)
* [Installing](#install)
* [Building](#build)
* [Testing](#test)
* [Contributing](#contrib) 

## <a name="setup"></a> Environment Setup

Pip.Services runtime uses Golang version 1.5 or higher.

## <a name="install"></a> Installing

After your Go workspace is ready you can get microservice source code from the Github repository:
```bash
go get github.com/pip-services/pip-clients-quotes-go
```

The go to the package folder and run:
```bash
go get ./...
```

## <a name="test"></a> Testing

Unit testing in the Golang client SDK doesn't run microservice and relies on external Node.js microservice instance.
So, follow instructions to install and run the microservice at https://github.com/pip-services/pip-services-quotes

Make sure you enable REST API endpoint in the microservice. The default microservice HTTP REST port is 8002.
Then check rest configuration in unit tests to match the microservice port. 

When Node.js microservices is up and running, go to **test/version1** folder of the package and run:
```bash
go test
```

## <a name="contrib"></a> Contributing

Developers interested in contributing should read the following instructions:

- [How to Contribute](http://www.pipservices.org/contribute/)
- [Guidelines](http://www.pipservices.org/contribute/guidelines)
- [Styleguide](http://www.pipservices.org/contribute/styleguide)
- [ChangeLog](CHANGELOG.md)

> Please do **not** ask general questions in an issue. Issues are only to report bugs, request
  enhancements, or request new features. For general questions and discussions, use the
  [Contributors Forum](http://www.pipservices.org/forums/forum/contributors/).

It is important to note that for each release, the [ChangeLog](CHANGELOG.md) is a resource that will
itemize all:

- Bug Fixes
- New Features
- Breaking Changes