package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/api/option"
	"google.golang.org/api/youtube/v3"
)

func main() {
	ctx := context.Background()
	apiKey := os.Getenv("YOUTUBE_API_KEY")
	playlistId := os.Getenv("YOUTUBE_PLAYLIST_ID")
	channelId := os.Getenv("YOUTUBE_CHANNEL_ID") // Assuming you have this

	if apiKey == "" || playlistId == "" || channelId == "" {
		log.Fatal("Missing API key, playlist ID, or channel ID in environment variables")
	}

	service, err := youtube.NewService(ctx, option.WithAPIKey(apiKey))
	if err != nil {
		log.Fatalf("Error creating YouTube service: %v", err)
	}
	// Fetch the playlist's title
	playlistCall := service.Playlists.List([]string{"snippet"}).Id(playlistId)
	playlistResponse, err := playlistCall.Do()
	if err != nil {
		log.Fatalf("Error fetching playlist details: %v", err)
	}
	playlistTitle := ""
	if len(playlistResponse.Items) > 0 {
		playlistTitle = playlistResponse.Items[0].Snippet.Title
	}

	fmt.Printf("Playlist: %s\n", playlistTitle) // Print the playlist title

	// Fetch playlist details
	call := service.PlaylistItems.List([]string{"contentDetails"}).PlaylistId(playlistId).MaxResults(50)
	response, err := call.Do()
	if err != nil {
		log.Fatalf("Error calling YouTube API: %v", err)
	}

	totalViews := uint64(0) // Initialize totalViews correctly as uint64
	for _, item := range response.Items {
		videoCall := service.Videos.List([]string{"statistics"}).Id(item.ContentDetails.VideoId)
		videoResponse, err := videoCall.Do()
		if err != nil {
			log.Fatalf("Error fetching video details: %v", err)
		}
		for _, video := range videoResponse.Items {
			totalViews += video.Statistics.ViewCount // Directly add the uint64 ViewCount to totalViews
		}
	}

	fmt.Printf("Total Playlist Views: %d\n\n", totalViews)

	// Fetch channel details
	channelCall := service.Channels.List([]string{"snippet", "statistics"}).Id(channelId)
	channelResponse, err := channelCall.Do()
	if err != nil {
		log.Fatalf("Error fetching channel details: %v", err)
	}
	for _, channel := range channelResponse.Items {
		fmt.Printf("Channel Title: %s\nSubscribers: %d\nTotal Views: %d\nVideo Count: %d\n\n",
			channel.Snippet.Title, channel.Statistics.SubscriberCount, channel.Statistics.ViewCount, channel.Statistics.VideoCount)
	}
}
