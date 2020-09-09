FROM golang:1.15.1-alpine as build
EXPOSE 4000

COPY . .
RUN CGO_ENABLED=0 go build -o /bin/consumer main.go

FROM scratch
COPY --from=build /bin/consumer /consumer
COPY /static /static

ENTRYPOINT ["/consumer"]
