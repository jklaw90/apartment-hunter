import aiohttp
import asyncio


async def worker(name, listing_queue):
    while True:
        listing = await listing_queue.get()
        async with aiohttp.ClientSession() as session:
            html = await fetch(session, listing.url)
            print(html)
            # TODO parse

        listing_queue.task_done()


async def fetch(session, url):
    async with session.get(url) as response:
        return await response.text()
