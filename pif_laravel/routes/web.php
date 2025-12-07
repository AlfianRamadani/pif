<?php

use App\Http\Controllers\PerformanceController;
use Illuminate\Support\Facades\Route;

Route::get('/test-performance', [PerformanceController::class, 'testPerformance']);
