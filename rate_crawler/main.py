import time

import redis

from rate_crawler import RateCrawler
from config import Config
conf = Config.load(env="db")


redis_pool = redis.ConnectionPool(
    host=conf.get("redis_endpoint"),
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
        for j in i[1].items():
            r.hset(i[0], j[0], j[1])
            r.expire(i[0], 86400)


if __name__ == "__main__":
    start_time = time.time()
    main()
    print(f"runtime === {round((time.time() - start_time)/60)} minutes"
        f" {round((time.time() - start_time)%60, 2)} seconds ===")