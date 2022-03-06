import requests
from datetime import datetime

def get_block_height_and_timestamp(tx_hash_str):
    res = requests.get(f"http://0.0.0.0:1317/txs/{tx_hash_str}")
    body = res.json()
    if "height" not in body:
        print(tx_hash_str)
        return
    height = int(body["height"])
    timestamp = datetime.strptime(body["timestamp"], "%Y-%m-%dT%H:%M:%SZ")
    return height, timestamp

def get_metrics():
    seen_heights = set([])
    with open("/home/tony_chen/outputs/test_new_chain_5", "r") as f:
        curr = f.readline()[:-1]
        while curr:
            height, ts = get_block_height_and_timestamp(curr)
            curr = f.readline()[:-1]
            seen_heights.add(height)
    print(seen_heights)

get_metrics()