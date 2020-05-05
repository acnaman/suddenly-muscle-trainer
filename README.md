# Suddenly Muscle Trainer

This program suddenly plays Muscle Traiing Video on your browser.
The Video 

## Quick starts
1. Download the zipfile from [releases](https://github.com/acnaman/suddenly-muscle-trainer/releases/)
2. Unzip the zipfile
3. Execute "SuddenlyMuscleTrainer" or "SuddenlyMuscleTrainer.exe"
4. After that, the program will start and youtube will play at random timing.

## Adjust Setting
Open "SettingTool" or "SettingTool.exe".
You can adjust following items

- Interval Time(min)
- Percentage
- Start Time
- End Time

## Install as a service

You can install this program as a service. (If you install the program, the program will start automatically when you reboot the OS.)

### How to Install and start service

1. Download the zipfile from [releases](https://github.com/acnaman/suddenly-muscle-trainer/releases/)
2. Move the zipfile to installation folder
   (You can choose any folder to install)
3. Unzip the zipfile
4. Run the following command to install service
   (Administrator authority required)
```
# suddenlyMuscleTraining --service install
```
5. To start service, Run the following command (or Reboot the OS)
```
# suddenlyMuscleTraining --service start
```

## for Development

### How to build
In order to build, you need need the following installation.

- Go
- Make

#### Mac

Execute the following command in the "suddenly-muscle-trainer" directory.

```
$ make pre
$ make
```

