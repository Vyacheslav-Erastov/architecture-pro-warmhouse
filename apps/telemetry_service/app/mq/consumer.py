import json
import requests
import pika
from app.repository.telemetry_repository import TelemetryRepository
from app.models.telemetry import TelemetryCreate
from app.mq.producer import EventProducer
from app.core.config import config as cfg
from app.database.session import SessionLocal


def start_consumer(amqp_url: str, event_producer: EventProducer):
    connection = pika.BlockingConnection(pika.URLParameters(amqp_url))
    channel = connection.channel()
    channel.queue_declare(queue="telemetry_queue")

    def callback(ch, method, properties, body):
        data = json.loads(body)
        device_id = data.get("device_id")

        # сохраняем в БД
        with SessionLocal as db:
            repo = TelemetryRepository(db)
            repo.save(
                TelemetryCreate(
                    device_id=device_id,
                    event_data=data["event_data"],
                )
            )
        # обновляем статус устройства через API
        if data["event_data"]:
            try:
                r = requests.patch(
                    f"{cfg.DEVICE_SERVICE_URL}/devices/{device_id}/status",
                    json={"status": data["event_data"]},
                    timeout=2,
                )
            except Exception as e:
                print(f"Failed to update device status: {e}")

        # публикуем событие в очередь системных событий
        event_producer.publish_event(
            {
                "source": "telemetry-service",
                "device_id": device_id,
                "event": data,
            }
        )

    channel.basic_consume(
        queue="telemetry_queue", on_message_callback=callback, auto_ack=True
    )
    print("Telemetry Consumer started. Waiting for messages...")
    channel.start_consuming()
