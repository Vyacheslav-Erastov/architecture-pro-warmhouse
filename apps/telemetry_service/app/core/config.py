"""Module with all FastAPI configuration based on Pydantic Settings"""

import os
from pathlib import Path
from pydantic import PostgresDsn, computed_field
from pydantic_core import MultiHostUrl
from pydantic_settings import BaseSettings, SettingsConfigDict


class Config(BaseSettings):
    """Class for store all settings and environment

    Args:
        BaseSettings (BaseSettings): Base Pydantic Settings

    Returns:
        Settings: Settings object with all configuration
    """

    model_config = SettingsConfigDict(
        env_file=(".env", os.path.join(Path(os.getcwd()).parent.absolute(), ".env")),
        env_ignore_empty=True,
        extra="ignore",
    )

    ENVIRONMENT: str = "LOCAL"

    API_HOST: str = "0.0.0.0"
    API_PORT: int = 8000

    POSTGRES_PORT: int = 5432
    POSTGRES_USER: str = "postgres"
    POSTGRES_PASSWORD: str = "postgres"
    POSTGRES_DB: str = "telemetry_service"

    RMQ_PORT: int = 5672
    RMQ_USER: str = "guest"
    RMQ_PASSWORD: str = "guest"

    DEVICE_SERVICE_HOST: str = "device-service"
    DEVICE_SERVICE_PORT: int = 8090

    @computed_field
    @property
    def RMQ_HOST(self) -> str:
        """Property method to calculate Postgres host

        Returns:
            str: Postgres host
        """
        if self.ENVIRONMENT != "CONTAINER":
            return "localhost"
        return "rabbitmq"

    @computed_field
    @property
    def POSTGRES_HOST(self) -> str:
        """Property method to calculate Postgres host

        Returns:
            str: Postgres host
        """
        if self.ENVIRONMENT != "CONTAINER":
            return "localhost"
        return "postgres"

    @computed_field
    @property
    def SQLALCHEMY_DATABASE_URL(self) -> PostgresDsn:
        """Property method to calculate Potgres URL

        Returns:
            PostgresDsn: Postgres URL
        """
        return MultiHostUrl.build(
            scheme="postgresql+psycopg",
            username=self.POSTGRES_USER,
            password=self.POSTGRES_PASSWORD,
            host=self.POSTGRES_HOST,
            port=self.POSTGRES_PORT,
            path=self.POSTGRES_DB,
        )

    @computed_field
    @property
    def AMQP_URL(self) -> str:
        """Property method to calculate AMQP URL

        Returns:
            str: AMQP URL
        """
        return f"amqp://{self.RMQ_USER}:{self.RMQ_PASSWORD}@{self.RMQ_HOST}:{self.RMQ_PORT}/"

    @computed_field
    @property
    def DEVICE_SERVICE_URL(self) -> str:
        """Property method to calculate Device Service URL

        Returns:
            str: Device Service URL
        """
        return f"http://{self.DEVICE_SERVICE_HOST}:{self.DEVICE_SERVICE_PORT}"


config = Config()
