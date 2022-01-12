import time

import redis

from rate_crawler import RateCrawler

redis_pool = redis.ConnectionPool(
    host="localhost",
    port=6379,
    db=0,
    decode_responses=True
)
r = redis.StrictRedis(connection_pool=redis_pool)


def main():
    crawler = RateCrawler()
    rate_dict = crawler.get_rate_dict()
    mapping_rate_dict = crawler.mapping_rate(rate_dict)
    for i in mapping_rate_dict.items():
        print(i[0])
        print(i[1])
        r.setex(i[0], 86400, i[1])


if __name__ == "__main__":
    start_time = time.time()
    main()
    print(f"runtime === {round((time.time() - start_time)/60)} "
        f"minutes {round((time.time() - start_time)%60, 2)} seconds ===")