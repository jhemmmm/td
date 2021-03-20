package entity

import (
	"io"
	"net/url"
	"strconv"
	"strings"

	"golang.org/x/net/html"
	"golang.org/x/xerrors"
)

type stackElem struct {
	offset int
	tag    string
	format Formatter
}

type htmlParser struct {
	tokenizer *html.Tokenizer
	builder   *Builder
	stack     []stackElem
	attr      map[string]string
}

func (p *htmlParser) fillAttrs() {
	// Clear old attrs.
	for k := range p.attr {
		delete(p.attr, k)
	}

	// Fill with new attributes.
	for {
		key, value, ok := p.tokenizer.TagAttr()
		p.attr[string(key)] = string(value)
		if !ok {
			break
		}
	}
}

func (p *htmlParser) startTag() error {
	const pre = "pre"

	var e stackElem
	tn, hasAttr := p.tokenizer.TagName()
	e.tag = string(tn)
	if hasAttr {
		p.fillAttrs()
	}

	e.offset = p.builder.message.Len()
	// See https://core.telegram.org/bots/api#html-style.
	switch e.tag {
	case "b", "strong":
		e.format = Bold()
	case "i", "em":
		e.format = Italic()
	case "u", "ins":
		e.format = Underline()
	case "s", "strike", "del":
		e.format = Strike()
	case "a":
		href, ok := p.attr["href"]
		if !ok {
			return xerrors.Errorf("tag %q must have attribute href", e.tag)
		}

		u, err := url.Parse(href)
		if err != nil {
			return xerrors.Errorf("href must be a valid URL, got %q", href)
		}

		if u.Scheme == "tg" && u.Host == "user" {
			id, err := strconv.Atoi(u.Query().Get("id"))
			if err != nil {
				return xerrors.Errorf("invalid user ID %q: %w", id, err)
			}

			e.format = MentionName(id)
		} else {
			e.format = TextURL(href)
		}
	case "code":
		e.format = Code()

		// BotAPI docs says:
		// > Use nested <pre> and <code> tags, to define programming language for <pre> entity.
		if len(p.stack) > 0 && p.stack[len(p.stack)-1].tag == pre {
			lang, ok := p.attr["class"]
			if ok {
				e.format = Pre(strings.TrimPrefix(lang, "language-"))
			}
		}
	case pre:
		e.format = Code()
	default:
		return xerrors.Errorf("unknown tag name %q", e.tag)
	}

	p.stack = append(p.stack, e)
	return nil
}

func (p *htmlParser) endTag() error {
	tn, _ := p.tokenizer.TagName()
	if len(p.stack) == 0 {
		return xerrors.Errorf("unexpected end tag %q", string(tn))
	}

	var s stackElem
	// Pop from SliceTricks.
	s, p.stack = p.stack[len(p.stack)-1], p.stack[:len(p.stack)-1]
	if s.tag != string(tn) {
		return xerrors.Errorf("expected tag %q, got %q", s.tag, string(tn))
	}

	length := ComputeLength(p.builder.message.String())
	p.builder.entities = append(p.builder.entities, s.format(s.offset, length-s.offset))
	return nil
}

func (p *htmlParser) parse() error {
	for {
		tt := p.tokenizer.Next()
		switch tt {
		case html.ErrorToken:
			if err := p.tokenizer.Err(); !xerrors.Is(err, io.EOF) {
				return err
			}
			return nil
		case html.TextToken:
			p.builder.message.Write(p.tokenizer.Text())
		case html.StartTagToken:
			if err := p.startTag(); err != nil {
				return err
			}
		case html.EndTagToken:
			if err := p.endTag(); err != nil {
				return err
			}
		}
	}
}

// HTML parses given input from reader and adds parsed entities to given builder.
//
// See https://core.telegram.org/bots/api#html-style.
func HTML(r io.Reader, b *Builder) error {
	p := htmlParser{
		tokenizer: html.NewTokenizer(r),
		builder:   b,
		attr:      map[string]string{},
	}

	return p.parse()
}
