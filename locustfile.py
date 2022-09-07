import random

from locust import HttpUser, task, between, constant_throughput


request_template = "?source=[src]&event_name=event&event_status=status&payout=[pyt]"
basic_url = "/"


def generate_random_value():
    return f"{random.randint(100,999)}"

class SendingEvent(HttpUser):
    # wait_time = between(1, 2)
    # wait_time = constant_throughput(500)

    @task
    def send_event(self):
        request_string = request_template.replace(
            "[src]", generate_random_value()
        ).replace(
            "[pyt]", generate_random_value()
        )
        response = self.client.get(basic_url + request_string)
