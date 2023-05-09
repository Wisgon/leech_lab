from urllib.parse import parse_qs, urlparse, unquote
from websockets.sync.client import connect


def get_query_parameters(url):
    url = unquote(url)
    parsed_url = urlparse(url)
    return parse_qs(
        parsed_url.query
    )  # format is like {"some_key": ["some_value, xxxx"]}


def connect_backend_decorator(func):
    def wrapper(*args, **kwargs):
        with connect("ws://localhost:8001/?user=env") as websocket:
            result = func(websocket, *args, **kwargs)
            return result

    return wrapper


if __name__ == "__main__":
    url = "/%3Fuser=back"
    param = get_query_parameters(url)
    print(param, type(param))
