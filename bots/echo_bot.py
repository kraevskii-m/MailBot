import botLibPy.bot_lib as bot_lib

TOKEN = ""
address = "echobot-mailbot@yandex.ru"

bot = bot_lib.Bot(TOKEN)
while True:
    messages = bot.get_updates()
    for message in messages:
        bot.send_message(bot_lib.Message(address, message.fr, message.subj, message.body))
