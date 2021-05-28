FROM nginx:stable
RUN rm /etc/nginx/conf.d/default.conf
COPY ./nginx/conf/ /etc/nginx/

COPY ./resources/ /usr/share/nginx/cdn

