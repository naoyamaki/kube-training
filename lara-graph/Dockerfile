FROM php:8.3-apache

ENV TZ=Asia/Tokyo
ENV LANG=ja_JP.UTF-8
ENV APACHE_DOCUMENT_ROOT='/var/www/public/'

RUN apt update
RUN apt install -y libicu-dev autoconf make g++ gcc ca-certificates wget curl
# for gd
RUN apt install -y libfreetype-dev libjpeg62-turbo-dev libpng-dev
# モジュール追加
RUN docker-php-ext-install gd intl pdo_mysql opcache

RUN sed -ri -e 's!/var/www/html!${APACHE_DOCUMENT_ROOT}!g' /etc/apache2/sites-available/*.conf \
  && sed -ri -e 's!/var/www/!${APACHE_DOCUMENT_ROOT}!g' /etc/apache2/apache2.conf /etc/apache2/conf-available/*.conf

COPY ./web/php.ini /usr/local/php/
COPY --chown=www-data:www-data ./src /var/www


RUN curl -L -O https://getcomposer.org/download/latest-stable/composer.phar && \
    chmod +x composer.phar && mv composer.phar /usr/local/bin/composer
