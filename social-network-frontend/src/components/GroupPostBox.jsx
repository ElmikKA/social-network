
const GroupPostBox = ({ posts }) => {
    return (
        <div className='groupEventDiv'>
            {posts && posts.length > 0 ? (
                posts.map((post) => (
                    <div className='groupEventBox' key={post.id}>
                        <p>title:{post.title}</p>
                        <p>when:{post.content}</p>
                    </div>
                ))
            ) : (<p>No posts</p>)}
        </div>
    )
}

export default GroupPostBox
