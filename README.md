# Personal Apps Hub

A central dashboard for monitoring the status and last updated information for all personal apps in the workspace.

## Features

- **Status Monitoring**: Shows which apps are online (port check) or offline
- **Last Updated**: Displays when each app was last updated based on git commit history
- **Auto-refresh**: Automatically refreshes status every 30 seconds
- **Modern UI**: Clean, responsive interface with dark mode support

## Architecture

- **Backend**: Go HTTP server (port 8518)
- **Frontend**: React + Vite + TypeScript + Tailwind CSS (port 8517)

## Setup

### Backend

From the `backend` directory:

```bash
cd backend
go run .
```

The backend will start on `http://localhost:8518` (or the port specified in the `PORT` environment variable).

### Frontend

From the `frontend` directory:

```bash
cd frontend
npm install
npm run dev
```

The frontend will start on `http://localhost:8517`.

You can configure the backend URL by creating a `.env` file:

```bash
cp .env.example .env
# Edit .env and set VITE_API_URL if needed (default: http://localhost:8518)
```

## How It Works

1. **Port Checking**: The backend checks if each app's configured port(s) are listening on localhost
2. **Git History**: For each app, the backend runs `git log` to get the last commit timestamp
3. **Status Aggregation**: All status information is combined and returned via the `/api/status` endpoint
4. **Frontend Display**: The React frontend fetches and displays the status in a card-based layout

## Monitored Apps

- **audiophile**: Ports 8000 (backend), 5173 (frontend)
- **blippy**: Port 6900 (Next.js)
- **contacts-app**: Ports 4001 (backend), 4000 (frontend)
- **digital-leatherman**: Ports 8100 (backend), 5273 (frontend)
- **slowtube**: Ports 3001 (backend), 5200 (frontend)

## Development

### Backend

The backend uses Go's standard library and follows the same patterns as `digital-leatherman/backend`:
- Standard `net/http` server
- Middleware for logging and recovery
- CORS support for local development

### Frontend

The frontend uses:
- React 19 with TypeScript
- Vite for build tooling
- Tailwind CSS v4 for styling
- Fetch API for backend communication

## Future Enhancements

- Add links to open each app
- Show more detailed information (uptime, health checks)
- Add filtering and sorting options
- Deploy as a standalone service
