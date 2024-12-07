import requests as req

url = 'https://7b40-110-175-160-104.ngrok-free.app/'


rsp = req.get(url)
print(rsp.text)

msg = input('enter msg: ')
req.put(url, data = msg, headers = {'Content-Type': 'text/plain'})
