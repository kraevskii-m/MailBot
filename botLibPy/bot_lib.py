import json
from dataclasses import dataclass
from typing import List

import requests
from dataclasses_json import dataclass_json


@dataclass
@dataclass_json
class Message:
    fr: str
    to: str
    subj: str
    body: str


class Bot:
    def __init__(self, token: str) -> None:
        self.token = token
        self.id = 0

    def get_updates(self) -> List[Message]:
        offset = self.id
        limit = 10
        address = "https://localhost:3000/bot{}/getupdates?offset={}&limit={}".format(self.token, offset, limit)
        response = requests.get(address)
        decoded = json.load(response)
        result = list(map(lambda x: Message.from_json(x), decoded))
        self.id += len(result)
        return result

    def send_message(self, message: Message) -> bool:
        address = "https://localhost:3000/bot{}/sendmessage".format(self.token)
        response = requests.post(address, data=message.to_json())
        if response.status_code == 200:
            return True
        return False
