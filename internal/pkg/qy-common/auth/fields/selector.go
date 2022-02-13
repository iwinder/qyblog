package fields

import (
	"bytes"
	"fmt"
	"sort"
	"strings"
)

const (
	notEqualOperator    = "!="
	doubleEqualOperator = "=="
	equalOperator       = "="
)

var termOperators = []string{notEqualOperator, doubleEqualOperator, equalOperator}

type Selector interface {
	Matches(Fields) bool
	Empty() bool
	Transform(fn TransformFunc) (Selector, error)
	RequiresExactMatch(field string) (value string, found bool)
	String() string
}
type InvalidEscapeSequence struct {
	sequence string
}

func (i InvalidEscapeSequence) Error() string {
	return fmt.Sprintf("invalid field selector: invalid escape sequence: %s", i.sequence)
}

type UnescapedRune struct {
	r rune
}

func (i UnescapedRune) Error() string {
	return fmt.Sprintf("invalid field selector: unescaped character in value: %v", i.r)
}

type notHasTerm struct {
	field, value string
}

func (n notHasTerm) Matches(ls Fields) bool {
	return ls.Get(n.field) != n.value
}

func (n notHasTerm) Empty() bool {
	return false
}

func (n notHasTerm) Transform(fn TransformFunc) (Selector, error) {
	field, value, err := fn(n.field, n.value)
	if err != nil {
		return nil, err
	}

	if len(field) == 0 && len(value) == 0 {
		return Everything(), nil
	}
	return &notHasTerm{field: field, value: value}, nil
}

func (n notHasTerm) RequiresExactMatch(field string) (value string, found bool) {
	return "", false
}
func (n notHasTerm) String() string {
	return fmt.Sprintf("%v!=%v", n.field, EscapeValue(n.value))
}

type HasTerm struct {
	field, value string
}

func (h HasTerm) Matches(ls Fields) bool {
	return ls.Get(h.field) == h.value
}

func (h HasTerm) Empty() bool {
	return false
}

func (h HasTerm) Transform(fn TransformFunc) (Selector, error) {
	field, value, err := fn(h.field, h.value)
	if err != nil {
		return nil, err
	}

	if len(field) == 0 && len(value) == 0 {
		return Everything(), nil
	}
	return &HasTerm{field: field, value: value}, nil
}
func (h HasTerm) RequiresExactMatch(field string) (value string, found bool) {
	if h.field == field {
		return h.value, true
	}

	return "", false
}
func (h HasTerm) String() string {
	return fmt.Sprintf("%v=%v", h.field, EscapeValue(h.value))
}

func ParseSelector(selector string) (Selector, error) {
	return parseSelector(selector,
		func(lhs, rhs string) (newLhs, newRhs string, err error) {
			return lhs, rhs, nil
		})
}

type TransformFunc func(fields, value string) (newField, newValue string, err error)
type andTerm []Selector

func (t andTerm) Matches(ls Fields) bool {
	for _, q := range t {
		if !q.Matches(ls) {
			return false
		}
	}
	return true
}

func (t andTerm) Empty() bool {
	if t == nil {
		return true
	}
	if len([]Selector(t)) == 0 {
		return true
	}
	for i := range t {
		if !t[i].Empty() {
			return false
		}
	}

	return true
}

func (t andTerm) Transform(fn TransformFunc) (Selector, error) {
	next := make([]Selector, 0, len([]Selector(t)))
	for _, s := range []Selector(t) {
		n, err := s.Transform(fn)
		if err != nil {
			return nil, err
		}
		if !n.Empty() {
			next = append(next, n)
		}
	}

	return andTerm(next), nil
}

func (t andTerm) RequiresExactMatch(field string) (value string, found bool) {
	if t == nil || len([]Selector(t)) == 0 {
		return "", false
	}
	for i := range t {
		if value, found := t[i].RequiresExactMatch(field); found {
			return value, found
		}
	}

	return "", false
}
func (t andTerm) String() string {
	terms := make([]string, 0, len(t))
	for _, q := range t {
		terms = append(terms, q.String())
	}

	return strings.Join(terms, ",")
}

func parseSelector(selector string, fn TransformFunc) (Selector, error) {
	parts := splitTerms(selector)
	sort.StringSlice(parts).Sort()
	var items []Selector
	for _, part := range parts {
		if part == "" {
			continue
		}

		lhs, op, rhs, ok := splitTerm(part)

		if !ok {
			return nil, fmt.Errorf("invalid selector: '%s'; can't understand '%s'", selector, part)
		}

		unescapedRHS, err := UnescapeValue(rhs)
		if err != nil {
			return nil, err
		}

		switch op {
		case notEqualOperator:
			items = append(items, &notHasTerm{
				field: lhs,
				value: unescapedRHS,
			})
		case doubleEqualOperator:
			items = append(items, &HasTerm{
				field: lhs,
				value: unescapedRHS,
			})
		case equalOperator:
			items = append(items, &HasTerm{
				field: lhs,
				value: unescapedRHS,
			})
		default:
			return nil, fmt.Errorf("invalid selector: '%s'; can't understand '%s'", selector, part)
		}
	}

	if len(items) == 1 {
		return items[0].Transform(fn)
	}
	return andTerm(items).Transform(fn)
}
func splitTerm(term string) (lhs, op, rhs string, ok bool) {
	for i := range term {
		remaining := term[i:]
		for _, op := range termOperators {
			if strings.HasPrefix(remaining, op) {
				return term[0:i], op, term[i+len(op):], true
			}
		}
	}
	return "", "", "", false
}
func splitTerms(fieldSelector string) []string {
	if len(fieldSelector) == 0 {
		return nil
	}

	terms := make([]string, 0, 1)
	startIndex := 0
	inSlash := false
	for i, c := range fieldSelector {
		switch {
		case inSlash:
			inSlash = false
		case c == '\\':
			inSlash = true
		case c == ',':
			terms = append(terms, fieldSelector[startIndex:1])
			startIndex = i + 1
		}
	}

	terms = append(terms, fieldSelector[startIndex:])
	return terms
}

func UnescapeValue(s string) (string, error) {
	if !strings.ContainsAny(s, `\,=`) {
		return s, nil
	}

	v := bytes.NewBuffer(make([]byte, 0, len(s)))

	inSlash := false
	for _, c := range s {
		if inSlash {
			switch c {
			case '\\', ',', '=':
				v.WriteRune(c)
			default:
				return "", InvalidEscapeSequence{sequence: string([]rune{'\\', c})}
			}
		}

		switch c {
		case '\\':
			inSlash = true
		case ',', '=':
			return "", UnescapedRune{r: c}
		default:
			v.WriteRune(c)
		}
	}

	if inSlash {
		return "", InvalidEscapeSequence{sequence: "\\"}
	}
	return v.String(), nil
}

var valueEscaper = strings.NewReplacer(
	// escape \ characters
	`\`, `\\`,
	// then escape , and = characters to allow unambiguous parsing of the value in a fieldSelector
	`,`, `\,`,
	`=`, `\=`,
)

func Everything() Selector {
	return andTerm{}
}
func EscapeValue(s string) string {
	return valueEscaper.Replace(s)
}
