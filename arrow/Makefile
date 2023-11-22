build-amd64:
	docker build . -t otelarrowcol-amd64 --platform linux/amd64

build-arm64:
	docker build . -t otelarrowcol-arm64 --platform linux/arm64

run-amd64:
	docker run -it -v ./config/:/config --entrypoint /otelarrowcol otelarrowcol-amd64 --config=/config/saas-collector.yaml

run-arm64:
	docker run -it -v ./config/:/config --entrypoint /otelarrowcol otelarrowcol-arm64 --config=/config/saas-collector.yaml