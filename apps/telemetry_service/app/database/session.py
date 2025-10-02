"""Module for DB initialization"""

from sqlmodel import Session, create_engine

from app.core.config import config as cfg


engine = create_engine(str(cfg.SQLALCHEMY_DATABASE_URL))

SessionLocal = Session(engine)
