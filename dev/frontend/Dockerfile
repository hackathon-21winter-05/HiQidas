FROM node:16.13.1-alpine AS builder
WORKDIR /build
RUN apk add git

COPY ./HiQidas_UI /build
RUN npm ci --unsafe-perm
RUN npm run build

FROM caddy:2.4.6-alpine AS runner
EXPOSE 80

COPY --from=builder /build/dist /usr/share/caddy
