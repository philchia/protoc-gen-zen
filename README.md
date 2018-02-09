# protoc-gen-zen

## How to

```sh
    go install
	protoc -I . --zen_out="plugins:." ./test/proto/*.proto
```