from app.predictions.SpectraPrediction import SpectraPrediction
from app.services.SpectraMicroservice import SpectraMicroservice
from app.configuration.envs import ApplicationEnvs

class MessageHandler:
    def __init__(self) -> None:
        self.model = SpectraPrediction()
        self.spectraService = SpectraMicroservice(
            ApplicationEnvs.SPECTRA_MICROSERVICE_HOST,
            f"Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImRpZWdvanNpbHZhYnJAZ21haWwuY29tIiwiZXhwIjoxNjM2OTMxNDMyLCJpZCI6IjBlZjg2ZDZjLTYxMTgtNDc3Mi1iNzNiLTY4OGZmODJjOTJjOSIsIm5hbWUiOiJEaWVnbyIsInByb2ZpbGUiOiJVU0VSIiwidG9rZW5LaW5kIjoiTE9HSU5fVVNFUiJ9.04JgVxDPcffWUDGFJWAWkoizlhk_cjhCu9ATknOzlQU"
        )
    
    def pass_message(self, message: str) -> None:
        self.spectra_id = message

        spectraData = self.spectraService.getSpectraData(message)

        self.model.predict(spectraData)
        
