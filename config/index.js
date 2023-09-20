const dotenv = require('dotenv');

dotenv.config();

const DB_URI = process.env.DB_URI;
const APP_PORT = process.env.APP_PORT;

module.exports = { DB_URI, APP_PORT }