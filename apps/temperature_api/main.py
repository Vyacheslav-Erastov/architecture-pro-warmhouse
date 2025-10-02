from fastapi import FastAPI
from api import router as temperature_router
import uvicorn

app = FastAPI()

app.include_router(temperature_router)


if __name__ == "__main__":
    uvicorn.run("main:app", host="0.0.0.0", port=8081, reload=True)
