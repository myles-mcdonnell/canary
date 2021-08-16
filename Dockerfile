FROM alpine
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*
RUN update-ca-certificates
WORKDIR /app
COPY ./svc ./svc
EXPOSE 8080
CMD ["./svc", "--port=8080", "--host=0.0.0.0"]