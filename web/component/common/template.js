const Mustache = require('mustache');
const fs = require('fs');

var templateCache = {};

async function loadTemplate(templatePath) {
    if (templatePath in templateCache) {
        return templateCache[templatePath];
    } else {
        return await new Promise((resolve) => {
            fs.readFile(templatePath, (err,data) => {
                const template = data.toString();
                templateCache[templatePath] = template;
                resolve(template);
            });
        });
    }
}

async function render(templatePath, attributes) {
    const template = await loadTemplate(templatePath);
    return Mustache.render(template, attributes);
}

async function renderInLayout(layoutPath, templatePath, attributes) {
    // Render the body.
    const body = await render(templatePath, attributes)
    // Generate a new set of attributes with the body as well.
    const attr = Object.assign({ body: body }, attributes)
    // Render the layout with the body on it.
    return await render(layoutPath, attr)
}

module.exports = {
    render: render,
    renderInLayout: renderInLayout,
}