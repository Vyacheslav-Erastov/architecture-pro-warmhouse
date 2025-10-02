import threading


from fastapi import FastAPI
from fastapi.middleware.cors import CORSMiddleware
from sqlmodel import SQLModel
from app.api.main import api_router
from app.database.session import engine
from app.mq.consumer import start_consumer
from app.mq.producer import EventProducer
from app.core.config import config as cfg

import uvicorn


# создаём таблицы
SQLModel.metadata.create_all(bind=engine)


def run_consumer():
    producer = EventProducer(cfg.AMQP_URL)
    start_consumer(cfg.AMQP_URL, producer)


app = FastAPI()


origins = ["*"]

app.add_middleware(
    CORSMiddleware,
    allow_origins=origins,
    allow_credentials=True,
    allow_methods=["*"],
    allow_headers=["*"],
)

app.include_router(api_router)


if __name__ == "__main__":
    # запускаем consumer в отдельном потоке
    threading.Thread(target=run_consumer, daemon=True).start()

    # запускаем API
    uvicorn.run("main:app", host=cfg.API_HOST, port=cfg.API_PORT, reload=True)
