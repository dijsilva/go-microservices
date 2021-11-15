FROM node:14.18-alpine3.12 as build

ENV PATH /app/node_modules/.bin:$PATH
RUN set -ex && apk update && \
    apk add --update --no-cache git
WORKDIR /app
COPY [ "./code/package.json", "." ]
RUN npm config set unsafe-perm true && \
    npm install && \
    npm install react-scripts@4.0.3 -g
COPY [ "./code", "." ]

RUN npm run build:local


# production environment
FROM nginx:1.21.4-alpine
COPY --from=build /app/build /usr/share/nginx/html
RUN rm /etc/nginx/conf.d/default.conf
COPY ./nginx.conf /etc/nginx/conf.d
COPY entrypoint.sh /
# RUN chmod +x /docker-entrypoint.sh
EXPOSE 80
CMD ["sh", "/entrypoint.sh"]
