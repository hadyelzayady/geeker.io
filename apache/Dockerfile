FROM debian:buster-slim

RUN apt update -y && apt upgrade -y && apt install apache2 -y

CMD ["httpd","-D"]
