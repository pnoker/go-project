#
#  Copyright 2019 Pnoker. All Rights Reserved.
#

FROM mongo:4.0-xenial
MAINTAINER pnoker <pnokers.icloud.com>

RUN ln -sf /usr/share/zoneinfo/Asia/Shanghai /etc/localtime \
    && mkdir -p /pnoker/mongo/config/

WORKDIR /pnoker/mongo

COPY ./mongo-init.js ./config/
COPY ./launch.sh ./

RUN chmod +x ./launch.sh

EXPOSE 27017

CMD [ "/pnoker/mongo/launch.sh"]