FROM golang:1.13 AS build

WORKDIR /go/src/github.com/cms-orbits/cms-sao
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -v \
    && chmod ugo+x cms-sao

FROM scratch

COPY --from=build /go/src/github.com/cms-orbits/cms-sao/cms-sao /opt/bin/cms-sao

CMD ["/opt/bin/cms-sao"]