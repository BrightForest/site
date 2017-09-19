FROM golang:latest
RUN mkdir /app
ADD . /app/
RUN useradd go
RUN chown -R go: /app
USER go
WORKDIR /app
RUN go build -o main .
CMD ["/app/main"]
EXPOSE 8080