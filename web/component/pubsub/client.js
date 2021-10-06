const {PubSub} = require('@google-cloud/pubsub');

const config = require('./config');
const projectId = config.GCP_PROJECT_ID

console.log('GCP Project ID:', projectId);

module.exports = new PubSub({ projectId });