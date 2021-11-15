from dotenv import load_dotenv
load_dotenv('../.env')
from flask import Flask
from flask_restful import Api

from app.routes.register import register_routes

from app.queue.Queue import Queue

app = Flask(__name__)

api = Api(app)

register_routes(api)

queue = Queue()
queue.receive()