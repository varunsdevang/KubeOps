import requests

def hit_endpoint(url, n):
    for i in range(n):
        try:
            response = requests.get(url)
            print(f"Request {i+1}: Status Code {response.status_code}")
        except requests.exceptions.RequestException as e:
            print(f"Request {i+1} failed: {e}")

if __name__ == "__main__":
    url = "http://127.0.0.1:59831/items/"  # Replace with your endpoint
    n = 100  # Number of times to hit the endpoint
    hit_endpoint(url, n)