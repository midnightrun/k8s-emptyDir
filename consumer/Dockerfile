FROM golang:1.15.1-alpine as build
EXPOSE 8080

COPY . .
RUN CGO_ENABLED=0 go build -o /bin/consumer main.go

FROM scratch
COPY /assets /assets
COPY --from=build /bin/consumer /consumer

ENTRYPOINT ["/consumer"]
