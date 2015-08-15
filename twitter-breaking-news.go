package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/ChimeraCoder/anaconda"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
)

//	Tweet holds everything about a single news item
type Tweet struct {
	Id         int64  `json:"id"`
	CreateTime int64  `json:"createtime"`
	Text       string `json:text`
	MediaUrl   string `json:url`
}

//	Set up our flags
var (
	port           = flag.Int("port", 3000, "The port to listen on")
	allowedOrigins = flag.String("allowedOrigins", "*", "A comma-separated list of valid CORS origins")
	consumerKey    = flag.String("consumerKey", "", "Your twitter consumer key")
	consumerSecret = flag.String("consumerSecret", "", "Your twitter consumer secret")
	authToken      = flag.String("authToken", "", "Your twitter auth token")
	authSecret     = flag.String("authSecret", "", "Your twitter auth secret")
)

func parseEnvironment() {
	//	Check for the listen port
	if env_port := os.Getenv("TWITTER_PORT"); env_port != "" {
		*port, _ = strconv.Atoi(env_port)
	}

	//	Check for allowed origins
	if env_origins := os.Getenv("TWITTER_ALLOWED_ORIGINS"); env_origins != "" {
		*allowedOrigins = env_origins
	}

	//	Check for consumer key
	if env_consumer_key := os.Getenv("TWITTER_CONSUMER_KEY"); env_consumer_key != "" {
		*consumerKey = env_consumer_key
	}

	//	Check for consumer secret
	if env_consumer_secret := os.Getenv("TWITTER_CONSUMER_SECRET"); env_consumer_secret != "" {
		*consumerSecret = env_consumer_secret
	}

	//	Check for auth token
	if env_auth_token := os.Getenv("TWITTER_AUTH_TOKEN"); env_auth_token != "" {
		*authToken = env_auth_token
	}

	//	Check for auth secret
	if env_auth_secret := os.Getenv("TWITTER_AUTH_SECRET"); env_auth_secret != "" {
		*authSecret = env_auth_secret
	}
}

func main() {

	//	Parse environment variables:
	parseEnvironment()

	//	Parse the command line for flags:
	flag.Parse()

	//	Create twitter api client:
	anaconda.SetConsumerKey(*consumerKey)
	anaconda.SetConsumerSecret(*consumerSecret)
	api := anaconda.NewTwitterApi(*authToken, *authSecret)

	r := mux.NewRouter()
	r.HandleFunc("/news/{twitterName}", func(w http.ResponseWriter, r *http.Request) {

		//	Parse the twitter name from the url
		twitterName := mux.Vars(r)["twitterName"]

		//	Our return values:
		tweets := []Tweet{}

		//	Set some url values:
		v := url.Values{}
		v.Set("screen_name", twitterName)
		v.Set("exclude_replies", "1")
		v.Set("include_rts", "false")
		v.Set("count", "50")

		timeline, _ := api.GetUserTimeline(v)

		//	Get the tweets with media (photos) and return them
		for _, tweet := range timeline {
			for _, media := range tweet.Entities.Media {

				//	If we found one with media, write out the
				//	tweet and the media and break out of the
				//	outer range loop
				tweetedTime, err := tweet.CreatedAtTime()

				if err == nil {
					tweets = append(tweets, Tweet{
						Id:         tweet.Id,
						CreateTime: tweetedTime.Unix(),
						Text:       tweet.Text,
						MediaUrl:   media.Media_url})
				}
			}
		}

		//	Set the content type header and return the JSON
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		json.NewEncoder(w).Encode(tweets)
	})

	//	CORS handler
	c := cors.New(cors.Options{
		AllowedOrigins:   strings.Split(*allowedOrigins, ","),
		AllowCredentials: true,
	})
	handler := c.Handler(r)

	//	Indicate what port we're starting the service on
	portString := strconv.Itoa(*port)
	fmt.Println("Allowed origins: ", *allowedOrigins)
	fmt.Println("Starting server on :", portString)
	http.ListenAndServe(":"+portString, handler)
}
