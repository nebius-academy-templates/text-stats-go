//go:build experimental_render

package main

import toolkit "example.com/course/text-document-toolkit"

// render is optional, so the ordinary test suite never imports it.
func (doc document) render() string {
	return toolkit.RenderDocument(doc.text)
}
