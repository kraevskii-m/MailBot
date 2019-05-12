from dataclasses import dataclass
from typing import List


@dataclass
class Message:
    fr: str
    to: str
    subj: str
    body: str


class bot:
    def __init__(self, token: str) -> None:
        self.token = token

    def get_updates(self) -> List[Message]:
        
