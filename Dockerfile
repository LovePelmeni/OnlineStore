FROM golang:1.8.13-bullseye
LABEL Creator=Kirill_Klimushin, Email=kirklimushin@gmail.com

CMD mkdir /proj/dir/
WORKDIR /proj/dir/
COPY . .

RUN go mod init github.com/LovePelmeni/OnlineStore/order_checkout/
RUN go mod tidy
ENV GO111MODULE=on 
RUN go get -u ./
RUN go build ./

ENTRYPOINT ["go", "run", "main.go"]
