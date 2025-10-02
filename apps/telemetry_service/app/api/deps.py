from typing import Annotated, Generator

from fastapi import Depends
from sqlmodel import Session

from app.database.session import SessionLocal


def get_db() -> Generator[Session, None, None]:
    """Function to genarte DB sessions

    Yields:
        Generator[Session, None, None]: DB Session generator
    """
    with SessionLocal as db:
        yield db


SessionDep = Annotated[Session, Depends(get_db)]
