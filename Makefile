build:
	protoc -I. --go_out=plugins=grpc:. proto/vessel.proto
	docker build -t vessel-service .
run:
	docker service create --name vessel-service --network consignment vessel-service