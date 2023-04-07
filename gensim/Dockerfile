FROM python:3.8.0-buster

WORKDIR /app

COPY requirements.txt ./

# Install Rust
RUN curl https://sh.rustup.rs -sSf | sh -s -- -y
ENV PATH="/root/.cargo/bin:${PATH}"

# Install Cython
RUN pip install Cython==0.29.33

# Install Python dependencies
RUN pip install --upgrade pip
RUN pip install -r requirements.txt

COPY . ./

CMD [ "python3", "gensim_main.py"]