build-NotificationLambda:
	GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -o bootstrap main.go
	mv bootstrap $(ARTIFACTS_DIR)/
