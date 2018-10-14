from collections import namedtuple
from datetime import datetime, timedelta
from urllib.request import urlopen
from queue import Queue

from bs4 import BeautifulSoup


BASE_URL = "https://losangeles.craigslist.org"
RESULTROW_CLASS = "result-row"
RESULTTITLE_CLASS = "result-title"
RESULTDATE_CLASS = "result-date"
RESULTPRICE_CLASS = "result-price"
RESULTIMAGE_CLASS = "thumb"
RANGETO_CLASS = "rangeTo"
TOTALCOUNT_CLASS = "totalcount"
NEXT_CLASS = "next"
EXPECTED_COUNT = 120

Listing = namedtuple("Listing", ["id", "url"])


def fetch_listing_urls(base_url, lastseen_dt, listing_queue):
    i = 0
    while True:
        listings_seen = 0
        url = base_url + "&s={}".format(str(i * EXPECTED_COUNT))

        html = urlopen(url).read()
        for listing in parse_listings(html, lastseen_dt):
            listings_seen += 1
            listing_queue.put_nowait(listing)

        if listings_seen < EXPECTED_COUNT:
            break
        i += 1


def parse_listings(html, min_time):
    try:
        soup = BeautifulSoup(html, "html.parser")
        for row in soup.find_all("li", {"class", RESULTROW_CLASS}):
            id = row["data-pid"]
            title_info = row.find("a", {"class", RESULTTITLE_CLASS})
            url = title_info["href"]
            time = row.find("time", {"class", RESULTDATE_CLASS})["datetime"]
            if to_datetime(time) <= min_time:
                continue
            yield Listing(id=id, url=url)
    except Exception as e:
        print(e)


def to_datetime(time):
    return datetime.strptime(time, "%Y-%m-%d %H:%M")


if __name__ == "__main__":
    queue = Queue()
    min_time = datetime.now() - timedelta(1)
    url = "https://losangeles.craigslist.org/search/wst/apa?sort=date&availabilityMode=0&max_price=3000"
    fetch_listing_urls(url, min_time, queue)
