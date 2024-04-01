FROM golang:1.22.0-alpine

RUN go install github.com/cosmtrek/air@latest

WORKDIR /app

COPY services/permissions/go.mod .
COPY services/permissions/go.sum .
RUN go mod download

COPY . .

CMD ["air"]
