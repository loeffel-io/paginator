# Paginator

[![Build Status](https://cloud.drone.io/api/badges/loeffel-io/paginator/status.svg)](https://cloud.drone.io/loeffel-io/paginator)

- Grpc support
- Go & JavaScript & TypeScript support

## Protobuf

```protobuf
syntax = "proto3";
package protobuf;

import "paginator.proto";

message CustomerInvoiceRequest{
  uint64 id = 1;
  loeffel_io.paginator.Paginator paginator = 2;
}
```

## Go

```
go get -u github.com/loeffel/paginator/go
```

## JavaScript & TypeScript

``` 
yarn loeffel-io/paginator 
```