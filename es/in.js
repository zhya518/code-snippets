{
    "query": {
        "bool": {
            "must": [
                {
                    "match": {
                        "age": 28
                    }
                },
                {
                    "terms": {
                        "work": [
                            "engineer",
                            "office lady",
                            "teacher"
                        ]
                    }
                }
            ]
        }
    }
}
