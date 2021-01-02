package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"
	"time"

	"github.com/spf13/cobra"
	"gitlab.com/zerok/zerokspot.com/pkg/bookscollection"
)

var booksCmd = &cobra.Command{
	Use: "books",
	RunE: func(cmd *cobra.Command, args []string) error {
		return nil
	},
}

var booksLintCmd = &cobra.Command{
	Use: "lint",
	RunE: func(cmd *cobra.Command, args []string) error {
		paths, err := filepath.Glob("content/reading/*.md")
		if err != nil {
			return err
		}
		for _, path := range paths {
			fp, err := os.Open(path)
			if err != nil {
				return err
			}
			_, err = bookscollection.ParseBook(fp)
			fp.Close()
			if err != nil {
				logger.Error().Err(err).Str("file", path).Msg("Failed to parse book file")
			}
		}
		return nil
	},
}

func generateBookStatsCommand() *cobra.Command {
	var year int
	var cmd = &cobra.Command{
		Use: "stats",
		RunE: func(cmd *cobra.Command, args []string) error {
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
					logger.Error().Err(err).Str("file", path).Msg("Failed to parse book file")
				}
				if book.FinishedDate == nil || book.FinishedDate.Year() != year {
					continue
				}
				books = append(books, *book)
			}
			fmt.Printf("Count: %d\n", len(books))
			totalPages := 0
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
					logger.Warn().Msgf("No genre set for %s", b.Title)
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
	booksCmd.AddCommand(generateBookStatsCommand())
}
