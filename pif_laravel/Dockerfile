# Menggunakan image FrankenPHP resmi
FROM dunglas/frankenphp

RUN apt-get update && apt-get add --no-cache git curl unzip

RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer

RUN install-php-extensions \
    pdo_mysql \
    mbstring \
    xml \
    opcache

WORKDIR /app

COPY . /app

RUN composer install --no-dev --optimize-autoloader

RUN chown -R frphp:frphp /app
