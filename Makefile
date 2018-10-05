
init:
	virtualenv -p python3.7 .venv;
	source ./venv/bin/activate;
	pip install -r requirements.txt;

gen:
	protoc -I proto/ proto/apthunter.proto --go_out=plugins=grpc:apartment-hunter/pkg/pb
	python -m grpc_tools.protoc -I./proto --python_out=./crawler --grpc_python_out=./crawler ./proto/apthunter.proto

