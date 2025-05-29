# Es la imagen base de golang, mas liviana que la oficial
FROM golang:1.23-alpine AS build 

WORKDIR /app

RUN apk add make

COPY . .

# RUN go mod download

RUN make build

FROM alpine:latest

WORKDIR /root/

COPY --from=build /app/main .

CMD ["./main"]