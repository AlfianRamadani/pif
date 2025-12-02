const express = require('express');
const app = express();
const port = 3000;

function fibonacci(n) {
    if (n <= 1) {
        return n;
    }
    return fibonacci(n - 1) + fibonacci(n - 2);
}

app.get('/test-performance', (req, res) => {
    const n = 10;

    const result = fibonacci(n);

    res.json({
        framework: 'Express.js',
        language: 'JavaScript (Node.js)',
        task: `Calculate Fibonacci(${n})`,
        result: result
    });
});

app.listen(port, () => {
    console.log(`Express app listening at http://localhost:${port}`);
});