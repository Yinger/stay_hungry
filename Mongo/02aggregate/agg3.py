from datetime import datetime, tzinfo, timezone
from pymongo import MongoClient
from pprint import pprint
# Requires the PyMongo package.
# https://api.mongodb.com/python/current

'''
实验三：计算月销量
计算前半年每个月的销售额和总订单数。

- 不算运费
- 不算取消（cancelled）状态的订单
'''
if __name__ == "__main__":
    uri = "mongodb://127.0.0.1:27018/"
    client = MongoClient(uri)
    pprint(client)
    result = client['mock']['orders'].aggregate([
        {
            # 步骤1：匹配条件
            '$match': {
                'status': {
                    '$ne': 'cancelled'
                }, 
                'orderDate': {
                    '$gte': datetime(2019, 1, 1, 0, 0, 0, tzinfo=timezone.utc), 
                    '$lt': datetime(2019, 7, 1, 0, 0, 0, tzinfo=timezone.utc)
                }
            }
        }, {
            # 步骤2：取出年月
            '$project': {
                'month': {
                    '$dateToString': {
                        'date': '$orderDate', 
                        'format': '%G年%m月'
                    }
                }, 
                'total': 1
            }
        }, {
            # 步骤3：按年月分组汇总
            '$group': {
                '_id': '$month', 
                'total': {
                    '$sum': '$total'
                }, 
                'count': {
                    '$sum': 1
                }
            }
        }
    ])

    pprint(list(result))
    '''
    [{'_id': '2019年03月', 'count': 8167, 'total': Decimal128('3574185')},
     {'_id': '2019年04月', 'count': 8038, 'total': Decimal128('3551291')},
     {'_id': '2019年02月', 'count': 7387, 'total': Decimal128('3258201')},
     {'_id': '2019年05月', 'count': 8163, 'total': Decimal128('3590503')},
     {'_id': '2019年01月', 'count': 8249, 'total': Decimal128('3620936')},
     {'_id': '2019年06月', 'count': 7942, 'total': Decimal128('3496645')}]
    '''

    client.close()