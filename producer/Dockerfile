FROM golang:1.15.1-alpine as build

WORKDIR /src/
COPY . .
RUN CGO_ENABLED=0 go build -o /bin/producer main.go

FROM scratch
COPY --from=build /bin/producer /bin/producer

ENTRYPOINT ["/bin/producer"]
