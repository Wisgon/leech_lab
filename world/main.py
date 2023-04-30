import asyncio
import websockets
import json

import utils
import env_handler

clients = {}
loop = asyncio.new_event_loop()


class QueryParamProtocol(websockets.WebSocketServerProtocol):
    async def process_request(self, path, headers):
        params = utils.get_query_parameters(path)
        self.params = params


async def handler(websocket):
    global loop, clients
    params = websocket.params
    users = params.get("user")
    print("here someone comes: user:", users)
    if users is None or len(users) == 0:
        print("users is not exists")
        await websocket.send(
            json.dumps(
                {"event": "error", "message": f"user {str(users)} is not exists"}
            )
        )
        return
    user = users[0]
    if not user in ["front", "back"]:
        print("user:", str(user), " username is wrong")
        await websocket.send(
            json.dumps({"event": "error", "message": f"user {str(user)} is wrong"})
        )
        return
    else:
        clients[user] = websocket
    if user == "back":
        loop.create_task(env_handler.send_env_info(clients[user]))
    while True:
        try:
            data = await websocket.recv()
            send_data = json.loads(data)
            # back send to front and front send to back
            if user == "back":
                send_websocket = clients.get("front")
            elif user == "front":
                send_websocket = clients.get("back")
            if send_websocket is None:
                print("front or back is not connect!")
                return
            await send_websocket.send(json.dumps(send_data))
        except Exception as e:
            if type(e) == websockets.exceptions.ConnectionClosedOK:
                print("client closing")
            else:
                print(e)
            return


async def server8001():
    # todo: restart when crash
    async with websockets.serve(handler, "", 8001, create_protocol=QueryParamProtocol):
        print("start webdocket~~~8001")
        await asyncio.Future()  # run forever


if __name__ == "__main__":
    asyncio.set_event_loop(loop)
    loop.create_task(server8001())
    loop.run_forever()
