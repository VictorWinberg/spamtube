FROM python:3.11.0-buster

WORKDIR /app

COPY requirements.txt ./

RUN pip install -r requirements.txt

COPY . ./

CMD [ "python3", "chatgpt.py"]