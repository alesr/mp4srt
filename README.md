# mp4srt
A command-line tool for generating subtitles (SRT files) for MP4 videos using Whisper AI API

With mp4srt, you can easily transcribe the audio from your videos and obtain subtitle files that can be used for various purposes like captioning, translation, or accessibility.

## Installation

### Mac OS X

To use mp4srt, you need to install the following dependencies:

Go programming language: mp4srt is written in Go, so you'll need to have Go installed on your system. If you don't have it installed, follow these steps:

a. Check if Homebrew (brew) is installed on your Mac by opening the Terminal application and running the following command:

```bash
brew --version
```

If brew is not installed, you will see a message indicating that the command was not found.

b. If brew is not installed, you can install it by running the following command in Terminal:

```bash
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

c. Once Homebrew is installed, you can install Go by running the following command in Terminal:

go

```bash
brew install go
```

FFmpeg: mp4srt uses FFmpeg for audio extraction and conversion. To install FFmpeg, run the following command in Terminal:

```bash
brew install ffmpeg
```

## Usage

Before running mp4srt, make sure you have your MP4 videos ready. Follow these steps:

1. Clone or download this repository to your computer.

```bash
git clone https://github.com/alesr/mp4srt.git
```

2. Open the Terminal application and navigate to the directory where you have the MP4SRT source code. You can use the cd command followed by the path to the directory. For example, if you downloaded the source code to your home directory, you can navigate to it using the following command:

```bash
cd ~/mp4srt
```


3. Place the `.mp4` files in the `data` directory of this repository.

4. Run the following command to compile and execute the program:

```bash
make srt
```

The application will automatically transcribe the audio from your MP4 videos and generate corresponding SRT subtitle files. The subtitle files will be saved in the `output` directory located within the mp4srt directory.

After the execution is complete, you can find the generated subtitle files in the "output" directory. These files will have the same name as the corresponding input MP4 files, but with the ".srt" extension.
