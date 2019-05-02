FROM golang:1.11-alpine3.9 as builder

WORKDIR /go/src/github.com/IhorBondartsov/microservices-with-ddd/client-api-ms

RUN apk add git dep
COPY ./Gopkg.toml ./Gopkg.lock ./
RUN dep ensure -vendor-only
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -ldflags '-extldflags "-static"'  -o /build/client-api-ms cmd/main.go


FROM alpine:3.9

COPY --from=builder /build/client-api-ms /app/
ENTRYPOINT ["/app/client-api-ms"]