FROM golang:alpine AS builder

RUN apk add --no-cache build-base wget git

ADD . /app 

WORKDIR /app

RUN go build -o article -ldflags '-w -s -linkmode external -extldflags "-static -fPIC"'

FROM scratch

WORKDIR /app

COPY --from=builder /app/article /app/article 

ENTRYPOINT ["/app/article"]
