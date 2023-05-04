FROM python:3.10.4-slim

WORKDIR /app

# install dependencies
RUN apt-get update
RUN apt-get install -y ffmpeg

COPY requirements.txt ./

RUN pip install -r requirements.txt

COPY . ./

CMD ["python", "magic.py"]
