gen:
	go install
	#rm ./test/proto/*.go
	rm -f ./test/proto/test.pb.go
	protoc -I . --zen_out="plugins:." ./test/proto/*.proto