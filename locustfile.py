import random

from locust import HttpUser, task


request_template = "?source=[src]&event_name=event&event_status=status&payout=[pyt]"
basic_url = "/"


def generate_random_value():
    return f"{random.randint(100,999)}"

class SendingEvent(HttpUser):
    @task
    def send_event(self):
        request_string = request_template.replace(
            "[src]", generate_random_value()
        ).replace(
            "[pyt]", generate_random_value()
        )
        response = self.client.get(basic_url + request_string)
