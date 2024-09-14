# social-network

## audit
https://github.com/01-edu/public/tree/master/subjects/social-network/audit

# database migrations

``goose -s create add new_table sql`` (make a new sql table file)

``goose up`` (goes up to latest version)

``goose up-by-one`` (goes up one table)

``goose down`` (drops the latest table)

``kooduser@pupil:~/js/social-network/backend$ GOOSE_DRIVER=sqlite3 GOOSE_DBSTRING=./db/sn.db goose -s -dir ./db/migrations/ create add_new_posts_table sql``


websocket
add websocket after login with /api/websocket ending

send private and group messages:

{
    "message":"message data here",
    "userId": userId of the person you're talking to (leave empty if you want to send group message),
    "groupId": groupId of the group you're talking in (leave empty if you want to send a private message)
}

