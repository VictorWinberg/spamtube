import os
import urllib
from docarray import Document

server_url = 'grpcs://dalle-flow.dev.jina.ai'
prompt = os.getenv('IMAGE_INPUT')
doc = Document(text=prompt).post(server_url, parameters={'num_images': 10})

for i, d in enumerate(doc.matches):
    response = urllib.request.urlopen(d.uri)
    image_id = str(i + 1).zfill(3)
    with open('out/' + image_id + '.png', 'wb') as f:
        f.write(response.file.read())

print(f'Successfully generated {len(doc.matches)} images')
