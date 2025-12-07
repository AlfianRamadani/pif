# Menggunakan image FrankenPHP resmi
FROM dunglas/frankenphp:latest-php8.5

RUN install-php-extensions \
    pdo_mysql \
    mbstring \
    xml \
    opcache

WORKDIR /app

COPY . /app

RUN composer install --no-dev --optimize-autoloader

RUN chown -R frphp:frphp /app
