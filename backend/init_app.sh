echo "Building database"
docker-compose up -d --build postgres-db && docker-compose up -d --build redis_auth_control && \
    docker-compose up -d --build mongo_espectra

echo "Building RabbitMQ"
docker-compose up -d --build rabbitmq

echo "Building microservices"
docker-compose up -d --build enroll && docker-compose up -d --build auth_control && \
    docker-compose up -d --build spectra
