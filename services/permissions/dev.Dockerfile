FROM golang:1.22.0-alpine

RUN go install github.com/cosmtrek/air@latest

WORKDIR /app

COPY proto/ ./proto
RUN cd /app/proto && go mod download && cd ..

COPY types/ ./types
COPY utils/ ./utils
RUN cd /app/utils && go mod download && cd ..

COPY services/permissions/go.mod ./services/permissions/
COPY services/permissions/go.sum ./services/permissions/

RUN cd /app/services/permissions && go mod download && cd ..

COPY go.work .
COPY go.mod .

COPY services/permissions/ ./services/permissions/


WORKDIR /app/services/permissions
