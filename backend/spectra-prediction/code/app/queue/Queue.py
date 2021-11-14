import pika
from app.configuration.envs import ApplicationEnvs

from app.predictions.Handler import MessageHandler


class Queue:
    def __init__(self) -> None:
        self.queue_name = ApplicationEnvs.SPECTRA_QUEUE_NAME
        self.message_handler = MessageHandler()
        
        self.connection = pika.BlockingConnection(
            pika.ConnectionParameters(
                host=ApplicationEnvs.RABBITMQ_HOST,
                port=ApplicationEnvs.RABBITMQ_PORT,
                credentials=pika.PlainCredentials(
                    username=ApplicationEnvs.RABBITMQ_USER,
                    password=ApplicationEnvs.RABBITMQ_PASS,
                ),
                heartbeat=600,
                blocked_connection_timeout=300,
            )
        )

    def receive(self):
        self.channel = self.connection.channel()
        self.channel.queue_declare(
            queue=self.queue_name,
            durable=True,
            exclusive=False,
        )

        def _callback(ch, method, _properties, body):
            self.message_handler.pass_message(body.decode('ascii'))
            ch.basic_ack(delivery_tag=method.delivery_tag)

        self.channel.basic_consume(
            queue=self.queue_name,
            on_message_callback=_callback,
            auto_ack=False
        )

        self.channel.start_consuming()
