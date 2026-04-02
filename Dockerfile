FROM goLang:1.22.12

WORKDIR /go
COPY . /go
RUn go build -o main .