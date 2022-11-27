# syntax=docker/dockerfile:1

RUN go build -o exercise-app/ cmd/api/main.go

FROM gcr.io/distroless/base-debian10

WORKDIR /cmd

EXPOSE 4040

USER nonroot:nonroot

ENTRYPOINT ["/api/exercise-app"]