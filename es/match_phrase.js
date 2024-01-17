{
    "query": {
        "bool": {
            "must": [
                {
                    "match_phrase": {
                        "name": {
                            "query": "房间"
                        }
                    }
                },
                {
                    "range": {
                        "update_time": {
                            "from": 1634882830,
                            "include_lower": true,
                            "include_upper": true,
                            "to": null
                        }
                    }
                },
                {
                    "range": {
                        "disband_expire_second": {
                            "from": null,
                            "include_lower": true,
                            "include_upper": true,
                            "to": 1650434830
                        }
                    }
                },
                {
                    "term": {
                        "channel_type": 2
                    }
                }
            ],
            "must_not": [
                {
                    "term": {
                        "app_name": "olaparty"
                    }
                },
                {
                    "term": {
                        "app_name": "abc"
                    }
                }
            ]
        }
    }
}
