# Suddenly Muscle Trainer

This program suddenly plays Muscle Traiing Video on your browser. The content of the video is ["Muscles for All!"](https://www.youtube.com/watch?v=HF7H6M4nzNY) on NHK WORLD-JAPAN uploaded to Youtube. The video is The video will be played at random timing. Let's do muscle training when the video is played! Eliminate lack of exercise!

## Quick starts
1. Download the zipfile from [releases](https://github.com/acnaman/suddenly-muscle-trainer/releases/)
2. Unzip the zipfile
3. Execute "SuddenlyMuscleTrainer" or "SuddenlyMuscleTrainer.exe"
4. After that, the program will start and youtube will be played at random timing.

## Adjust Setting
Open "SettingTool" or "SettingTool.exe".
You can adjust following items.

| Item | meaning | Default |
----|----|----
| Interval Time(min) | You can get chance that the video will be played every "Interval Time" minutes | 30 |
| Percentage | The probability that the video will be played in one chance | 5 |
| Start Time | The video will not be played before "Start Time" | 09:00 |
| End Time | The video will not be played after "End Time" | 17:30 |

If `Start Time` and `End Time` have the same value, you can get the chance all day.

### Example

If you set the item values below, every 5 minutes there is a 1％ chance that the video will be played between 10:00 and 19:00. 

| Item | Value |
----|----
| Interval Time(min) | 5 |
| Percentage | 1 |
| Start Time | 10:00 |
| End Time | 19:00 |


## Install as a service

You can install this program as a service. If you install the program, the program will start automatically when you reboot the OS.

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

