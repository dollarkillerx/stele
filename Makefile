Generate:
	@echo 'Build GRPC'
	protoc -I rpc/generate/ rpc/generate/*.proto --go_out=plugins=grpc:rpc/generate/.

SSLKey:
	@echo 'SSLKey'
	openssl genrsa -out cert/server.key 2048
	openssl ecparam -genkey -name secp384r1 -out cert/serveryek
	openssl req -new -x509 -sha256 -key cert/serveryek -out cert/servermpe -days 3650