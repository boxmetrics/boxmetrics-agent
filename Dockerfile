FROM golang:1.12.4-stretch
RUN apt update 
RUN apt install -y make
RUN mkdir /app 
WORKDIR /app 
COPY . .
RUN make test
RUN make build
EXPOSE 4455 5544
CMD ["./bin/boxmetrics-agent"]