package main

import (
	"github.com/dghubble/go-twitter/twitter"
)

const (
	FILENAME = "tweets"
)

func main() {
	config := oauth1.NewConfig("conskey", "conssecret")
	token := oauth.NewToken()
	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

}

func parseTweets() []string {
	tweets := []string{}
	f := os.Open(FILENAME)
	reader := bufio.NewReader(f)
	for tweet, isPrefix, err := reader.ReadLine(); tweet != nil && err != nil; tweet, isPrefix, err := reader.ReadLine() {
		if isPrefix {
			return
			//not sure
		}
		if validTweet(tweet) {
			append(tweets, tweet)
		}
	}
	return tweets
}

func validTweet(tweet string) {
	if len(tweet) > 140 {
		return false
	}
	return true
}
