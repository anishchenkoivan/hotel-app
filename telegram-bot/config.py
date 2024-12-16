from dotenv import load_dotenv
from os import getenv


load_dotenv()
TOKEN = getenv("BOT_TOKEN")
if TOKEN is None:
    raise ValueError("No BOT_TOKEN variable set in .env")
BOT_HOST = getenv("BOT_HOST")
if BOT_HOST is None:
    raise ValueError("No BOT_HOSTNAME variable set in .env")
BOT_PORT = int(getenv("BOT_PORT"))
if BOT_PORT is None:
    raise ValueError("No BOT_PORT variable set in .env")
BOOKING_SERVICE_HOST = getenv("BOOKING_SERVICE_HOST")
if BOOKING_SERVICE_HOST is None:
    raise ValueError("No BOOKING_SERVICE_HOST variable set in .env")
BOOKING_SERVICE_PORT = int(getenv("BOOKING_SERVICE_PORT"))
if BOOKING_SERVICE_PORT is None:
    raise ValueError("No BOOKING_SERVICE_PORT variable set in .env")
HOTEL_SERVICE_HOST = getenv("HOTEL_SERVICE_HOST")
if HOTEL_SERVICE_HOST is None:
    raise ValueError("No HOTEL_SERVICE_HOST variable set in .env")
HOTEL_SERVICE_PORT = int(getenv("HOTEL_SERVICE_PORT"))
if HOTEL_SERVICE_PORT is None:
    raise ValueError("No HOTEL_SERVICE_PORT variable set in .env")