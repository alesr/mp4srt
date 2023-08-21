# mp4srt

A command-line tool to extract audio from MP4 videos, transcribe the audio, and generate subtitles (SRT files) using the Whisper AI API.

## Overview

The `mp4srt` tool transcribes audio from MP4 videos and generates corresponding SRT subtitle files. Once the transcription is complete, the subtitle files are saved in the `output` directory, located within the `mp4srt` directory. The generated SRT files will have the same name as the input MP4 files, but with the `.srt` extension.

## Prerequisites

- Set up the OpenAI API key as an environment variable named `OPENAI_API_KEY`.

## Usage

1. **Prepare the Videos**: Place the MP4 videos you wish to transcribe in the `data` directory.
2. **Run the Tool**: Use the command below to transcribe the videos:

```bash
go run main.go
```

This will process MP4 files in the data directory, extract their audio using ffmpeg, and transcribe them. The default sample rate is set to 3000, which, although low, is optimized for generating WAV files within the 25MB limit of the Whisper API. This sample rate is typically sufficient for accurate results.

- **Customizing Sample Rate**: If you wish to use a different sample rate, use the `--sample_rate` flag followed by the desired value. For example:

```bash
go run main.go --sample_rate=8000
```

- **Generating Textual Transcripts**: In addition to SRT files, if you want a textual transcription of the audio, use the `--text_transcript=true` flag:

```bash
go run main.go --text_transcript=true
```

## Notes

Due to the Whisper API's 25MB file size limit, a low sample rate is used by default to ensure the generated WAV file stays within this limit. Adjusting the sample rate can affect the transcription quality, so choose a value that balances file size and transcription accuracy.
