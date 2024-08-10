package component

import "encoding/json"

const (
	ClickOpenURL        = "open_url"
	ClickRunCommand     = "run_command"
	ClickSuggestCommand = "suggest_command"
	ClickCopyCommand    = "copy_command"

	HoverShowText = "show_text"
)

type TextComponent struct {
	Type          string          `json:"type"`
	Text          string          `json:"text"`
	Color         Color           `json:"color,omitempty"`
	Bold          bool            `json:"bold,omitempty"`
	Italic        bool            `json:"italic,omitempty"`
	Underlined    bool            `json:"underlined,omitempty"`
	Strikethrough bool            `json:"strikethrough,omitempty"`
	Obfuscated    bool            `json:"obfuscated,omitempty"`
	ClickEvent    *ClickEvent     `json:"clickEvent,omitempty"`
	HoverEvent    *HoverEvent     `json:"hoverEvent,omitempty"`
	Extras        []TextComponent `json:"extra,omitempty"`
}

type TranslatedTextComponent struct {
	Type      string          `json:"type"`
	Translate string          `json:"translate"`
	With      []TextComponent `json:"with,omitempty"`
}

type ClickEvent struct {
	Action ClickEventAction `json:"action"`
	Value  string           `json:"value"`
}

type HoverEvent struct {
	Action HoverEventAction `json:"action"`
	Value  string           `json:"value"`
}

type ClickEventAction string
type HoverEventAction string

func NewTextComponent(text string) *TextComponent {
	return &TextComponent{Type: "text", Text: text}
}

func NewTranslatable(text string) *TextComponent {
	return &TextComponent{Type: "translate", Text: text}
}

func (c *TextComponent) WithColor(color Color) *TextComponent {
	c.Color = color
	return c
}

func (c *TextComponent) WithBold(bold bool) *TextComponent {
	c.Bold = bold
	return c
}

func (c *TextComponent) WithItalic(italic bool) *TextComponent {
	c.Italic = italic
	return c
}

func (c *TextComponent) WithUnderlined(underlined bool) *TextComponent {
	c.Underlined = underlined
	return c
}

func (c *TextComponent) WithStrikethrough(strikethrough bool) *TextComponent {
	c.Strikethrough = strikethrough
	return c
}

func (c *TextComponent) WithObfuscate(obfuscated bool) *TextComponent {
	c.Obfuscated = obfuscated
	return c
}

func (c *TextComponent) WithClickEvent(action ClickEventAction, value string) *TextComponent {
	c.ClickEvent = newClickEvent(action, value)
	return c
}

func (c *TextComponent) WithHoverEvent(action HoverEventAction, value string) *TextComponent {
	c.HoverEvent = newHoverEvent(action, value)
	return c
}

func (c *TextComponent) WithExtras(extras ...TextComponent) *TextComponent {
	c.Extras = append(c.Extras, extras...)
	return c
}

func newClickEvent(action ClickEventAction, value string) *ClickEvent {
	return &ClickEvent{Action: action, Value: value}
}

func newHoverEvent(action HoverEventAction, value string) *HoverEvent {
	return &HoverEvent{Action: action, Value: value}
}

func (c *TextComponent) Serialize() (string, error) {
	val, err := json.Marshal(c)
	return string(val), err
}
