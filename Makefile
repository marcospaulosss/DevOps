.PHONY: proto stress

proto:
	rm proto/*.pb.go
	protoc proto/common.proto --go_out=plugins=grpc:${PWD}
	protoc proto/elearning.proto --proto_path=./proto --go_out=plugins=grpc:${PWD}/proto
	protoc proto/accounts.proto --proto_path=./proto --go_out=plugins=grpc:${PWD}/proto
	protoc proto/ecommerce.proto --proto_path=./proto --go_out=plugins=grpc:${PWD}/proto

tests: proto
	@cd apps/severino && make test && cd ../elearning && make test && cd ../accounts && make test

test: tests

stress:
	docker run -it --rm --net backend_default -e SEVERINO_URL=http://severino:4001 --link severino -v ${PWD}/stress:/stress loadimpact/k6 run --vus 150 --duration 20s /stress/elearning.js
	# docker-compose up --exit-code-from stress --force-recreate

