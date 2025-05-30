FROM golang:1.24

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"

RUN go install github.com/spf13/cobra-cli@latest && \
    go install github.com/golang/mock/mockgen@v1.5.0 && \
    go install github.com/spf13/cobra-cli@latest

RUN apt-get update && apt-get upgrade -y && apt-get install --no-install-recommends sqlite3 -y && rm -rf /var/lib/apt/lists/*

RUN usermod -u 1000 www-data
RUN mkdir -p /var/www/.cache
RUN chown -R www-data:www-data /go
RUN chown -R www-data:www-data /var/www/.cache
USER www-data

CMD ["tail", "-f", "/dev/null"]