
```bash
conda create -n myenv python=3.12
conda activate myenv
```



### gRPC
```bash
python -m grpc_tools.protoc -I internal/protobuf/protos/ --python_out=internal/protobuf/ --pyi_out=internal/protobuf/ --grpc_python_out=./internal/protobuf  internal/protobuf/protos/emotion.proto
```