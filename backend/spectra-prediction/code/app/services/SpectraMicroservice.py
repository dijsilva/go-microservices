
from app.utils.call import Call

class SpectraMicroservice(Call):
    def __init__(self, url: str, token: str) -> None:
        super().__init__(url, token)


    def getSpectraData(self, spectra_id: str) -> dict:
        spectraResponse = self._get(f"spectra/{spectra_id}")
        return spectraResponse