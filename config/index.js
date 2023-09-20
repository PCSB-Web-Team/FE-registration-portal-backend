const dotenv = require('dotenv');

dotenv.config();

const DB_URI = process.env.DB_URI;
const APP_PORT = process.env.APP_PORT;
const ADMIN_TOKEN = process.env.ADMIN_TOKEN;

module.exports = { DB_URI, APP_PORT, ADMIN_TOKEN }