'use strict'

// Load templating engine.
const path = require('path');
const templates = require(path.join(process.env.PROJECT_ROOT_DIR, '/component/common/template'));
// Construct the path to the layout template.
const layout = path.join(process.env.PROJECT_ROOT_DIR, '/component/common/templates/layout.mustache');
const home = path.join(process.env.PROJECT_ROOT_DIR, '/component/about/templates/about.mustache');

async function handle(req, res) {
    // Render template.
    const data = await templates.renderInLayout(layout, home, {
        title: 'About',
        visitcount: res.locals.visitCount,
    });

    res.send(data);
}

module.exports = {
    handle: handle
};