import os
import asyncio
from datetime import datetime, timedelta

import grpc
import apthunter_pb2 as pb
from search import fetch_listing_urls
from listing import worker
from apthunter_pb2_grpc import WriterStub


async def main(addr, url):
    # add 7 hours diff for pst... TODO fix
    lastseen_dt = datetime.now() - timedelta(hours=1+7)

    listing_queue = asyncio.Queue()
    parsed_queue = asyncio.Queue()
    fetch_listing_urls(url, lastseen_dt, listing_queue)

    print("listings found:", listing_queue.qsize())
    start_dt = datetime.now()
    listing_tasks = []
    for i in range(1000):
        task = asyncio.create_task(worker(f"worker-{i}", listing_queue, parsed_queue))
        listing_tasks.append(task)

    await listing_queue.join()

    for task in listing_tasks:
        task.cancel()

    # Wait until all worker tasks are cancelled.
    await asyncio.gather(*listing_tasks, return_exceptions=True)
    total_time = datetime.now() - start_dt
    print("done", total_time, parsed_queue.qsize())

    with grpc.insecure_channel(addr) as channel:
        stub = WriterStub(channel)

        while not parsed_queue.empty():
            info = await parsed_queue.get()
            response = stub.CreateOrUpdate(
                pb.CreateOrUpdateRequest(
                    id=info.id,
                    url=info.url,
                    price=info.price,
                    title=info.title,
                    images=info.images,
                    body=info.body,
                    details=info.details,
                    lng=info.lng,
                    lat=info.lat,
                )
            )
            parsed_queue.task_done()

print("Starting crawler")
addr = os.environ.get("WRITER_ADDR", "localhost:9000")
url = "https://losangeles.craigslist.org/search/wst/apa?sort=date&availabilityMode=0&max_price=3000"
asyncio.run(main(addr, url))
