FROM golang AS builder
LABEL stage=builder

#disable crosscompiling 
ENV CGO_ENABLED=0

#compile linux only
ENV GOOS=linux

WORKDIR /build
COPY . .

RUN go build -a -installsuffix cgo -o main .

FROM alpine
COPY --from=builder /build/main .
CMD ./main
EXPOSE 8080