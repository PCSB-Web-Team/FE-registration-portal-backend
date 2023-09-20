const express = require("express");
const { User, validateUser } = require("./models/userSchema");
const connectDatabase = require("./config/db");
const cors = require("cors");

const app = express();
app.use(cors());
app.use(express.json({ limit: "25mb" }));
app.use(express.urlencoded({ limit: "25mb", extended: true }));

app.get("/", (req, res) => {
  res.send("Hello from Drive API server");
});

app.post("/register", async (req, res) => {
  const { error } = validateUser(req.body);

  if (error) {
    return res.status(400).json({ error: error.details[0].message });
  }

  try {
    const user = new User(req.body);

    const savedUser = await user.save();

    res
      .status(201)
      .json({ message: "User data saved successfully", user: savedUser });
  } catch (error) {
    console.error("Error saving user data:", error);
    res.status(500).json({ error: "Internal server error" });
  }
});

app.get("/users", async (req, res) => {
  const page = parseInt(req.query.page) || 1; // Default to page 1
  const pageSize = parseInt(req.query.pageSize) || 10; // Default page size

  try {
    const skip = (page - 1) * pageSize;
    const totalUsers = await User.countDocuments(); // Get the total number of users
    const users = await User.find().skip(skip).limit(pageSize);

    const totalPages = Math.ceil(totalUsers / pageSize);

    res.status(200).json({
      users,
      pageInfo: {
        currentPage: page,
        pageSize,
        totalPages,
        totalUsers,
      },
    });
  } catch (error) {
    console.error("Error retrieving paginated users:", error);
    res.status(500).json({ error: "Internal server error" });
  }
});

module.exports = app;
