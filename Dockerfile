FROM golang:1.18-bullseye 

LABEL Author=Klimushin_Kirill, Email=kirklimushin@gmail.com 
CMD mkdir /project/dir/ 
WORKDIR /project/dir/ 

RUN go init github.com/LovePelmeni/StoreProject 
RUN go get ./ 
RUN go build ./ 
RUN go test ./tests
ENTRYPOINT [ "go", "run", "./main/main.go" ]



