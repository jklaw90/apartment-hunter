import asyncio
import random
import time
from datetime import datetime, timedelta

from search import fetch_listing_urls
from listing import worker


async def main():
    url = "https://losangeles.craigslist.org/search/wst/apa?sort=date&availabilityMode=0&max_price=3000"
    lastseen_dt = datetime.now() - timedelta(hours=1)

    listing_queue = asyncio.Queue()
    parsed_queue = asyncio.Queue()
    fetch_listing_urls(url, lastseen_dt, listing_queue)

    start_dt = datetime.now() 
    listing_tasks = []
    for i in range(16):
        task = asyncio.create_task(worker(f"worker-{i}", listing_queue, parsed_queue))
        listing_tasks.append(task)

    await listing_queue.join()

    for task in listing_tasks:
        task.cancel()

    # Wait until all worker tasks are cancelled.
    await asyncio.gather(*listing_tasks, return_exceptions=True)
    total_time = datetime.now() - start_dt
    print("done", total_time, parsed_queue.qsize())
    while not parsed_queue.empty():
        info = await parsed_queue.get()
        print(info)
        parsed_queue.task_done()


asyncio.run(main())
