FROM golang:1.15-alpine as builder

WORKDIR /gitlab.com/datanest-engine/dce/gin-boilerplate
ADD . /gitlab.com/datanest-engine/dce/gin-boilerplate

ENV CGO_ENABLED=0
RUN go build -mod=vendor -tags=jsoniter -o main

FROM alpine:3.8
RUN apk add --no-cache tzdata ca-certificates 
COPY --from=builder /gitlab.com/datanest-engine/dce/gin-boilerplate/main /app/
COPY --from=builder /gitlab.com/datanest-engine/dce/gin-boilerplate/config.json /app/
WORKDIR /app
CMD ["./main"]