const mongoose = require("mongoose");
const Joi = require("joi");

const userSchema = new mongoose.Schema({
  email: {
    type: String,
    required: true,
    unique: true,
  },
  phone: {
    type: String,
    required: true,
    validate: {
      validator: (value) => /^\d{10}$/.test(value),
      message: "Phone number must be a 10-digit number",
    },
  },
  first_name: {
    type: String,
    maxlength: 50,
  },
  middleName: {
    type: String,
    maxlength: 50,
  },
  last_name: {
    type: String,
    maxlength: 50,
  },
  year: {
    type: String,
    enum: ["FE", "SE", "TE", "BE"],
  },
  div: {
    type: String,
    maxlength: 10,
  },
  roll_no: {
    type: String,
    maxlength: 20,
  },
  department: {
    type: String,
    maxlength: 50,
  },
  expectation: {
    type: String,
    maxlength: 500,
  },
  payment_id: {
    type: String,
    default: "",
  },
  trId: {
    type: String,
    maxlength: 50,
  },
  receiptImage: {
    type: String,
  },
});

const User = mongoose.model("User", userSchema);

function validateUser(user) {
  const schema = Joi.object({
    email: Joi.string().required(),
    phone: Joi.string()
      .regex(/^\d{10}$/)
      .required(),
    first_name: Joi.string().max(50),
    middleName: Joi.string().max(50).allow(""),
    last_name: Joi.string().max(50),
    year: Joi.string().valid("FE", "SE", "TE", "BE"),
    div: Joi.string().max(10),
    roll_no: Joi.string().max(20),
    department: Joi.string().max(50),
    expectation: Joi.string().max(500).allow(""),
    payment_id: Joi.string().allow(""),
    trId: Joi.string().max(50),
    receiptImage: Joi.string().required(), // Adjust this for file validation
  });

  return schema.validate(user);
}

module.exports = {
  User,
  validateUser,
};
