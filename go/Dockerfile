# syntax=docker/dockerfile:1

# ---- Build stage: compile a static binary ----
FROM golang:1.23-alpine AS build
WORKDIR /app
COPY go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /server .

# ---- Runtime stage: minimal image, non-root ----
FROM alpine:3.21
RUN adduser -D appuser
USER appuser
COPY --from=build /server /server
EXPOSE 8090
ENTRYPOINT ["/server"]
