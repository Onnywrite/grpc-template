GEN=./gen
PROTOPATH=./proto
PROTOS=**/*.proto

protoc:
	protoc --go_out=${GEN} --go_opt=paths=source_relative \
    --go-grpc_out=${GEN} --go-grpc_opt=paths=source_relative \
	--proto_path=${PROTOPATH} ${PROTOS}
cert:
	./certs/cert-gen.sh