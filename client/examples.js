// This file contains examples of scenarios implementation using
// the SDK for channels management.

const channels = require('./channels/client');

const client = channels.Client('http://localhost:8080');

// Scenario 1: Display available channels.
client.listChannels()
    .then((list) => {
        console.log('=== Scenario 1 ===');
        console.log('Available channels:');
        list.forEach((c) => console.log(c.name));
    })
    .catch((e) => {
        console.log(`Problem listing available channels: ${e.message}`);
    });

// Scenario 2: Create new channel.
client.createChannel('my-new-channel')
    .then((resp) => {
        console.log('=== Scenario 2 ===');
        console.log('Create channel response:', resp);
        return client.listChannels()
            .then((list) => list.map((c) => c.name).join(', '))
            .then((str) => {
                console.log(`Current channels: ${str}`);
            })
    })
    .catch((e) => {
        console.log(`Problem creating a new channel: ${e.message}`);
    });