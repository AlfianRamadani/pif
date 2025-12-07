FROM dunglas/frankenphp


RUN apt-get update && apt-get install -y git curl unzip \
    && rm -rf /var/lib/apt/lists/*

RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer

RUN install-php-extensions \
    pdo_mysql \
    mbstring \
    xml \
    opcache

WORKDIR /app

COPY composer.json composer.lock ./

RUN composer install --no-dev --optimize-autoloader

COPY . /app 

RUN chown -R www-data:www-data /app

