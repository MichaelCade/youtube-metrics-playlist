package main

import (
	"context"
	"flag"
	"fmt"
	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
	"log"
	"net/http"
	"os"
)

func main() {
	ctx := context.Background()
	apiKeyFlag := flag.String("apiKey", "", "YouTube API key")
	playlistIdFlag := flag.String("playlistId", "", "YouTube Playlist ID")
	channelIdFlag := flag.String("channelId", "", "YouTube Channel ID")
	flag.Parse()

	apiKey := *apiKeyFlag
	if apiKey == "" {
		apiKey = os.Getenv("YOUTUBE_API_KEY")
	}
	playlistId := *playlistIdFlag
	if playlistId == "" {
		playlistId = os.Getenv("YOUTUBE_PLAYLIST_ID")
	}
	channelId := *channelIdFlag
	if channelId == "" {
		channelId = os.Getenv("YOUTUBE_CHANNEL_ID")
	}

	if apiKey == "" || playlistId == "" || channelId == "" {
		log.Fatal("Missing API key, playlist ID, or channel ID. Please provide them as command-line arguments or environment variables.")
	}

	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error creating YouTube service: %v", err)
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "index.html")
	})

	http.HandleFunc("/metrics", func(w http.ResponseWriter, r *http.Request) {
		playlistTitle, totalViews, err := fetchPlaylistDetails(service, playlistId)
		if err != nil {
			http.Error(w, "Failed to fetch playlist details", http.StatusInternalServerError)
			return
		}
		channelTitle, subscribers, channelViews, videoCount, err := fetchChannelDetails(service, channelId)
		if err != nil {
			http.Error(w, "Failed to fetch channel details", http.StatusInternalServerError)
			return
		}
		fmt.Fprintf(w, "<p>Playlist: %s<br>Total Playlist Views: %d<br>Channel: %s<br>Subscribers: %d<br>Total Channel Views: %d<br>Video Count: %d</p>",
			playlistTitle, totalViews, channelTitle, subscribers, channelViews, videoCount)
	})

	fmt.Println("Starting server at :8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(http.ListenAndServe(":8080", nil))
	}
}

func fetchPlaylistDetails(service *youtube.Service, playlistId string) (string, uint64, error) {
	playlistCall := service.Playlists.List([]string{"snippet"}).Id(playlistId)
	playlistResponse, err := playlistCall.Do()
	if err != nil {
		return "", 0, err
	}
	playlistTitle := ""
	if len(playlistResponse.Items) > 0 {
		playlistTitle = playlistResponse.Items[0].Snippet.Title
	}

	call := service.PlaylistItems.List([]string{"contentDetails"}).PlaylistId(playlistId).MaxResults(50)
	response, err := call.Do()
	if err != nil {
		return "", 0, err
	}

	totalViews := uint64(0)
	for _, item := range response.Items {
		videoCall := service.Videos.List([]string{"statistics"}).Id(item.ContentDetails.VideoId)
		videoResponse, err := videoCall.Do()
		if err != nil {
			return "", 0, err
		}
		for _, video := range videoResponse.Items {
			totalViews += video.Statistics.ViewCount
		}
	}

	return playlistTitle, totalViews, nil
}

func fetchChannelDetails(service *youtube.Service, channelId string) (string, uint64, uint64, uint64, error) {
	channelCall := service.Channels.List([]string{"snippet", "statistics"}).Id(channelId)
	channelResponse, err := channelCall.Do()
	if err != nil {
		return "", 0, 0, 0, err
	}
	channelTitle, subscribers, channelViews, videoCount := "", uint64(0), uint64(0), uint64(0)
	for _, channel := range channelResponse.Items {
		channelTitle = channel.Snippet.Title
		subscribers = channel.Statistics.SubscriberCount
		channelViews = channel.Statistics.ViewCount
		videoCount = channel.Statistics.VideoCount
	}
	return channelTitle, subscribers, channelViews, videoCount, nil
}
