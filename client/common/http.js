const request = require('request');

const Client = (baseUrl) => {

    const respHandler = (resp) => {
        if (resp.ok) {
            return resp.json();
        }
        throw new Error(`Unexpected response from the server ${resp.status} ${resp.statusText}`)
    };

    return {
        get: (path) => {
            return new Promise((resolve, reject) => {
                request(`${baseUrl}${path}`, {json: true}, (err, res, body) => {
                    if (err) {
                        reject(err);
                        return;
                    }
                    resolve(body);
                });
            });
        },
        post: async (path, data) => {
            return new Promise((resolve, reject) => {
                request(`${baseUrl}${path}`, {json: true, method: 'POST', body: data}, (err, res, body) => {
                    if (err) {
                        reject(err);
                        return;
                    }
                    resolve(body);
                });
            });
        }
    };
};

module.exports = { Client };