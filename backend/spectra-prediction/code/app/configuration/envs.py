
import os

def getEnvs(envName: str) -> str:
    try:
        return os.environ[envName]
    except KeyError:
        raise KeyError(f"Env {envName} not defined")

class ApplicationEnvs(object):
    PORT                            = getEnvs('APP_PORT')
    SPECTRA_QUEUE_NAME              = getEnvs('SPECTRA_QUEUE_NAME')
    RABBITMQ_HOST                   = getEnvs('RABBITMQ_HOST')
    RABBITMQ_USER                   = getEnvs('RABBITMQ_USER')
    RABBITMQ_PASS                   = getEnvs('RABBITMQ_PASS')
    RABBITMQ_PORT                   = getEnvs('RABBITMQ_PORT')
    SPECTRA_MICROSERVICE_HOST       = getEnvs('SPECTRA_MICROSERVICE_HOST')
    

