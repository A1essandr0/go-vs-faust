import logging

from aiokafka import AIOKafkaProducer


KAFKA_BROKER = "kafka-go-vs-faust:9092"
KAFKA_USERNAME = "username"
KAFKA_PASSWORD = "password"
KAFKA_NO_SASL = True
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
        bootstrap_servers = KAFKA_BROKER
        sasl_plain_username = KAFKA_USERNAME
        sasl_plain_password = KAFKA_PASSWORD
        logger.info(f"Starting kafka producer: {bootstrap_servers} {sasl_plain_username} {sasl_plain_password}")

        if KAFKA_NO_SASL:
            logger.info("Disabled SASL in kafka connection")
            producer = AIOKafkaProducer(bootstrap_servers=bootstrap_servers)
        else:
            logger.info("enabled SASL in kafka connection")
            producer = AIOKafkaProducer(
                bootstrap_servers=bootstrap_servers,
                sasl_plain_username=sasl_plain_username,
                sasl_plain_password=sasl_plain_password,
                sasl_mechanism="PLAIN",
                security_protocol="SASL_PLAINTEXT",
            )
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
        logger.info("Stop kafka producer")
        await self.producer.stop()