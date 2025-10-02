from fastapi import APIRouter
from app.repository.telemetry_repository import TelemetryRepository
from app.models.telemetry import Telemetry
from app.api.deps import SessionDep

router = APIRouter()


@router.get("/{device_id}", response_model=list[Telemetry])
def get_telemetry(device_id: int, db: SessionDep):
    repo = TelemetryRepository(db)
    return repo.get_by_device(device_id)
