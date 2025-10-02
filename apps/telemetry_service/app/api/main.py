"""Module with initialize main API router"""

from fastapi import APIRouter

from app.api.routes import telemetry

api_router = APIRouter()
api_router.include_router(telemetry.router, prefix="/telemetry", tags=["telemetry"])
