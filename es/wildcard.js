{
  "query": {
    "wildcard": {
      "name": {
        "value": "test*",
        "boost": 1.0,
        "rewrite": "constant_score"
      }
    }
  }
}
