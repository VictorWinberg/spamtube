FROM python:3.8.0-buster

WORKDIR /app

# install dependencies
RUN apt-get -y update
RUN apt-get -y upgrade
RUN apt-get install -y ffmpeg libsm6 libxext6

COPY requirements.txt ./

RUN pip install -r requirements.txt

COPY . ./

CMD [ "python3", "main_swinir.py", "--task", "real_sr", "--scale", "4", "--model_path", "model_zoo/swinir/003_realSR_BSRGAN_DFO_s64w8_SwinIR-M_x4_GAN.pth", "--folder_lq", "data"]
