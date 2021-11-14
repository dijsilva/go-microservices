"""Run the application."""

import os, sys
sys.path.append(
    os.path.abspath(
        os.path.join(os.path.dirname(__file__),
        '..')
    ))

import os
from app import app
from app.configuration.envs import ApplicationEnvs

if __name__ == "__main__":
    app.run(debug=True, host='0.0.0.0', port=int(ApplicationEnvs.PORT))