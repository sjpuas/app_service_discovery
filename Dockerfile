FROM golang:latest AS builder
ADD . /app
WORKDIR /app
RUN go get go.mongodb.org/mongo-driver/mongo
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -o /main .

FROM alpine:latest
RUN apk --no-cache add ca-certificates
COPY --from=builder /main ./
RUN chmod +x ./main
ENTRYPOINT ["./main"]
EXPOSE 3030