package detector

import (
	"fmt"
	"strings"
)

// hyvaThemePath is the theme registered by hyva-themes/magento2-default-theme.
const hyvaThemePath = "Hyva/default"

// themeIDConfigPath is the core_config_data path holding the default theme.
const themeIDConfigPath = "design/theme/theme_id"

// themeIDQuery returns the SQL that looks up a theme's numeric ID.
func themeIDQuery(themePath string) string {
	return fmt.Sprintf("SELECT theme_id FROM theme WHERE theme_path='%s'", themePath)
}

// parseThemeID extracts the theme ID from a mysql -N query result. An empty
// or non-numeric result must fail the install: writing it to
// design/theme/theme_id would silently leave the store on Luma.
func parseThemeID(out, themePath string) (string, error) {
	id := strings.TrimSpace(out)
	if id == "" {
		return "", fmt.Errorf("theme %q not found in database", themePath)
	}
	if !isNumeric(id) {
		return "", fmt.Errorf("unexpected theme ID %q for theme %q", id, themePath)
	}
	return id, nil
}

func isNumeric(s string) bool {
	for _, r := range s {
		if r < '0' || r > '9' {
			return false
		}
	}
	return true
}
