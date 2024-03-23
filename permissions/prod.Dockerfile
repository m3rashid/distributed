# Development image
FROM golang:1.22.0-alpine as builder
WORKDIR /app

COPY go.mod .
COPY go.sum .
RUN go mod download

COPY . .

RUN go build -o /bin/permissions ./main.go


# Production image
FROM scratch
COPY --from=builder /bin/permissions /bin/permissions
CMD ["/bin/permissions"]
