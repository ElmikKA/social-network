# social-network
A **social network platform** similar to Facebook or Instagram where users can create profiles, follow each other, post updates, and participate in group chats. The platform includes features like notifications, posts, comments, and real-time messaging. The project is built with a **React frontend** and a **Go (Golang) backend**, containerized with Docker for ease of deployment.

## Audit
https://github.com/01-edu/public/tree/master/subjects/social-network/audit

# Database migrations

``goose -s create add new_table sql`` (make a new sql table file)

``goose up`` (goes up to latest version)

``goose up-by-one`` (goes up one table)

``goose down`` (drops the latest table)

``kooduser@pupil:~/js/social-network/backend$ GOOSE_DRIVER=sqlite3 GOOSE_DBSTRING=./db/sn.db goose -s -dir ./db/migrations/ create add_new_posts_table sql``

## Clone the repository
```bash
   git clone https://github.com/your-username/social-network.git
```

## Running the Project with Docker

1. Build and run Container
```bash
   docker-compose up --build
```

2. Access the application
- Frontend: http://localhost

3. Stop the container
```bash
   docker-compose down
```


## Running the project without the Docker
Open two terminales

### Backend Setup
1. In one of the terminals navigate to backend directory:
```bash
   cd backend
```

2. Run the backend:
```bash
   go run .
```

### Frontend Setup
1. In one of the terminals navigate to frontend directory:
```bash
   cd social-network-frontend
```

2. Install dependencies:
```bash
   npm install
```

3. Install dependencies:
```bash
   npm run dev
```

- The frontend will be running at http://localhost:5173.