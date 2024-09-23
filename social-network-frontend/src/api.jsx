import { useState, useEffect, createContext, useContext, useRef } from "react"
import { Navigate, useNavigate } from "react-router-dom"
import { closeSocket, InitSocket } from "./WebSocket"



export const useLogin = () => {
    const [loginData, setLoginData] = useState({
        name: '',
        password: ''
    })

    const navigate = useNavigate()

    const handleSubmit = (e) => {
        e.preventDefault()

        const requestOptions = {
            method: 'POST',
            headers: { 'Content-Type': 'application/json' },
            credentials: 'include',
            body: JSON.stringify(loginData)
        }

        const submitLogin = async () => {
            try {
                const response = await fetch('http://localhost:8080/api/login', requestOptions)
                const data = await response.json()
                console.log(data)
                if (data.response === "success") {
                    console.log("logged in, moving to home")
                    navigate('/home')
                }

            } catch (err) {
                console.log(err)
            }
        }

        
        submitLogin()
    }

    const handleChange = (e) => {
        const { id, value } = e.target
        setLoginData(prevState => ({
            ...prevState,
            [id]: value
        }))
    }
    return {
        loginData,
        handleChange,
        handleSubmit
    }
}

export const useLogOut = () => {
    const [logOut, setLogOut] = useState(false)
    const navigate = useNavigate()


    useEffect(() => {
        const sendLogOut = async () => {
            const requestOptions = {
                method: "DELETE",
                credentials: 'include'
            }
            try {
                const response = await fetch('http://localhost:8080/api/logout', requestOptions)
                const data = await response.json()
                console.log(data)
                if (data.response === "success") {
                    navigate('/login')
                    closeSocket()
                }
            } catch (err) {
                console.log(err)
            }
        }
        if (logOut) {
            sendLogOut()
            setLogOut(false)
        }
    }, [logOut, navigate])
    const handleLogOut = () => {
        setLogOut(true)
    }
    return handleLogOut

}

export const useRegister = () => {

    const [registerData, setRegisterData] = useState({
        name: '',
        email: '',
        password: '',
        firstName: '',
        lastName: '',
        dateOfBirth: '',
        privacy: 'public',
        avatar: null,
        nickname: '',
        aboutMe: ''
    })

    const handleChange = (e) => {
        const { id, value } = e.target
        setRegisterData(prevState => ({
            ...prevState,
            [id]: value
        }))
    }

    const handleFileChange = (e) => {
        const file = e.target.files[0];
        if (file) {
            setRegisterData(prevState => ({
                ...prevState,
                avatar: file
            }))
        }
    }
    const navigate = useNavigate()

    const handleSubmit = async (e) => {
        e.preventDefault()

        const formData = new FormData()

        Object.keys(registerData).forEach(key => {
            if (registerData[key] !== null) {
                formData.append(key, registerData[key])
            }
        })

        try {
            const response = await fetch('http://localhost:8080/api/register', {
                method: 'POST',
                body: formData,
                mode: 'cors'
            })
            const data = await response.json()
            console.log(data)
            if (data.response === "success") {
                navigate('/login')
            }

        } catch (err) {
            console.log(err)
        }
    }

    return {
        registerData,
        handleChange,
        handleFileChange,
        handleSubmit
    }

}

export const useCreatePost = (groupId) => {
    const [postData, setPostdata] = useState({
        title: '',
        content: '',
        avatar: null,
        privacy: 'public',
        groupId: groupId
    })

    const navigate = useNavigate()
    const handleSubmit = async (e) => {
        e.preventDefault()
        const formData = new FormData()

        formData.append('title', postData.title)
        formData.append('content', postData.content)
        formData.append('privacy', postData.privacy)
        formData.append('groupId', postData.groupId)

        if (postData.avatar) {
            formData.append('avatar', postData.avatar)
        }

        const requestOptions = {
            method: 'POST',
            body: formData,
            mode: 'cors',
            credentials: 'include',
        }

        try {
            const response = await fetch('http://localhost:8080/api/addPost', requestOptions)
            const data = await response.json()
            console.log(data)
            if (!data.loggedIn) {
                navigate('/login')
            }
        } catch (err) {
            console.log(err)
        }
    }

    const handleChange = (e) => {
        const { id, value } = e.target
        setPostdata(prevState => ({
            ...prevState,
            [id]: value
        }))
    }

    const handleFileChange = (e) => {
        const file = e.target.files[0]
        if (file) {
            setPostdata(prevState => ({
                ...prevState,
                avatar: file
            }))
        }
    }

    return {
        postData,
        handleChange,
        handleFileChange,
        handleSubmit
    }
}

