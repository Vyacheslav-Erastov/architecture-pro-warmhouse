from datetime import datetime
import random
from fastapi import APIRouter
from pydantic import BaseModel, Field

router = APIRouter(prefix="/temperature", tags=["temperature"])


class TemperatureResponse(BaseModel):
    value: float
    unit: str = "°C"
    timestamp: datetime = Field(
        default_factory=lambda: datetime.now().isoformat() + "Z"
    )
    location: str
    status: str = "active"
    sensor_id: str
    sensor_type: str = "temperature"
    description: str = ""


@router.get("/", response_model=TemperatureResponse)
async def get_temperature_by_location(location: str):
    sensor_id = "0"
    match location:
        case "Living Room":
            sensor_id = "1"
        case "Bedroom":
            sensor_id = "2"
        case "Kitchen":
            sensor_id = "3"
    value = random.randint(0, 30)
    return TemperatureResponse(value=value, location=location, sensor_id=sensor_id)


@router.get("/{sensor_id}", response_model=TemperatureResponse)
async def get_temperature_by_sensor_id(sensor_id: str):
    location = "Unknown"
    match sensor_id:
        case "1":
            location = "Living Room"
        case "2":
            location = "Bedroom"
        case "3":
            location = "Kitchen"
    value = random.randint(0, 30)
    return TemperatureResponse(value=value, location=location, sensor_id=sensor_id)
