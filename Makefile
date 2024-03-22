gen-proto:
	protoc --go_out=proto --go-grpc_out=require_unimplemented_servers=false:proto proto/authz/group.proto proto/authz/pagination.proto proto/authz/permission.proto

remove-gen-proto:
	rm -rf proto/generated/authz
