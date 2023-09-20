const express = require("express");
const { User, validateUser } = require("./models/userSchema");
const cors = require("cors");
const { ADMIN_TOKEN } = require("./config");

const app = express();
app.use(cors());
app.use(express.json({ limit: "25mb" }));
app.use(express.urlencoded({ limit: "25mb", extended: true }));

app.get("/", (req, res) => {
  res.send("Hello from PCSB Registration Portal API server");
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

app.get("/users/:token", async (req, res) => {
  const token = req.params.token;
  if (!token || token !== ADMIN_TOKEN) {
    return res.status(401).json({ error: "Unauthorized" });
  }

  const page = parseInt(req.query.page) || 1; // Default to page 1
  const pageSize = parseInt(req.query.pageSize) || 10; // Default page size

  try {
    const skip = (page - 1) * pageSize;
    const totalUsers = await User.countDocuments(); // Get the total number of users
    const users = await User.find({}, "-receiptImage").skip(skip).limit(pageSize);

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

app.get("/users/:id/:token", async (req, res) => {
  if (!req.params.token || req.params.token !== ADMIN_TOKEN) {
    return res.status(401).json({ error: "Unauthorized" });
  }

  try {
    const user = await User.findById(req.params.id);

    if (!user) {
      return res.status(404).json({ error: "User not found" });
    }

    res.status(200).json(user);
  } catch (error) {
    console.error("Error retrieving user:", error);
    res.status(500).json({ error: "Internal server error" });
  }
});

module.exports = app;
