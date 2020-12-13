const { Client } = require("./accounts/client");

const client = Client("http://localhost:15000");

(async () => {
    // Scenario 1: Display a list of accounts with a balance on them and the time of the last transaction performed on the account
   let accounts;
    console.log("=== Scenario 1 ===");
    try {
        accounts = await client.fetchAccounts();
        console.log("Accounts with a balance on them:");
        console.table(accounts);
    } catch (err) {
        console.log(`Problem listing accounts: `, err);
    }
})();

(async () => {
    // Scenario 2: Transfer the specified amount from one account to another.
    console.log("=== Scenario 2 ===");
    try {
        const transfer = await client.transferMoney(
            1000,
            2,
            3
        );
    } catch (err) {
        console.log(`Problem with transfer: `, err);
    }
})();