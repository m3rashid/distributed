.PHONY gen-proto:
	protoc \
	--go_out=proto \
	--go-grpc_out=require_unimplemented_servers=false:proto \
	proto/permissions/group.proto proto/permissions/pagination.proto proto/permissions/permission.proto

.PHONY remove-gen-proto:
	rm -rf proto/generated/permissions

.PHONY gen-types:
	cd types && go run main.go

.PHONY remove-gen-types:
	cd types && \
	rm -rf generated && \
	mkdir generated && \
	cd generated && \
	touch .gitkeep && \
	cd .. && \
	rm -rf backup && \
	mkdir backup && \
	cd backup && \
	touch .gitkeep && \
	cd -
