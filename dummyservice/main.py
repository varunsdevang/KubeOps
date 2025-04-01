# main.py
from fastapi import FastAPI, Request, HTTPException
from pydantic import BaseModel
from fastapi.responses import JSONResponse

import time
import grpc
import asyncio
import ratelimit_pb2 as rl_pb
import ratelimit_pb2_grpc as rl_pb_grpc

app = FastAPI()


GRPC_SERVER_ADDRESS = "localhost:3001"
SERVICE = "dummyService"

async def check_rate_limit(ip_address: str) -> bool:
    """Calls the gRPC rate limiter API and returns a boolean."""
    async with grpc.aio.insecure_channel(GRPC_SERVER_ADDRESS) as channel:
        stub = rl_pb_grpc.RateLimitServiceStub(channel)
        request = rl_pb.RateLimitRequest(ip=ip_address, service=SERVICE)
        try:
            response = await stub.IsRequestAllowed(request)
            return response.isAllowed
        except grpc.aio.AioRpcError as e:
            print(f"gRPC error: {e}") #handle errors more gracefully in prod
            return False #Default to rate limited on error.

@app.middleware("http")
async def rate_limit_middleware(request: Request, call_next):
    ip_address = request.client.host #or use request.headers.get("X-Forwarded-For")

    allowed = await check_rate_limit(ip_address)

    if not allowed:
        return JSONResponse(status_code=429, content={"message": "Rate limit exceeded"})

    response = await call_next(request)
    return response

class Item(BaseModel):
    name: str
    description: str = None
    price: float
    tax: float = None

items = {}

@app.post("/items/{item_id}")
async def create_item(item_id: int, item: Item):
    if item_id in items:
        raise HTTPException(status_code=400, detail="Item already exists")
    items[item_id] = item
    return {"item_id": item_id, **item.dict()}

@app.get("/items/{item_id}")
async def read_item(item_id: int):
    if item_id not in items:
        raise HTTPException(status_code=404, detail="Item not found")
    return {"item_id": item_id, **items[item_id].dict()}

@app.get("/items/")
async def read_items():
    return items

@app.put("/items/{item_id}")
async def update_item(item_id: int, item: Item):
    if item_id not in items:
        raise HTTPException(status_code=404, detail="Item not found")
    items[item_id] = item
    return {"item_id": item_id, **item.dict()}

@app.delete("/items/{item_id}")
async def delete_item(item_id: int):
     if item_id not in items:
        raise HTTPException(status_code=404, detail="Item not found")
     del items[item_id]
     return {"deleted": True}