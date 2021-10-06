const Mustache = require('mustache');
const fs = require('fs');

function render(templatePath, attributes) {
    const template = fs.readFileSync(templatePath, 'utf8');
    return Mustache.render(template, attributes);
}

function renderInLayout(layoutPath, templatePath, attributes) {
    // Render the body.
    const body = render(templatePath, attributes)
    // Generate a new set of attributes with the body as well.
    const attr = Object.assign({ body: body }, attributes)
    // Render the layout with the body on it.
    return render(layoutPath, attr)
}

module.exports = {
    render: render,
    renderInLayout: renderInLayout,
}