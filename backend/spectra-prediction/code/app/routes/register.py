
from app.controllers.HealthController import HealtController

def register_routes(api_instance):
    api_instance.add_resource(HealtController, '/')