#!/bin/bash
APP_ENV=${APP_ENV:-local}
echo "[`date`] Running entrypoint using '${APP_ENV}' environment..."
echo "[`date`] Starting server..."
./server