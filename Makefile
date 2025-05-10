PROTO_DIR=./proto
OUT_TASK_DIR=$(PROTO_DIR)/task
OUT_USER_DIR=$(PROTO_DIR)/user

PROTOC=protoc

all: task user

task:
	$(PROTOC) \
		--proto_path=$(PROTO_DIR) \
		--go_out=paths=source_relative:$(OUT_TASK_DIR) \
		--go-grpc_out=paths=source_relative:$(OUT_TASK_DIR) \
		$(PROTO_DIR)/task.proto

user:
	$(PROTOC) \
		--proto_path=$(PROTO_DIR) \
		--go_out=paths=source_relative:$(OUT_USER_DIR) \
		--go-grpc_out=paths=source_relative:$(OUT_USER_DIR) \
		$(PROTO_DIR)/user.proto


clean:
	rm -f $(OUT_TASK_DIR)/*.pb.go $(OUT_USER_DIR)/*.pb.go