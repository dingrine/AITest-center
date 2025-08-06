build:
	@time -p go build -o ./bin/exam-center main.go
	
model:
	@time -p cd internal/chaindata/models && gormodel

certs:
	cd ./scripts && make gen-certs
	cat ./scripts/cert/certs/root.crt > ./src/certs/root.crt
	cat ./scripts/cert/certs/root.key > ./src/certs/root.key
	cat ./scripts/cert/certs/server.crt > ./src/certs/server.crt
	cat ./scripts/cert/certs/server.key > ./src/certs/server.key
	cat ./scripts/cert/certs/client.crt > ./src/certs/client.crt
	cat ./scripts/cert/certs/client.key > ./src/certs/client.key

swag:
	@swag fmt
	@swag init