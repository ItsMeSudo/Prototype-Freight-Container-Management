# REST API Documentation

This documentation provides information about the REST API endpoints available in my system.

## Base URL

The base URL for all API endpoints is `https://example.com/api`.

## API Versioning

Currently, there are two versions of the API available: v1 and v2.

- v1 endpoints: `/api`
- v2 endpoints: `/api/v2`

## Endpoints

### v1 Endpoints

#### Insert JSON Data
- **Endpoint:** `/api/containers`
- **Method:** POST
- **Description:** This endpoint allows you to insert JSON data into the database.

#### Handle CSV Binary Upload
- **Endpoint:** `/api/containers/import`
- **Method:** POST
- **Description:** This endpoint handles the upload of CSV files in binary format.

#### Get Block Information (GET)
- **Endpoint:** `/api/blocks/stat`
- **Method:** GET
- **Description:** This endpoint retrieves statistical information about blocks.

### v2 Endpoints

#### Get All Containers
- **Endpoint:** `/api/v2/getall`
- **Method:** POST
- **Description:** This endpoint retrieves all containers.

#### Handle CSV Form Upload
- **Endpoint:** `/api/v2/importcsv`
- **Method:** POST
- **Description:** This endpoint handles the upload of CSV files via form data.

#### Get Block Information (POST)
- **Endpoint:** `/api/v2/getstat`
- **Method:** POST
- **Description:** This endpoint retrieves statistical information about blocks.

#### Insert JSON File Data
- **Endpoint:** `/api/v2/importjson`
- **Method:** POST
- **Description:** This endpoint handles the upload of JSON files via form data.

## Response Format

All endpoints return responses in JSON format. ((Except critical errors))

## Error Handling

Errors are returned with appropriate HTTP status codes and error messages in the response body.
