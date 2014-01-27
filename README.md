go-scrub-mp3-id3-tags
=====================
A quick hack to scrub out camelot wheel keys and energy from mp3 id3 title tag. In order for this to work I had to first convert all of my mp3's to ID3v2.3. It seems that iTunes or some of my music was in ID3v2.4 and the library I'm depending on ([http://github.com/mikkyang/id3-go](http://github.com/mikkyang/id3-go)) only supports reading / editing upto ID3v2.3.

Just a way I found to easily fix my mistake of configuring Mixed in Key incorrectly. My apologies in advance if it's not up to your programming standard.

Setup / Run
-----------
1. go get "github.com/mikkyang/id3-go"
1. Edit the code if you want to modify your mp3 tags you'll have to "set_flag" to true.
1. go run scrub.go "/path/to/music/"
