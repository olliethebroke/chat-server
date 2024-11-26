LOCAL_BIN = "$(CURDIR)/bin"

install-deps:
	@echo "Installing dependencies to $(LOCAL_BIN)..."
	GOBIN = $(LOCAL_BIN) go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
	GOBIN = $(LOCAL_BIN) go install -mod=mod google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

get-deps:
	go get -u google.golang.org/protobuf/cmd/protoc-gen-go
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc

generate:
	make generate-chat-api

#proto_path - путь к proto файлу относительно makefile
#go_out - путь, куда положить скомпилированные в go файлы для используемых Protocol Buffers
#go_opt - говорит компилятору генерировать Go-код с относительными путями, основываясь на местоположении источника `.proto` файла.
#plugin=protoc-gen-go=bin/protoc-gen-go - указывает на путь к плагину для генерации Go-кода. Он может находиться в любой указанной директории.
#go-grpc_out - путь, куда положить скомпилированные в go файлы для gRPC
#api/note_v1/note.proto - путь к самому файлу `.proto`, который будет обработан
generate-chat-api:
	protoc --proto_path=api/chat_v1 \
           	--go_out=pkg/chat_v1 \
           	--go_opt=paths=source_relative \
           	--go-grpc_out=pkg/chat_v1 \
           	--go-grpc_opt=paths=source_relative \
           	api/chat_v1/chat.proto