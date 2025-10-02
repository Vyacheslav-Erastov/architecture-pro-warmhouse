import pika
import json
import time
from datetime import datetime

import requests


def continuous_telemetry_sender(amqp_url, payloads, interval=3):
    """
    Непрерывно отправляет тестовые телеметрические данные
    """
    connection = pika.BlockingConnection(pika.URLParameters(amqp_url))
    channel = connection.channel()
    channel.queue_declare(queue="telemetry_queue")

    try:
        while True:
            for payload in payloads:

                event = {
                    **payload,
                    "timestamp": datetime.now().isoformat(),
                }

                channel.basic_publish(
                    exchange="",
                    routing_key="telemetry_queue",
                    body=json.dumps(event),
                    properties=pika.BasicProperties(
                        delivery_mode=2, content_type="application/json"
                    ),
                )

                print(payload)
                time.sleep(interval)

            print(f"Waiting {interval} seconds...\n")
            time.sleep(interval)

    except KeyboardInterrupt:
        print("\nStopping telemetry sender...")
    finally:
        connection.close()


if __name__ == "__main__":
    amqp_url = "amqp://guest:guest@localhost:5672/"
    r = requests.post(
        "http://localhost:8090/devices",
        json={
            "name": "Лампа прихожая",
            "type": "light_bulb",
            "status": {"power": "off"},
            "location": "Прихожая",
        },
    )
    id = int(r.text)
    payloads = [
        {"device_id": id, "event_data": {"power": "on"}},
        {"device_id": id, "event_data": {"power": "off"}},
    ]
    continuous_telemetry_sender(amqp_url, payloads, interval=3)
