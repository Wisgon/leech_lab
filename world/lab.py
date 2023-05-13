# neure science lab, to do some experiment like habituation or sensitization and so on.

import time
import json

import utils


@utils.connect_backend_decorator
def habituation_experiment(websocket=None, stimulate_times=20, duration_time=5):
    # 这个实验模仿习惯化，当某种无害刺激刺激越多次，生物体的反应对其越弱化
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


@utils.connect_backend_decorator
def sensitization_experiment(websocket=None):
    # 这个实验模仿敏感化，在习惯化后，对其进行电击，之前习惯化的神经元以及其他全部感觉神经元和运动神经元的连接会得到强化，任何再次触发的哪怕最微小的感觉都会引起运动
    # for i in range(0, 20):
    #     # 大约重复刺激20次能使其习惯化，这是衰减函数决定的
    #     print("sending habituation action****")
    #     websocket.send(
    #         json.dumps(
    #             {
    #                 "event": "experiment",
    #                 "message": {
    #                     "action": "stimulate",
    #                     "action_detail": {
    #                         "stimulate_skin_prefix": "skin_common_biggerPress_leftMiddleUp",
    #                         "stimulate_skin_neure_number": 100,  # how many neure activate in this stimulate
    #                         "stimulate_later_skin_prefix": "",  # for habituation, here is ""
    #                         "stimulate_later_skin_number": 0,
    #                     },
    #                 },
    #             }
    #         )
    #     )
    #     time.sleep(5)

    # 发送点击信号
    # websocket.send(
    #     json.dumps(
    #         {
    #             "event": "experiment",
    #             "message": {
    #                 "action": "stimulate",
    #                 "action_detail": {
    #                     "stimulate_skin_prefix": "skin_common_extremelyPress_rightBackUp",
    #                     "stimulate_skin_neure_number": 100,  # how many neure activate in this stimulate
    #                     "stimulate_later_skin_prefix": "",  # for habituation, here is ""
    #                     "stimulate_later_skin_number": 0,
    #                 },
    #             },
    #         }
    #     )
    # )
    # time.sleep(5)
    # 此时再发送一次普通触碰的信号，会观察到实验体敏感化的行动
    websocket.send(
        json.dumps(
            {
                "event": "experiment",
                "message": {
                    "action": "stimulate",
                    "action_detail": {
                        "stimulate_skin_prefix": "skin_common_normalPress_leftMiddleUp",
                        "stimulate_skin_neure_number": 50,  # how many neure activate in this stimulate
                        "stimulate_later_skin_prefix": "",  # for habituation, here is ""
                        "stimulate_later_skin_number": 0,
                    },
                },
            }
        )
    )


if __name__ == "__main__":
    # habituation_experiment(stimulate_times=1, duration_time=2)
    sensitization_experiment()
