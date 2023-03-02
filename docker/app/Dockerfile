FROM golang:1.19.5
ENV TZ=Asia/Tokyo

RUN mkdir /app
WORKDIR /app
ADD ./ /app/

RUN go install github.com/cosmtrek/air@v1.40.0
COPY ./ ./app/
RUN go mod download

EXPOSE 8080
CMD ["air", "-c", ".air.toml"]