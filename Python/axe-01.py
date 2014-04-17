#!/usr/bin/env python
# -*- coding: utf-8 -*-
import json
import requests
from pyquery import PyQuery as pq

page = requests.get('http://axe-level-1.herokuapp.com/')
html = pq(page.content)

data = []
subjects = [u'國語', u'數學', u'自然', u'社會', u'健康教育']
for row in html('table.table tr')[1:]:
    elements = [element.text for element in row]
    data.append({
        'name': elements[0],
        'grades': dict(zip(subjects, map(int, elements[1:])))
    })

print(json.dumps(data, ensure_ascii=False).encode('utf-8'))
