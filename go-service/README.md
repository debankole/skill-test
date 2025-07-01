# Go PDF Report Microservice

> **Note:** The Node.js service does not provide the actual student retrieval endpoint, so for the purposes of this microservice, the following JSON structure was assumed for the student response:
>
> ```json
> {
>   "id": 1,
>   "name": "John Doe",
>   "email": "john.doe@example.com",
>   "systemAccess": true,
>   "phone": "123-456-7890",
>   "gender": "Male",
>   "dob": "2005-09-15",
>   "class": "10",
>   "section": "A",
>   "roll": "23",
>   "fatherName": "Richard Doe",
>   "fatherPhone": "111-222-3333",
>   "motherName": "Jane Doe",
>   "motherPhone": "444-555-6666",
>   "guardianName": "Uncle Bob",
>   "guardianPhone": "777-888-9999",
>   "relationOfGuardian": "Uncle",
>   "currentAddress": "123 Main St, Springfield",
>   "permanentAddress": "456 Elm St, Springfield",
>   "admissionDate": "2020-06-01",
>   "reporterName": "Admin User"
> }
> ```

## Overview

This microservice generates PDF reports for students by consuming data from an existing Node.js backend API. It is designed as a standalone service and exposes a REST API endpoint to generate and download student reports as PDFs.

## General Approach

- **API Integration:** The service fetches student data from the Node.js backend's `/api/v1/students/:id` endpoint.
- **PDF Generation:** The service uses the [gofpdf](https://github.com/phpdave11/gofpdf) library to generate PDF reports from the fetched JSON data.
- **Modular Design:**
  - **Handler:** Handles HTTP requests and responses.
  - **Service:** Contains business logic for fetching data and generating PDFs.
  - **Client:** Responsible for communicating with the Node.js backend. An interface and a mock implementation are provided for testing.
  - **Middleware:** Includes authentication and logging middleware.
- **Mock Support:** For testing and development, a mock client can be used to serve example student data from an embedded JSON file.

## Environment Variables

- `PORT` - The port on which the Go service will run (default: `8080`).
- `NODE_BACKEND_URL` - The base URL of the Node.js backend (default: `http://localhost:3000`).
- `API_KEY` - The API key required in the `X-API-Key` header for authentication.
- `USE_MOCK` - If set to `true`, the service will use the mock client with embedded example data instead of calling the real Node.js backend. This is useful for local development and testing.

Example `.env`:
```env
PORT=8080
NODE_BACKEND_URL=http://localhost:3000
API_KEY=your-secure-api-key
USE_MOCK=true
```

## Running the Service

1. Ensure the Node.js backend and PostgreSQL database are running (unless using the mock).
2. Set the required environment variables (see above).
3. Start the Go service:
   ```sh
   go run main.go
   ```
4. Make a request to generate a PDF report:
   ```sh
   curl -H "X-API-Key: your-secure-api-key" http://localhost:8080/api/v1/students/1/report --output student_report.pdf
   ```

## Running the Integration Test

The integration test (`main_test.go`) starts the Go service as a subprocess, sends a real HTTP request to the PDF report endpoint, and verifies the response.

To run the test:
```sh
go test -v
```

The test sets `USE_MOCK=true` to ensure the service uses the embedded mock data for predictable results.

---

**Note:**
- The mock client loads example student data from `client/testdata/student_example.json` using Go's `embed` feature.
- You can modify this JSON file to change the mock data returned during tests or when `USE_MOCK=true`. 