FROM python:3.8-slim
RUN \
  apt-get update && \
  apt-get install build-essential -y && \
  pip --no-cache-dir install -U pip

RUN apt-get install \
  supervisor \
  python3-dev -y

COPY ./code/requirements.txt /requirements.txt
RUN python3 -m pip install -r requirements.txt

EXPOSE 80
WORKDIR /app
COPY ./code/app /app

RUN ["chmod", "+x", "./entrypoint.sh"]
ENTRYPOINT ["./entrypoint.sh"]