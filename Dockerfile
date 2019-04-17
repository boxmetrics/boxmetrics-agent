FROM golang:1.12.4-stretch
RUN apt update 
RUN apt install -y make
RUN mkdir /app 
WORKDIR /app 
COPY . .
RUN make build
EXPOSE 8080 9090
CMD ["./boxagent"]