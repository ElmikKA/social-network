FROM node:18.17.1-alpine AS build-stage

WORKDIR /social-network-frontend

COPY package.json package-lock.json ./

RUN npm install

COPY . .

RUN npm run build

FROM nginx:alpine AS production-stage
COPY --from=build-stage /social-network-frontend/dist /usr/share/nginx/html
COPY nginx.conf /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]
