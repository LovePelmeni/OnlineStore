FROM golang:1.18.2-bullseye
LABEL Creator=Kirill_Klimushin, Email=kirklimushin@gmail.com

CMD mkdir /proj/dir/
WORKDIR /proj/dir/
COPY . .

RUN go mod init github.com/LovePelmeni/OnlineStore/OrderCheckout
RUN go mod tidy
ENV GO111MODULE=on 
RUN go get -u ./
RUN go build github.com/LovePelmeni/OnlineStore/OrderCheckout/main/

ENTRYPOINT ["go", "run", "main.go"]
