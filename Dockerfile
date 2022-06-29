FROM golang:1.18-bullseye 
LABEL Creator=Klimushin_Kirill Email=kirklimushin@gmail.com

CMD mkdir /project/dir/ 
WORKDIR /project/dir/ 

RUN go get ./
RUN go build ./ 

RUN chmod +x entrypoint.sh 
ENTRYPOINT ["sh", "entrypoint.sh"]

