const http = require('../common/http');

const Client = (baseUrl) => {

    const client = http.Client(baseUrl);

    return {
        listChannels: () => client.get('/channels'),
        createChannel: (name) => client.post('/channels', { name })
    }

};

module.exports = { Client };