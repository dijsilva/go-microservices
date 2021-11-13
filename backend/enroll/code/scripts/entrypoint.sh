#!/bin/bash -e

APP_ENV=${APP_ENV:-local}
echo "[`date`] Running entrypoint using '${APP_ENV}' environment..."
CONFIG_FILE=/app/config/${APP_ENV}.yml
if [[ -z ${APP_DSN} ]]; then
  export APP_DSN=`sed -n 's/^dsn:[[:space:]]*"\(.*\)"/\1/p' ${CONFIG_FILE}`
fi
echo "[`date`] Running DB migrations..."
/usr/local/bin/migrate -database "${APP_DSN}" -path /app/migrations/ up

echo "[`date`] Starting server..."
./server