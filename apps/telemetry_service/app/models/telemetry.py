from datetime import datetime
from sqlmodel import JSON, Field, SQLModel


class TelemetryBase(SQLModel):
    event_data: dict = Field(sa_type=JSON)
    timestamp: datetime = Field(default_factory=datetime.now)


class TelemetryCreate(TelemetryBase):
    device_id: int


class Telemetry(TelemetryCreate, table=True):
    __tablename__ = "telemetry"

    id: int | None = Field(default=None, primary_key=True)
