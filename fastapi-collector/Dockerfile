FROM python:3.10.5-slim as base

# Dependencies
ENV PKGS_DIR=/install \
    PIP_NO_CACHE_DIR=off \
    PIP_DISABLE_PIP_VERSION_CHECK=on \
    PIP_DEFAULT_TIMEOUT=100

FROM base as builder

RUN mkdir $PKGS_DIR
RUN mkdir /code
COPY requirements.txt /code/

WORKDIR /code

RUN pip3 install --upgrade pip
RUN pip3 install --target=$PKGS_DIR -r ./requirements.txt


# Main image with service
FROM base
ARG SRC_PATH=.

ENV PYTHONPATH=/usr/local
ENV SERVICE_PORT=5005

COPY --from=builder /install /usr/local
COPY . /code
WORKDIR /code

# Run service
CMD python main.py
