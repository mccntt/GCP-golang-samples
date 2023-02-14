// ╰─❯ GOOGLE_APPLICATION_CREDENTIALS=/Users/fangchih/Downloads/hok-my-project-47979-gcloud01-7aa03be33cf2.json go run stt/normal/model-adaptation-sync.go

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
		// AudioSource: &speechpb.RecognitionAudio_Uri{Uri: "gs://gvoice_test/audio-files/hok.pcm"},
		AudioSource: &speechpb.RecognitionAudio_Uri{Uri: "gs://stt-model-adaptation-2023/hok.pcm"},
	}

	config := speechpb.RecognitionConfig{
		Encoding:        speechpb.RecognitionConfig_LINEAR16,
		SampleRateHertz: 44100,
		LanguageCode:    "en-US",
		Adaptation: &speechpb.SpeechAdaptation{
			PhraseSetReferences: []string{"projects/481263715628/locations/global/phraseSets/test-phrase-set-1"},
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
	fmt.Println("hok resp:", resp)
	for _, result := range resp.Results {
		for _, alt := range result.Alternatives {
			fmt.Printf("%v\n", alt.Transcript)
		}
	}
}
