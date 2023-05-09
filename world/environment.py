import time
import json

from websockets.sync.client import connect


def send_env_info():
    time.sleep(5)
    with connect("ws://localhost:8001/?user=env") as websocket:
        while True:
            print("sending env info****")
            websocket.send(
                json.dumps(
                    {
                        "event": "env_info",
                        "message": {"temperature": 34.2, "touch_press": "soft"},
                    }
                )
            )
            time.sleep(3)


if __name__ == "__main__":
    send_env_info()
