package lookup

import (
	"fmt"

	"github.com/mans82/ldoce-cli/internal/commons"
	"github.com/mans82/ldoce-cli/pkg/parser"
)

const baseURL = "https://www.ldoceonline.com"

func LookupDefault(word string) (*parser.QueryResult, error) {

	resp, err := commons.HTTPGet(baseURL + "/dictionary/" + word, false)

	if err != nil {
		return nil, fmt.Errorf("error sending request: %v", err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != 302 && resp.StatusCode != 200 {
		return nil, fmt.Errorf("error: %v", resp.Status)
	}

	if resp.StatusCode == 302 {

		spellcheckUrl := resp.Header.Get("Location")
		spellcheckResp, err := commons.HTTPGet(baseURL + spellcheckUrl, true)

		if err != nil {
			return nil, fmt.Errorf("error sending request: %v", err)
		}

		spellCheck, err := parser.ParseSpellcheck(spellcheckResp.Body)

		if err != nil {
			return nil, fmt.Errorf("error parsing spellcheck list: %v", err)
		}

		resp, err = commons.HTTPGet(spellCheck.Suggestions[0].Url, true)

		if err != nil {
			return nil, fmt.Errorf("error sending request: %v", err)
		}
	}

	entries, err := parser.ParseEntry(resp.Body)

	if err != nil {
		return nil, fmt.Errorf("error parsing entry: %v", err)
	}

	return entries, nil

}
