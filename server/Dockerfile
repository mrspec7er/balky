FROM golang:1.21.6-bookworm

WORKDIR /api

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go install github.com/zzwx/fresh@latest

RUN chmod +x ./start.sh

EXPOSE 8080

CMD ["./start.sh"]