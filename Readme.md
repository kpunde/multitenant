# Multi-Tenant Application Setup

This repository contains a multi-tenant React frontend application and a Go backend. The application dynamically renders content based on the `tenant_id` extracted from the subdomain (e.g., `tenant1.localhost` or `tenant2.localhost`).

## Prerequisites

- [Node.js](https://nodejs.org/en/download/) (for running the React app)
- [Go](https://golang.org/doc/install) (for running the backend)
- A local machine (Linux, MacOS, or Windows) where you can modify the `/etc/hosts` file.

## Setup Instructions

### Step 1: Modify `/etc/hosts` File

To test subdomain-based routing locally, you need to modify the `/etc/hosts` file on your machine. This will map the subdomains `tenant1.localhost` and `tenant2.localhost` to `127.0.0.1`.

1. Open a terminal and edit the `/etc/hosts` file:

   ```bash
   sudo nano /etc/hosts
   ```

2. Add the following lines to the file:

   ```plaintext
   127.0.0.1 tenant1.localhost
   127.0.0.1 tenant2.localhost
   ```

3. Save and close the file (`Ctrl + O`, then `Ctrl + X`).

If you are on Windows, edit the `C:\Windows\System32\drivers\etc\hosts` file instead.

### Step 2: Install Dependencies

Both the Go backend and the React frontend need to have their dependencies installed before starting.

#### Backend (Go)

1. Navigate to the backend directory:

   ```bash
   cd backend
   ```

2. Sync Go modules (install dependencies):
   ```bash
   go mod tidy
   ```

#### Frontend (React)

1. Navigate to the frontend directory:

   ```bash
   cd frontend
   ```

2. Install the React app dependencies:
   ```bash
   npm install
   ```

### Step 3: Start the Applications

#### Starting the Go Backend

1. In the `backend` directory, start the Go backend server:

   ```bash
   go run main.go
   ```

   The Go backend will now be running on `http://localhost:8080`.

#### Starting the React Frontend

1. In the `frontend` directory, start the React development server:

   ```bash
   npm start
   ```

   The React app will now be running on `http://localhost:3000`.

### Step 4: Access the Application in the Browser

Now that the backend and frontend are running, you can access the multi-tenant app using the subdomains you added to the `/etc/hosts` file.

1. Open your browser and go to:
   - `http://tenant1.localhost:3000` to simulate the first tenant.
   - `http://tenant2.localhost:3000` to simulate the second tenant.

The application should dynamically load content based on the tenant ID (`tenant1` or `tenant2`), which is extracted from the subdomain.

### Script for Automating Setup and Testing

You can automate the setup process (starting Go backend, React frontend, and modifying `/etc/hosts`) by using the provided `setup.sh` script. Make sure to run the script with superuser privileges to modify `/etc/hosts`.

#### Script: `setup.sh`

```bash
#!/bin/bash

# Modify /etc/hosts for tenant1 and tenant2 subdomains
echo "Modifying /etc/hosts to add tenant1.localhost and tenant2.localhost..."
sudo -- sh -c -e "echo '127.0.0.1 tenant1.localhost' >> /etc/hosts"
sudo -- sh -c -e "echo '127.0.0.1 tenant2.localhost' >> /etc/hosts"
echo "Hosts file updated."

# Start Go backend
echo "Starting Go backend..."
cd backend
go mod tidy
go run main.go &
GO_PID=$!

# Start React frontend
echo "Starting React frontend..."
cd ../frontend
npm install
npm start &
REACT_PID=$!

# Wait for user to hit CTRL+C to stop both processes
trap "kill $GO_PID; kill $REACT_PID" INT

wait
```

### How to Use the Script

1. Make the script executable:

   ```bash
   chmod +x setup.sh
   ```

2. Run the script:
   ```bash
   ./setup.sh
   ```

This script will:

- Add the `tenant1.localhost` and `tenant2.localhost` entries to your `/etc/hosts` file.
- Start both the Go backend and React frontend in parallel.

To stop the services, press `CTRL + C`, and both the Go and React apps will stop.

### Troubleshooting

- Ensure the Go backend is running on `http://localhost:8080` before trying to access the frontend.
- Check the Network tab in the browser's Developer Tools to ensure the `tenant_id` header is being passed in requests to the backend.
- If the tenant subdomains do not work, double-check that the `/etc/hosts` file is correctly updated and saved.

### License

This project is licensed under the MIT License.
