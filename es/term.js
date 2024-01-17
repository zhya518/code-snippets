{
    "query": {
        "bool": {
            "must": {
                "term": {
                    "vcid": "ja693184"
                }
            },
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
