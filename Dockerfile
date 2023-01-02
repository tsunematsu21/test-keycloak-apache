FROM httpd:2.4

COPY ./httpd.conf /usr/local/apache2/conf/httpd.conf
COPY ./index.html /usr/local/apache2/htdocs/index.html

RUN apt-get update \
 && apt-get install -y --no-install-recommends \
    libapache2-mod-auth-openidc \
 && apt-get -y clean \
 && rm -rf /var/lib/apt/lists/*
