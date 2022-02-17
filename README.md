# Currency Converter [under development]
Convert your amount to multiple other currencies. 

<summary><h2 style="display: inline-block">Table of Contents</h2></summary>
<ol>
  <li><a href="#overview">Overview</a></li>
  <li><a href="#feature">Feature</a></li>
  <li><a href="#prerequisites">Prerequisites</a></li>
  <li><a href="#usage">Usage</a></li>
  <li><a href="#contact">Contact</a></li>
</ol>

<br>

## Overview
Currency Converter is a tool implemented in Go and Python, it provides 8 kinds of currency to exchange from one to another.

* TWD (Taiwan Dollar)
* USD (American Dollar)
* JPY (Japanese Yen)
* KRW (Korean Won)
* CNY (China Yen)
* HKD (HongKong Dollar)
* SGD (Singapore Dollar)
* EUR (Euro)


Foreign exchange rate reference:   
https://rate.bot.com.tw/xrt?Lang=en-US

## Feature
+ [x] rate crawler collect data to Redis
+ [x] currency rate Get method for getting currency rate data.
+ [x] currency rate POST method for convert amount from current currency to another currency.


## Prerequisites
Python3 and Golang runtime environment   
Redis server   


## Build and Run
Create Python virtualenv for isolating environment. (Optional)
```
virtualenv .
. bin/activate
```

Install python dependencies.
```
pip3 install -r rate_crawler/requirements.txt
```

Run python crawler and make sure your Redis server is working.
The rate data will be stored in Redis a day and expire after 24 hours later.
Add this crawling mission to cronjob if you want to run this project in the long term.
```
python rate_crawler/main.py
```

Build application and API serving.
```
go build -o currency_converter
./currency_converter
```


## Usage
1. Get one kind of currency with its exchange_rate list.
    ```
    curl localhost:8080/rate/TWD

    # example response
    {
        "currency_type": "TWD",
        "exchange_rate": {
            "CNY": "0.234",
            "EUR": "0.033",
            "HKD": "0.293",
            "JPY": "4.286",
            "KRW": "46.795",
            "SGD": "0.05",
            "TWD": "1.0",
            "USD": "0.037"
        },
        "target_type": ""
    }
    ```

    If you only need specify currency exchange_rate, add target_type in query behind.
    ```
    curl localhost:8080/rate/TWD?target_type=USD

    # example response
    {"currency_type":"TWD","exchange_rate":"0.037","target_type":"USD"}
    ```

2. convert amount of one currency to another currency.
    ```
    curl -X POST localhost:8080/rate -H "Content-Type: application/json" -d '{"currency_type": "TWD","target_type": "USD", "amount": "10000"}'

    # example response
    {"message":"exchange complete.","result":370}
    ```



## Contact
---
email: adelberteng@gmail.com