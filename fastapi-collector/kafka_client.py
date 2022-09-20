import logging

from aiokafka import AIOKafkaProducer


KAFKA_BROKER = "kafka-go-vs-faust:9092"
KAFKA_EVENTS_TOPIC = 'test-events-from'


logger = logging.getLogger(__name__)


class MetaSingleton(type):
    _instances = {}
    def __call__(cls, *args, **kwargs):
        if cls not in cls._instances:
            cls._instances[cls] = super(MetaSingleton, cls).__call__(*args, **kwargs)
        return cls._instances[cls]


class KafkaClient(metaclass=MetaSingleton):
    def __init__(self):
        self.producer = None

    async def producer_start(self):
        producer = AIOKafkaProducer(bootstrap_servers=KAFKA_BROKER)
        await producer.start()
        return producer


    async def send_message(self, topic: str, message: bytes):
        if not self.producer:
            logger.info("Kafka producer is not running")
            self.producer = await self.producer_start()

        try:
            logger.info(f"Sending to {topic=}")
            await self.producer.send_and_wait(topic, message)

        except Exception as exc:
            logger.error(f"Can't send message to kafka; {exc}")
            raise exc

    async def stop(self):
        await self.producer.stop()