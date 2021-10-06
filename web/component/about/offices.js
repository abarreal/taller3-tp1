'use strict'

// Load templating engine.
const path = require('path');
const templates = require(path.join(process.env.PROJECT_ROOT_DIR, '/component/common/template'));
// Construct the path to the layout template.
const layout = path.join(process.env.PROJECT_ROOT_DIR, '/component/common/templates/layout.mustache');
const home = path.join(process.env.PROJECT_ROOT_DIR, '/component/about/templates/offices.mustache')

function handle(req, res) {
    // Render template.
    const data = templates.renderInLayout(layout, home, {
        title: 'Offices',
        visitcount: res.locals.visitCount,
    });

    res.send(data);
}

module.exports = {
    handle: handle
};