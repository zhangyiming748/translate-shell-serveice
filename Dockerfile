# docker run -dit --name service -p 8192:8192 -v /Users/zen/github/translate-shell-serveice:/app golang:1.24.2-alpine3.21 ash
FROM golang:1.24.2-alpine3.21 AS build
WORKDIR /app
COPY . .
RUN go build -o /usr/local/bin/service main.go

FROM alpine:3.21.3
COPY --from=build /usr/local/bin/service /usr/local/bin/service
RUN apk add translate-shell
EXPOSE 8192
CMD ["service"]
# docker build -t translate-shell-service .
# docker run -d --name service -p 8192:8192 translate-shell-service