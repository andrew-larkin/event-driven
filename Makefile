PHONY: generate
generate:
		mkdir -p pkg/order_v1
		protoc  --proto_path vendor.protogen --proto_path api/order_v1 \
				--go_out=pkg/order_v1 --go_opt=paths=source_relative \
				--go-grpc_out=pkg/order_v1 --go-grpc_opt=paths=source_relative \
				--grpc-gateway_out=pkg/order_v1 \
				--grpc-gateway_opt=logtostderr=true \
				--grpc-gateway_opt=paths=source_relative \
				api/order_v1/order.proto


PHONY: vendor-proto
vendor-proto: .vendor-proto

PHONY: .vendor-proto
.vendor-proto:
		@if [ ! -d vendor.protogen/google ]; then \
			git clone https://github.com/googleapis/googleapis vendor.protogen/googleapis &&\
			mkdir -p  vendor.protogen/google/ &&\
			mv vendor.protogen/googleapis/google/api vendor.protogen/google &&\
			rm -rf vendor.protogen/googleapis ;\
		fi