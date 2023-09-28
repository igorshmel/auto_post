### Auto_post
This application enables automatic scheduling and random posting to a VK.ru group. The posts consist of artistic works sourced from platforms like ArtStation. The process is as follows: an administrator browses new artworks on ArtStation. If they like a particular piece and want to publish it on their VK profile group, they click the browser plugin button. The image is then sent to the application's database along with links to the original work and a description. The application subsequently determines the optimal time to create a post on the social network, without any further intervention from the administrator.

The application is built on a pure DDD architecture and an event-based interaction model between domains is used

### Features
1. Add images along with original work links and descriptions to the application's database.
2. Automatically schedule and post artworks to a VK.ru group.

### Getting Started
These instructions
1. install go v1.17
2. fill out the configuration file config.yaml, registering social network tokens and database connections there.
3. compile and run app
4. the necessary tables in the database will be created automatically

### Usage
```
go run app/cmd/auto_post/main.go
```
For example: 
create task request 
```
curl -X POST --location "http://localhost:<port>/api/v1/init/" \
    -H "Content-Type: application/json" \
    -d "{
          \"url\": \"http://www.file.url/file.jpg\",
          \"auth_url\": \"artstation/shmel\",
          \"service\": \"artstation\"
        }"
```
### Contributing
If you'd like to contribute to this project, please follow these steps:

1. Fork the repository
2. Create a new branch (git checkout -b feature/your-feature)
3. Make your changes and commit them (git commit -m 'Add some feature')
4. Push to the branch (git push origin feature/your-feature)
5. Create a new Pull Request

### Credits
This application was created by Igor Shmel and is maintained by the community.

Special thanks to contributors for their invaluable input.

### Disclaimer
This project is not affiliated with or endorsed by VK.ru or ArtStation. It is an independent project created for educational and personal use.

### License
This project is licensed under the Apache License 2.0