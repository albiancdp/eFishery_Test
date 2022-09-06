import sqlite3 from 'sqlite3';
sqlite3.verbose();

const DB_SOURCE = 'database';

let db = new sqlite3.Database(DB_SOURCE, (err) => {
  if (err) {
    // Cannot open database
    // console.error(err.message);
    throw err;
  } else {
    // console.log('Connected to the SQLite database.');
    db.run(`CREATE TABLE user (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name text, 
            email text UNIQUE, 
            password text, 
            CONSTRAINT email_unique UNIQUE (email)
            )`,
    (err) => {
      if (err) {
        // Table already created
      } else {
        // create success
      }
    });
  };
});


export default db;
