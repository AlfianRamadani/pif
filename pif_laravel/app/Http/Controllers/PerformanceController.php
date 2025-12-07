<?php

namespace App\Http\Controllers;

use Illuminate\Http\Request;

class PerformanceController extends Controller
{
    public function fibonacci(int $n): int
    {
        if ($n <= 1) {
            return $n;
        }
        return $this->fibonacci($n - 1) + $this->fibonacci($n - 2);
    }

    public function testPerformance(Request $request)
    {
        $min_n = 5;
        $max_n = 30; 

        $n = mt_rand($min_n, $max_n);

        $result = $this->fibonacci($n);

        return response()->json([
            'framework' => 'Laravel',
            'language' => 'PHP',
            'task' => "Calculate Fibonacci($n)", 
            'result' => $result
        ]);
    }
}
