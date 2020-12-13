const http = require("../common/http");

const Client = (baseUrl) => (
  (client = http.Client(baseUrl)),
  {
    fetchAccounts: () => client.get("/accounts"),
    transferMoney: () => client.post("/accounts"),
  }
);

module.exports = { Client };