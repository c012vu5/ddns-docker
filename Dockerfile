FROM golang AS build
WORKDIR /go/src
COPY ddns.go /go/src
ARG ACC PASS
ENV CGO_ENABLED=0
RUN go build -ldflags "-X main.ACC=${ACC:-id} -X main.PASS=${PASS:-password}" ddns.go

FROM scratch
LABEL maintainer="c012vu5"
LABEL description="Notify global IP to mydns"
COPY --from=build /go/src/ddns /bin/ddns
ENTRYPOINT ["/bin/ddns"]
