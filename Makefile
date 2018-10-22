default: build

init:
	virtualenv -p python3.7 .venv;
	source ./venv/bin/activate;
	pip install -r requirements.txt;

gen:
	protoc -I proto/ proto/apthunter.proto --go_out=plugins=grpc:apartment-hunter/pkg/pb
	python -m grpc_tools.protoc -I./proto --python_out=./crawler --grpc_python_out=./crawler ./proto/apthunter.proto

build:
	GOARCH=amd64 GOOS=linux CGO_ENABLED=0 go build -a -o ./apartment-hunter/cmd/writer/bin/writer ./apartment-hunter/cmd/writer

docker-up: build
	docker-compose up --build

docker-remove:
	docker-compose stop
	docker-compose down --volumes --remove-orphans
