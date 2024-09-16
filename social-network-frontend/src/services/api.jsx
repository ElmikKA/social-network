import { useState, useEffect } from "react"

export const useLogin = () => {


    const [loginData, setLoginData] = useState({
        name: '',
        password: ''
    })

    const requestOptions = {
        method: 'POST',
        headers: { 'Content-Type': 'application/json' },
        credentials: 'include',
        body: JSON.stringify(loginData)
    }

    const handleSubmit = (e) => {
        e.preventDefault()
        const submitLogin = async () => {
            try {

                fetch('http://localhost:8080/api/login', requestOptions)
                    .then(Response => Response.json())
                    .then(data => console.log(data))

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
        avatar: null,
        // avatarMimeType: '',
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

    const handleSubmit = (e) => {
        e.preventDefault()

        const formData = new FormData()

        Object.keys(registerData).forEach(key => {
            if (registerData[key] !== null) {
                formData.append(key, registerData[key])
            }
        })

        try {
            fetch('http://localhost:8080/api/register', {
                method: 'POST',
                body: formData,
                mode: 'cors'
            })
                .then(Response => Response.json())
                .then(data =>  {
                    console.log(data)
                    return data
        })
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

    const handleSubmit = (e) => {
        e.preventDefault()
        const formData = new FormData()

        formData.append('title', postData.title)
        formData.append('content', postData.content)
        formData.append('privacy', postData.privacy)
        formData.append('groupId', postData.groupId)

        if (postData.avatar) {
            console.log("adding avatar")
            formData.append('avatar', postData.avatar)
        }


        try {
            fetch('http://localhost:8080/api/addPost', {
                method: 'POST',
                body: formData,
                mode: 'cors',
                credentials: 'include'
            })
                .then(Response => {
                    if (Response.ok) {
                        Response.json()
                    }
                }
                )
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


export const useGetUser = () => {

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
                let response = await fetch('http://localhost:8080/api/getUser', requestOptions)
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
                console.log(data)
                setPostData(data.post)
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


    const handleSubmit = (e) => {
        e.preventDefault()
        const formData = new FormData()

        formData.append('content', commentData.content)
        formData.append('groupId', commentData.groupId)
        formData.append('postId', postId)

        if (commentData.avatar) {
            formData.append('avatar', commentData.avatar)
        }

        try {
            fetch('http://localhost:8080/api/addComment', {
                method: 'POST',
                body: formData,
                credentials: 'include'
            })
                .then(Response => {
                    if (Response.ok) {
                        Response.json()
                    }
                }
                )
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
            try {
                await fetch(`http://localhost:8080/api/addFollow/${followId}`, {
                    credentials: 'include',
                    method: "GET"
                })
            } catch (err) {
                console.log("error adding follower", err)
                return err
            }
        }
        addFollow()
    }, [])
    return null
}

export const RespondFollow = (toUserId, pending) => {
    // toUserId is the id of the person who sent you the follow request
    // pending is "completed" or "rejected" depending of if you want to accept or reject the follow

    useEffect(() => {
        const acceptRejectFollow = async () => {
            try {
                const response = await fetch(`http://localhost:8080/api/respondFollow/${toUserId}`, {
                    credentials: 'include',
                    method: "POST",
                    body: JSON.stringify({ pending: `${pending}` })
                })
                console.log(response)
            } catch (err) {
                console.log("error confirm/reject follow", err)
                return err
            }
        }
        acceptRejectFollow()
    }, [])
    return null
}