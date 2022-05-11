FROM alpine:3.15 AS runtime
RUN mkdir -p /app
WORKDIR /app
ADD ./src/sample-api-docker /app/sample-api-docker

CMD ["/app/sample-api-docker"]
