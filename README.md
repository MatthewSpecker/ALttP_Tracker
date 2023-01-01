# ALttP_Tracker

Purpose:
The purpose of this application is to run as a tracker while you play "A Link to the Past Randomizer". You can learn more about ALttPR and play it at: https://alttpr.com/en. As a player collects randomly located items in the game, and discovers where locations are randomly located, they can mark the information on this tracker to help them keep track.

There are other trackers for ALttPR that are quite good. However, all trackers always lack some features or can be improved. I wanted to make a tracker for fun and because I want a tracker that works exactly the way I want. The eventual goal of my tracker is to accomplish 2 goals that I believe most other trackers don't:
1. Hotkey/keyboard input tracking. The popular trackers rely on mouse input by the user to know what info the user wants to update on the tracker. Since the user has limited time to take a hand off their game controller and look away from the video game screen, mouse input is not ideal. Especially in modes where locations are randomized in addition to items. Well designed hotkey/keyboard input tracking combined with mouse input should allow users to focus more on the game and less on their tracking.
2. Provide a clean looking UI that can be easily adjusted on the fly to user preference. Most trackers don't allow for small personal preferences that allow for a better experience. And some of the trackers with the best tracking tools are just not pleasant to look at (IMO).

How to Install:
Eventually, a downloadable installer will be provided. Windows, Mac, and Linux will all be supported. For now, one must download and install Golang, this repository, and Fyne: https://fyne.io/, in a single directory. If one does this, they can run this program by navigating to the main directory of the repository and running the "go run ." command.

Technical Info:
This application is written entirely in Golang using the Fyne GUI. I chose Golang because I wanted to learn it and it's just a pleasant language to use with lots of tools and packages made by others online. I chose Fyne for the GUI for a couple reasons. One, it's one of the most popular GUI options for Golang and is open source and free to use. Two, Fyne does most of the work for your program to work on Windows, Mac, Linux, and mobile. You can simply make installers for each option and Fyne will do all the work to make sure your program will run as desired. Viper is used for managing the files I save user states and preferences in. Viper makes adding defaults and reading/writing files easy. The golang-design/hotkey package is used to enable hotkey functionality on Windows, Mac, and Linux.
-Fyne: https://developer.fyne.io/
-Viper: https://github.com/spf13/viper
-Hotkey: https://github.com/golang-design/hotkey

Next Steps:
-Currently refactoring the code to ease the burden of future development
-Adding unit testing to ensure updates won't affect user experience in the future
-Cleaning up UI issues/making small UI improvements
-Fully flesh out the hotkey tracking to work for item tracking. This has been tested, but is disabled right now.
-Add map tracking. This will allow users to track what locations they have been to in the game and this is when my application will start to have use to some in the ALttP Rando community.
-When the above are done, I would like to add entrance tracking and auto tracking.

Other:
I stream ALttPR ladder races at: https://www.twitch.tv/speckyspecks. I use my tracker on my streams if you wish to see it in action.


