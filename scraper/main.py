#!/usr/bin/python3

from dataclasses import dataclass
from typing import Optional
from mako.template import Template

import yaml


@dataclass
class Action:
    name: str
    description: str
    payload: Optional[dict]
    response_to_group: Optional[dict]
    response_to_user: Optional[dict]

    def __init__(self, name: str, data: dict) -> None:
        self.name = name
        self.description = data.get('description')
        self.payload = data.get('payload')
        self.response_to_group = data.get('response_to_group')
        self.response_to_user = data.get('response_to_user')

    def html(self) -> str:
        t = Template(filename='scraper/template.html')
        return t.render(action=self)


WS_FILE = 'examples/chat-ws.yaml'

with open(WS_FILE, 'r') as f:
    doc = yaml.load(f)


if __name__ == '__main__':
    actions = [
        Action(key, value)
        for key, value in doc.items()
    ]

    print(actions[0].html())
