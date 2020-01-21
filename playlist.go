package main

import "fmt"

type Song struct {
	Name     string
	Artist   string
	NextSong *Song
}

type Playlist struct {
	Name       string
	Start      *Song
	NowPlaying *Song
}

func createPlaylist(name string) *Playlist {
	return &Playlist{
		Name: name,
	}
}

func (p *Playlist) addSong(name string, artist string) error {
	song := &Song{
		Name:   name,
		Artist: artist,
	}

	if p.Start == nil {
		p.Start = song
	} else {
		currentSong := p.Start
		for currentSong.NextSong != nil {
			currentSong = currentSong.NextSong
		}
		currentSong.NextSong = song
	}
	fmt.Printf("Song '%s' added to playlist.\n", name)
	return nil
}

func (p *Playlist) showPlaylist() error {
	currentSong := p.Start
	if currentSong == nil {
		fmt.Println("Playlist is empty")
		return nil
	}
	fmt.Printf("%s by %s\n", currentSong.Name, currentSong.Artist)
	for currentSong.NextSong != nil {
		currentSong = currentSong.NextSong
		fmt.Printf("%s by %s\n", currentSong.Name, currentSong.Artist)
	}
	return nil
}

func (p *Playlist) startPlaying() error {
	if p.Start != nil {
		p.NowPlaying = p.Start
	}
	return nil
}

func (p *Playlist) playNextSong() error {
	p.NowPlaying = p.NowPlaying.NextSong
	return nil
}

func (p *Playlist) removeSong(name string) error {
	currentSong := p.Start
	if currentSong == nil {
		fmt.Println("Playlist is empty")
		return nil
	}
	if currentSong.Name == name {
		p.Start = currentSong.NextSong
	}
	previousSong := currentSong
	for currentSong.NextSong != nil {
		currentSong = currentSong.NextSong
		if currentSong.Name == name {
			previousSong.NextSong = currentSong.NextSong
			break
		}
		previousSong = currentSong
	}
	return nil
}

func main() {
	//create playlist
	playlist := createPlaylist("Playlist 1")
	fmt.Println("Playlist created")
	fmt.Println("Adding songs to playlist")
	playlist.addSong("song 1", "author 1")
	playlist.addSong("song 2", "author 2")
	playlist.addSong("song 3", "author 3")
	playlist.addSong("song 4", "author 4")
	playlist.addSong("song 5", "author 5")
	fmt.Println("Playlist")
	playlist.showPlaylist()
	fmt.Println("start playing playlist")
	playlist.startPlaying()
	fmt.Printf("Now streaming %s by %s \n", playlist.NowPlaying.Name, playlist.NowPlaying.Artist)
	fmt.Println("Changing to next song")
	playlist.playNextSong()
	fmt.Printf("Now streaming %s by %s \n", playlist.NowPlaying.Name, playlist.NowPlaying.Artist)
	fmt.Println("Changing to next song")
	playlist.playNextSong()
	fmt.Printf("Now streaming %s by %s \n", playlist.NowPlaying.Name, playlist.NowPlaying.Artist)
	fmt.Println("remove songs")
	playlist.removeSong("song 3")
	fmt.Println("Playlist")
	playlist.showPlaylist()
}
