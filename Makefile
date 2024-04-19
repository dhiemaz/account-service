.PHONY: clean protoc test grpc build

clean:
	@echo "--- cleanup all build and generated files ---"
	@rm -rf infrastructure/server/grpc/proto

protoc: clean
	@echo "--- preparing proto output directories ---"
	@mkdir -p infrastructures/server/grpc/proto/access_svc
	@mkdir -p infrastructures/server/grpc/proto/account_svc
	@mkdir -p infrastructures/server/grpc/proto/office_svc
	@mkdir -p infrastructures/server/grpc/proto/role_svc

	@echo "--- compiling all proto files ---"
	@cd ./shared/proto/access_svc && protoc -I. --go-grpc_out=../../../infrastructures/server/grpc/proto/access_svc *.proto
	@cd ./shared/proto/account_svc && protoc -I. --go-grpc_out=../../../infrastructures/server/grpc/proto/account_svc *.proto
	@cd ./shared/proto/office_svc && protoc -I. --go-grpc_out=../../../infrastructures/server/grpc/proto/office_svc *.proto
	@cd ./shared/proto/role_svc && protoc -I. --go-grpc_out=../../../infrastructures/server/grpc/proto/role_svc *.proto


