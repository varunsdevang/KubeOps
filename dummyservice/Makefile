PORT = 3002

run:
	uvicorn main:app --port $(PORT)

buildClient:
	python -m grpc_tools.protoc \
    -I ../ratelimiterapi/proto \
    --python_out=. \
    --grpc_python_out=. \
    ratelimit.proto

dockerbuild:
	docker build -t dummyservice:latest .

dockerrun:
	docker run -p 3002:3002 dummyservice:latest