db.createUser(
    {
        user: "juanse1801",
        pwd: "12345juanse",
        roles: [
            {
                role: "readWrite",
                db: "naranja_bot"
            }
        ]
    }
);