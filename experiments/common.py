# from test_python_pkg import *


def test_yield():
    print("before yield")
    yield 3


def test_decorator(func):
    def wrapper(*args, **kwargs):
        a = 5
        func(a, *args, **kwargs)

    return wrapper


@test_decorator
def decorator_caller(a):
    print(a)


if __name__ == "__main__":
    # print("!!!aaa:", pkg1.aaa)
    # a = test_yield()
    # print(next(a))
    # print("end")
    decorator_caller()
