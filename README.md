# drive-share
A lightweight Go client for generateing Google Drive share links quickly. This is still very much in pre-alpha version and only created for personal use at the moment. That said, if you wish to use this tool, you will find instructions for installing and running it below.

## Using drive-share
Running `drive-share -name [name-of-file]` where the `name-of-file` is the name of some file in your Google Drive that has been set to shared mode will automatically copy a link to your clipboard that, when sent to anybody, will allow them to view or download that file.

At the moment, this does not take into files with the same names, but in different folders. It will print all the files with the matching name and their web URLS and copy the last item in the list to your clipboard. 

## Setting up drive-share

#### Getting a Google Drive API
1. Go to [Google's developer console](https://console.developers.google.com/) and create a new project. 
2. Open the project and search for the Drive API in the library. 
3. Enable the API for your project. 
4. Generate your OAUTH2 credentials for it. 
5. Download the `json` file generated.

#### Making an installation folder
1. Create a folder somewhere on your computer.
2. Move over the `json` file containing your credentials to this folder. 
3. Rename it `client_id.json`. 
4. Download the `drive-share` executable from the `build/OS` directory of this repo corresponding to your OS. 

#### Running the executable for the first time
1. Open up a terminal prompt in the folder.
2. Run `drive-share` in this prompt. It will generate a link to an authorization page.
3. Follow this link. Approve the permissions and copy the ID given to you. 
4. Paste the ID back into the terminal prompt.

Future runs of the executable will not require authorization. A token will automatically be placed in your folder that caches the auuthorization header.


## To-Do's
- Implement a flag to take into account folder locations
- Implement functionality to give shortcuts to file names that you often share
- Implement more searching flags (instead of exact file name match, allow contains or time range matching)
- Sort out client secret so that every user doesn't have to start up their own Drive API instance. 
