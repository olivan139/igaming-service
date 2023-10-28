# igaming-service

## Start
To start server, one should write this in command line in the root of the project:
```
make build
```
To launch tests, one should write:
```
make test
```

The environment file should contain **PORT** value. Otherwise the server will start on port 8888.

## Handlers
There are two handlers:
### /payoff
Type: POST


Body:
```
{
    "reels" : [
        ["A", "B", "C", "D", "E"],
        ["F", "A", "F", "B", "C"],
        ["D", "E", "A", "G", "A"]
    ],

    "lines" : [
        {
            "line": 1, 
            "positions": [
                {"row": 0, "col": 0},
                {"row": 1, "col": 1},
                {"row": 2, "col": 2},
                {"row": 1, "col": 3},
                {"row": 0, "col": 4}
            ]
        }
    ],
    "payouts" : [
        {
            "symbol": "A",
            "payout": [0, 0, 50, 100, 200]
        }
    ]
}
```
Response:
```
{
  "lines": [
    {
      "line": 1,
      "payout": 50
    }
  ],
  "total": 50
}
```
### /metric
Prometheus metrics:
1. **total_request_count** - total number of requests
2. **http_request_duration_seconds** - duration of requests
3. **bad_request_count** - count of bad requests
