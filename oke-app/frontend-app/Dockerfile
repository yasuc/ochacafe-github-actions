FROM node:16 AS builder

WORKDIR /app
COPY package*.json ./
COPY . .
RUN npm install && npm run build

FROM node:16-slim AS runner
ENV NODE_ENV=production

WORKDIR /app

COPY --from=builder /app/next.config.js ./
COPY --from=builder /app/.next/static ./.next/static

COPY --from=builder /app/.next/standalone ./
CMD ["node", "server.js"]