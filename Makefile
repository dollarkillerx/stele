Generate:
	@echo 'Build GRPC'
	protoc -I rpc/generate/ rpc/generate/*.proto --go_out=plugins=grpc:rpc/generate/.