from app.predictions.SpectraPrediction import SpectraPrediction
from app.services.SpectraMicroservice import SpectraMicroservice
from app.configuration.envs import ApplicationEnvs
from datetime import datetime

class MessageHandler:
    def __init__(self) -> None:
        self.model = SpectraPrediction()
        self.spectra_service = SpectraMicroservice(
            ApplicationEnvs.SPECTRA_MICROSERVICE_HOST,
        )
    
    def pass_message(self, message: str) -> None:
        self.spectra_id = message

        spectraData = self.spectra_service.get_spectra_data(message)

        prediction, prediction_number = self.model.predict(spectraData)

        prediction_data = {
            'prediction_date': datetime.isoformat(datetime.now()),
            'prediction_string': prediction,
            'prediction_number': int(prediction_number),
        }

        print(f'sending prediction {prediction_data}')

        self.spectra_service.inform_prediction(self.spectra_id, data=prediction_data)
        
