from typing import Optional

import faust
from faust.serializers import codecs
from pydantic import BaseModel


class RawEvent(BaseModel):
    source: Optional[str] = None
    event_name: Optional[str] = None
    event_status: Optional[str] = None
    created: Optional[str] = None
    payout: Optional[str] = None

class PydanticSerializer(codecs.Codec):
    def __init__(self, cls_type):
        self.cls_type = cls_type
        super(self.__class__, self).__init__()

    def _dumps(self, cls) -> bytes:
        return cls.json().encode()

    def _loads(self, s: bytes):
        cls_impl = self.cls_type.parse_raw(s)
        return cls_impl


app = faust.App(
    f"python-sender",
    broker="kafka-go-vs-faust:9092",
    web_host="0.0.0.0",
    web_port="6050",
)

topic_from = app.topic(
    "test-events-from",
    partitions=2,
    value_serializer=PydanticSerializer(RawEvent)
)
topic_to = app.topic(
    "test-events-to",
    partitions=2,
    value_serializer=PydanticSerializer(RawEvent)
)


@app.agent(topic_from)
async def on_event(stream) -> None:
    async for msg_key, event in stream.items():
        print(f"Received :: {event=}")
        await topic_to.send(key="key", value=event)
        print("...sent further...")