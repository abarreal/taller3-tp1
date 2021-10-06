'use strict';

process.env.PROJECT_ROOT_DIR = process.env.PROJECT_ROOT_DIR || __dirname;

// Read environment configuration.
const PORT = process.env.PORT || 8000;
const HOST = process.env.HOST || '0.0.0.0';

// Initialize express application.
const express = require('express');
const app = express();

// Handle static files.
app.use('/static', express.static('static'));

// Add middleware.
const visitcounter = require('./component/visits/middleware')
app.use(visitcounter);

// Use custom router.
const router = require('./router');
app.use(router);

// Initialize server.
const server = app.listen(PORT, HOST, () => {
    console.log(`Express server listening on port ${PORT}`)
});

// Implement graceful shutdown.
function shutdown() {
    console.log('Server shutting down');
    server.close(() => {
        console.log('Program execution finished');
        process.exit(0);
    })
}

['SIGTERM', 'SIGINT'].forEach(signal => {
    process.on(signal, shutdown);
});