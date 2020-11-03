FROM golang:1.15.2 as develop

WORKDIR /go/src/lottery_back
COPY . .

EXPOSE 8080
CMD /bin/bash


FROM golang:1.15.2 as builder

WORKDIR /go/src/lottery_back
COPY ./* ./
COPY ./cmd/ ./cmd/
COPY ./pkg/ ./pkg/
COPY ./conf/ ./conf/
COPY ./public/ ./public/
COPY ./private/ ./private/

RUN CGO_ENABLED=0 go build -o /go/bin/lottery_back ./cmd/lottery_app/


FROM alpine:3.12 as app
COPY --from=builder /go/bin/lottery_back .
COPY --from=builder /go/src/lottery_back/conf/ ./conf/
COPY --from=builder /go/src/lottery_back/private/ ./private/
ENTRYPOINT [ "./lottery_back" ]

EXPOSE 8080
