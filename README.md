# Authentication Service

This project implements a simple authentication service with two REST endpoints for issuing and refreshing authentication tokens.

## Task Description

### REST Endpoints

1. **Issue Tokens** (`/issue_tokens`)
   - Issues a pair of Access and Refresh tokens for a user identified by a GUID provided as a request parameter.

2. **Refresh Tokens** (`/refresh_tokens`)
   - Performs a refresh operation using an Access and Refresh token pair.

## Requirements

- **Access Token**:
  - JWT format.
  - Uses SHA512 algorithm.
  - Must not be stored in the database.

- **Refresh Token**:
  - Arbitrary format, transmitted in base64.
  - Stored in the database as a bcrypt hash.
  - Must be protected against client-side modifications and reuse attacks.

- **Token Association**:
  - Access and Refresh tokens must be linked.
  - A refresh operation can only be performed using the Refresh token issued together with the Access token.

- **Client IP Validation**:
  - The token payload must contain the client’s IP address.
  - If the client’s IP address changes during the refresh operation, an email warning must be sent to the user (mocked for this implementation).

## API Endpoints

### 1. Issue Tokens

**Request:**
```http
GET /issue_tokens?id={GUID}
```

**Response:**
```json
{
  "access_token": "<JWT_TOKEN>",
  "refresh_token": "<BASE64_REFRESH_TOKEN>"
}
```

### 2. Refresh Tokens

**Request:**
```http
POST /refresh_tokens
Content-Type: application/json

{
  "access_token": "<JWT_TOKEN>",
  "refresh_token": "<BASE64_REFRESH_TOKEN>"
}
```

**Response:**
```json
{
  "access_token": "<NEW_JWT_TOKEN>",
  "refresh_token": "<NEW_BASE64_REFRESH_TOKEN>"
}
```
