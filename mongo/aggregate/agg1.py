from pymongo import MongoClient
from pprint import pprint

'''
实验一：总销量
计算到目前为止的总销售额

无论订单状态
不限制时间范围
不算运费
'''

if __name__ == "__main__":
    uri = "mongodb://127.0.0.1:27018/"
    client = MongoClient(uri)
    pprint(client)

    db = client["mock"]
    orders = db["orders"]

    result = orders.aggregate([
        {
            "$group":{
                "_id": "null",
                "total": {"$sum": "$total"}
            }
        }
    ]
    )

    pprint(list(result))
    # [{'_id': 'null', 'total': Decimal128('44019609')}]

    client.close()