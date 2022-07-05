FROM --platform=linux/amd64 golang:1.17-alpine
LABEL Creator=Kirill_Klimushin, Email=kirklimushin@gmail.com

RUN echo "Settings Up Environment Variables...."
ENV PATH="/go/bin:${PATH}"
ENV GO111MODULE=on 
ENV CGO_ENABLED=1
ENV GOARCH=amd64 
ENV GOOS=linux
ENV GIN_MODE=release
ENV PKG_CONFIG_PATH=/usr/local/pkg-go

CMD mkdir /proj/dir/
WORKDIR /proj/dir/

COPY ./go.mod ./go.sum ./ 

RUN go get -tags musl ./... 
RUN go mod download

RUN echo "Compling Sources.."
RUN apk update && apk upgrade && apk add pkgconf git bash build-base
RUN git clone https://github.com/edenhill/librdkafka.git && cd librdkafka && ./configure --prefix /usr && make && make install  
RUN export PKG_CONFIG_PATH=$PKG_CONFIG_PATH:/usr/lib/pkgconfig/

RUN set -ex &&\
    apk add --no-progress --no-cache \
    gcc \
    musl-dev
COPY . .
RUN echo "Building Project..."
RUN go build -tags musl -o main ./main/main.go
ENTRYPOINT ["go", "run", "./main/main.go"]
