FROM golang
WORKDIR /app

COPY go.mod ./
COPY go.sum ./
RUN go mod download
COPY ./ ./
WORKDIR /app/cmd
RUN go build -o member_club

WORKDIR /app
CMD ["/app/cmd/member_club"]
