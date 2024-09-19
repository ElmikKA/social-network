export const getInitials = (name) => {
    const nameParts = name.split(' ');
    const firstInitial = nameParts[0] ? nameParts[0][0] : '';
    const lastInitial = nameParts[1] ? nameParts[1][0] : '';
    return `${firstInitial}${lastInitial}`.toUpperCase();
};

export function timeAgo(dateString) {
    const now = new Date();
    const past = new Date(dateString);
    const diffInMs = now - past; // Difference in milliseconds

    const diffInSeconds = Math.floor(diffInMs / 1000);
    const diffInMinutes = Math.floor(diffInSeconds / 60);
    const diffInHours = Math.floor(diffInMinutes / 60);
    const diffInDays = Math.floor(diffInHours / 24);

    if (diffInSeconds < 60) {
        return `${diffInSeconds} seconds ago`;
    } else if (diffInMinutes < 60) {
        return `${diffInMinutes} minutes ago`;
    } else if (diffInHours < 24) {
        return `${diffInHours} hours ago`;
    } else if (diffInDays < 30) {
        return `${diffInDays} days ago`;
    } else {
        // If the date is more than 30 days ago, return a formatted date
        return past.toLocaleDateString();
    }
}