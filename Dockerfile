FROM golang:1.7.3
WORKDIR /app_code
COPY main.go .
RUN go get ./...
RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
WORKDIR /root/
COPY --from=0 /app_code/app .
CMD "./app" --bind 0.0.0.0:$PORT
