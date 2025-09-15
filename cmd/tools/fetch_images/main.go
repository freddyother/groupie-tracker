package main

import (
	"fmt"
	"io"
	"log"
	"mime"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

type Item struct {
	Name string
	URL  string // URL original de upload.wikimedia.org
}

var items = []Item{
	{"The Beatles", "https://upload.wikimedia.org/wikipedia/commons/d/df/The_Fabs.JPG"},
	{"Queen", "https://upload.wikimedia.org/wikipedia/commons/0/0c/QueenPerforming1977.jpg"},
	{"Pink Floyd", "https://upload.wikimedia.org/wikipedia/commons/4/4e/Pink_Floyd_1973.jpg"},
	{"Led Zeppelin", "https://upload.wikimedia.org/wikipedia/commons/7/7e/Led_Zeppelin_1977.jpg"},
	{"The Rolling Stones", "https://upload.wikimedia.org/wikipedia/commons/1/12/Rolling_Stones_1972.jpg"},
	{"Nirvana", "https://upload.wikimedia.org/wikipedia/commons/9/9d/Nirvana_around_1992.jpg"},
	{"Radiohead", "https://upload.wikimedia.org/wikipedia/commons/a/a7/Radiohead_2008.jpg"},
	{"Coldplay", "https://upload.wikimedia.org/wikipedia/commons/2/20/Coldplay_2011.jpg"},
	{"U2", "https://upload.wikimedia.org/wikipedia/commons/0/0f/U2_2009.jpg"},
	{"Metallica", "https://upload.wikimedia.org/wikipedia/commons/2/2c/Metallica_live_2003.jpg"},
	{"AC/DC", "https://upload.wikimedia.org/wikipedia/commons/a/a0/ACDC_In_Tacoma_2009.jpg"},
	{"The Doors", "https://upload.wikimedia.org/wikipedia/commons/8/87/The_Doors_1968.JPG"},
	{"Arctic Monkeys", "https://upload.wikimedia.org/wikipedia/commons/f/f1/Arctic_Monkeys_2012.jpg"},
	{"Foo Fighters", "https://upload.wikimedia.org/wikipedia/commons/e/e5/Foo_Fighters_2018.jpg"},
	{"Guns N' Roses", "https://upload.wikimedia.org/wikipedia/commons/f/f9/Guns_n_Roses_2017.jpg"},
}

const outDir = "web/static/img"
const width = 800 // px

var client = &http.Client{Timeout: 30 * time.Second}

func main() {
	if err := os.MkdirAll(outDir, 0o755); err != nil {
		log.Fatalf("no pude crear %s: %v", outDir, err)
	}

	for _, it := range items {
		dst, err := fetchThumb(it)
		if err != nil {
			log.Printf("❌ %s: %v", it.Name, err)
			continue
		}
		fmt.Printf("✅ %s -> %s\n", it.Name, dst)
		fmt.Printf("   Usa en tu struct: Image: \"/static/img/%s\"\n", filepath.Base(dst))
	}
}

func fetchThumb(it Item) (string, error) {
	orig, err := url.Parse(it.URL)
	if err != nil {
		return "", fmt.Errorf("parse URL: %w", err)
	}
	// Debe ser upload.wikimedia.org/wikipedia/commons/<h1>/<h2>/<filename>
	if !strings.Contains(orig.Host, "upload.wikimedia.org") {
		return "", fmt.Errorf("URL no es de upload.wikimedia.org: %s", it.URL)
	}
	segs := strings.Split(strings.TrimPrefix(orig.Path, "/"), "/")
	// segs = ["wikipedia","commons","h1","h2","FileName.ext"]
	if len(segs) < 5 || segs[0] != "wikipedia" || segs[1] != "commons" {
		return "", fmt.Errorf("ruta no coincide con /wikipedia/commons/h1/h2/filename: %s", orig.Path)
	}
	h1, h2, filename := segs[2], segs[3], segs[len(segs)-1]

	// Construir URL de thumb:
	// /wikipedia/commons/thumb/<h1>/<h2>/<filename>/<width>px-<filename>
	thumbPath := fmt.Sprintf("/wikipedia/commons/thumb/%s/%s/%s/%dpx-%s", h1, h2, filename, width, filename)
	thumbURL := url.URL{Scheme: "https", Host: "upload.wikimedia.org", Path: thumbPath}

	req, _ := http.NewRequest("GET", thumbURL.String(), nil)
	req.Header.Set("User-Agent", "GroupieTracker/1.0 (+https://localhost)")
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("http: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		body := readPreview(resp.Body, 256)
		return "", fmt.Errorf("status %d body %q", resp.StatusCode, body)
	}
	ct := resp.Header.Get("Content-Type")
	if !strings.HasPrefix(ct, "image/") {
		body := readPreview(resp.Body, 256)
		return "", fmt.Errorf("no es imagen (Content-Type=%q) body %q", ct, body)
	}

	ext := pickExtFrom(ct, filename)
	name := slugify(it.Name) + ext

	tmp := filepath.Join(outDir, name+".part")
	out := filepath.Join(outDir, name)

	f, err := os.Create(tmp)
	if err != nil {
		return "", fmt.Errorf("create: %w", err)
	}
	if _, err := io.Copy(f, resp.Body); err != nil {
		f.Close()
		_ = os.Remove(tmp)
		return "", fmt.Errorf("write: %w", err)
	}
	_ = f.Close()
	if err := os.Rename(tmp, out); err != nil {
		return "", fmt.Errorf("rename: %w", err)
	}
	return out, nil
}

func pickExtFrom(ct, fallback string) string {
	if exts, _ := mime.ExtensionsByType(ct); len(exts) > 0 {
		e := strings.ToLower(exts[0])
		if e == ".jpeg" {
			e = ".jpg"
		}
		return e
	}
	e := strings.ToLower(filepath.Ext(fallback))
	if e == ".jpeg" {
		e = ".jpg"
	}
	if e == "" {
		e = ".jpg"
	}
	return e
}

var nonWord = regexp.MustCompile(`[^a-z0-9]+`)

func slugify(s string) string {
	s = strings.ToLower(s)
	s = strings.ReplaceAll(s, "&", "and")
	s = strings.ReplaceAll(s, "/", "_")
	s = strings.ReplaceAll(s, "'", "")
	s = nonWord.ReplaceAllString(s, "_")
	return strings.Trim(s, "_")
}

func readPreview(r io.Reader, n int) string {
	b, _ := io.ReadAll(io.LimitReader(r, int64(n)))
	return string(b)
}
