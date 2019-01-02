v?=latest

local:
	export JAEGER_SERVICE_NAME=bbb && go run cmd/app/main.go
start:
	bash scripts/start.sh
stop:
	bash scripts/stop.sh
docker:
	docker build -f build/Dockerfile -t todo/auth:$(v) .
	docker tag todo/auth:$(v) dreambo8563docker/todo-auth:$(v)
	docker push dreambo8563docker/todo-auth:$(v)
rpc:
	protoc -I internal/adapter/http/rpc/auth/ internal/adapter/http/rpc/auth/auth.proto --go_out=plugins=grpc:internal/adapter/http/rpc/auth
clean:
	rm -f auth-log cmd/app/*log app