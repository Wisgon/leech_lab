import pymysql


def get_db():
    db = pymysql.connect(
        host="localhost",
        port=3306,
        user="root",
        password="123456",
        database="brain",
    )
    return db
