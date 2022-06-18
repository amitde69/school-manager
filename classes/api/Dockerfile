
FROM golang:1.18 as builder
WORKDIR /usr/src/app
ARG GITHUB_TOKEN=""
RUN git config --global url."https://${GITHUB_TOKEN}:x-oauth-basic@github.com/".insteadOf "https://github.com/"
ENV GOPRIVATE="github.com/amitde69"
COPY ./go.mod ./
COPY ./go.sum ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -ldflags "-s -w" -o server .

FROM scratch as runner
COPY --from=builder /usr/src/app/server /opt/app/
CMD ["/opt/app/server"]