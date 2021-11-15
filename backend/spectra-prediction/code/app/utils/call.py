from json.decoder import JSONDecodeError
import requests
import urllib

class Call:
    def __init__(self, url: str) -> None:
        self.url = url

    def _get(self, endpoint, data = None) -> dict:
        try:
            response = requests.get(
                f"{self.url}/{endpoint}",
                data=data if data else {},
                headers={
                    "Content-Type": "application/json",
                }
            )
            if response.status_code == 401:
                raise ValueError("Não autorizado")

            data = response.json()
            return data
        except JSONDecodeError as Err:
            raise ConnectionError(f"Error to call endpoint {endpoint} - {Err}")
        except ValueError as Err:
            raise ConnectionError(f"Error to call endpoint {endpoint} - {Err}")

    def _post(self, endpoint, data = None) -> dict:
        try:
            response = requests.post(
                f"{self.url}/{endpoint}",
                json=data if data else {},
                headers={
                    "Content-Type": "application/json",
                }
            )
            if response.status_code == 401:
                raise ValueError("Não autorizado")

            data = response.json()
            return data
        except JSONDecodeError as Err:
            raise ConnectionError(f"Error to call endpoint {endpoint} - {Err}")
        except ValueError as Err:
            raise ConnectionError(f"Error to call endpoint {endpoint} - {Err}")