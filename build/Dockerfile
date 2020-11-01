FROM golang:1.15.2 as develop

WORKDIR /go/src/lottery_back
COPY . .

CMD /bin/bash


FROM golang:1.15.2 as builder

WORKDIR /go/src/lottery_back
COPY ./* ./
COPY ./cmd/ ./cmd/
COPY ./pkg/ ./pkg/
COPY ./conf/ ./conf/
COPY ./public/ ./public/

RUN CGO_ENABLED=0 go build -o /go/bin/lottery_back ./cmd/lottery_back_app/


FROM alpine:3.12 as app
COPY --from=builder /go/bin/lotery-back .
COPY --from=builder /go/src/lottery_back/conf/ ./conf/
COPY --from=builder /go/src/lottery_back/public/ ./public/
ENTRYPOINT [ "./lotery-back" ]
