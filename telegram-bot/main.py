import sys
import logging

from fastapi import FastAPI
from telebot import TeleBot, types
from telebot.apihelper import ApiTelegramException
import multiprocessing
import uvicorn
import requests

from models import NotifyRequestModel
from config import TOKEN, BOT_HOST, BOT_PORT, BOOKING_SERVICE_HOST, BOOKING_SERVICE_PORT, HOTEL_SERVICE_HOST, \
    HOTEL_SERVICE_PORT

multiprocessing.set_start_method('spawn', True)

logging.basicConfig(level=logging.INFO, stream=sys.stdout)
logger = logging.getLogger(__name__)
bot = TeleBot(token=TOKEN)
app = FastAPI()


@bot.message_handler(commands=['start'])
def start_handler(message: types.Message):
    bot.send_message(message.from_user.id, f"Привет, чтобы забронировать отель введи /book")


user_data = {}


@bot.message_handler(commands=['book'])
def book_room_handler(message):
    user_data[message.from_user.id] = {}
    types.InlineKeyboardMarkup()
    res = requests.get(f"http://{HOTEL_SERVICE_HOST}:{HOTEL_SERVICE_PORT}/hotel-service/api/room")
    if res.status_code != 200:
        bot.send_message(message.chat.id, "Сервис временно не доступен")
        logger.info(f"Hotel service responsed with code {res.status_code}: {res.content}")
        return
    rooms = res.json()
    if len(rooms) == 0:
        bot.send_message(message.chat.id, "К сожалению, нет доступных комнат")
        return
    bot.send_message(message.chat.id,
                     "Введите id комнаты, которую хотите забронировать:\n" + "\n".join([room['id'] for room in rooms]))


@bot.message_handler(func=lambda message: True)
def handle_user_input(message):
    user_id = message.from_user.id

    if user_id not in user_data:
        start_handler(message)
        return

    if 'roomId' not in user_data[user_id]:
        user_data[user_id]['roomId'] = message.text
        bot.send_message(user_id, "Теперь введите дату и время заезда в формате дд.мм.гггг")
    elif 'inTime' not in user_data[user_id]:
        user_data[user_id]['inTime'] = message.text
        bot.send_message(user_id, "Теперь введите дату и время выезда в формате дд.мм.гггг")
    elif 'outTime' not in user_data[user_id]:
        user_data[user_id]['outTime'] = message.text
        bot.send_message(user_id, "Теперь введите ваше имя")
    elif 'clientName' not in user_data[user_id]:
        user_data[user_id]['clientName'] = message.text
        bot.send_message(user_id, "Теперь введите вашу фамилию")
    elif 'clientSurname' not in user_data[user_id]:
        user_data[user_id]['clientSurname'] = message.text
        bot.send_message(user_id, "Теперь введите ваш номер телефона")
    elif 'clientPhone' not in user_data[user_id]:
        user_data[user_id]['clientPhone'] = message.text
        bot.send_message(user_id, "Теперь введите вашу почту")
    elif 'clientEmail' not in user_data[user_id]:
        user_data[user_id]['clientEmail'] = message.text
        response = requests.post(
            f"http://{BOOKING_SERVICE_HOST}:{BOOKING_SERVICE_PORT}/booking-service/api/add-reservation",
            json=user_data[user_id]
        )
        if response.status_code != 200:
            bot.send_message(user_id, "Сервис временно не доступен")
            logger.info("Booking service responded with status" + str(response.status_code) + ". Content:" + str(
                response.content))
            return
        del user_data[user_id]
        try:
            url = response.json()['PaymentUrl']
        except:
            bot.send_message(user_id, "Сервис временно не доступен")
            logger.info("Booking service responded with status" + str(response.status_code) + ". Content:" + str(
                response.content))
            return
        user_data[user_id] = {'paymentUrl': url}
        markup = types.InlineKeyboardMarkup(
            keyboard=[
                [
                    types.InlineKeyboardButton(
                        text='Оплатить',
                        callback_data='pay'
                    )
                ]
            ]
        )
        bot.send_message(user_id, "Ссылка для оплаты\n" + str(url), reply_markup=markup)
    else:
        bot.send_message(user_id, "Вы уже отправили данные. Оплатите бронь")


@bot.callback_query_handler(func=lambda call: call == 'pay')
def pay_callback(callback: types.CallbackQuery):
    user_id = callback.from_user
    url = user_data[user_id]['paymentUrl']
    response = requests.post(url)
    logger.info(response.status_code)
    if response.status_code == 200:
        bot.edit_message_text(chat_id=callback.message.chat.id, message_id=callback.message.message_id, text="Номер успешно забронирован")
    else:
        logger.info(f"Payment system responded with status {response.status_code}: {response.content}")
        bot.edit_message_text(chat_id=callback.message.chat.id, message_id=callback.message.message_id, text="Сервис временно не доступен")


@app.post("/notify")
async def create_item(request: NotifyRequestModel):
    try:
        bot.send_message(request.tgId, request.message)
    except ApiTelegramException:
        return {"status": "error", "message": f"no chat with user {request.tgId}"}
    return {"status": "success", "message": "Message sent"}


def start_bot():
    logging.info("Starting bot")
    bot.polling(none_stop=True, interval=0)


def start_server():
    uvicorn.run(app, host=BOT_HOST, port=BOT_PORT)


if __name__ == "__main__":
    bot_process = multiprocessing.Process(target=start_bot)
    server_process = multiprocessing.Process(target=start_server)
    bot_process.start()
    server_process.start()
    bot_process.join()
    server_process.join()
