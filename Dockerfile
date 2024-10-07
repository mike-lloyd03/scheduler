##### Backend Builder #####
FROM golang:1.23 AS backend_build

WORKDIR /app

COPY backend/go.* .
RUN go mod download

COPY backend/main.go .
COPY backend/migrations ./migrations
COPY backend/tests ./tests

RUN CGO_ENABLED=0 GOOS=linux go build -o backend

##### Frontend Builder #####
FROM node:22-alpine AS frontend_build

RUN npm install -g pnpm

WORKDIR /app

COPY frontend/pnpm-lock.yaml frontend/package.json ./

RUN pnpm install

COPY frontend/ ./

RUN pnpm run build

##### Final #####
FROM node:22-alpine

COPY --from=backend_build /app/backend /backend/backend
COPY --from=frontend_build /app /frontend
COPY run_app /run_app

ENV ORIGIN=http://localhost:3000

EXPOSE 8090
EXPOSE 3000

VOLUME /backend/pb_data

CMD ["/run_app"]
