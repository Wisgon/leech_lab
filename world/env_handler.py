import time
import json


async def send_env_info(websocket):
    while True:
        print("sending env info****")
        await websocket.send(
            json.dumps(
                {
                    "event": "env_info",
                    "message": {"temperature": 34.2, "touch_press": "soft"},
                }
            )
        )
        time.sleep(600)
