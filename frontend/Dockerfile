# syntax=docker/dockerfile:1
FROM nginxinc/nginx-unprivileged:alpine-perl AS base
USER nginx
EXPOSE 8080

FROM node:lts AS dependencies
WORKDIR /src
COPY package.json package-lock.json ./
RUN --mount=type=cache,target=/root/.npm npm ci

FROM dependencies AS build
WORKDIR /src
COPY . .
RUN npm run build

FROM base AS final
COPY --chown=nginx:nginx --from=build /src/dist/frontend/browser /usr/share/nginx/html
COPY nginx.conf /etc/nginx/nginx.conf
ENTRYPOINT ["nginx", "-c", "/etc/nginx/nginx.conf"]
CMD ["-g", "daemon off;"]
