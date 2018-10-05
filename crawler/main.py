import asyncio
import random
import time
from datetime import datetime, timedelta

from search import fetch_listing_urls
from listing import worker


async def main():
    url = "https://losangeles.craigslist.org/search/wst/apa?sort=date&availabilityMode=0&max_price=3000"
    lastseen_dt = datetime.now() - timedelta(1)

    listing_queue = asyncio.Queue()
    fetch_listing_urls(url, lastseen_dt, listing_queue)

    listing_tasks = []
    for i in range(16):
        task = asyncio.create_task(worker(f"worker-{i}", listing_queue))
        listing_tasks.append(task)

    await listing_queue.join()

    for task in listing_tasks:
        task.cancel()

    # Wait until all worker tasks are cancelled.
    await asyncio.gather(*listing_tasks, return_exceptions=True)
    print("done")


asyncio.run(main())
