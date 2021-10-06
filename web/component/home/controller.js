'use strict'

// Load templating engine.
const path = require('path');
const random = require(path.join(process.env.PROJECT_ROOT_DIR, '/component/common/random'));
const templates = require(path.join(process.env.PROJECT_ROOT_DIR, '/component/common/template'));
// Construct the path to the layout template.
const layout = path.join(process.env.PROJECT_ROOT_DIR, '/component/common/templates/layout.mustache');
const home = path.join(process.env.PROJECT_ROOT_DIR, '/component/home/templates/home.mustache')

function handle(req, res) {
    // Consume time.
    random.randomBusyWait(800,400);
    // Render template.
    const data = templates.renderInLayout(layout, home, {
        title: 'Home',
        visitcount: res.locals.visitCount,
    });
    res.send(data);
}

module.exports = {
    handle: handle
};