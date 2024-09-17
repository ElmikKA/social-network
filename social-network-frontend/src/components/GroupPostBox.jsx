
const GroupPostBox = ({ posts }) => {
    return (
        <div className='groupEventDiv'>
            {posts && posts.length > 0 ? (
                posts.map((post) => (
                    <div className='postBox' key={post.id}>
                        <p>title:{post.title}</p>
                        <p>when:{post.content}</p>
                        {post.avatar &&
                            <img
                                src={`http://localhost:8080/api/avatars/${post.avatar}`}
                                alt="Post Avatar"
                                style={{ width: '150px', height: '150px', borderRadius: '50%' }}
                            />
                        }
                    </div>
                ))
            ) : (<p>No posts</p>)}
        </div>
    )
}

export default GroupPostBox
