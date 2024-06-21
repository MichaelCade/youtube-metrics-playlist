```markdown
## Getting Started

### Step 1: Clone the Repository
Start by cloning this repository to your local machine.
```
git clone https://github.com/michaelcade/youtube-metrics-playlist.git
cd youtube-metrics-playlist
```

### Step 2: Obtain YouTube API Key, Playlist ID, and Channel ID
1. **YouTube API Key**: Follow the instructions [here](https://developers.google.com/youtube/registering_an_application) to obtain your YouTube API key.
2. **Playlist ID**: Find the playlist ID from the playlist URL. It's usually after the `list=` parameter.
3. **Channel ID**: You can find the channel ID in the channel's URL or in the advanced settings of the channel's account.

### Step 3: Set Environment Variables or Use Command-Line Arguments
You have two options to provide the necessary information to the application:

#### Option A: Environment Variables
Set the following environment variables:
- `YOUTUBE_API_KEY`: Your YouTube API key.
- `YOUTUBE_PLAYLIST_ID`: The ID of the playlist you want to analyze.
- `YOUTUBE_CHANNEL_ID`: The ID of the channel you're interested in.

On macOS or Linux, you can set these variables in your terminal like so:
```
export YOUTUBE_API_KEY='your_youtube_api_key_here'
export YOUTUBE_PLAYLIST_ID='your_playlist_id_here'
export YOUTUBE_CHANNEL_ID='your_channel_id_here'
```

For Windows, use the following in Command Prompt:
```
set YOUTUBE_API_KEY=your_youtube_api_key_here
set YOUTUBE_PLAYLIST_ID=your_playlist_id_here
set YOUTUBE_CHANNEL_ID=your_channel_id_here
```

#### Option B: Command-Line Arguments
Alternatively, you can provide the information as command-line arguments when running the application:
```
go run main.go -apiKey="your_youtube_api_key_here" -playlistId="your_playlist_id_here" -channelId="your_channel_id_here"
```

### Step 4: Run the Application
Depending on your choice in Step 3, navigate to the project directory and run the application using the environment variables or command-line arguments as shown above.
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.