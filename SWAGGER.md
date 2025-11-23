# Social Backend API Documentation

## Swagger API Documentation

This project includes comprehensive Swagger/OpenAPI documentation for all endpoints.

### Accessing the Documentation

Once the server is running, you can access the interactive API documentation at:

**🌐 Swagger UI**: <http://localhost:3000/swagger/index.html>

### Available Documentation Formats

- **Interactive UI**: `/swagger/index.html` - Full interactive documentation with "Try it out" functionality
- **JSON Format**: `/swagger/doc.json` - Raw OpenAPI JSON specification  
- **YAML Format**: `/swagger/swagger.yaml` - OpenAPI YAML specification

### API Overview

#### Authentication Endpoints

- `POST /api/register` - Register a new user account
- `POST /api/login` - Login and receive JWT token

#### Post Management

- `GET /api/posts` - Get all posts (paginated)
- `POST /api/posts` - Create a new post (authenticated)
- `GET /api/posts/{id}` - Get a specific post
- `PUT /api/posts/{id}` - Update a post (authenticated, owner only)
- `DELETE /api/posts/{id}` - Delete a post (authenticated, owner only)
- `GET /api/users/{userId}/posts` - Get posts by specific user

#### Like System

- `POST /api/posts/{id}/like` - Toggle like on a post (authenticated)
- `GET /api/posts/{id}/like` - Get like status and count

#### Comment System  

- `GET /api/posts/{postId}/comments` - Get comments for a post (paginated)
- `POST /api/posts/{postId}/comments` - Create a comment (authenticated)
- `GET /api/comments/{id}` - Get a specific comment
- `PUT /api/comments/{id}` - Update a comment (authenticated, owner only)
- `DELETE /api/comments/{id}` - Delete a comment (authenticated, owner only)
- `GET /api/users/{userId}/comments` - Get comments by specific user

### Authentication

Most endpoints require authentication using JWT Bearer tokens. To authenticate:

1. **Register** or **Login** to get a JWT token
2. **Include the token** in the `Authorization` header: `Bearer YOUR_JWT_TOKEN`
3. **Use the "Authorize" button** in Swagger UI to set your token globally

### Response Format

All API responses follow a consistent format:

```json
{
  "success": true,
  "message": "Operation successful", 
  "data": { /* response data */ },
  "meta": { /* optional metadata like pagination */ }
}
```

### Error Responses

Error responses include appropriate HTTP status codes and descriptive messages:

```json
{
  "success": false,
  "error": "Error description"
}
```

### Generating Documentation

If you make changes to the API annotations, regenerate the documentation:

```bash
# Install swag if not already installed
go install github.com/swaggo/swag/cmd/swag@latest

# Generate/update documentation
swag init -g cmd/server/main.go
```

### Development Workflow

1. **Add Swagger annotations** to controller methods
2. **Update DTO structs** with example tags for better documentation
3. **Run `swag init`** to regenerate docs
4. **Test endpoints** using the Swagger UI

### Example Usage

1. **Start the server**: `./bin/server`
2. **Open Swagger UI**: <http://localhost:3000/swagger/index.html>
3. **Register a user** using the `/api/register` endpoint
4. **Login** to get your JWT token
5. **Click "Authorize"** and enter `Bearer YOUR_TOKEN`
6. **Try out authenticated endpoints** like creating posts or comments

Happy coding! 🚀
