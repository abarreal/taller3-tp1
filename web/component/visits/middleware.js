const config = require('./config');

let http = null

if (config.IS_PRODUCTION) {
    console.log('Using HTTPS for API connectivity');
    http = require('https');
} else {
    console.log('Using HTTP for API connectivity');
    http = require('http');
}

// Import Pub/Sub client.
const path = require('path');
const pubsubpath = path.join(process.env.PROJECT_ROOT_DIR, 'component/pubsub/client');
const pubsub = require(pubsubpath);

console.log('Topic ID:', config.GCP_PUBSUB_TOPIC_ID);

function visitCounterMiddleware(req, res, next) {
    // Do not count visits if a static resource is requested.
    if (req.path === "/" || req.path === "/favicon.ico" || req.path.startsWith('/static')) {
        console.log('Static resource requested: not counting visit');
        return next();
    }

    // Perform HTTP request to get visit counter.
    const options = {
        hostname: config.API_HOST,
        path: config.API_ENDPOINT_GET_TOTAL_VISITS,
        port: config.API_PORT,
        method: 'GET',
    };
    http.request(options, httpResponse => {
        console.log('Requesting visit count from API');
        httpResponse.on('data', data => {
            // Parse JSON response. Set the visit counter to be what we retrieved plus one
            // to account for this visit.
            const raw = data.toString()
            console.log('Received data:', raw);
            const obj = JSON.parse(raw)
            res.locals.visitCount = obj.Count + 1;

            // Register the new visit. Construct notification message.
            const jsondata = JSON.stringify({
                path: req.path,
            });
            const msgbuffer = Buffer.from(jsondata, 'utf8');
            // Push notification about new visit.
            console.log('Sending visit notification via Pub/Sub')
            pubsub.topic(config.GCP_PUBSUB_TOPIC_ID).publish(msgbuffer);

            // Proceed to handle request.
            console.log('Proceeding with request handling')
            next();
        });
    })
    .on('error', error => {
        console.error(error);
        // Set visit count to -1 to indicate that there was a
        // problem retrieving the counter.
        res.locals.visitCount = -1
        next();
    })
    .end();
}

module.exports = visitCounterMiddleware;