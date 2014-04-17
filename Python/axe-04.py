#!/usr/bin/env python
# -*- coding: utf-8 -*-
import json
import requests
from pyquery import PyQuery as pq

data = []
columns = ['town', 'village', 'name']
headers = {
    'User-Agent': 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_1) AppleWebKit/537.73.11 (KHTML, like Gecko) Version/7.0.1 Safari/537.73.11',
    'Accept': 'text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8',
    'Referer': 'http://axe-level-4.herokuapp.com/lv4/?page=1'
}

for i in range(1, 25):
    page = requests.get(
        'http://axe-level-4.herokuapp.com/lv4/?page={0}'.format(i),
        headers=headers
    )
    html = pq(page.content)
    for row in html('table.table tr')[1:]:
        elements = [element.text for element in row]
        data.append(dict(zip(columns, elements)))

print(json.dumps(data, ensure_ascii=False).encode('utf-8'))