export const useGetAllUsers = () => {

    const [allUsers, setAllUsers] = useState([])
    const [loading, setLoading] = useState(true)

    const requestOptions = {
        method: "GET",
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include'
    }
    const navigate = useNavigate()

    useEffect(() => {
        const getAllUsers = async () => {
            try {
                let response = await fetch('http://localhost:8080/api/getAllUsers', requestOptions)
                let data = await response.json()
                console.log(data)
                if (!data.loggedIn) {
                    navigate('/')
                }
                setAllUsers(data.getAllUsers)
            } catch (err) {
                console.log(err)
                return null
            } finally {
                setLoading(false)
            }
        }
        getAllUsers()
    }, [])

    return {
        allUsers,
        loading
    }
}



export const useGetUser = (id, refreshTrigger) => {
    const [userData, setUserData] = useState(null);
    const [loading, setLoading] = useState(true);
    const [error, setError] = useState(null);
    const navigate = useNavigate()

    useEffect(() => {
        const getProfile = async () => {
            setLoading(true)
            const requestOptions = {
                method: 'GET',
                headers: { 'Content-Type': 'application/json' },
                credentials: 'include',
            };

            try {
                let response = await fetch(`http://localhost:8080/api/getUser/${id}`, requestOptions);
                if (!response.ok) {
                    throw new Error('Failed to fetch user data');
                }
                let data = await response.json();
                if (!data.loggedIn) {
                    navigate('/login')
                    return
                }
                setUserData(data);
            } catch (err) {
                setError(err.message);
            } finally {
                setLoading(false);
            }
        };
        getProfile();
    }, [id, refreshTrigger, navigate]);

    return { userData, loading, error };
};


export const useGetAllPosts = (refreshTrigger) => {

    const [allPosts, setAllposts] = useState([])
    const navigate = useNavigate()

    useEffect(() => {
        const FetchAllPosts = async () => {
            try {
                const response = await fetch('http://localhost:8080/api/getAllPosts', {
                    method: 'GET',
                    credentials: 'include'
                })
                const data = await response.json()
                console.log(data)
                if (!data.loggedIn) {
                    navigate('/login')
                }
                if (data) {
                    setAllposts(data.getAllPostsPrivate)
                }
            } catch (err) {
                console.log(err)
                return null
            }
        }
        FetchAllPosts()
    }, [refreshTrigger])

    return { allPosts }

}
export const useGetGroupData = (groupId, refreshTrigger) => {

    const [groupData, setGroupData] = useState([])
    const [loading, setLoading] = useState(true)
    const navigate = useNavigate()

    useEffect(() => {
        const fetchGroupData = async () => {

            const requestOptions = {
                method: 'POST',
                credentials: 'include',
                body: JSON.stringify({ "groupId": groupId })
            }
            try {
                const response = await fetch('http://localhost:8080/api/getGroupData', requestOptions)
                const data = await response.json()
                console.log(data)
                if (!data.loggedIn) {
                    navigate('/login')
                }
                if (data) {
                    setGroupData(data)
                }
            } catch (err) {
                console.log(err)
            } finally {
                setLoading(false)
            }
        }
        fetchGroupData()
    }, [groupId, refreshTrigger])

    return { groupData, loading }

}

export const useGetOnePost = (id) => {
    const navigate = useNavigate()

    useEffect(() => {
        const fetchPost = async () => {
            const requestOptions = {
                method: "GET",
                credentials: 'include',
                body: JSON.stringify({ "id": id })
            }
            try {
                const response = await fetch('http://localhost:8080/api/getPost', requestOptions)
                const data = await response.json()
                console.log(data)
                if (!data.loggedIn) {
                    navigate('/login')
                }
            } catch (err) {
                console.log(err)
            }
        }
        fetchPost()
    }, [id])

}


export const useSetComment = (postId, refreshComments) => {

    const [commentData, setCommentData] = useState({
        content: '',
        avatar: null,
    })

    const navigate = useNavigate()

    const handleSubmit = async (e) => {
        e.preventDefault()
        const formData = new FormData()

        formData.append('content', commentData.content)
        formData.append('postId', postId)

        if (commentData.avatar) {
            formData.append('avatar', commentData.avatar)
        }

        const requestOptions = {
            method: 'POST',
            body: formData,
            credentials: 'include'
        }

        try {
            const response = await fetch('http://localhost:8080/api/addComment', requestOptions)
            const data = await response.json()
            console.log(data)
            if (!data.loggedIn) {
                navigate('/login')
            }
            if (data.response === "success") refreshComments()
        } catch (err) {
            console.log(err)
        }
    }

    const handleChange = (e) => {
        const { id, value } = e.target
        setCommentData(prevState => ({
            ...prevState,
            [id]: value
        }))
    }

    const handleFileChange = (e) => {
        const file = e.target.files[0]
        if (file) {
            setCommentData(prevState => ({
                ...prevState,
                avatar: file
            }))
        }
    }
    return {
        commentData,
        handleFileChange,
        handleChange,
        handleSubmit
    }
}

export const useAddFollow = (followId, useContactCreated) => {

    const [isFollowing, setIsFollowing] = useState(false)
    const navigate = useNavigate()
    const addFollow = async () => {
        const requestOptions = {
            credentials: 'include',
            method: "POST",
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ "id": Number(followId) })
        }
        console.log(requestOptions)
        console.log(followId)

        try {
            const response = await fetch('http://localhost:8080/api/addFollow', requestOptions)
            const data = await response.json()
            console.log(data)
            if (!data.loggedIn) {
                navigate('/login')
            }
            if (data.response === "success") {
                if (useContactCreated) useContactCreated()
                setIsFollowing(true)
            }
        } catch (err) {
            console.log("error adding follower", err)
            return err
        }
    }
    return {
        isFollowing,
        addFollow
    }
}

export const useRespondNotification = (refreshSidebar) => {

    const navigate = useNavigate()

    const sendNotificationResponse = async (idRef, type, response) => {
        console.log(idRef)
        const responseData = {
            type: type,
            idRef: idRef,
            response: response,
        }
        const requestOptions = {
            method: "POST",
            credentials: "include",
            body: JSON.stringify(responseData),
            header: { 'Content-Type': 'application/json' }
        }
        try {
            const response = await fetch('http://localhost:8080/api/respondNotification', requestOptions)
            const data = await response.json()
            console.log(data)
            if (!data.loggedIn) {
                navigate('/login')
            }
            if (refreshSidebar) refreshSidebar()
        } catch (err) {
            console.log(err)
        }
    }
    return sendNotificationResponse

}

export const useGetNotifications = (refreshTrigger) => {
    const [notificationData, setNotificationData] = useState(null)
    const [loading, setLoading] = useState(true)
    const navigate = useNavigate()

    useEffect(() => {
        const fetchNotifications = async () => {
            setLoading(true)
            const requestOptions = {
                method: "GET",
                credentials: "include"
            }
            try {
                const response = await fetch('http://localhost:8080/api/getNotifications', requestOptions)
                const data = await response.json()
                // console.log(data)
                if (!data.loggedIn) {
                    navigate('/login')
                }
                if (data.response === "success") {
                    setNotificationData(data)
                }
                return data
            } catch (err) {
                console.log(err)
                return err
            } finally {
                setLoading(false)
            }
        }
        fetchNotifications()
    }, [navigate, refreshTrigger])

    return {
        notificationData,
        loading
    }

}

export const useCreateGroup = (onGroupCreated) => {

    const [groupData, setGroupData] = useState({
        title: '',
        description: ''
    })

    const navigate = useNavigate()
    const handleChange = (e) => {
        const { id, value } = e.target
        setGroupData(prevState => ({
            ...prevState,
            [id]: value
        }))
    }
    const handleSubmit = (e) => {
        e.preventDefault()

        const requestOptions = {
            method: "POST",
            credentials: "include",
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(groupData)
        }

        const AddGroup = async () => {
            try {
                const response = await fetch('http://localhost:8080/api/createGroup', requestOptions)
                const data = await response.json()
                console.log(data)
                if (!data.loggedIn) {
                    navigate('/login')
                }
                if (data.response === "success") {
                    if (onGroupCreated) onGroupCreated()
                    navigate(`/group/${data.groupId}`)
                }
            } catch (err) {
                console.log(err)
            }
        }
        AddGroup()
    }

    return {
        groupData,
        handleChange,
        handleSubmit
    }
}



export const useSendGroupJoinRequest = () => {
    const navigate = useNavigate();

    const sendGroupJoinRequest = async (groupId) => {
        const requestOptions = {
            method: 'POST',
            credentials: 'include',
            body: JSON.stringify({ "id": groupId })
        }
        try {
            const response = await fetch('http://localhost:8080/api/requestGroupJoin', requestOptions)
            const data = await response.json()
            console.log(data)
            if (!data.loggedIn) {
                navigate('/login')
            }
        } catch (err) {
            console.log(err)
        }
    }

    return sendGroupJoinRequest;
}


export const useCreateEvent = (groupId) => {

    const [eventData, setEventData] = useState({
        groupId: groupId,
        title: '',
        description: '',
        time: ''
    })
    const navigate = useNavigate()

    const handleChange = (e) => {

        const { id, value } = e.target
        setEventData(prevState => ({
            ...prevState,
            [id]: value,
        }))
    }

    const handleSubmit = async (e) => {
        console.log("time:", eventData.time)
        e.preventDefault()
        const requestOptions = {
            method: 'POST',
            credentials: 'include',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify(eventData),
        }

        try {
            const response = await fetch('http://localhost:8080/api/createEvent', requestOptions)
            const data = await response.json()
            console.log(data)
            if (!data.loggedIn) {
                navigate('/login')
            }
        } catch (err) {
            console.log(err)
        }
    }

    return {
        eventData,
        handleChange,
        handleSubmit
    }
}

export const useGetContacts = (refreshTrigger) => {
    const [contacts, setContacts] = useState(null)
    const [loading, setLoading] = useState(true)
    const navigate = useNavigate()

    useEffect(() => {
        const getContacts = async () => {
            const requestOptions = {
                method: "GET",
                credentials: "include"
            }
            try {
                const response = await fetch('http://localhost:8080/api/getContacts', requestOptions)
                const data = await response.json()
                console.log(data)
                if (!data.loggedIn) {
                    navigate('/login')
                }
                if (data.response === "success") {
                    setContacts(data)
                }
            } catch (err) {
                console.log(err)
            } finally {
                setLoading(false)
            }

        }
        getContacts()
    }, [navigate, refreshTrigger])
    return {
        contacts,
        loading
    }
}


export const useGetComments = (postId, refreshCommentTrigger) => {
    const [comments, setComments] = useState([])
    const [loading, setLoading] = useState(true)
    const navigate = useNavigate()
    useEffect(() => {
        const getComments = async () => {
            const requestOptions = {
                method: "POST",
                credentials: "include",
                body: JSON.stringify({ "id": Number(postId) })
            }
            try {
                const response = await fetch("http://localhost:8080/api/getComments", requestOptions)
                const data = await response.json()
                console.log("getcomment:", data)
                if (!data.loggedIn) {
                    navigate('/login')
                }
                if (data.response === "success") {
                    setComments(data.comments || [])
                }
            } catch (err) {
                console.log(err)
            } finally {
                setLoading(false)
            }
        }
        getComments()
    }, [postId, navigate, refreshCommentTrigger])
    return { comments, loading }
}

export const useGetMessages = (userId) => {
    const groupId = 0
    const navigate = useNavigate()
    const [messages, setMessages] = useState([])
    const [loading, setLoading] = useState(true)
    useEffect(() => {
        const getMessages = async () => {
            const requestOptions = {
                method: "POST",
                credentials: 'include',
                body: JSON.stringify({ "userId": userId, "groupId": groupId })
            }
            try {
                const response = await fetch('http://localhost:8080/api/getMessages', requestOptions)
                const data = await response.json()
                console.log(data)
                if (!data.loggedIn) {
                    navigate('/login')
                }
                if (data.response === "success") {
                    setMessages(data)
                }

            } catch (err) {
                console.log(err)
            } finally {
                setLoading(false)
            }
        }
        getMessages()
    }, [userId])
    return {
        messages,
        loading
    }
}
export const useGetGroupMessages = (groupId) => {
    const userId = 0
    const navigate = useNavigate()
    const [messages, setMessages] = useState([])
    const [loading, setLoading] = useState(true)
    useEffect(() => {
        const getMessages = async () => {
            const requestOptions = {
                method: "POST",
                credentials: 'include',
                body: JSON.stringify({ "userId": userId, "groupId": groupId })
            }
            try {
                const response = await fetch('http://localhost:8080/api/getMessages', requestOptions)
                const data = await response.json()
                console.log(data)
                if (!data.loggedIn) {
                    navigate('/login')
                }
                if (data.response === "success") {
                    setMessages(data)
                }
            } catch (err) {
                console.log(err)
            } finally {
                setLoading(false)
            }
        }
        getMessages()
    }, [groupId])
    return {
        messages,
        loading
    }
}

