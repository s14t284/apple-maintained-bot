import os
import logging
from urllib.request import urlopen
from urllib.error import URLError, HTTPError


def lambda_handler(event, context):
    url = os.getenv("URL")
    if url == "":
        logging.error("cannot load url from environment variables")
        return
    try:
        urlopen(url)
        logging.info("request success")
    except HTTPError as e:
        logging.error("cannot request server")
        logging.error("error code: " + str(e.code))
    except URLError as e:
        logging.error("cannot request server")
        logging.error("reason:" + e.reason)


if __name__ == "__main__":
    lambda_handler('', '')
