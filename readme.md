```markdown
# YouTube Channel Metrics Dashboard

This Go application provides a simple web server that serves a dashboard displaying YouTube channel metrics, including playlist details and channel statistics. It's designed to demonstrate how to integrate with the YouTube Data API v3 and serve dynamic content using Go's built-in HTTP server capabilities.

## Prerequisites

Before you begin, ensure you have met the following requirements:
- You have installed [Go](https://golang.org/dl/) (version 1.13 or later recommended).
- You have a basic understanding of Go programming.
- You have access to a Google account for using the YouTube Data API v3.

## Getting Started

Follow these steps to get your YouTube Channel Metrics Dashboard running:

### Step 1: Clone the Repository

Clone this repository to your local machine using:

```bash
git clone https://github.com/michaelcade/youtube-channel-metrics.git
cd youtube-channel-metrics
```

### Step 2: Obtain a YouTube Data API Key

1. Visit the [Google Developers Console](https://console.developers.google.com/).
2. Create a new project.
3. Navigate to `Credentials`, then click on `Create credentials` > `API key`.
4. Enable the YouTube Data API v3 for your project.

### Step 3: Set Up Environment Variables

Set the following environment variables with your API key, YouTube playlist ID, and channel ID:

```bash
export YOUTUBE_API_KEY='YOUR_API_KEY'
export YOUTUBE_PLAYLIST_ID='YOUR_PLAYLIST_ID'
export YOUTUBE_CHANNEL_ID='YOUR_CHANNEL_ID'
```

### Step 4: Run the Application

Run the application with:

```bash
go run main.go
```

The server will start on `localhost:8080`. Open your web browser and navigate to `http://localhost:8080` to view the dashboard.

## Contributing

Contributions are welcome! Feel free to open an issue or submit a pull request.

## License

This project is licensed under the MIT License - see the LICENSE.md file for details.
```