FROM python:3.8.0-buster

WORKDIR /app

# install dependencies
RUN apt-get -y update
RUN apt-get -y upgrade
RUN apt-get install -y ffmpeg libsm6 libxext6

COPY requirements.txt ./

RUN pip install -r requirements.txt

COPY . ./

CMD [ "python", "main_swinir.py", "--task", "lightweight_sr", "--scale", "4", "--model_path", "model_zoo/swinir/002_lightweightSR_DIV2K_s64w8_SwinIR-S_x4.pth", "--folder_lq", "data", "--folder_gt", "data"]