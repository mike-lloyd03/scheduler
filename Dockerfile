##### Backend Builder #####
FROM golang:1.23 AS backend_build

WORKDIR /app

COPY backend/go.* .
RUN go mod download

COPY backend/main.go .
COPY backend/migrations ./migrations
COPY backend/tests ./tests

RUN CGO_ENABLED=0 GOOS=linux go build -o backend

##### Final #####
FROM alpine:3.20

COPY --from=backend_build /app/backend /backend/backend

VOLUME /backend/pb_data

CMD ["/backend/backend", "serve", "--http=0.0.0.0:8090"]
