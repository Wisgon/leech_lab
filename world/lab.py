# neure science lab, to do some experiment like habituation or sensitization and so on.

import time
import json

from websockets.sync.client import connect
import utils


@utils.connect_backend_decorator
def habituation_experiment(websocket=None, stimulate_times=10, duration_time=5):
    for i in range(stimulate_times):
        print("sending habituation_experiment info****")
        websocket.send(
            json.dumps(
                {
                    "event": "experiment",
                    "message": {
                        "action": "stimulate",
                        "action_detail": {
                            "stimulate_skin_prefix": "skin_common_biggerPress_leftMiddleUp",
                            "stimulate_skin_neure_number": 100,  # how many neure activate in this stimulate
                            "stimulate_later_skin_prefix": "",  # for habituation, here is ""
                            "stimulate_later_skin_number": 0,
                        },
                    },
                }
            )
        )
        time.sleep(duration_time)


if __name__ == "__main__":
    habituation_experiment(stimulate_times=1, duration_time=2)
