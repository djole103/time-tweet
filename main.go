package main

import (
	//"bufio"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"os"
)

const (
	FILENAME = "tweets"
)

func PostTweet(client *twitter.Client, tweet string) error {
	_, _, err := client.Statuses.Update(tweet, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func main() {
	consumer_key := os.Getenv("TWITTER_CONS_KEY")
	consumer_secret := os.Getenv("TWITTER_CONS_SECRET")
	access_token := os.Getenv("TWITTER_ACCESS_TOKEN")
	access_secret := os.Getenv("TWITTER_ACCESS_SECRET")

	config := oauth1.NewConfig(consumer_key, consumer_secret)
	token := oauth1.NewToken(access_token, access_secret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)

	PostTweet(client, "testing2")
}

/*func parseTweets() []string {
	tweets := []string{}
	f, _ := os.Open(FILENAME)
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
}*/

func validTweet(tweet string) bool {
	if len(tweet) > 140 {
		return false
	}
	return true
}
