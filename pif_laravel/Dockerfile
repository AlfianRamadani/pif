FROM php:8.5-fpm

# Instalasi Sistem, Composer, dan Ekstensi PHP
RUN apt-get update && apt-get install -y \
    git \
    curl \
    unzip \
    libzip-dev \
    libsqlite3-dev \
    libpng-dev \
    nano \
    libonig-dev \
    # Hapus file cache setelah instalasi
    && rm -rf /var/lib/apt/lists/* \
    && docker-php-ext-install pdo pdo_mysql opcache bcmath gd mbstring \
    && apt-get clean

COPY --from=composer:latest /usr/bin/composer /usr/bin/composer

WORKDIR /var/www/html

COPY . .

RUN composer install --no-dev --optimize-autoloader --no-scripts

RUN chown -R www-data:www-data /var/www/html \
    && chmod -R 777 /var/www/html/storage \
    && chmod -R 777 /var/www/html/bootstrap/cache

EXPOSE 9000

CMD ["php-fpm"]