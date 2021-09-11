FROM golang:1.17 as base

WORKDIR /build

COPY . .

RUN go get -d -v .
# RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app .


FROM alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=base /build/app ./

EXPOSE 9000

CMD ["./app"]