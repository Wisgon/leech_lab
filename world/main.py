import asyncio

import websockets


async def handler(websocket):
    while True:
        try:
            message = await websocket.recv()
            print(message)
            await websocket.send(f"From server {message}")
        except Exception as e:
            if type(e) == websockets.exceptions.ConnectionClosedOK:
                print("client closing")
            else:
                print(e)
            return


async def main():
    async with websockets.serve(handler, "", 8001):
        print("start webdocket~~~")
        await asyncio.Future()  # run forever


if __name__ == "__main__":
    asyncio.run(main())
