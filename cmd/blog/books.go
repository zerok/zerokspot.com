package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"time"

	"log/slog"

	"github.com/spf13/cobra"
	"gitlab.com/zerok/zerokspot.com/pkg/bookscollection"
)

var booksCmd = &cobra.Command{
	Use: "books",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

func generateGenOPMLCommand() *cobra.Command {
	var cmd = &cobra.Command{
		Use: "gen-opml",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			ctx = findParentTrace(ctx)
			ctx, span := tracer.Start(ctx, "cmd:books/gen-opml")
			defer span.End()
			collection, err := bookscollection.LoadBooks(ctx)
			if err != nil {
				return err
			}

			books := make([]*bookscollection.Book, 0, 10)

			for _, book := range collection {
				if book.StartedDate != nil && book.FinishedDate == nil {
					books = append(books, book)
				}
			}

			return bookscollection.GenerateOPML(ctx, "public/opml/books/current.opml", "Horst Gutmann's current reading list", books)
		},
	}
	return cmd
}

var booksLintCmd = &cobra.Command{
	Use: "lint",
	RunE: func(cmd *cobra.Command, args []string) error {
		ctx := cmd.Context()
		log := slog.With(slog.String("component", "lint"))
		paths, err := filepath.Glob("content/reading/*.md")
		if err != nil {
			return err
		}
		numWarnings := 0
		for _, path := range paths {
			l := log.With(slog.String("file", path))
			fp, err := os.Open(path)
			if err != nil {
				return err
			}
			b, err := bookscollection.ParseBook(fp)
			fp.Close()
			if err != nil {
				l.ErrorContext(ctx, "Failed to parse book file", slog.Any("err", err))
			}
			if b.OpenLibraryID == "" {
				numWarnings++
				l.WarnContext(ctx, "No OpenLibraryID set")
			}
		}
		log.InfoContext(ctx, "%d warnings found", numWarnings)
		return nil
	},
}

func generateEnrichBookDataCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use: "enrich",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			failed := false
			for _, path := range args {
				logger := slog.With(slog.String("path", path))
				if err := bookscollection.EnrichData(cmd.Context(), path); err != nil {
					logger.ErrorContext(ctx, "Failed to enrich data", slog.Any("err", err))
					failed = true
				}
			}
			if failed {
				return fmt.Errorf("failed to enrich at least one file")
			}
			return nil
		},
	}
	return cmd
}

func generateBookStatsCommand() *cobra.Command {
	var year int
	var cmd = &cobra.Command{
		Use: "stats",
		RunE: func(cmd *cobra.Command, args []string) error {
			ctx := cmd.Context()
			var books []bookscollection.Book
			paths, err := filepath.Glob("content/reading/*.md")
			if err != nil {
				return err
			}
			for _, path := range paths {
				fp, err := os.Open(path)
				if err != nil {
					return err
				}
				book, err := bookscollection.ParseBook(fp)
				fp.Close()
				if err != nil {
					slog.ErrorContext(ctx, "Failed to parse book", slog.String("file", path), slog.Any("err", err))
				}
				if book.FinishedDate == nil || book.FinishedDate.Year() != year {
					continue
				}
				books = append(books, *book)
			}
			var totalPages int64 = 0
			genreCount := make(map[string]int)
			var maxPagesBook bookscollection.Book
			var minPagesBook bookscollection.Book
			var maxDurBook bookscollection.Book
			var minDurBook bookscollection.Book
			for _, b := range books {
				totalPages += b.Pages
				if b.Pages > maxPagesBook.Pages {
					maxPagesBook = b
				}
				if minPagesBook.Pages == 0 || (b.Pages != 0 && b.Pages < minPagesBook.Pages) {
					minPagesBook = b
				}
				bdur := b.ReadingDur()
				if bdur > maxDurBook.ReadingDur() {
					maxDurBook = b
				}
				if minDurBook.FinishedDate == nil || (bdur > time.Duration(0.0) && bdur < minDurBook.ReadingDur()) {
					minDurBook = b
				}
				if b.Genre != "" {
					c := genreCount[b.Genre]
					genreCount[b.Genre] = c + 1
				} else {
					slog.WarnContext(ctx, "No genre set", slog.String("book", b.Title))
				}
			}
			fmt.Printf("Total pages: %d\n", totalPages)
			fmt.Printf("Max pages: %d (%s)\n", maxPagesBook.Pages, maxPagesBook.Title)
			fmt.Printf("Min pages: %d (%s)\n", minPagesBook.Pages, minPagesBook.Title)
			fmt.Printf("Max duration: %s (%s)\n", humanizeDuration(maxDurBook.ReadingDur()), maxDurBook.Title)
			fmt.Printf("Min duration: %s (%s)\n", humanizeDuration(minDurBook.ReadingDur()), minDurBook.Title)
			fmt.Printf("Genres:\n")
			for k, v := range genreCount {
				fmt.Printf("  %s: %d\n", k, v)
			}
			return nil
		},
	}
	cmd.Flags().IntVar(&year, "year", 2020, "Year")
	return cmd
}

func humanizeDuration(d time.Duration) string {
	days := int(math.Floor(d.Hours() / 24))
	weeks := int(math.Floor(float64(days) / 7.0))
	hours := int(math.Floor(d.Hours())) - (days * 24)
	days = days - (weeks * 7)
	return fmt.Sprintf("%d week(s) %d day(s) %d hour(s)", weeks, days, hours)
}

func init() {
	rootCmd.AddCommand(booksCmd)
	booksCmd.AddCommand(booksLintCmd)
	booksCmd.AddCommand(generateGenOPMLCommand())
	booksCmd.AddCommand(generateEnrichBookDataCommand())
	booksCmd.AddCommand(generateBookStatsCommand())
}
