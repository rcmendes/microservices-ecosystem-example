FROM golang:1.17 as base

WORKDIR /build

COPY go.mod go.sum /build/
RUN go mod download -x

COPY . .

RUN go get -d -v .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app .


FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=base /build/app ./

EXPOSE 9001

CMD ["./app"]