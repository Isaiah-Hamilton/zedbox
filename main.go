package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strings"
)

type Style string

const (
	Material Style = "material"
	Mix      Style = "mix"
	Original Style = "original"
)

type Strength string

const (
	Soft   Strength = "soft"
	Medium Strength = "medium"
	Hard   Strength = "hard"
)

type Mode string

const (
	Dark  Mode = "dark"
	Light Mode = "light"
)

type Variant struct {
	Style    Style
	Strength Strength
	Mode     Mode
}

type Colors struct {
	BgDim          string `json:"bg_dim"`
	Bg0            string `json:"bg0"`
	Bg1            string `json:"bg1"`
	Bg2            string `json:"bg2"`
	Bg3            string `json:"bg3"`
	Bg4            string `json:"bg4"`
	Bg5            string `json:"bg5"`
	BgStatusline1  string `json:"bg_statusline1"`
	BgStatusline2  string `json:"bg_statusline2"`
	BgStatusline3  string `json:"bg_statusline3"`
	BgVisualRed    string `json:"bg_visual_red"`
	BgVisualYellow string `json:"bg_visual_yellow"`
	BgVisualGreen  string `json:"bg_visual_green"`
	BgVisualBlue   string `json:"bg_visual_blue"`
	BgVisualPurple string `json:"bg_visual_purple"`
	BgDiffRed      string `json:"bg_diff_red"`
	BgDiffGreen    string `json:"bg_diff_green"`
	BgDiffBlue     string `json:"bg_diff_blue"`
	BgCurrentWord  string `json:"bg_current_word"`
	Fg0            string `json:"fg0"`
	Fg1            string `json:"fg1"`
	Red            string `json:"red"`
	Orange         string `json:"orange"`
	Yellow         string `json:"yellow"`
	Green          string `json:"green"`
	Aqua           string `json:"aqua"`
	Blue           string `json:"blue"`
	Purple         string `json:"purple"`
	BgRed          string `json:"bg_red"`
	BgGreen        string `json:"bg_green"`
	BgYellow       string `json:"bg_yellow"`
	Grey0          string `json:"grey0"`
	Grey1          string `json:"grey1"`
	Grey2          string `json:"grey2"`
}

func main() {
	variant := Variant{
		Style:    Material,
		Strength: Hard,
		Mode:     Dark,
	}

	colors, err := readColors(variant)
	if err != nil {
		log.Fatalf("Error loading colors: %v", err)
	}

	scheme, err := os.ReadFile("./src/scheme.json")
	if err != nil {
		log.Fatalf("Error reading scheme.json: %v", err)
	}

	processed := string(scheme)

	// Replace color placeholders
	for key, value := range colors.Iter() {
		placeholder := fmt.Sprintf("{{%s}}", key)
		processed = strings.ReplaceAll(processed, placeholder, value)
	}

	// Replace variant placeholders
	processed = strings.ReplaceAll(processed, "{{style}}", strings.ToUpper(string(variant.Style))[:1]+string(variant.Style)[1:])
	processed = strings.ReplaceAll(processed, "{{strength}}", strings.ToUpper(string(variant.Strength))[:1]+string(variant.Strength)[1:])
	processed = strings.ReplaceAll(processed, "{{Mode}}", strings.ToUpper(string(variant.Mode))[:1]+string(variant.Mode)[1:])
	processed = strings.ReplaceAll(processed, "{{mode}}", string(variant.Mode))

	var themes any
	if err := json.Unmarshal([]byte(processed), &themes); err != nil {
		log.Fatalf("Invalid JSON after substitution: %v", err)
	}

	output := map[string]any{
		"$schema": "https://zed.dev/schema/themes/v0.2.0.json",
		"name":    "Zedbox",
		"author":  "isaiah hamilton <isaiah-hamilton@gmail.com>",
		"themes":  []any{themes},
	}

	// Marshal with pretty printing
	pretty, err := MarshalIndentNoEscape(output, "", "  ")
	if err != nil {
		log.Fatalf("Error marshaling JSON: %v", err)
	}

	filePath := "./themes/zedbox.json"
	if err := os.WriteFile(filePath, pretty, 0644); err != nil {
		log.Fatalf("Error writing file: %v", err)
	}

	fmt.Printf("'%s' has been created or overwritten successfully.", filePath)
}

func readColors(variant Variant) (*Colors, error) {
	fileName := fmt.Sprintf("./src/colors/%s_%s_%s.json",
		string(variant.Style),
		string(variant.Strength),
		string(variant.Mode),
	)

	data, err := os.ReadFile(fileName)
	if err != nil {
		return nil, fmt.Errorf("failed to read colors file %s: %w", fileName, err)
	}

	var colors Colors
	if err := json.Unmarshal(data, &colors); err != nil {
		return nil, fmt.Errorf("failed to parse colors JSON: %w", err)
	}

	return &colors, nil
}

func MarshalIndentNoEscape(v any, prefix, indent string) ([]byte, error) {
	var buf bytes.Buffer
	enc := json.NewEncoder(&buf)
	enc.SetEscapeHTML(false)
	enc.SetIndent(prefix, indent)
	err := enc.Encode(v)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (c *Colors) Iter() map[string]string {
	return map[string]string{
		"bg_dim":           c.BgDim,
		"bg0":              c.Bg0,
		"bg1":              c.Bg1,
		"bg2":              c.Bg2,
		"bg3":              c.Bg3,
		"bg4":              c.Bg4,
		"bg5":              c.Bg5,
		"bg_statusline1":   c.BgStatusline1,
		"bg_statusline2":   c.BgStatusline2,
		"bg_statusline3":   c.BgStatusline3,
		"bg_visual_red":    c.BgVisualRed,
		"bg_visual_yellow": c.BgVisualYellow,
		"bg_visual_green":  c.BgVisualGreen,
		"bg_visual_blue":   c.BgVisualBlue,
		"bg_visual_purple": c.BgVisualPurple,
		"bg_diff_red":      c.BgDiffRed,
		"bg_diff_green":    c.BgDiffGreen,
		"bg_diff_blue":     c.BgDiffBlue,
		"bg_current_word":  c.BgCurrentWord,
		"fg0":              c.Fg0,
		"fg1":              c.Fg1,
		"red":              c.Red,
		"orange":           c.Orange,
		"yellow":           c.Yellow,
		"green":            c.Green,
		"aqua":             c.Aqua,
		"blue":             c.Blue,
		"purple":           c.Purple,
		"bg_red":           c.BgRed,
		"bg_green":         c.BgGreen,
		"bg_yellow":        c.BgYellow,
		"grey0":            c.Grey0,
		"grey1":            c.Grey1,
		"grey2":            c.Grey2,
	}
}
