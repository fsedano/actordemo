FROM golang:1.23 AS build
WORKDIR /app
COPY go.* /app/
RUN go mod download
COPY cmd/ /app/cmd/
COPY api/ /app/api/
RUN CGO_ENABLED=0 go build -o pub ./cmd/pub 
RUN CGO_ENABLED=0 go build -o sub ./cmd/sub 
FROM scratch
WORKDIR /app/
COPY --from=build /app/pub /app/
COPY --from=build /app/sub /app/
CMD ["/app/sub"]