import json

import requests

url = 'http://localhost:3000'


class Mail:

    def __init__(self, address, body, subject):
        self.body = body
        self.address = address
        self.subject = subject


resp = requests.get(url=url + "/updates")
data = resp.json()

letters = []
for d in data:
    letters.append(Mail(d[0], d[1], d[2]))

r = requests.post(url=url + "/send", data=json.dumps(data))
print(r.status_code, r.reason)
