# Resum-API
A golang API with some hard coded values that make up my resume. v1 is super simplistic, but as I find more motivation I will begin pulling information dynamically through linkedin and other sources.

## How do I use this thing?
So bascially you can curl the endpoints or navigate to them via a web browser. Everything will be text so even a cli based web based browser is fair game.

### Endpoints
```
joey - displays a message
contact - displays contact information including twitter and linkedin
jobs - see my job history including accomplishments
skills - a list of hardcoded endorsements I have received on linkedin
default - directs you to this repo for the documentation
```

### Example:
```
$ curl joeygrasso.com/api/contact
{
 "email": "me@joeygrasso.com",
 "github": "github.com/joeygrasso",
 "linkedin": "linkedin.com/in/grassojoey",
 "twitter": "@grassojoey",
 "web": "JoeyGrasso.com"
}

```

## Future Features?
I'd like to make this a bit more dynamic to pull data from other sources. I'd like to also add a couple more endpoints including a project section to show off some interesting projects as well. I'll also include the ability to filter output like get_recent_job to get only my current job. If I integrate linkedin, the skills portion could include the count of endorsements and then be sorted asc|desc or whatever. We'll see what happens!
