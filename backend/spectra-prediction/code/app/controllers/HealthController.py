from flask_restful import Resource

class HealtController(Resource):
    def get(self):
        return {'status': 'running'}, 200