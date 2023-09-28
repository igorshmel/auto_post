### Auto_post
This application enables automatic scheduling and random posting to a VK.ru group. The posts consist of artistic works sourced from platforms like ArtStation. The process is as follows: an administrator browses new artworks on ArtStation. If they like a particular piece and want to publish it on their VK profile group, they click the browser plugin button. The image is then sent to the application's database along with links to the original work and a description. The application subsequently determines the optimal time to create a post on the social network, without any further intervention from the administrator.

### Features
Automatically schedule and post artworks to a VK.ru group.
Add images along with original work links and descriptions to the application's database.
Automatically determine the best time for posting on social media.
Support for YouTube videos as post content.

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
1. Access the application in your preferred web browser at http://localhost:5000.
2. Log in with your VK.ru account credentials.
3. Start adding artworks or YouTube videos to the database using the browser plugin.
4. Configure the posting schedule and preferences in the application interface.
5. Sit back and let the application handle the rest!

### Contributing
If you'd like to contribute to this project, please follow these steps:

1. Fork the repository
2. Create a new branch (git checkout -b feature/your-feature)
3. Make your changes and commit them (git commit -m 'Add some feature')
4. Push to the branch (git push origin feature/your-feature)
5. Create a new Pull Request

License
This project is licensed under the MIT License.