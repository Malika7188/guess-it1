FROM golang

#Metadata
LABEL maintainer = "masman, clomollo"
LABEL version ="1.0"
LABEL description ="This is an ascii-art-web application with go"

WORKDIR /app/

COPY . .

RUN go mod tidy

EXPOSE 8080

CMD ["go","run","main.go"]
