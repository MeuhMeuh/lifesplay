FROM golang:alpine

RUN apk --no-cache add curl git gcc libc6-compat
RUN curl -fsSL -o /usr/local/bin/dep https://github.com/golang/dep/releases/download/v0.5.0/dep-linux-amd64 && chmod +x /usr/local/bin/dep

RUN mkdir -p /go/src/github.com/meuhmeuh/lifesplay
RUN mkdir -p /go/bin

ADD . /go/src/github.com/meuhmeuh/lifesplay

WORKDIR /go/src/github.com/meuhmeuh/lifesplay

RUN go get github.com/cortesi/modd/cmd/modd

RUN dep ensure -vendor-only

CMD ["modd"]
