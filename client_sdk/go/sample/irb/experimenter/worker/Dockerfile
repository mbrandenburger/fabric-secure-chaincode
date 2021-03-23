# Copyright IBM Corp. All Rights Reserved.
#
# SPDX-License-Identifier: Apache-2.0

FROM python:3.8

# set the working directory in the container
WORKDIR /code

RUN apt-get update -q \
    && apt-get install -y -q \
    libnss-mdns \
    libnss-myhostname \
    lsb-release

# copy experiment code
COPY src/ .

RUN pip install torchvision pillow
RUN python download-pretrained-model.py

# our go service
COPY ./app /app

EXPOSE 5050
ENTRYPOINT /app