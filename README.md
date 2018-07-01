# Go client for Emotions API by Microsoft

## Use

Create API client

`emo := emotions.NewClient(emotionsAPIKey)`

Get emotions from URL

`faces, err := emo.GetEmotions(photoURL)`

## Install

Add to import

`"github.com/aafanasev/ms-emotions-go"`
