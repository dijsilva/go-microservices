FROM python:3.9.8-alpine3.14

RUN \
  apk update && \
  apk add bash py-pip && \
  pip --no-cache-dir install -U pip

RUN apk add --update \
  python3-dev \
  build-base \
  linux-headers

COPY ./code/requirements.txt /requirements.txt
RUN pip3 install -r requirements.txt

EXPOSE 80
WORKDIR /app
COPY ./code/app /app

RUN ["chmod", "+x", "./entrypoint.sh"]
ENTRYPOINT ["./entrypoint.sh"]