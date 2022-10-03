FROM golang:1.15 as builder
#WORKDIR $GOPATH/src/github.com/gregsidelinger/pulse-oximeter
WORKDIR /src
COPY . .

RUN make mod amd64 \
  && ls -l bin \
  && chmod a+rx bin/pulse-oximeter.linux.amd64 

FROM scratch
EXPOSE 9100

COPY --from=builder /src/bin/pulse-oximeter.linux.amd64 /pulse-oximeter

ENTRYPOINT ["/pulse-oximeter"]

