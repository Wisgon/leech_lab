# from test_python_pkg import *


def test_yield():
    print("before yield")
    yield 3


if __name__ == "__main__":
    # print("!!!aaa:", pkg1.aaa)
    a = test_yield()
    print(next(a))
    print("end")
