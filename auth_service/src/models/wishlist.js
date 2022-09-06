const mongoose = require('mongoose');
const Schema = mongoose.Schema;

const tableSchema = new Schema({
  bookId: {
    type: String,
    required: true,
  },
  title: {
    type: String,
    unique: true,
    required: true,
    trim: true,
  },
  thumbnail: {
    type: String,
    // required: true,
  },
  author: {
    type: Array,
    // required: true,
    trim: true,
  },
  rating: {
    type: Number,
    default: 0,
    required: true,
  },
  description: {
    type: String,
    // required: true,
  },
  link: {
    type: String,
    // required: true,
  },
  createdAt: {
    type: Date,
    default: Date.now
  },
  updatedAt: {
    type: Date,
    default: Date.now
  }
}, {
  timestamps: true
});

tableSchema.method('toJSON', function () {
  const { _id, ...object } = this.toObject();
  object.id = _id;
  delete object.__v;
  return object;
});

module.exports = mongoose.model('wishlist', tableSchema);
