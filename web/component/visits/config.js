const ENV_PRODUCTION = 'production'

module.exports = {
    IS_PRODUCTION: (process.env.ENV == ENV_PRODUCTION),
    API_HOST: process.env.API_HOST,
    API_PORT: process.env.API_PORT,
    API_ENDPOINT_GET_TOTAL_VISITS: process.env.API_ENDPOINT_GET_TOTAL_VISITS,
    GCP_PUBSUB_TOPIC_ID: process.env.GCP_PUBSUB_TOPIC_ID,
}