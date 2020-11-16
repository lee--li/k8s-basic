import os
import math
from flask import Flask, request
app = Flask(__name__)

port = 8080
if "PORT" in os.environ:
    port = os.environ["PORT"]

@app.route('/hello')
def hello():
    #version = os.environ.get('SERVICE_VERSION')
    return 'Hello python \n'


@app.route('/health')
def health():
    return 'Helloworld is healthy', 200


if __name__ == "__main__":
    app.run(host='0.0.0.0', port=port,threaded=True)