FROM golang:1.20

RUN apt update && apt upgrade

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN go build -o main .

CMD ["./main", "-task=runcrons"]
