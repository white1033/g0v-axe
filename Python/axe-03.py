#!/usr/bin/env python
# -*- coding: utf-8 -*-
import json
import requests
from pyquery import PyQuery as pq

columns = ['town', 'village', 'name']
data = []
session = requests.Session()

for i in range(1, 77):
    if i == 1:
        url = 'http://axe-level-1.herokuapp.com/lv3/'
    else:
        url = 'http://axe-level-1.herokuapp.com/lv3/?page=next'

    page = session.get(url)
    html = pq(page.content)
    for row in html('table.table tr')[1:]:
        elements = [element.text for element in row]
        data.append(dict(zip(columns, elements)))

print(json.dumps(data, ensure_ascii=False).encode('utf-8'))
