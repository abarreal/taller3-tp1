'use strict'

// Load templating engine.
const path = require('path');
const random = require(path.join(process.env.PROJECT_ROOT_DIR, '/component/common/random'));
const templates = require(path.join(process.env.PROJECT_ROOT_DIR, '/component/common/template'));
// Construct the path to the layout template.
const layout = path.join(process.env.PROJECT_ROOT_DIR, '/component/common/templates/layout.mustache');
const home = path.join(process.env.PROJECT_ROOT_DIR, '/component/jobs/templates/jobs.mustache')

async function handle(req, res) {
    // Consume time.
    await random.simulateActivity(1300,200);
    // Render template.
    const data = await templates.renderInLayout(layout, home, {
        title: 'Jobs',
        visitcount: res.locals.visitCount,
    });
    res.send(data);
}

module.exports = {
    handle: handle
};