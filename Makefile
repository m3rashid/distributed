gen-proto:
	protoc \
	-I . --grpc-gateway_out ./proto/generated \
  --grpc-gateway_opt paths=source_relative \
  --grpc-gateway_opt generate_unbound_methods=true \
	--go_out=. \
	--go-grpc_out=require_unimplemented_servers=false:proto \
	proto/source/*.proto 

remove-gen-proto:
	rm -rf proto/generated && \
	cd proto && \
	mkdir generated && \
	cd generated && \
	touch .gitkeep && \
	cd -

gen-types:
	cd types && go run main.go

remove-gen-types:
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

.PHONY: gen-proto remove-gen-proto gen-types remove-gen-types
