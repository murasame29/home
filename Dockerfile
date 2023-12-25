FROM node:lts AS builder
WORKDIR /build
COPY . .

RUN npm install
RUN npm run build

FROM node:lts-alpine AS runner
COPY --from=builder /build/dist .

ENV HOST=0.0.0.0
ENV PORT=8080
EXPOSE 8080

CMD node dist/server/entry.mjs