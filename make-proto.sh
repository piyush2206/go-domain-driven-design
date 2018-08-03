#! /bin/bash

# compile proto files to go extension for grpc server
protoc \
-I /usr/local/include \
-I $GOPATH/src \
-I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
-I admin \
--go_out=plugins=grpc:admin \
admin/admin.proto

protoc \
-I /usr/local/include \
-I $GOPATH/src \
-I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
-I admin \
-I report \
--go_out=plugins=grpc:report \
report/report.proto

#############################################################################
# compile proto files to go gateway extension for grpc reverse proxy server
protoc \
-I /usr/local/include \
-I $GOPATH/src \
-I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
-I admin \
--grpc-gateway_out=logtostderr=true:admin \
admin/admin.proto

protoc \
-I /usr/local/include \
-I $GOPATH/src \
-I $GOPATH/src/github.com/grpc-ecosystem/grpc-gateway/third_party/googleapis \
-I admin \
-I report \
--grpc-gateway_out=logtostderr=true:report \
report/report.proto