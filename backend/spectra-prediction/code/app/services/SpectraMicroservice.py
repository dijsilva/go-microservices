
from app.utils.call import Call

class SpectraMicroservice(Call):
    def __init__(self, url: str) -> None:
        super().__init__(url)


    def get_spectra_data(self, spectra_id: str) -> dict:
        spectra_response = self._get(f"spectra/{spectra_id}")
        return spectra_response

    def inform_prediction(self, spectra_id: str, data: dict) -> dict:
        response = self._post(f"prediction/{spectra_id}", data)
        return response