NAME=terraform-provider-jenkins


build:
	go build -o $(NAME) cmd/provider/main.go