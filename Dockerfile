FROM golang:alpine AS builder

ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

WORKDIR /build

COPY go.mod ./

RUN go mod download

COPY . .

RUN go build -o main .

WORKDIR /dist

RUN cp /build/main .
RUN cp -r /build/static .

FROM scratch
COPY --from=builder /dist/main /
COPY --from=builder /dist/static/ /static
ENTRYPOINT [ "/main" ]