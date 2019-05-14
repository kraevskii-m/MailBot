import json
import os
import time
from dataclasses import dataclass
from pathlib import Path

import botLibPy.bot_lib as bot_lib

TOKEN = ""
address = "zavuchbot@yandex.ru"


@dataclass
class Task:
    task_id: str
    task_level: int
    task_text: str
    correct_answer: int


@dataclass
class User:
    address: str
    tasks: list


path = str(Path().absolute()) + "\\exercises"
os.chdir(path)

tasks = []
for name in os.listdir(path):
    with open(name, encoding='utf-8') as file:
        text = file.read().replace('\0', '')
        data = json.loads(text)
        task = Task(data["task_id"], data["task_level"], data["task_text"], data["correct_answer"])
        tasks.append(task)
tasks = sorted(tasks, key=lambda x: x.task_level)

users = []

bot = bot_lib.Bot(TOKEN)


class NoRemainingTasksException(Exception):
    pass


def get_task(last_answer: bool, user_tasks: list) -> Task:
    remaining_tasks = list(sorted(set(tasks).difference(user_tasks), key=lambda x: x.task_level))
    if len(remaining_tasks) == 0:
        raise NoRemainingTasksException
    if last_answer:
        for i, t in enumerate(remaining_tasks):
            if t.task_level > user_tasks[-1].task_level:
                return t
        return remaining_tasks[-1]
    for i, t in enumerate(reversed(remaining_tasks)):
        if t.task_level < user_tasks[-1].task_level:
            return t
    return remaining_tasks[0]


def send_task(user_id: int, answer: str):
    us = users[user_id]
    if len(us.tasks) == 0:
        us.tasks.append(tasks[len(tasks) // 2])
        bot.send_message(bot_lib.Message(address, us.address, "Новая задача", us.tasks[-1].task_text))
        return
    elif answer == us.tasks[-1].correct_answer:
        bot.send_message(bot_lib.Message(address, us.address, "Ответ верный!", ""))
        us.tasks.append(get_task(True, us.tasks))
    else:
        bot.send_message(
            bot_lib.Message(address, us.address, "Неправильно!", "Правильный ответ: " + us.tasks[-1].correct_answer))
        us.tasks.append(get_task(False, us.tasks))
    bot.send_message(bot_lib.Message(address, us.address, "Новая задача", us.tasks[-1].task_text))


while True:
    time.sleep(5)
    messages = bot.get_updates()
    for message in messages:
        sent = False
        for index, u in enumerate(users):
            if u.address == message.fr:
                send_task(index, message.body)
                sent = True
                break
        if not sent:
            u = User(message.fr, [])
            users.append(u)
            send_task(len(users) - 1, "")
