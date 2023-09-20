const mongoose = require("mongoose");
const { DB_URI } = require("./index");

const connectDatabase = async () => {
  await mongoose
    .connect(DB_URI, {
      useNewUrlParser: true,
      useUnifiedTopology: true,
      //   useCreateIndex: true,
    })
    .then((data) => {
      console.log(`✅ Mongodb connected with server: ${data.connection.host}`);
    });
};

module.exports = connectDatabase;
