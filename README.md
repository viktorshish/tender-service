# tender-service

Сервис проведения тендеров

## Installation

```bash
# create .env and fill your data
$ cp .env.dist .env

# run app 
$ make start

# migrate
$ make migrate-up

# enjoy!
```

## Tender Management API Documentation

A concise RESTful API for managing tenders within organizations. Built with Go, Gin, and GORM, this API allows
authorized users to create and publish tenders, while enabling all users to view available tenders.

### Ping

Endpoint: *GET /ping*

Description:
Health check endpoint to verify that the API is running.

Request:

	•	Method: GET
	•	URL: /ping

Response:

	•	Status Code: 200 OK
	•	Body:

```json
"pong"
```

### Create Tender

Endpoint: *POST /api/tenders/new*

Description:
Create a new tender associated with the user’s organization. Only users who are trusted representatives of the
organization can create tenders.

Access:
Protected (Requires creatorUsername in the request body to verify responsibility)

Headers:

•	Content-Type: application/json

Request Body:

```json
{
    "name": "Tender Name",
    "description": "Detailed description of the tender",
    "serviceType": "Service Type",
    "organizationId": "uuid-organization-id",
    "creatorUsername": "creator_username"
}
```

Fields:

•	name (string, required): Name of the tender.

•	description (string, optional): Description of the tender.

•	serviceType (string, optional): Type of service related to the tender.

•	organizationId (UUID, required): ID of the organization.

•	creatorUsername (string, required): Username of the tender creator.

Response:

•	Status Code: 200 OK

•	Body:

```json
{
    "id": "uuid-tender-id",
    "name": "Tender Name",
    "description": "Detailed description of the tender",
    "serviceType": "Service Type",
    "organizationId": "uuid-organization-id",
    "creatorUsername": "creator_username",
    "status": "Created"
}
```

### Publish Tender

Endpoint: **PUT /api/tenders/:id/publish**

Description:
Publish a tender, making it visible to all users. Only trusted representatives of the organization can publish tenders.

Access:
Protected (Requires creatorUsername in the request body to verify responsibility)

Headers:

	•	Content-Type: application/json

Path Parameters:

	•	id (UUID): ID of the tender to publish.

```json
{
    "creatorUsername": "creator_username"
}
```

Fields:

•	creatorUsername (string, required): Username of the user attempting to publish the tender.

Response:

•	Status Code: 200 OK

•	Body:

```json
{
    "id": "uuid-tender-id",
    "name": "Tender Name",
    "description": "Detailed description of the tender",
    "serviceType": "Service Type",
    "organizationId": "uuid-organization-id",
    "creatorUsername": "creator_username",
    "status": "Published"
}
```

### Get Tenders

Endpoint: **GET /api/tenders**

Description:
Retrieve a list of tenders. Authenticated users can see their own tenders as well as all published tenders. All users
can view published tenders without authentication.

Access:
Public

Query Parameters:

•	serviceType (string, optional): Filter tenders by service type.

Response:

•	Status Code: 200 OK

•	Body:

```json
[
    {
        "id": "uuid-tender-id-1",
        "name": "Tender Name 1",
        "description": "Detailed description of tender 1",
        "serviceType": "Service Type 1",
        "status": "Published",
        "organizationName": "Organization Name 1"
    },
    {
        "id": "uuid-tender-id-2",
        "name": "Tender Name 2",
        "description": "Detailed description of tender 2",
        "serviceType": "Service Type 2",
        "status": "Published",
        "organizationName": "Organization Name 2"
    }
    // More tenders...
]
```
