from sqlmodel import Session
from app.models.telemetry import Telemetry, TelemetryCreate


class TelemetryRepository:
    def __init__(self, db: Session):
        self.db = db

    def save(self, data: TelemetryCreate) -> Telemetry:
        db_obj = Telemetry.model_validate(data)
        self.db.add(db_obj)
        self.db.commit()
        self.db.refresh(db_obj)
        return db_obj

    def get_by_device(self, device_id: int):
        return self.db.query(Telemetry).filter(Telemetry.device_id == device_id).all()
