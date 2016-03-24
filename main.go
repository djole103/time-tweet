package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/dghubble/go-twitter/twitter"
	"github.com/dghubble/oauth1"
	"os"
)

func PostTweet(client *twitter.Client, tweet string) error {
	_, _, err := client.Statuses.Update(tweet, nil)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines, scanner.Err()
}

// writeLines writes the lines to the given file.
func writeLines(lines []string, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func LoadTweet(FILENAME string) (string, error) {
	f, err := os.Open(FILENAME)
	if err != nil {
		fmt.Printf("Couldn't open file: %s\n", FILENAME)
	}
	reader := bufio.NewReader(f)
	tweet, isPrefix, err := reader.ReadLine()
	if isPrefix {
		fmt.Println("It's a prefex")
		return "", nil
	}
	if err != nil {
		fmt.Println("Failed to read line")
		fmt.Println(err)
		return "", err
	}
	return string(tweet), nil
}

func CheckValid(tweet string) error {
	if len(tweet) > 140 {
		return fmt.Errorf("More than 140 chars")
	}
	return nil
}

func main() {
	consumer_key := os.Getenv("TWITTER_CONS_KEY")
	consumer_secret := os.Getenv("TWITTER_CONS_SECRET")
	access_token := os.Getenv("TWITTER_ACCESS_TOKEN")
	access_secret := os.Getenv("TWITTER_ACCESS_SECRET")
	FILENAME := flag.String("tweets", "tweets", "Path to tweets")
	flag.Parse()
	var err error
	var tweets []string

	config := oauth1.NewConfig(consumer_key, consumer_secret)
	token := oauth1.NewToken(access_token, access_secret)

	httpClient := config.Client(oauth1.NoContext, token)
	client := twitter.NewClient(httpClient)
	fmt.Println(*FILENAME)
	tweets, err = readLines(*FILENAME)

	if err != nil {
		fmt.Println("Failed to load tweets")
		fmt.Println(err)
	}
	if len(tweets) == 0 {
		fmt.Println("Refill your tweets dawg")
		os.Exit(1)
	}
	err = CheckValid(tweets[0])
	if err != nil {
		fmt.Println("Failed to validate tweet")
		fmt.Println(err)
		os.Exit(1)
	}
	PostTweet(client, tweets[0])
	writeLines(tweets[1:], *FILENAME)
}
