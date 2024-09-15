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


# /:
    if not logged in /login
    if logged in /home

# /login:
    button to /register
    const { loginData, handleChange, handleSubmit } = useLogin()
    if response: "success", go to /home

# /register:
    button to /login
    const { registerData, handleChange, handleFileChange, handleSubmit } = useRegister()
    if response: "success", go to /login

# /home:
   ## nabar:
        logo to /home

        notifications (todo)

        logout button (todo)

        search?

 ##   left sidebar:
###        button to /home

###        button to profile /user/id:
            const { userData } = useGetUser(userId) (wip) (private/public)
 ###       button to fetch all users: (todo)

###        button to fetch all groups: (todo)
   ###         button to request to join a group:
                useSendGroupJoinRequest(groupId) (change url to body)

###            button to create a group:
                const { groupData, handleChange, handleSubmit } = useCreateGroup()
            
   ###         button to group page /group/id:
                fetch group data (todo)

  ##   main body:
###        create post: 
            const { postData, handleChange, handleFileChange, handleSubmit } = useCreatePost()
            
   ###     fetch all public posts including posts from your own groups:
            const { allPosts } = useGetAllPosts() (wip)
        
  ##   notification bar: (todo)
        get all notifications

        add trigger to send new notification if user is online

        follow requests

        group requests

        events
    
  ##   right sidebar:
###        message buttons for people who are following you or you are following:
            fetch all messages (todo)

            send message button (todo)

###        group chat buttons for each group you are in:
            fetch all messages (todo)

            send message button (todo)









