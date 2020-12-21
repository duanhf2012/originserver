protoc.exe --go_out=./proto/rpc ./proto/rpcproto/*.proto
protoc.exe --go_out=./proto/msg ./proto/msgproto/*.proto
PAUSE