package requests

var ActionMap map[ActionID]func = {
	TWEET:      sendTweet
	BURSTTWEET: sendBurstTweets
}

func sendTweet(tweet) {
	//
}

func sendBurstTweets(tweets){
	//
}
