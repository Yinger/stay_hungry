from datetime import datetime, tzinfo, timezone
from pymongo import MongoClient
from pprint import pprint

'''
实验五：统计SKU销售件数
统计每个sku在第一季度销售的次数。

- 不算取消（cancelled）状态的订单；
- 按销售数量降序排列；
'''
if __name__ == "__main__":
    uri = "mongodb://127.0.0.1:27018/"
    client = MongoClient(uri)
    pprint(client)
    result = client['mock']['orders'].aggregate(
        [
            {
                # 步骤1：匹配条件
                '$match': {
                    'orderDate': {
                        '$gte': datetime(2019, 1, 1, 0, 0, 0, tzinfo=timezone.utc), 
                        '$lt': datetime(2019, 4, 1, 0, 0, 0, tzinfo=timezone.utc)
                    }, 
                    'status': {
                        '$ne': 'cancelled'
                    }
                }
            }, {
                # 步骤2：按订单行展开
                '$unwind': {
                    'path': '$orderLines'
                }
            }, {
                # 步骤3：按sku汇总
                '$group': {
                    '_id': '$orderLines.sku', 
                    'count': {
                        '$sum': '$orderLines.qty'
                    }
                }
            }, {
                '$sort': {
                    'count': -1
                }
            }
        ]
    )

    pprint(list(result))
    '''
    [{'_id': '4751', 'count': 2115},
    {'_id': '798', 'count': 1945},
    {'_id': '3863', 'count': 1913},
    {'_id': '2558', 'count': 1896},
    {'_id': '2049', 'count': 1859},
    {'_id': '7239', 'count': 1854},
    {'_id': '6687', 'count': 1847},
    {'_id': '3844', 'count': 1807},
    {'_id': '7602', 'count': 1805},
    {'_id': '7519', 'count': 1803}]
    ...
    '''
    client.close()