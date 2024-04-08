package spotify

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"testing"
)

func TestUserHasTracks(t *testing.T) {
	client, server := testClientString(http.StatusOK, `[ false, true ]`)
	defer server.Close()

	contains, err := client.UserHasTracks(context.Background(), "0udZHhCi7p1YzMlvI4fXoK", "55nlbqqFVnSsArIeYSQlqx")
	if err != nil {
		t.Error(err)
	}
	if l := len(contains); l != 2 {
		t.Error("Expected 2 results, got", l)
	}
	if contains[0] || !contains[1] {
		t.Error("Expected [false, true], got", contains)
	}
}

func TestAddTracksToLibrary(t *testing.T) {
	client, server := testClientString(http.StatusOK, "")
	defer server.Close()

	err := client.AddTracksToLibrary(context.Background(), "4iV5W9uYEdYUVa79Axb7Rh", "1301WleyT98MSxVHPZCA6M")
	if err != nil {
		t.Error(err)
	}
}

func TestAddTracksToLibraryFailure(t *testing.T) {
	client, server := testClientString(http.StatusUnauthorized, `
{
  "error": {
    "status": 401,
    "message": "Invalid access token"
  }
}`)
	defer server.Close()
	err := client.AddTracksToLibrary(context.Background(), "4iV5W9uYEdYUVa79Axb7Rh", "1301WleyT98MSxVHPZCA6M")
	if err == nil {
		t.Error("Expected error and didn't get one")
	}
}

func TestAddTracksToLibraryWithContextCancelled(t *testing.T) {
	client, server := testClientString(http.StatusOK, ``)
	defer server.Close()

	ctx, done := context.WithCancel(context.Background())
	done()

	err := client.AddTracksToLibrary(ctx, "4iV5W9uYEdYUVa79Axb7Rh", "1301WleyT98MSxVHPZCA6M")
	if !errors.Is(err, context.Canceled) {
		t.Error("Expected error and didn't get one")
	}
}

func TestRemoveTracksFromLibrary(t *testing.T) {
	client, server := testClientString(http.StatusOK, "")
	defer server.Close()

	err := client.RemoveTracksFromLibrary(context.Background(), "4iV5W9uYEdYUVa79Axb7Rh", "1301WleyT98MSxVHPZCA6M")
	if err != nil {
		t.Error(err)
	}
}

func TestGetSavedTracksFromLibrary(t *testing.T) {
	client, server := testClientFile(http.StatusOK, "test_data/user_saved_tracks.json")
	defer server.Close()

	tracks, err := client.GetSavedTracksFromLibrary(context.Background())
	if err != nil {
		t.Error(err)
	}

	if tracks.Total != 3 {
		t.Errorf("Got %d tracks, expected 3", tracks.Total)
	}
	expected := "55nlbqqFVnSsArIeYSQlqx"
	fmt.Println(tracks.Tracks)

	if tracks.Tracks[0].ID.String() != expected {
		t.Errorf("Got track ID of %s, expected first track ID in the response to be %s", tracks.Tracks[0].ID, expected)
	}

	if tracks.Tracks[0].Album.Name != "Love In The Future" {
		t.Errorf("Got album name of %s, expected first track album name in the response to be Love In The Future", tracks.Tracks[0].Album.Name)
	}

	if tracks.Tracks[0].Name != "You & I (Nobody In The World)" {
		t.Errorf("Got track name of %s, expected first track name in the response to be You & I (Nobody In The World)", tracks.Tracks[0].Name)
	}
}

func TestUserHasAlbums(t *testing.T) {
	client, server := testClientString(http.StatusOK, `[ false, true ]`)
	defer server.Close()

	contains, err := client.UserHasAlbums(context.Background(), "0udZHhCi7p1YzMlvI4fXoK", "55nlbqqFVnSsArIeYSQlqx")
	if err != nil {
		t.Error(err)
	}
	if l := len(contains); l != 2 {
		t.Error("Expected 2 results, got", l)
	}
	if contains[0] || !contains[1] {
		t.Error("Expected [false, true], got", contains)
	}
}

func TestAddAlbumsToLibrary(t *testing.T) {
	client, server := testClientString(http.StatusOK, "")
	defer server.Close()

	err := client.AddAlbumsToLibrary(context.Background(), "4iV5W9uYEdYUVa79Axb7Rh", "1301WleyT98MSxVHPZCA6M")
	if err != nil {
		t.Error(err)
	}
}

func TestAddAlbumsToLibraryFailure(t *testing.T) {
	client, server := testClientString(http.StatusUnauthorized, `
{
  "error": {
    "status": 401,
    "message": "Invalid access token"
  }
}`)
	defer server.Close()
	err := client.AddAlbumsToLibrary(context.Background(), "4iV5W9uYEdYUVa79Axb7Rh", "1301WleyT98MSxVHPZCA6M")
	if err == nil {
		t.Error("Expected error and didn't get one")
	}
}

func TestRemoveAlbumsFromLibrary(t *testing.T) {
	client, server := testClientString(http.StatusOK, "")
	defer server.Close()

	err := client.RemoveAlbumsFromLibrary(context.Background(), "4iV5W9uYEdYUVa79Axb7Rh", "1301WleyT98MSxVHPZCA6M")
	if err != nil {
		t.Error(err)
	}
}
