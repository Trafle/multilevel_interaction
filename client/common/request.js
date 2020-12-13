const { parse } = require("url");

module.exports = async (url, method = "GET", reqData) => {
  url = parse(url);
  url.protocol = url.protocol.slice(0, -1);

  const params = {
    method,
    host: url.hostname,
    port: url.port || (url.protocol === "https" ? 443 : 80),
    path: url.path,
  };

  return new Promise((resolve, reject) => {
    const req = require(url.protocol).request(params, (res) => {
      const data = [];
      res.on("data", (chunk) => {
        data.push(chunk);
      });
      res.on("end", () => resolve(data.toString()));
    });
    req.on("error", reject);

    reqData && req.write(JSON.stringify(reqData));

    req.end();
  });
};
