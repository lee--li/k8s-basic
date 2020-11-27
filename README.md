# k8s-basic
# test test
## grpc
### go go get -u google.golang.org/grpc  go get -u github.com/golang/protobuf/protoc-gen-go 
### protoc -I goserver/ goserver/pb/demo.proto --go_out=plugins=grpc:goserver
### python pip install grpcio pip install protobuf pip install grpcio-tools
### python -m grpc_tools.protoc -I pythonclient  --python_out=./pythonclient --grpc_python_out=./pythonclient pythonclient/pb/demo.proto
### nodejs 