package main

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	"gitlab.com/zerok/zerokspot.com/pkg/searchdoc"

	"github.com/gomarkdown/markdown"
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
	"github.com/spf13/cobra"
	"gitlab.com/zerok/zerokspot.com/pkg/bloggraph"
)

type doc struct {
	ObjectID string `json:"objectID"`
	Content  string `json:"content"`
	File     string `json:"file"`
}

type PostPathElement struct {
	ContentID string `json:"objectID"`
	Degree    int    `json:"degree"`
	Title     string `json:"title"`
}

type PostPaths struct {
	Up   []PostPathElement `json:"up"`
	Down []PostPathElement `json:"down"`
}

func decodeMetadata(ctx context.Context, path string) (*searchdoc.SearchDoc, error) {
	fp, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer fp.Close()
	var d searchdoc.SearchDoc
	if err := json.NewDecoder(fp).Decode(&d); err != nil {
		return nil, err
	}
	return &d, nil
}

var buildGraphCmd = &cobra.Command{
	Use: "build-graph",
	RunE: func(command *cobra.Command, args []string) error {
		if err := os.MkdirAll("data", 0700); err != nil {
			return err
		}
		ctx := logger.WithContext(context.Background())
		bg := bloggraph.NewGraph()

		allContentIDs := make([]string, 0, 100)
		mapping := make(map[string]PostPaths)

		// Pass 1
		logger.Info().Msg("Pass 1")
		if err := filepath.Walk("public", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() || info.Name() != "index.json" {
				return nil
			}
			if !(strings.HasPrefix(path, "public/weblog") || strings.HasPrefix(path, "public/notes")) {
				return nil
			}
			doc, err := decodeMetadata(ctx, path)
			if err != nil {
				return err
			}
			allContentIDs = append(allContentIDs, doc.ObjectID)
			bg.GetOrCreateNode(bloggraph.Node{
				ContentID: doc.ObjectID,
				Title:     doc.Title,
			})
			return nil
		}); err != nil {
			return err
		}

		// Pass 2
		logger.Info().Msg("Pass 2")
		if err := filepath.Walk("public", func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if info.IsDir() || info.Name() != "index.json" {
				return nil
			}
			if !(strings.HasPrefix(path, "public/weblog") || strings.HasPrefix(path, "public/notes")) {
				return nil
			}
			doc, err := decodeMetadata(ctx, path)
			if err != nil {
				return err
			}
			parents, err := findParents(ctx, doc.File)
			if err != nil {
				return err
			}
			targetNode := bg.GetOrCreateNode(bloggraph.Node{
				ContentID: doc.ObjectID,
			})
			for _, parent := range parents {
				if strings.HasPrefix(parent, "/") {
					sourceNode := bg.GetOrCreateNode(bloggraph.Node{
						ContentID: parent,
					})
					bg.CreateEdge(sourceNode, targetNode, "inspired")
				}
			}
			return nil
		}); err != nil {
			return err
		}

		fmt.Printf("%d edges found.\n", bg.NumEdges())
		for _, id := range allContentIDs {
			mapping[id] = PostPaths{
				Up:   make([]PostPathElement, 0, 5),
				Down: make([]PostPathElement, 0, 5),
			}
		}
		for _, id := range allContentIDs {
			bg.WalkDown(id, "inspired", 0, func(g *bloggraph.Graph, node *bloggraph.Node, degree int) {
				m := mapping[id]
				if !containsContentID(m.Down, node.ContentID) {
					m.Down = append(m.Down, PostPathElement{Title: node.Title, ContentID: node.ContentID, Degree: degree})
				}
				mapping[id] = m
			})
			bg.WalkUp(id, "inspired", 0, func(g *bloggraph.Graph, node *bloggraph.Node, degree int) {
				m := mapping[id]
				if !containsContentID(m.Up, node.ContentID) {
					m.Up = append(m.Up, PostPathElement{Title: node.Title, ContentID: node.ContentID, Degree: degree})
				}
				mapping[id] = m
			})
		}
		data, err := json.Marshal(mapping)
		if err != nil {
			return err
		}
		return ioutil.WriteFile(filepath.Join("data", "postpaths.json"), data, 0600)
	},
}

func containsContentID(haystack []PostPathElement, needle string) bool {
	for _, u := range haystack {
		if u.ContentID == needle {
			return true
		}
	}
	return false
}

func findParents(ctx context.Context, path string) ([]string, error) {
	content, err := ioutil.ReadFile(filepath.Join("content", path))
	if err != nil {
		return nil, err
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

func init() {
	rootCmd.AddCommand(buildGraphCmd)
}
