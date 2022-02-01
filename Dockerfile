FROM golang:latest
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN GOARCH=arm go build -o main .
CMD ["/app/main"]

EXPOSE 8080