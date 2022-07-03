FROM --platform=linux/amd64 golang:1.18.2-alpine
LABEL Creator=Kirill_Klimushin, Email=kirklimushin@gmail.com

ENV GO111MODULE=on 
ENV CGO_ENABLED=1
ENV PKG_CONFIG_PATH=/usr/local/pgk-go

COPY . .
RUN apk add git
RUN apk add librdkafka-dev --version 1.9.0-RC9
RUN apk add pkgconf 
RUN apk add build-base
RUN export PKG_CONFIG_PATH=$PKG_CONFIG_PATH:/usr/lib/pkgconfig/

CMD mkdir /proj/dir/
WORKDIR /proj/dir/

COPY . .
RUN go mod init github.com/LovePelmeni/OnlineStore/OrderCheckout
RUN go mod tidy 
RUN go build -tags musl -o main ./main/main.go
ENTRYPOINT ["go", "run", "./main/main.go"]



