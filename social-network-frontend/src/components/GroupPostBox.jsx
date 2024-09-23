import CommentToggle from "./ui/CommentToggle"
import { PostItem } from "./PostBox"

const GroupPostBox = ({ posts }) => {
    return (
        <div className='groupPostMain'>
            {posts && posts.length > 0 ? (
                posts.map((post) => (
                    <PostItem key={post.id} post={post}></PostItem>
                ))
            ) : (
                <div className='no-posts-or-private-profile'>
                    <p>No posts</p>
                </div>
            )}
        </div>
    )
}

export default GroupPostBox
