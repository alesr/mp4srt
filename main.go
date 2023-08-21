package main

import (
	"flag"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/alesr/httpclient"
	"github.com/alesr/mp4srt/whisperclient"
)

const (
	defaultModel      string = "whisper-1"
	defaultSampleRate string = "3000"
	inputDir          string = "data" + string(os.PathSeparator)
	inputExt          string = ".mp4"
	outputDir         string = "output" + string(os.PathSeparator)
	outputExt         string = ".srt"
	tmpDir            string = "tmp" + string(os.PathSeparator)
	wavExt            string = ".wav"
)

func main() {
	defer clearTmpDir()

	sampleRate := flag.String("sample_rate", defaultSampleRate, "sample rate")
	textTranscript := flag.Bool("text_transcription", false, "text transcription")

	flag.Parse()

	_ = textTranscript

	apiKey := os.Getenv("OPENAI_API_KEY")
	if apiKey == "" {
		log.Fatalln("missing OPENAI_API_KEY")
	}

	client := whisperclient.New(httpclient.New(), apiKey, defaultModel)

	filepath.Walk(inputDir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			log.Fatalf("could not to get stats of '%s' file: %s\n", path, err)
		}

		inputFileName := info.Name()
		wavFileName := strings.ReplaceAll(inputFileName, inputExt, wavExt)
		outputFileName := strings.ReplaceAll(inputFileName, inputExt, outputExt)

		if !strings.HasSuffix(inputFileName, inputExt) {
			return nil
		}

		if info.IsDir() {
			return nil
		}

		cmd := exec.Command(
			"ffmpeg", "-y", "-i", path, "-vn", "-acodec", "pcm_s16le", "-ar", *sampleRate,
			"-ac", "1", "-b:a", "32k", tmpDir+wavFileName)

		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if err := cmd.Run(); err != nil {
			log.Fatalf("could not to convert '%s' file: %s\n", path, err)
		}

		wavFile, err := os.Open(tmpDir + wavFileName)
		if err != nil {
			log.Fatalf("could not to open '%s' file: %s\n", wavFileName, err)
		}
		defer wavFile.Close()

		str, err := client.TranscribeAudio(whisperclient.TranscribeAudioInput{
			Name: wavFileName,
			Data: wavFile,
		})
		if err != nil {
			log.Fatalf("could not to transcribe audio: %s\n", err)
		}

		outputFile := outputDir + outputFileName

		outFile, err := os.Create(outputFile)
		if err != nil {
			log.Fatalf("could not to create '%s' output file: %s\n", outputFile, err)
		}
		defer outFile.Close()

		if _, err := outFile.Write(str); err != nil {
			log.Fatalf("could not to write to '%s' output file: %s\n", outputFile, err)
		}
		return nil
	})
}

func clearTmpDir() {
	filepath.Walk(tmpDir, func(path string, info os.FileInfo, err error) error {
		if strings.HasSuffix(path, wavExt) {
			if err := os.Remove(path); err != nil {
				log.Fatalf("could not to remove '%s' file: %s\n", path, err)
			}
		}
		return nil
	})
}
