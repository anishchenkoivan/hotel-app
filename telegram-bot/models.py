from pydantic import BaseModel


class NotifyRequestModel(BaseModel):
    tgId: str
    message: str
