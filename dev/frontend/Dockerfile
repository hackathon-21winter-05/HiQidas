FROM node:16.13.1-alpine AS builder
WORKDIR /build
RUN apk add git

RUN git clone https://github.com/hackathon-21winter-05/HiQidas_UI.git
WORKDIR /build/HiQidas_UI
RUN npm ci --unsafe-perm
RUN npm run build

FROM caddy:2.4.6-alpine AS runner
EXPOSE 80

COPY --from=builder /build/HiQidas_UI/dist /usr/share/caddy
