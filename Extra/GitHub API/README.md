Call the GitHub API. 

If I run curl https://api.github.com/users/alextldr in the terminal, I get the below info:

{
  "login": "AlexTLDR",
  "id": 66307303,
  "node_id": "MDQ6VXNlcjY2MzA3MzAz",
  "avatar_url": "https://avatars.githubusercontent.com/u/66307303?v=4",
  "gravatar_id": "",
  "url": "https://api.github.com/users/AlexTLDR",
  "html_url": "https://github.com/AlexTLDR",
  "followers_url": "https://api.github.com/users/AlexTLDR/followers",
  "following_url": "https://api.github.com/users/AlexTLDR/following{/other_user}",
  "gists_url": "https://api.github.com/users/AlexTLDR/gists{/gist_id}",
  "starred_url": "https://api.github.com/users/AlexTLDR/starred{/owner}{/repo}",
  "subscriptions_url": "https://api.github.com/users/AlexTLDR/subscriptions",
  "organizations_url": "https://api.github.com/users/AlexTLDR/orgs",
  "repos_url": "https://api.github.com/users/AlexTLDR/repos",
  "events_url": "https://api.github.com/users/AlexTLDR/events{/privacy}",
  "received_events_url": "https://api.github.com/users/AlexTLDR/received_events",
  "type": "User",
  "site_admin": false,
  "name": "Alex Badragan",
  "company": null,
  "blog": "",
  "location": "Stuttgart",
  "email": null,
  "hireable": null,
  "bio": "Aspiring Gopher\r\n",
  "twitter_username": null,
  "public_repos": 11,
  "public_gists": 0,
  "followers": 4,
  "following": 1,
  "created_at": "2020-06-02T10:14:21Z",
  "updated_at": "2022-08-25T11:18:45Z"
}

Create a function that requests the Login, Name and Number of Repos:
