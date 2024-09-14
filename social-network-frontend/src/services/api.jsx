import { useState, useEffect } from "react"

export const useLogin = () => {
    const [loginData, setLoginData] = useState({
        name: '',
        password: ''
    })

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

export const useCreatePost = () => {
    const [postData, setPostdata] = useState({
        title: '',
        content: '',
        avatar: null,
        privacy: 'public',
        groupId: '0'
    })

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

    const requestOptions = {
        method: "GET",
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include'
    }

    useEffect(() => {
        const getAllUsers = async () => {
            try {
                let response = await fetch('http://localhost:8080/api/getAllUsers', requestOptions)
                let data = await response.json()
                setAllUsers(data.getAllUsers)
            } catch (err) {
                console.log(err)
                return null
            }
        }
        getAllUsers()
    }, [])

    return {
        allUsers
    }
}


export const useGetUser = (id) => {
    // missing all user made posts and followers/following

    const [userData, setUserData] = useState({
        name: '',
        email: '',
        password: '',
        firstName: '',
        lastName: '',
        dateOfBirth: '',
        avatar: '',
        avatarMimeType: '',
        nickname: '',
        aboutMe: ''
    })

    const requestOptions = {
        method: 'GET',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include'
    }

    useEffect(() => {
        const getProfile = async () => {
            try {
                let response = await fetch(`http://localhost:8080/api/getUser/${id}`, requestOptions)
                let data = await response.json()
                console.log(data)
                if (data) {
                    setUserData(data.getUser)
                }
            } catch (err) {
                console.log(err)
                return null
            }
        }
        getProfile()
    }, [])

    return {
        userData
    }
}

export const useGetAllPosts = () => {

    const [allPosts, setAllposts] = useState([])

    useEffect(() => {
        const FetchAllPosts = async () => {
            try {
                const response = await fetch('http://localhost:8080/api/getAllPosts', {
                    method: 'GET',
                    credentials: 'include'
                })
                const data = await response.json()
                if (data) {
                    setAllposts(data.getAllPosts)
                }
            } catch (err) {
                console.log(err)
                return null
            }
        }
        FetchAllPosts()
    }, [])

    return { allPosts }

}

export const useGetOnePost = (id) => {


    const [postData, setPostData] = useState({
        id: '',
        userId: '',
        groupId: '',
        creator: '',
        title: '',
        content: '',
        avatar: '',
        createdAt: '',
        pricacy: ''
    })

    useEffect(() => {
        const fetchPost = async () => {
            try {
                const response = await fetch(`http://localhost:8080/api/getPost/${id}`, {
                    method: "GET",
                    credentials: 'include'
                })
                const data = await response.json()
                setPostData(data.post)
                console.log(data)
            } catch (err) {
                console.log(err)
            }
        }
        fetchPost()
    }, [])


    return {
        postData
    }

}


export const useSetComment = (postId) => {

    const [commentData, setCommentData] = useState({
        content: '',
        avatar: null,
    })


    const handleSubmit = async (e) => {
        e.preventDefault()
        const formData = new FormData()

        formData.append('content', commentData.content)
        formData.append('groupId', commentData.groupId)
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

export const useAddFollow = (followId) => {
    useEffect(() => {
        const addFollow = async () => {
            const requestOptions = {
                credentials: 'include',
                method: "GET"
            }

            try {
                const response = await fetch(`http://localhost:8080/api/addFollow/${followId}`, requestOptions)
                const data = await response.json()
                console.log(data)
            } catch (err) {
                console.log("error adding follower", err)
                return err
            }
        }
        addFollow()
    }, [])
    return null
}

export const useRespondNotification = (idRef, type, response) => {

    const responseData = {
        type: type,
        idRef: idRef,
        response: response,
    }
    useEffect(() => {

        const sendNotificationResponse = async () => {
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
            } catch (err) {
                console.log(err)
            }
        }
        sendNotificationResponse()

    }, [])
}

export const useCreateGroup = () => {

    const [groupData, setGroupData] = useState({
        title: '',
        description: ''
    })

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

export const useSendGroupJoinRequest = (groupId) => {
    useEffect(() => {
        const sendGroupJoinRequest = async () => {
            const requestOptions = {
                method: 'POST',
                credentials: 'include',
            }
            try {
                const response = await fetch(`http://localhost:8080/api/requestGroupJoin/${groupId}`, requestOptions)
                const data = await response.json()
                console.log(data)
            } catch (err) {
                console.log(err)
            }
        }
        sendGroupJoinRequest()
    }, [])
}


export const useCreateEvent = (groupId) => {

    const [eventData, setEventData] = useState({
        groupId: groupId,
        title: '',
        description: '',
        time: ''
    })

    const handleChange = (e) => {

        const { id, value } = e.target
        setEventData(prevState => ({
            ...prevState,
            [id]: value,
        }))
    }

    const handleSubmit = async (e) => {
        e.preventDefault()
        const requestOptions = {
            method: 'POST',
            credentials: 'include',
            header: { 'Content-Type': 'application/json' },
            body: JSON.stringify(eventData),
        }

        try {
            const response = await fetch('http://localhost:8080/api/createEvent', requestOptions)
            const data = await response.json()
            console.log(data)
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