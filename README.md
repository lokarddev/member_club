# member_club

You can run project both from source code or docker container
First, you should clone repository source code to your machine, then from 
project root directory run one of the following: 

source - `go run cmd/main.go`
docker - `docker build -t club . && docker run -p 8000:8000 club`

Server listens to localhost:8000 by default
You can execute simple test cases with `go test ./...` command inside project directory