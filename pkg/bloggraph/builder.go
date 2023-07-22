package bloggraph

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
	"github.com/rs/zerolog"
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
	logger := zerolog.Ctx(ctx)
	allContentIDs := make([]string, 0, 100)
	mapping := make(map[string]PostPaths)
	bg := NewGraph()

	// Pass 1
	logger.Info().Msg("Pass 1")
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
		})
		return nil
	}); err != nil {
		return nil, err
	}

	// Pass 2
	logger.Info().Msg("Pass 2")
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
				bg.CreateEdge(sourceNode, targetNode, "inspired")
			}
		}
		return nil
	}); err != nil {
		return nil, err
	}

	logger.Info().Msgf("%d edges found.\n", bg.NumEdges())
	for _, id := range allContentIDs {
		mapping[id] = PostPaths{
			Up:   make([]PostPathElement, 0, 5),
			Down: make([]PostPathElement, 0, 5),
		}
	}
	for _, id := range allContentIDs {
		visited := make(map[string]struct{})
		bg.WalkDown(visited, id, "inspired", 0, func(g *Graph, node *Node, degree int) {
			m := mapping[id]
			if !containsContentID(m.Down, node.ContentID) {
				m.Down = append(m.Down, PostPathElement{Title: node.Title, ContentID: node.ContentID, Degree: degree})
			}
			mapping[id] = m
		})
		visited = make(map[string]struct{})
		bg.WalkUp(visited, id, "inspired", 0, func(g *Graph, node *Node, degree int) {
			m := mapping[id]
			if !containsContentID(m.Up, node.ContentID) {
				m.Up = append(m.Up, PostPathElement{Title: node.Title, ContentID: node.ContentID, Degree: degree})
			}
			mapping[id] = m
		})
	}

	return mapping, nil
}

func decodeMetadata(ctx context.Context, path string) (*searchdoc.SearchDoc, error) {
	logger := zerolog.Ctx(ctx)
	logger.Debug().Msgf("Loading data from %s", path)
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

func findParents(ctx context.Context, rootPath, path string) ([]string, error) {
	logger := zerolog.Ctx(ctx)
	logger.Debug().Msgf("Finding parents of %s", path)
	contentPath := filepath.Join(rootPath, "content", path)
	content, err := ioutil.ReadFile(contentPath)
	if err != nil {
		return nil, fmt.Errorf("failed to open `%s`: %w", contentPath, err)
	}
	extensions := parser.CommonExtensions | parser.AutoHeadingIDs
	p := parser.NewWithExtensions(extensions)
	tree := markdown.Parse(content, p)
	var nv NodeVisitor
	ast.Walk(tree, &nv)
	return nv.Parents, nil
}

type NodeVisitor struct {
	Parents []string
}

func NewNodeVisitor() *NodeVisitor {
	return &NodeVisitor{
		Parents: make([]string, 0, 10),
	}
}

func (v *NodeVisitor) Visit(node ast.Node, entering bool) ast.WalkStatus {
	if link, ok := node.(*ast.Link); ok {
		dest := string(link.Destination)
		if strings.HasPrefix(dest, "https://zerokspot.com") {
			dest = strings.TrimPrefix(dest, "https://zerokspot.com")
		}
		if !(strings.HasPrefix(dest, "/weblog/") || strings.HasPrefix(dest, "/notes/")) {
			return ast.GoToNext
		}
		for _, p := range v.Parents {
			if p == dest {
				return ast.GoToNext
			}
		}
		v.Parents = append(v.Parents, dest)
	}
	return ast.GoToNext
}
