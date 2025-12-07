FROM dunglas/frankenphp:latest-php8.5

# Instalasi Sistem, Composer, dan Ekstensi PHP
RUN apt-get update && apt-get install -y git curl unzip \
    && rm -rf /var/lib/apt/lists/*
RUN curl -sS https://getcomposer.org/installer | php -- --install-dir=/usr/local/bin --filename=composer
RUN install-php-extensions pdo_mysql mbstring xml opcache

WORKDIR /app

# Salin SEMUA file aplikasi terlebih dahulu agar artisan dan bootstrap/app.php tersedia saat composer install
COPY . /app

# Instal dependensi Composer (folder vendor akan dibuat di sini)
# Composer akan menjalankan package:discover dengan file aplikasi yang lengkap
RUN composer install --no-dev --optimize-autoloader

# Atur izin user web
RUN chown -R www-data:www-data /app