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


## /:
    if not logged in /login
    if logged in /home

## /login:
    button to /register
    const { loginData, handleChange, handleSubmit } = useLogin()
    if response: "success", go to /home

## /register:
    button to /login
    const { registerData, handleChange, handleFileChange, handleSubmit } = useRegister()
    if response: "success", go to /login

## /home:
    nabar:
        logo to /home

        search?

        notifications
            useGetNotifications() 

        logout button
            useLogout()


    left sidebar:
        button to /home

        button to own profile /user/id

        button to all users /users

        button to all groups /groups

        button to create a group (todo)

    main body:
        create post: 
            const { postData, handleChange, handleFileChange, handleSubmit } = useCreatePost()
            
        fetch all public posts:
            const { allPosts } = useGetAllPosts()
        
    notification bar:
        get all notifications
            useGetNotifications()

        add trigger to send new notification if user is online (todo)


    right sidebar:

    TODO: add message boxes, with send message function

        message buttons for people who are following you or you are following:

            fetch all id, username, avatar of follows/followees
                useGetContacts()

            fetch all messages:
                useGetMesages(userId, 0) 

            send message button (todo) (use the websocket.send thing)

        group chat buttons for each group you are in:
            fetch all messages:
                useGetMessages(0, groupId)

            send message button (todo) (use websocket.send thing)


## /user/id:
    follow button:
        useAddFollow(userId)

    public profile:
        fetch user info:
            useGetUser(userId)

    private profile: 
        fetch private user info:
            useGetUser(userId) (only display username, avatar)
            username
            avatar


## /group/id:
    fetch group data: 
        getAllGroupPosts(groupId)

        if part of the group:

            group id, title, description, owner, all members, all posts, all events

            crete event button

        if not part of the group:
        response: "failure"
        message: "User isn't part of the group"
            join group button:
                useSendGroupJoinRequest(groupId)


## /groups:
    searchbar

    fetch all group names, id, member status
    
    join group button:
        useSendGroupJoinRequest(groupId)


## /users:
    searchbar

    fetch all names, id, following status



## /post/id:
    useGetOnePost(postId)
    useGetComment(postId)



