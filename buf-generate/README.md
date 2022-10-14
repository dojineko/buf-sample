# buf generate

This is an example of applying a setting to buf that produces the same result as a build by protoc command line using protoc-gen-go and protoc-gen-go-grpc and protoc-gen-doc.

- protoc-gen-go & protoc-gen-grpc
  - https://grpc.io/docs/languages/go/quickstart/
- protoc-gen-doc
  - https://github.com/pseudomuto/protoc-gen-doc

## Command line example using protoc

Execute the following command line under the protobuf directory.

**HINT:** If you use protoc, you must uncomment the `option go_package ...` line in the proto file must be uncommented.

```sh
# move to workspace
cd protobuf

# create output dir before generate
mkdir -p gen/{docs/protobuf,go}

# generate using protoc
protoc -I proto \
  --doc_out=gen/docs/protobuf \
  --doc_opt=markdown,README.md,source_relative \
  --go_out=gen/go \
  --go_opt=paths=source_relative \
  --go-grpc_out=gen/go \
  --go-grpc_opt=paths=source_relative \
  $(find . -name '*.proto')
```

## Command line example using buf

Just execute the following command line under the protobuf directory.

```sh
# move to workspace
cd protobuf

# just do it
buf generate
```
