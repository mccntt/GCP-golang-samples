// Copyright 2019 Google LLC
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     https://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v3poc

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/GoogleCloudPlatform/golang-samples/internal/testutil"
	"github.com/google/uuid"
)

func TestTranslateTextWithGlossary(t *testing.T) {
	tc := testutil.SystemTest(t)

	location := "us-central1"
	sourceLang := "en"
	targetLang := "ja"
	text := "你好天气"
	glossaryID := fmt.Sprintf("translate_text_with_glossary-%v", uuid.New().ID())
	// glossaryInputURI := "gs://cloud-samples-data/translation/glossary_ja.csv"
	glossaryInputURI := "gs://testing-translation-v3-glossary/glossary_zh.csv"

	// Create a glossary.
	var buf bytes.Buffer
	if err := createGlossary(&buf, tc.ProjectID, location, glossaryID, glossaryInputURI); err != nil {
		t.Fatalf("createGlossary: %v", err)
	}
	defer deleteGlossary(&buf, tc.ProjectID, location, glossaryID)

	start := time.Now()
	// Translate text.
	if err := translateTextWithGlossary(&buf, tc.ProjectID, location, sourceLang, targetLang, text, glossaryID); err != nil {
		t.Fatalf("translateTextWithGlossary: %v", err)
	}
	duration := time.Since(start)
	fmt.Printf("buf: %s, duration: %d", buf.String(), duration)

	if got, want1, want2 := buf.String(), "アカウント", "yaoxibaga"; !strings.Contains(got, want1) && !strings.Contains(got, want2) {
		t.Errorf("translateTextWithGlossary got:\n----\n%s----\nWant to contain:\n----\n%s\n----\nOR\n----\n%s\n----", got, want1, want2)
	}
}

func TestTranslateTextWithGlossaryEU(t *testing.T) {
	tc := testutil.SystemTest(t)

	location := "europe-west1"
	sourceLang := "zh"
	targetLang := "ja"
	text := "你好天气"
	glossaryID := fmt.Sprintf("translate_text_with_glossary-%v", uuid.New().ID())
	//glossaryInputURI := "gs://cloud-samples-data/translation/glossary_ja.csv"
	glossaryInputURI := "gs://testing-translation-v3-glossary/glossary_zh.csv"

	// Create a glossary.
	var buf bytes.Buffer
	if err := createGlossaryEU(&buf, tc.ProjectID, location, glossaryID, glossaryInputURI); err != nil {
		t.Fatalf("createGlossary: %v", err)
	}
	defer deleteGlossary(&buf, tc.ProjectID, location, glossaryID)

	start := time.Now()
	// Translate text.
	if err := translateTextWithGlossaryEU(&buf, tc.ProjectID, "europe-west1", sourceLang, targetLang, text, glossaryID); err != nil {
		t.Fatalf("translateTextWithGlossary: %v", err)
	}
	duration := time.Since(start)
	fmt.Print("duration:")
	fmt.Println(duration)

	if got, want1, want2 := buf.String(), "アカウント", "口座"; !strings.Contains(got, want1) && !strings.Contains(got, want2) {
		t.Errorf("translateTextWithGlossary got:\n----\n%s----\nWant to contain:\n----\n%s\n----\nOR\n----\n%s\n----", got, want1, want2)
	}
}
