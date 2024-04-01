gen-proto:
	protoc \
	-I proto \
	--grpc-gateway_out ./proto/generated/ \
  --grpc-gateway_opt paths=source_relative \
  --grpc-gateway_opt generate_unbound_methods=true \
	--openapiv2_out ./proto/swagger/ \
	--go_out=./proto/ \
	--go-grpc_out=require_unimplemented_servers=false:proto \
	proto/*.proto 

remove-gen-proto:
	rm -rf proto/generated && \
	cd proto && \
	mkdir generated && \
	cd generated && \
	touch .gitkeep && \
	cd .. && \
	rm -rf swagger && \
	mkdir swagger && \
	cd swagger && \
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

up:
	docker compose -f ./dev.docker-compose.yml up

down:
	docker compose -f ./dev.docker-compose.yml down

.PHONY: gen-proto remove-gen-proto gen-types remove-gen-types up down
