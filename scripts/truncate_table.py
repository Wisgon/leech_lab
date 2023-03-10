from utils.db import get_db


def truncate(table_name: str):
    try:
        with get_db() as db:
            with db.cursor() as cursor:
                cursor.execute(f"TRUNCATE TABLE {table_name}")
    except Exception as e:
        print("fail:", str(e))


if __name__ == "__main__":
    truncate("un")
