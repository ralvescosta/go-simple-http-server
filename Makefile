run:
	go build main.go && \
	ENVIRONMENT=local SERVICE_NAME=go-simple-http-server ./main
