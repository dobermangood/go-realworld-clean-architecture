test:
	go test -cover -covermode=atomic -race -p 10 `go list ./...`

generate_swagger:
	swag init
