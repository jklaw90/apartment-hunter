import aiohttp
from collections import namedtuple

from bs4 import BeautifulSoup


MAP_ID = "map"
MAP_LNG = "data-longitude"
MAP_LAT = "data-latitude"
DATE_AVAILABLE = "data-date"

WD_HOOKUPS = "w/d hookups"


ListingDetails = namedtuple(
    "ListingDetails",
    ["id", "url", "price", "title", "images", "body", "details", "lng", "lat"],
)

derp_set = set()


async def worker(name, listing_queue, parsed_queue):
    while True:
        listing = await listing_queue.get()
        async with aiohttp.ClientSession() as session:
            html = await fetch(session, listing.url)
            parsed_queue.put_nowait(parse_listing(html, listing))
        listing_queue.task_done()
        print(derp_set)


async def fetch(session, url):
    async with session.get(url) as response:
        return await response.text()


def parse_listing(html, listing):
    soup = BeautifulSoup(html, "html.parser")
    title = soup.title.string
    price = get_price(soup)
    images = get_images(soup)
    body = get_body(soup)
    details = get_details(soup)
    lng, lat = get_cords(soup)

    return ListingDetails(
        listing.id, listing.url, price, title, images, body, details, lng, lat
    )


def get_price(soup):
    prices = soup.findAll("span", {"class": "price"})

    if len(prices) is 1:
        price = prices[0].string.strip("$")
        return float(price)

    # shouldn't happen based on listings ive seen
    return 0


def get_images(soup):
    images = []
    thumbs = soup.find(id="thumbs")

    if thumbs:
        images = [x.get("href") for x in thumbs.findAll("a")]

    return images


def get_body(soup):
    body = soup.find(id="postingbody")

    if body:
        return body.text

    return ""


def get_cords(soup):
    map_div = soup.find(id=MAP_ID)

    if map_div:
        lng = map_div.get(MAP_LNG)
        lat = map_div.get(MAP_LAT)
        return float(lng), float(lat)

    return None, None


Details = namedtuple(
    "Details",
    ["bedrooms", "bathrooms", "", "title", "images", "body", "details", "lng", "lat"],
)


def get_details(soup):
    attr_groups = soup.findAll("p", {"class": "attrgroup"})
    details = [y.text.strip() for x in attr_groups[1:] for y in x.findAll("span")]
    return


# TODO
[
    "cats are OK - purrr",
    "dogs are OK - wooof",
    "furnished",
    "apartment",
    "condo",
    "cabin",
    "duplex",
    "flat",
    "house",
    "law",
    "loft",
    "townhouse",
    "manufactured",
    "living",
    "land",
    "w/d in unit",
    "w/d hookups",
    "laundry in bldg",
    "laundry on site",
    "no laundry on site",
    "off-street parking",
    "valet parking",
    "street parking",
    "no parking",
    "carport",
    "attached garage",
    "detached garage",
]
