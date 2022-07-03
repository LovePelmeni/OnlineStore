FROM --platform=linux/amd64 golang:1.18.2-alpine
LABEL Creator=Kirill_Klimushin, Email=kirklimushin@gmail.com

ENV GO111MODULE=on 
ENV CGO_ENABLED=1
ENV GOPATH=/usr/local/go
ENV PKG_CONFIG_PATH=/usr/local/pgk-go

COPY . .
RUN apk add git
RUN git clone https://github.com/edenhill/librdkafka.git
RUN cd librdkafka
RUN ./configure
RUN make
RUN make install
RUN export PKG_CONFIG_PATH=$PKG_CONFIG_PATH:/usr/lib/pkgconfig

CMD mkdir /proj/dir/
WORKDIR /proj/dir/

COPY ./go.mod . 
COPY ./go.sum . 

COPY . .
RUN go mod download 
RUN go build -tags dynamic -o main ./main/main.go
ENTRYPOINT ["go", "run", "./main/main.go"]



