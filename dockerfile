# Choose whatever you want, version >= 1.16
FROM golang:1.23-alpine




ENV PORT=8080
ENV DB_HOST=localhost
ENV DB_USER=admim
ENV DB_PASS=admim1234
ENV DB_NAME=localhost
ENV DB_PORT=5432
ENV SUPA_API_URL=supabase.co
ENV SUPA_API_KEY=SUPA_API_KEY
ENV SUPA_BUCKET_NAME=bucket_name
ENV JWT_SECRET_KEY=jwt_secret_key
ENV API_MODEL=api_model
ENV SENDING_EMAIL=sending_email
ENV EMAIL_PASSWORD=email_password

WORKDIR /app

RUN go install github.com/air-verse/air@latest

COPY go.mod go.sum ./
RUN go mod download

COPY . .




CMD ["air"]