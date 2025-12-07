const express = require('express');
const app = express();
const port = 3000;

function fibonacci(n) {
    if (n <= 1) {
        return n;
    }
    return fibonacci(n - 1) + fibonacci(n - 2);
}

function getRandomInt(min, max) {
    min = Math.ceil(min);
    max = Math.floor(max);
    return Math.floor(Math.random() * (max - min + 1)) + min;
}

app.get('/test-performance', (req, res) => {
    const min_n = 5;
    const max_n = 30; 

    const n = getRandomInt(min_n, max_n);

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