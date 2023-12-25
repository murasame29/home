FROM node:lts AS base
WORKDIR /app

COPY package.json package-lock.json ./

FROM base AS prod-deps
RUN npm install --production

FROM base AS build-deps
RUN npm install --production=false

FROM build-deps AS build
COPY . .
RUN npm run build

FROM base AS runtime
COPY --from=prod-deps /app/node_modules ./node_modules
COPY --from=build /app/dist ./dist

ENV HOST=0.0.0.0
ENV PORT=8080
ENV NR_ACCOUNT_ID=
ENV NR_TRUST_KEY=
ENV NR_AGENT_ID=
ENV NR_LICENSE_KEY=
ENV NR_APPLICATION_ID=
EXPOSE 8080
CMD node ./dist/server/entry.mjs