package main

import (
	"context"
	"fmt"
	"log"

	speech "cloud.google.com/go/speech/apiv1p1beta1"
	speechpb "cloud.google.com/go/speech/apiv1p1beta1/speechpb"
	// speech "cloud.google.com/go/speech/apiv1"
	// speechpb "cloud.google.com/go/speech/apiv1/speechpb"
)

func main() {
	ctx := context.Background()

	// Instantiates a client
	client, err := speech.NewClient(ctx)
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}
	defer client.Close()

	// The path to the remote audio file to transcribe
	audio := speechpb.RecognitionAudio{
		AudioSource: &speechpb.RecognitionAudio_Uri{Uri: "gs://mclab2023-speech-to-text/audio-files/tencent_HOK_test_20230209_163850.mp3"},
	}

	config := speechpb.RecognitionConfig{
		Encoding:          speechpb.RecognitionConfig_MP3,
		SampleRateHertz:   48000,
		AudioChannelCount: 1,
		LanguageCode:      "zh",
		Model:             "default",
		// EnableWordTimeOffsets: true,
		EnableAutomaticPunctuation: true,
		Adaptation: &speechpb.SpeechAdaptation{
			PhraseSetReferences: []string{"projects/480714232090/locations/global/phraseSets/test-phrase-set-2"},
		},
	}

	// Detects speech in the audio file
	op, err := client.LongRunningRecognize(ctx, &speechpb.LongRunningRecognizeRequest{
		Config: &config,
		Audio:  &audio,
	})
	if err != nil {
		log.Fatalf("failed to recognize: %v", err)
	}
	resp, err := op.Wait(ctx)
	if err != nil {
		log.Fatalf("failed to wait for long-running operation: %v", err)
	}
	// Prints the results
	for _, result := range resp.Results {
		for _, alt := range result.Alternatives {
			fmt.Printf("%v\n", alt.Transcript)
		}
	}
}
