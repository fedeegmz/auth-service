FROM golang:1.25-alpine

WORKDIR /app

COPY . /app/

EXPOSE 1323

CMD ["go", "run", "cmd/api/api.go"]
