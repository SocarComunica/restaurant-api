FROM golang:1.21

WORKDIR /app

# ENV GIN_MODE release
ENV PORT 8080
ENV HOST 0.0.0.0

COPY . .

RUN go build -o restaurant-api

EXPOSE 8080

CMD ["./restaurant-api"]

