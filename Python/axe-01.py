#!/usr/bin/env python
# -*- coding: utf-8 -*-
import json
import requests
import lxml.html

page = requests.get('http://axe-level-1.herokuapp.com/')
page.encoding = 'utf-8'
tree = lxml.html.fromstring(page.text)
rows = tree.cssselect('table.table tr')[1:]

data = []
subjects = [u'國語', u'數學', u'自然', u'社會', u'健康教育']
for row in rows:
    elements = [element.text_content() for element in row]
    data.append({
        'name': elements[0],
        'grades': dict(zip(subjects, map(int, elements[1:])))
    })

print(json.dumps(data, ensure_ascii=False).encode('utf-8'))
