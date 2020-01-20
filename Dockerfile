FROM golang:1.12-buster
RUN mkdir /app 
WORKDIR /app 
COPY . .
EXPOSE 4455 5544
CMD ["./bin/boxmetrics-agent"]
