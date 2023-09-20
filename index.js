const app = require("./app");
const { APP_PORT } = require("./config");

const connectDatabase = require("./config/db");

// Handeling Uncaught Exceptions
process.on("uncaughtException", (err) => {
  console.log(`Error: ${err.message}`);
  console.log("Shutting Down Server due to uncaught exception");
  process.exit(1);
});

// config

// connecting database
connectDatabase();

app.listen(APP_PORT || 8080, () => {
  console.log(`server is listening on port ${APP_PORT || 8080}`);
});

// Unhandled Promise Rejection
process.on("unhandledRejection", (err) => {
  console.log(`Error: ${err.message}`);
  console.log("shutting down the server due to Unhandled Promise Rejection");

  server.close(() => {
    process.exit(1);
  });
});
