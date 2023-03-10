FROM registry.access.redhat.com/ubi8/go-toolset:1.18 as build
COPY main.go .
COPY go.mod .
COPY go.sum .
COPY backend ./backend


RUN go mod download && \
    go build main.go

FROM ubi8/ubi-micro
COPY --from=build /opt/app-root/src/main .
CMD ./main