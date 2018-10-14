import aiohttp
import re
from collections import namedtuple

from bs4 import BeautifulSoup


MAP_ID = "map"
MAP_LNG = "data-longitude"
MAP_LAT = "data-latitude"
DATE_AVAILABLE = "data-date"

ListingDetails = namedtuple(
    "ListingDetails",
    [
        "id",
        "address",
        "url",
        "price",
        "title",
        "images",
        "body",
        "bedrooms",
        "bathrooms",
        "sqft",
        "available_date",
        "cats",
        "dogs",
        "housing_type",
        "wd_type",
        "parking_type",
        "lng",
        "lat",
    ],
)


async def worker(name, listing_queue, parsed_queue):
    while True:
        listing = await listing_queue.get()
        async with aiohttp.ClientSession() as session:
            html = await fetch(session, listing.url)
            parsed_queue.put_nowait(parse_listing(html, listing))
        listing_queue.task_done()


async def fetch(session, url):
    async with session.get(url) as response:
        return await response.text()


def parse_listing(html, listing):
    soup = BeautifulSoup(html, "html.parser")
    address = get_address(soup)
    title = soup.title.string
    price = get_price(soup)
    images = get_images(soup)
    body = get_body(soup)
    details = get_details(soup)
    lng, lat = get_cords(soup)

    return ListingDetails(
        listing.id,
        address,
        listing.url,
        price,
        title,
        images,
        body,
        details.bedrooms,
        details.bathrooms,
        details.sqft,
        details.available_date,
        details.cats,
        details.dogs,
        details.housing_type,
        details.wd_type,
        details.parking_type,
        lng,
        lat,
    )


def get_address(soup):
    address = soup.find("div", {"class": "mapaddress"}).text
    return address


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


HOUSING_TYPES = [
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
]
WD_TYPES = [
    "w/d in unit",
    "w/d hookups",
    "laundry in bldg",
    "laundry on site",
    "no laundry on site",
]
PARKING_TYPES = [
    "off-street parking",
    "valet parking",
    "street parking",
    "no parking",
    "carport",
    "attached garage",
    "detached garage",
]


Details = namedtuple(
    "Details",
    [
        "bedrooms",
        "bathrooms",
        "sqft",
        "available_date",
        "cats",
        "dogs",
        "housing_type",
        "wd_type",
        "parking_type",
    ],
)


def get_details(soup):
    attr_groups = soup.findAll("p", {"class": "attrgroup"})

    # lazy and bad programming here
    major_details = [y.text.strip() for x in attr_groups[:1] for y in x.findAll("span")]
    bedrooms = 0
    bathrooms = 0.0
    sqft = 0.0
    available_date = ""
    for detail in major_details:
        if "BR" in detail:
            bed_bath = get_numbers(detail)
            if len(bed_bath) == 2:
                bedrooms = int(bed_bath[0])
                bathrooms = float(bed_bath[1])
        elif "ft2" in detail:
            sqft = float(get_numbers(detail)[0])
        else:
            available_date = detail

    details = [y.text.strip() for x in attr_groups[1:] for y in x.findAll("span")]
    cats = "cats are OK - purrr" in details
    dogs = "dogs are OK - wooof" in details
    housing = ""
    washer_dryer = ""
    parking = ""

    for detail in details:
        if detail in HOUSING_TYPES:
            housing = detail
        elif detail in WD_TYPES:
            washer_dryer = detail
        elif detail in PARKING_TYPES:
            parking = detail

    return Details(
        bedrooms,
        bathrooms,
        sqft,
        available_date,
        cats,
        dogs,
        housing,
        washer_dryer,
        parking,
    )


def get_numbers(string):
    numberfromstring = re.findall("\d+\.\d+|\d+", string)
    return numberfromstring

