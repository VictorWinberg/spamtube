FROM python:3.8.0-buster

WORKDIR /app

# install dependencies
RUN apt-get -y update
RUN apt-get -y upgrade
RUN apt-get install -y ffmpeg libsm6 libxext6

COPY requirements.txt ./

RUN pip install -r requirements.txt

COPY . ./

CMD ["sh", "make-subs.sh"]
