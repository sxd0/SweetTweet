FROM golang:1.24.1-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd/user

RUN go build -o user

CMD ["/app/cmd/user/user"]
