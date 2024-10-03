db.createUser(
    {
        user: "dbadmin",
        pwd: "despitia",
        roles: [
            {
                role: "readWrite",
                db: "transactionsdb"
            }
        ]
    }
);