FROM ubuntu:23.04 AS BUILDER

RUN apt update && apt upgrade -y
RUN apt install -y golang sqlite3 ca-certificates
WORKDIR /app
COPY . .
RUN go mod download
RUN go build main.go
CMD ["./main"]
