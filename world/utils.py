from urllib.parse import parse_qs, urlparse, unquote


def get_query_parameters(url):
    url = unquote(url)
    parsed_url = urlparse(url)
    return parse_qs(
        parsed_url.query
    )  # format is like {"some_key": ["some_value, xxxx"]}


if __name__ == "__main__":
    url = "/%3Fuser=back"
    param = get_query_parameters(url)
    print(param, type(param))