export const useCheckLoggedIn = () => {

    const [userData, setUserData] = useState(null);
    const [loading, setLoading] = useState(true);
    const navigate = useNavigate()
    const requestOptions = {
        method: 'GET',
        credentials: 'include',
    }

    useEffect(() => {
        const checkLoggedIn = async () => {
            try {
                const result = await fetch('http://localhost:8080/api/checkLogin', requestOptions)
                const data = await result.json()
                console.log(data)
                if (!data.loggedIn) {
                    navigate('/login')
                }
                if (data.response === "success") {
                    setUserData(data)
                }
            } catch (err) {
                console.log(err)
            } finally {
                setLoading(false)
            }
        }
        checkLoggedIn()
    }, [navigate])

    return {
        userData,
        loading
    }
}

export const useGetAllGroups = () => {
    const [groupData, setGroupData] = useState([])
    const [loading, setLoading] = useState(true)
    const navigate = useNavigate()

    useEffect(() => {
        const fetchGroupData = async () => {
            const requestOptions = {
                method: 'GET',
                credentials: 'include'
            }
            try {
                const response = await fetch('http://localhost:8080/api/getAllGroups', requestOptions)
                const data = await response.json()
                console.log(data)
                if (!data.loggedIn) {
                    navigate('/login')
                }
                if (data.response === "success") {
                    setGroupData(data)
                }
            } catch (err) {
                console.log(err)
            } finally {
                setLoading(false)
            }
        }
        fetchGroupData()
    }, [])
    return {
        groupData,
        loading
    }
}

export const UnFollow = (userId, onContactCreated) => {
    const unFollow = async () => {
        const requestOptions = {
            method: 'DELETE',
            credentials: 'include',
            body: JSON.stringify({ "id": userId }),
        }
        try {
            const response = await fetch('http://localhost:8080/api/unFollow', requestOptions)
            const data = await response.json()
            console.log(data)
            if (data.response === "success") onContactCreated()
        } catch (err) {
            console.log(err)
        }
    }
    console.log("calling unfollow")
    unFollow()
}

export const useChangePrivacy = () => {
    const navigate = useNavigate()
    const changePrivacy = async (privacy) => {
        const requestOptions = {
            method: 'POST',
            credentials: 'include',
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ 'privacy': privacy })
        }
        console.log(privacy)
        try {
            const response = await fetch("http://localhost:8080/api/changePrivacy", requestOptions)
            const data = await response.json()
            console.log(data)
            if (!data.loggedIn) {
                navigate('/login')
            }
            if (data.response === "success") {
            }
        } catch (err) {
            console.log(err)
        }
    }
    return changePrivacy
}

export const useGetGroupInviteUsers = (groupId) => {
    const [users, setUsers] = useState([])
    const [loading, setLoading] = useState(false)
    const navigate = useNavigate()

    const fetchUsers = async () => {
        setLoading(true)
        const requestOptions = {
            method: "POST",
            credentials: "include",
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ "groupId": groupId })
        }
        try {
            const response = await fetch('http://localhost:8080/api/getGroupInviteUsers', requestOptions)
            const data = await response.json()
            console.log(data)
            if (!data.loggedIn) {
                navigate('/login')
            }
            if (data.response === "success") {
                setUsers(data.users)
            }
        } catch (err) {
            console.log(err)
        } finally {
            setLoading(false)
        }
    }
    return {
        users, loading, fetchUsers
    }
}

export const useSendGroupInvite = () => {
    const navigate = useNavigate()
    const sendInvite = async (groupId, userId) => {
        const requestOptions = {
            method: 'POST',
            credentials: "include",
            headers: { 'Content-Type': 'application/json' },
            body: JSON.stringify({ "groupId": groupId, 'userId': userId })
        }
        try {
            const response = await fetch('http://localhost:8080/api/sendGroupInvite', requestOptions)
            const data = await response.json()
            console.log(data)
            if (!data.loggedIn) {
                navigate('/login')
            }
        } catch (err) {
            console.log(err)
        }
    }
    return { sendInvite }
}