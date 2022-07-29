##
## Build
##
FROM golang:1.18-buster AS build

ENV APP_HOME /go/src/myapp
WORKDIR $APP_HOME

COPY . .

RUN go mod download
RUN go mod verify



RUN GOOS=linux CGO_ENABLED=0 GOARCH=amd64 go build -o /myapp

##
## Deploy
##
FROM gcr.io/distroless/base-debian11

WORKDIR /

COPY --from=build /myapp /myapp

EXPOSE 8080

#ENV TZ Europe/Moscow
#RUN ln -snf /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

USER nonroot:nonroot

ENTRYPOINT ["/myapp"]

