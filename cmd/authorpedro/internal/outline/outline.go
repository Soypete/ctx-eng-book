package outline

import (
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

type Book struct {
	Title    string    `yaml:"title"`
	Slug     string    `yaml:"slug"`
	Chapters []Chapter `yaml:"chapters"`
}

type Chapter struct {
	Number  int      `yaml:"number"`
	Title   string   `yaml:"title"`
	Slug    string   `yaml:"slug"`
	Modules []Module `yaml:"modules"`
}

type Module struct {
	Number int    `yaml:"number"`
	Title  string `yaml:"title"`
	Slug   string `yaml:"slug"`
}

func ParseChaptersMarkdown(path string) (Book, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return Book{}, err
	}

	book := Book{
		Title:    "Context Engineering: Building Reliable AI Systems",
		Slug:     "ctx-eng-book",
		Chapters: []Chapter{},
	}

	lines := strings.Split(string(data), "\n")
	var currentChapter *Chapter

	chapterRE := regexp.MustCompile(`^##\s+Chapter\s+(\d+):\s+(.+)$`)
	moduleRE := regexp.MustCompile(`^###\s+(.+)$`)
	partRE := regexp.MustCompile(`^#\s+Part\s+([IVX]+):\s+(.+)$`)

	for _, line := range lines {
		line = strings.TrimSpace(line)

		if partRE.MatchString(line) {
			continue
		}

		if chapterRE.MatchString(line) {
			matches := chapterRE.FindStringSubmatch(line)
			chapterNum := parseChapterNumber(matches[1])
			chapterTitle := strings.TrimSpace(matches[2])
			chapterSlug := toSlug(chapterTitle)

			chapter := Chapter{
				Number:  chapterNum,
				Title:   chapterTitle,
				Slug:    chapterSlug,
				Modules: []Module{},
			}
			book.Chapters = append(book.Chapters, chapter)
			currentChapter = &book.Chapters[len(book.Chapters)-1]
			continue
		}

		if moduleRE.MatchString(line) && currentChapter != nil {
			matches := moduleRE.FindStringSubmatch(line)
			moduleTitle := strings.TrimSpace(matches[1])

			if !strings.HasPrefix(moduleTitle, "Why") &&
				!strings.HasPrefix(moduleTitle, "Beyond") &&
				!strings.HasPrefix(moduleTitle, "The") &&
				!strings.HasPrefix(moduleTitle, "Retrieval") &&
				!strings.HasPrefix(moduleTitle, "Vector") &&
				!strings.HasPrefix(moduleTitle, "Knowledge") &&
				!strings.HasPrefix(moduleTitle, "Agents") &&
				!strings.HasPrefix(moduleTitle, "Memory") &&
				!strings.HasPrefix(moduleTitle, "Context") &&
				!strings.HasPrefix(moduleTitle, "Authorization") &&
				!strings.HasPrefix(moduleTitle, "Semantic") &&
				!strings.HasPrefix(moduleTitle, "Tool") &&
				!strings.HasPrefix(moduleTitle, "Observability") &&
				!strings.HasPrefix(moduleTitle, "Evaluation") &&
				!strings.HasPrefix(moduleTitle, "Building") {
				continue
			}

			moduleSlug := toSlug(moduleTitle)
			module := Module{
				Number: len(currentChapter.Modules) + 1,
				Title:  moduleTitle,
				Slug:   moduleSlug,
			}
			currentChapter.Modules = append(currentChapter.Modules, module)
		}
	}

	return book, nil
}

func parseChapterNumber(s string) int {
	num := 0
	for _, c := range s {
		if c >= '0' && c <= '9' {
			num = num*10 + int(c-'0')
		}
	}
	if num == 0 {
		for _, c := range s {
			if c >= 'A' && c <= 'Z' {
				num = num*26 + int(c-'A') + 1
			}
		}
	}
	return num
}

func toSlug(s string) string {
	s = strings.ToLower(s)
	s = regexp.MustCompile(`[^a-z0-9]+`).ReplaceAllString(s, "-")
	s = strings.Trim(s, "-")
	return s
}

func (b Book) ResolvePath(chapterSlug, moduleSlug string) (string, error) {
	for _, ch := range b.Chapters {
		if ch.Slug != chapterSlug {
			continue
		}
		for _, mod := range ch.Modules {
			if mod.Slug == moduleSlug {
				return filepath.Join(
					"chapters",
					formatSlug(ch.Number, ch.Slug),
					"modules",
					formatSlug(mod.Number, mod.Slug)+".md",
				), nil
			}
		}
	}
	return "", nil
}

func formatSlug(num int, slug string) string {
	return slug
}

func (b Book) AllPaths() []struct {
	Chapter  string
	Module   string
	FullPath string
} {
	var paths []struct {
		Chapter  string
		Module   string
		FullPath string
	}

	for _, ch := range b.Chapters {
		for _, mod := range ch.Modules {
			path, _ := b.ResolvePath(ch.Slug, mod.Slug)
			paths = append(paths, struct {
				Chapter  string
				Module   string
				FullPath string
			}{
				Chapter:  ch.Slug,
				Module:   mod.Slug,
				FullPath: path,
			})
		}
	}

	return paths
}
