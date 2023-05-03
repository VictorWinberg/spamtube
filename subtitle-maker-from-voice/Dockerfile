FROM python:3.10.4-alpine

WORKDIR /app

# install dependencies
RUN apt-get -y update
RUN apt-get -y upgrade
RUN apt-get install -y ffmpeg libsm6 libxext6

COPY requirements.txt ./

RUN pip install -r requirements.txt

COPY . ./

CMD ["python", "magic.py"]
