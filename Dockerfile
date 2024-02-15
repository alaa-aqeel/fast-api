FROM alpine:latest

WORKDIR /app

RUN apk add --no-cache go

ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN go install github.com/cosmtrek/air@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

EXPOSE 8080

# CMD ["go", "test", "./test"]
# CMD ["go", "run", "./cmd/api/main.go"]
CMD ["air", "-c", ".air.toml"]
