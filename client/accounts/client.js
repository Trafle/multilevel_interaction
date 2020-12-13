const http = require("../common/http");

const Client = (baseUrl) => (
  (client = http.Client(baseUrl)),
  {
    fetchAccounts: () => client.get("/fetch"),
    transferMoney: () => client.post("/transfer"),
  }
);

module.exports = { Client };