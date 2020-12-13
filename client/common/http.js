const request = require("./request");

const Client = (baseUrl) => ({
  get: (path) => request(baseUrl + path),
  post: (path, data) => request(baseUrl + path, "POST", data),
  patch: (path, data) => request(baseUrl + path, "PATCH", data),
});

module.exports = { Client };