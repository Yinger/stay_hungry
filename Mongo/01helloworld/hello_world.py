from pymongo import MongoClient
from pprint import pprint

if __name__ == "__main__":
    uri = "mongodb://127.0.0.1:27018/"
    client = MongoClient(uri)
    pprint(client)


    db = client["eshop"]
    user_coll = db["users"]

    new_user = {
        "username": "tina",
        "password": "xxxx",
        "email": "123456@chiiki.com"
    }
    result = user_coll.insert_one(new_user)
    pprint(result)

    result = user_coll.find_one()
    pprint(result)

    result = user_coll.update_one({
        "username": "nina"
    }, {
        "$set": {
            "phone": "123456789"
        }
    })

    result = user_coll.find_one({ "username": "tina" })
    pprint(result)
    
    client.close()
