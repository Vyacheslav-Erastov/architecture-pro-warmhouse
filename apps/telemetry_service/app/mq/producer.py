import json
import pika


class EventProducer:
    def __init__(self, amqp_url: str):
        self.connection = pika.BlockingConnection(pika.URLParameters(amqp_url))
        self.channel = self.connection.channel()
        self.exchange = "system_events"
        self.channel.exchange_declare(exchange=self.exchange, exchange_type="fanout")

    def publish_event(self, event: dict):
        body = json.dumps(event)
        self.channel.basic_publish(exchange=self.exchange, routing_key="", body=body)
