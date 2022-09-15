### Built With

This section should list any major frameworks/libraries used to project.

* [NodeJS v16.14.0](https://nodejs.org/)
* [MongoDB](https://www.mongodb.com/)
* [Express.js](https://expressjs.com/)

<p align="right">(<a href="#top">back to top</a>)</p>

### Installation

1. Clone the repo
   ```sh
   git clone https://github.com/albiancdp/eFishery_Test.git
   ```
2. Install NPM packages
   ```sh
   npm install
   ```
3. Rename .env.example to .env
4. Run project for development
   ```sh
   npm run start:dev
   ```
5. Run project 
   ```sh
   npm run start
   ```
<p align="right">(<a href="#top">back to top</a>)</p>

### Description Response

#### Response Code
Code | Description |
--- | --- |
200 | Success |
400 | Bad Request |
402 | Request Validation Error |
404 | Not Found |
500 | Internal Server Error |

#### Response Error Code
Code | Description |
--- | --- |
validate_error | Validation Error |
create_error | Error Create Data |
update_error | Error Update Data |
delete_error | Error Delete Data |
not_found | Data Not Found |
server_error | Internal Server Error |

### Documentation

* [Postman](https://www.getpostman.com/collections/88efa3896bd0437bf002)

<p align="right">(<a href="#top">back to top</a>)</p>
