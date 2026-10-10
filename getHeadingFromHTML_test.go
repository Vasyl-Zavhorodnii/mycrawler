package main

import (
	"testing"
)

func TestGetHeadingFromHTMLBasic(t *testing.T) {
	inputBody := "<html><body><h1>Test Title</h1></body></html>"
	inputBody1 := "<html><body><h2>Test Title</h2></body></html>"
	actual := getHeadingFromHTML(inputBody)
	actual1 := getHeadingFromHTML(inputBody1)
	expected := "Test Title"

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
	if actual1 != expected {
		t.Errorf("expected %q, got %q", expected, actual1)
	}
}

func TestGetFirstParagraphFromHTMLBasic(t *testing.T) {
	inputBody := `<html><body>
		<p>Outside paragraph.</p>
		<main>
			<p>Main paragraph.</p>
		</main>
	</body></html>`
	actual := getFirstParagraphFromHTML(inputBody)
	expected := "Main paragraph."

	if actual != expected {
		t.Errorf("expected %q, got %q", expected, actual)
	}
}
