'use strict';

const express = require('express');
const router = express.Router();

router.get('/home', (req, res) => {
    const controller = require('./component/home/controller');
    return controller.handle(req, res);
});

router.get('/jobs', (req, res) => {
    const controller = require('./component/jobs/controller');
    return controller.handle(req, res);
});

router.get('/about', (req, res) => {
    const controller = require('./component/about/about');
    return controller.handle(req, res);
});

router.get('/about/legals', (req, res) => {
    const controller = require('./component/about/legals');
    return controller.handle(req, res);
});

router.get('/about/offices', (req, res) => {
    const controller = require('./component/about/offices');
    return controller.handle(req, res);
});

module.exports = router;