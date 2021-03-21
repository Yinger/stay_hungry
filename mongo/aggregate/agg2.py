from pymongo import MongoClient
from pprint import pprint
from datetime import datetime,timezone

'''
实验二：订单金额汇总

查询2019年第一季度（1月1日~3月31日）订单中
已完成（completed）状态的总金额和总数量
'''

if __name__ == "__main__":
    uri = "mongodb://127.0.0.1:27018/"
    client = MongoClient(uri)
    pprint(client)

    db = client["mock"]
    orders = db["orders"]

    result = orders.aggregate([
    {
        # 步骤1：匹配条件
        '$match': {
            'status': 'completed', 
            'orderDate': {
                '$gte': datetime(2019, 1, 1, 0, 0, 0, tzinfo=timezone.utc), 
                '$lt': datetime(2019, 4, 1, 0, 0, 0, tzinfo=timezone.utc)
            }
        }
    }, {
        # 步骤二：聚合订单总金额、总运费、总数量
        '$group': {
            '_id': None, 
            'total': {
                '$sum': '$total'
            }, 
            'shippingFee': {
                '$sum': '$shippingFee'
            }, 
            'count': {
                '$sum': 1
            }
        }
    }, {
        # 计算总金额
        '$project': {
            '_id': 0, 
            'grandTotal': {
                '$add': [
                    '$total', '$shippingFee'
                ]
            }, 
            'count': 1
        }
    }])

    pprint(list(result))
    # [{'count': 5875, 'grandTotal': Decimal128('2636376.00')}]

    client.close()