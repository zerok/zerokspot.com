package bloggraph

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log/slog"
	"os"
	"path/filepath"
	"strings"

	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/renderer"
	"gitlab.com/zerok/zerokspot.com/pkg/searchdoc"
)

type PostPathElement struct {
	ContentID string `json:"objectID"`
	Degree    int    `json:"degree"`
	Title     string `json:"title"`
}

type PostPaths struct {
	Up   []PostPathElement `json:"up"`
	Down []PostPathElement `json:"down"`
}

func BuildMapping(ctx context.Context, rootPath string) (map[string]PostPaths, error) {
	allContentIDs := make([]string, 0, 100)
	mapping := make(map[string]PostPaths)
	bg := NewGraph()

	// Pass 1
	slog.InfoContext(ctx, "Pass 1")
	if err := filepath.Walk(filepath.Join(rootPath, "public"), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || info.Name() != "index.json" {
			return nil
		}
		if !(strings.HasPrefix(path, filepath.Join(rootPath, "public/weblog")) || strings.HasPrefix(path, filepath.Join(rootPath, "public/notes"))) {
			return nil
		}
		doc, err := decodeMetadata(ctx, path)
		if err != nil {
			return err
		}
		allContentIDs = append(allContentIDs, doc.ObjectID)
		bg.GetOrCreateNode(Node{
			ContentID: doc.ObjectID,
			Title:     doc.Title,
			Date:      doc.Date,
		})
		return nil
	}); err != nil {
		return nil, err
	}

	// Pass 2
	slog.InfoContext(ctx, "Pass 2")
	if err := filepath.Walk(filepath.Join(rootPath, "public"), func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if info.IsDir() || info.Name() != "index.json" {
			return nil
		}
		if !(strings.HasPrefix(path, filepath.Join(rootPath, "public/weblog")) || strings.HasPrefix(path, filepath.Join(rootPath, "public/notes"))) {
			return nil
		}
		doc, err := decodeMetadata(ctx, path)
		if err != nil {
			return err
		}
		parents, err := findParents(ctx, rootPath, doc.File)
		if err != nil {
			return err
		}
		targetNode := bg.GetOrCreateNode(Node{
			ContentID: doc.ObjectID,
		})
		for _, parent := range parents {
			if strings.HasPrefix(parent, "/") {
				sourceNode := bg.GetOrCreateNode(Node{
					ContentID: parent,
				})
				if sourceNode.ContentID == targetNode.ContentID {
					continue
				}
				bg.CreateEdge(sourceNode, targetNode, "inspired")
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}

	slog.InfoContext(ctx, fmt.Sprintf("%d edges found", bg.NumEdges()))
	for _, id := range allContentIDs {
		mapping[id] = PostPaths{
			Up:   make([]PostPathElement, 0, 5),
			Down: make([]PostPathElement, 0, 5),
		}
	}
	for _, id := range allContentIDs {
		visited := make(map[string]struct{})
		thisNode := bg.GetOrCreateNode(Node{
			ContentID: id,
		})
		bg.WalkDown(visited, id, "inspired", 0, func(g *Graph, node *Node, degree int) {
			m := mapping[id]
			// Prevent circles here that might arise due to higher-degree connections:
			if node.ContentID == id {
				return
			}
			if !containsContentID(m.Down, node.ContentID) {
				// Never add anything to "Down" that is older than thisNode
				if node.Date < thisNode.Date {
					return
				}
				m.Down = append(m.Down, PostPathElement{Title: node.Title, ContentID: node.ContentID, Degree: degree})
			}
			mapping[id] = m
		})
		visited = make(map[string]struct{})
		bg.WalkUp(visited, id, "inspired", 0, func(g *Graph, node *Node, degree int) {
			m := mapping[id]
			// Prevent circles here that might arise due to higher-degree connections:
			if node.ContentID == id {
				return
			}
			if !containsContentID(m.Up, node.ContentID) {
				// Never add anything to Up that is newer
				if node.Date > thisNode.Date {
					return
				}
				m.Up = append(m.Up, PostPathElement{Title: node.Title, ContentID: node.ContentID, Degree: degree})
			}
			mapping[id] = m
		})
	}

	return mapping, nil
}

func decodeMetadata(ctx context.Context, path string) (*searchdoc.SearchDoc, error) {
	fp, err := os.Open(path)
	if err != nil {
		return nil, fmt.Errorf("failed to open `%s`: %w", path, err)
	}
	defer fp.Close()
	var d searchdoc.SearchDoc
	if err := json.NewDecoder(fp).Decode(&d); err != nil {
		return nil, err
	}
	return &d, nil
}

func containsContentID(haystack []PostPathElement, needle string) bool {
	for _, u := range haystack {
		if u.ContentID == needle {
			return true
		}
	}
	return false
}

type markdownRenderer struct {
	parents []string
}

func newMarkdownRenderer() *markdownRenderer {
	return &markdownRenderer{
		parents: make([]string, 0, 5),
	}
}

func (r *markdownRenderer) Render(w io.Writer, source []byte, n ast.Node) error {
	foundParents := make(map[string]struct{})
	ast.Walk(n, func(n ast.Node, entering bool) (ast.WalkStatus, error) {
		if !entering && n.Kind() == ast.KindLink {
			link := n.(*ast.Link)
			dest := string(link.Destination)
			dest = strings.TrimPrefix(dest, "https://zerokspot.com")
			if !(strings.HasPrefix(dest, "/weblog/") || strings.HasPrefix(dest, "/notes/")) {
				return ast.WalkContinue, nil
			}
			foundParents[dest] = struct{}{}
		}
		return ast.WalkContinue, nil
	})

	result := make([]string, 0, len(foundParents))
	for p := range foundParents {
		result = append(result, p)
	}
	r.parents = result
	return nil
}

func (r *markdownRenderer) AddOptions(...renderer.Option) {}

func findParents(ctx context.Context, rootPath, path string) ([]string, error) {
	contentPath := filepath.Join(rootPath, "content", path)
	content, err := os.ReadFile(contentPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open `%s`: %w", contentPath, err)
	}

	mdRenderer := newMarkdownRenderer()
	if err := goldmark.New(goldmark.WithRenderer(mdRenderer)).Convert(content, io.Discard); err != nil {
		return nil, fmt.Errorf("failed to parse content: %w", err)
	}
	return mdRenderer.parents, nil
}
