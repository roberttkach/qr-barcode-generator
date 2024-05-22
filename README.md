This is a Go-based API application that generates QR codes and barcodes from provided data. It uses the Gin web framework and various third-party libraries for QR code and barcode generation.

## Features

- Generate QR codes with customizable size and error correction level
- Generate barcodes (Code 128 format) with customizable size
- Logging to a file (`server.log`)
- Session management using cookies
- Error handling middleware
- Dockerized deployment

## Prerequisites

- Go (version 1.16 or later)
- Docker (for containerized deployment)

## Installation

1. Clone the repository:

```
git clone https://github.com/your-username/qr-barcode-api.git
```

2. Change to the project directory:

```
cd qr-barcode-api
```

3. Install the dependencies:

```
go get ./...
```

## Configuration

1. Create a `config.json` file in the project root directory with the following content:

```json
{
  "port": "8080"
}
```

You can change the port number if desired.

2. Set the `SESSION_SECRET` environment variable with a secret key for session management:

```
export SESSION_SECRET=your_secret_key
```

Replace `your_secret_key` with a secure random string.

## Running the Application

### Without Docker

```
go run main.go
```

The API will be available at `http://localhost:8080`.

### With Docker

1. Build the Docker image:

```
docker build -t qr-barcode-api .
```

2. Run the Docker container:

```
docker run -p 8080:8080 -e SESSION_SECRET=your_secret_key qr-barcode-api
```

Replace `your_secret_key` with the same secret key you used earlier.

The API will be available at `http://localhost:8080`.

## API Endpoints

### Generate QR Code

```
GET /create_qr
```

Query parameters:

- `data` (required): The data to be encoded in the QR code.
- `size` (optional, default: `256`): The size of the QR code image in pixels.
- `level` (optional, default: `Medium`): The error correction level (`Low`, `Medium`, `High`, or `Highest`).

Response: PNG image of the generated QR code.

### Generate Barcode

```
GET /create_barcode
```

Query parameters:

- `data` (required): The data to be encoded in the barcode.
- `size` (optional, default: `200`): The height of the barcode image in pixels.

Response: PNG image of the generated barcode.

## Testing

The project includes basic test cases for the QR code and barcode generation endpoints. To run the tests, execute the following command:

```
go test
```

## Logging

The application logs to a file named `server.log` in the project root directory. The log file will be created if it doesn't exist, and new log entries will be appended to the file.

## Contributing

Contributions are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request.
