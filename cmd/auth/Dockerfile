FROM golang:1.24-alpine

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

WORKDIR /app/cmd/auth

RUN go build -o auth

CMD ["/app/cmd/auth/auth"]
