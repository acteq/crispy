FROM golang:1.13 as builder
WORKDIR /build
COPY src .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags '-w -s' -o server


FROM alpine:latest  
RUN apk --no-cache add ca-certificates
WORKDIR /app
COPY --from=builder /build  .
CMD ["/app/server"] 
