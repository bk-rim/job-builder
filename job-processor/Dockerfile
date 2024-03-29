FROM golang:1.18-alpine AS build

WORKDIR /app
COPY . .

RUN go mod download
RUN go build -o /app/job-processor

FROM alpine:latest

WORKDIR /app
COPY --from=build /app/job-processor .

CMD ["./job-processor"]