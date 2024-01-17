{
    "query": {
        "bool": {
            "must": [
                {
                    "range": {
                        "LastOnlineTime": {
                            "gte": 100,
                            "lte": 1000
                        }
                    }
                },
                {
                    "match": {
                        "AppID": "5"
                    }
                }
            ]
        }
    },
    "sort": {
        "LastOnlineTime": "desc"
    },
    "size": 20
}
