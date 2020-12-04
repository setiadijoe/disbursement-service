build: 
	@GOOS=linux GOARCH=amd64
	@echo ">> Building TRANSPORT..."
	@go build -o main-app ./cmd
	@echo ">> Finished"

test-case:
	@go test ./usecase/ -coverprofile=./usecase/coverage.out & go tool cover -html=./usecase/coverage.out