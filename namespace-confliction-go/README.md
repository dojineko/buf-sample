# namespace-confliction-go

Reproductions and examples of fixes for Protocol Buffer namespace conflicts that occur when protoc-gen-go is used are exhibited.

**NOTE:** buf is used for simplicity, but in principle the same error occurs with protoc.

ref: https://developers.google.com/protocol-buffers/docs/reference/go/faq#namespace-conflict

## ⛔ Reproduction version

Reproduce namespace conflicts when using proto-gen-go.

```sh
cd reproduction
make dev

# You will probably get a message like this.
go run main.go
panic: proto: file "dummy.proto" is already registered
        previously from: ".../reproduction/baz/protobuf/gen/go"
        currently from:  ".../reproduction/foo/protobuf/gen/go"
See https://developers.google.com/protocol-buffers/docs/reference/go/faq#namespace-conflict
```

## ✅ Fixed version

This version changes the structure of the foo and baz proto directories to a structure consisting of package name and version number to eliminate namespace conflicts.

```sh
cd reproduction
make dev

# You will probably get a message like this.
go run main.go
gRPC Server started! Press Ctrl+C to exit.
```

The general changes are as follows

```sh
# before (reproduction)
./baz
└── protobuf
    └── proto
        └── dummy.proto (namespace: dummy.proto)
./foo
└── protobuf
    └── proto
        └── dummy.proto (namespace: dummy.proto)

# after (fixed)
./baz
└── protobuf
    └── proto
        └── baz
            └── v1
                └── dummy.proto (namespace: baz/v1/dummy.proto)
./foo
└── protobuf
    └── proto
        └── foo
            └── v1
                └── dummy.proto (namespace: foo/v1/dummy.proto)
```
