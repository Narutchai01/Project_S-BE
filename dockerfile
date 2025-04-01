# Choose whatever you want, version >= 1.16
FROM golang:1.23-alpine

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .

ENV PORT=""
ENV DB_HOST=""
ENV DB_USER=""
ENV DB_PASS=""
ENV DB_NAME=""
ENV DB_PORT=""
ENV SUPA_API_URL=""
ENV SUPA_API_KEY=""
ENV SUPA_BUCKET_NAME=""
ENV JWT_SECRET_KEY=""
ENV API_MODEL=""
ENV SENDING_EMAIL=""
ENV EMAIL_PASSWORD=""

CMD ["air"]