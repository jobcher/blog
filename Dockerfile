FROM hub.jobcher.com/base/base-hugo:latest as build
COPY . /blog
RUN cd /blog \
    && hugo -t LoveIt -D

FROM nginx:1.21
LABEL maintainer="nb@nbtyfood.com"
COPY --from=build /blog/public/ /usr/share/nginx/html
COPY ./ads.txt /usr/share/nginx/html