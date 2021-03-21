from datetime import datetime, tzinfo, timezone
from pymongo import MongoClient
from pprint import pprint

'''
实验四：地区销量top1
计算第一季度每个州（state）销量最多的sku第一名

- 只算complete订单
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
                    'status': 'completed'
                }
            }, {
                # 步骤2：按订单行展开
                '$unwind': {
                    'path': '$orderLines'
                }
            }, {
                # 步骤3：按sku汇总
                '$group': {
                    '_id': {
                        'state': '$state',
                        'sku': '$orderLines.sku'
                    }, 
                    'count': {
                        '$sum': '$orderLines.qty'
                    }
                }
            }, {
                # 步骤4：按州和销量排序
                '$sort': {
                    '_id.state': 1, 
                    'count': -1
                }
            }, {
                # 步骤4：取每个州top1
                '$group': {
                    '_id': '$_id.state', 
                    'sku': {
                        '$first': '$_id.sku'
                    }, 
                    'count': {
                        '$first': '$count'
                    }
                }
            }, {
                '$sort': {
                    '_id': -1
                }
            }
        ]
    )

    pprint(list(result))

    '''
    [{'_id': 'Wyoming', 'count': 183, 'sku': '8181'},
    {'_id': 'Wisconsin', 'count': 195, 'sku': '9684'},
    {'_id': 'West Virginia', 'count': 170, 'sku': '9376'},
    {'_id': 'Washington', 'count': 185, 'sku': '422'},
    {'_id': 'Virginia', 'count': 173, 'sku': '1570'},
    {'_id': 'Vermont', 'count': 181, 'sku': '4049'},
    {'_id': 'Utah', 'count': 180, 'sku': '4525'},
    {'_id': 'Texas', 'count': 198, 'sku': '1584'},
    {'_id': 'Tennessee', 'count': 207, 'sku': '2839'},
    {'_id': 'South Dakota', 'count': 199, 'sku': '9941'},
    {'_id': 'South Carolina', 'count': 161, 'sku': '4525'},
    {'_id': 'Rhode Island', 'count': 177, 'sku': '9657'},
    {'_id': 'Pennsylvania', 'count': 193, 'sku': '9614'},
    {'_id': 'Oregon', 'count': 254, 'sku': '2065'},
    {'_id': 'Oklahoma', 'count': 184, 'sku': '1474'},
    {'_id': 'Ohio', 'count': 192, 'sku': '7490'},
    {'_id': 'North Dakota', 'count': 243, 'sku': '2411'},
    {'_id': 'North Carolina', 'count': 205, 'sku': '2850'},
    {'_id': 'New York', 'count': 204, 'sku': '6056'},
    {'_id': 'New Mexico', 'count': 170, 'sku': '5966'},
    {'_id': 'New Jersey', 'count': 232, 'sku': '3254'},
    {'_id': 'New Hampshire', 'count': 191, 'sku': '126'},
    {'_id': 'Nevada', 'count': 180, 'sku': '1341'},
    {'_id': 'Nebraska', 'count': 171, 'sku': '8556'},
    {'_id': 'Montana', 'count': 259, 'sku': '4521'},
    {'_id': 'Missouri', 'count': 192, 'sku': '8049'},
    {'_id': 'Mississippi', 'count': 184, 'sku': '8100'},
    {'_id': 'Minnesota', 'count': 164, 'sku': '1225'},
    {'_id': 'Michigan', 'count': 179, 'sku': '5963'},
    {'_id': 'Massachusetts', 'count': 196, 'sku': '3778'},
    {'_id': 'Maryland', 'count': 187, 'sku': '3679'},
    {'_id': 'Maine', 'count': 233, 'sku': '7011'},
    {'_id': 'Louisiana', 'count': 213, 'sku': '484'},
    {'_id': 'Kentucky', 'count': 226, 'sku': '2061'},
    {'_id': 'Kansas', 'count': 182, 'sku': '5056'},
    {'_id': 'Iowa', 'count': 167, 'sku': '7960'},
    {'_id': 'Indiana', 'count': 268, 'sku': '5396'},
    {'_id': 'Illinois', 'count': 188, 'sku': '2244'},
    {'_id': 'Idaho', 'count': 193, 'sku': '4670'},
    {'_id': 'Hawaii', 'count': 191, 'sku': '6256'},
    {'_id': 'Georgia', 'count': 192, 'sku': '7106'},
    {'_id': 'Florida', 'count': 158, 'sku': '1715'},
    {'_id': 'Delaware', 'count': 183, 'sku': '769'},
    {'_id': 'Connecticut', 'count': 165, 'sku': '9189'},
    {'_id': 'Colorado', 'count': 182, 'sku': '7298'},
    {'_id': 'California', 'count': 195, 'sku': '905'},
    {'_id': 'Arkansas', 'count': 171, 'sku': '2669'},
    {'_id': 'Arizona', 'count': 189, 'sku': '1615'},
    {'_id': 'Alaska', 'count': 195, 'sku': '9847'},
    {'_id': 'Alabama', 'count': 166, 'sku': '781'}]
    '''

    client.close()