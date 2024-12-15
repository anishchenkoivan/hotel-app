from pydantic import BaseModel


class NotifyRequestModel(BaseModel):
    tgId: int
    message: str
