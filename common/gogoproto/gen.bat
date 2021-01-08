protoc.exe --plugin=protoc-gen-go=.\protoc-gen-gogofast.exe --go_out=.\gogorpc\ rpcproto\*.proto
protoc.exe --plugin=protoc-gen-go=.\protoc-gen-gogofast.exe --go_out=.\gogomsg\ msgproto\*.proto
PAUSE