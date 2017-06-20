#!/bin/bash
cd `dirname $0`
set -e
export PATH=$PATH:~/grpc/bins/opt/protobuf:~/grpc/bins/opt/

which protoc
which grpc_ruby_plugin

arr=(`find schema -name '*.proto'`)
echo "files:"
for p in "${arr[@]}"
do
  echo "$p"
  protoc -I./schema -I"$PROTOBUF_SRC" --ruby_out=lib --grpc_out=lib --plugin=protoc-gen-grpc=`which grpc_ruby_plugin` "$p"
done

