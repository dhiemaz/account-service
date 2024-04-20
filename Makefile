.PHONY: clean protoc test grpc build

clean:
	@echo "--- cleanup all build and generated files ---"
	@rm -rf infrastructures/server/grpc

protoc: clean
	@echo "--- preparing proto output directories ---"
	@mkdir -p infrastructures/server/grpc/proto/access_svc
	@mkdir -p infrastructures/server/grpc/proto/account_svc
	@mkdir -p infrastructures/server/grpc/proto/office_svc
	@mkdir -p infrastructures/server/grpc/proto/role_svc

	@echo "--- compiling all proto files ---"
	@protoc -I./shared/proto/access_svc        \
       --go_out=./infrastructures/server/grpc/proto/access_svc      \
       --go-grpc_out=./infrastructures/server/grpc/proto/access_svc \
       shared/proto/access_svc/*.proto
	@protoc -I./shared/proto/account_svc        \
       --go_out=./infrastructures/server/grpc/proto/account_svc      \
       --go-grpc_out=./infrastructures/server/grpc/proto/account_svc \
       shared/proto/account_svc/*.proto
	@protoc -I./shared/proto/office_svc        \
       --go_out=./infrastructures/server/grpc/proto/office_svc      \
       --go-grpc_out=./infrastructures/server/grpc/proto/office_svc \
       shared/proto/office_svc/*.proto
	@protoc -I./shared/proto/role_svc        \
       --go_out=./infrastructures/server/grpc/proto/role_svc      \
       --go-grpc_out=./infrastructures/server/grpc/proto/role_svc \
       shared/proto/role_svc/*.proto


