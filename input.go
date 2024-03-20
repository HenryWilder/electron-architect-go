package main

import rl "github.com/gen2brain/raylib-go/raylib"

type Input int32

const (
	CreateNode Input = iota
	ConnectNodes
	RemoveHovered
)

type InputBindTag uint8

const (
	// Keyboard key binding
	KeyBind InputBindTag = iota

	// Mouse button binding
	BtnBind

	// Scroll wheel binding
	WhlBind
)

type InputBind struct {
	// KeyBind => KeyVal
	// BtnBind => BtnVal
	// WhlBind => WhlVal
	Tag InputBindTag

	// KeyboardKey
	KeyVal int32

	// MouseButton
	BtnVal int32

	// true = positive; false = negative
	WhlVal bool
}

func KeyboardKeyBinding(key int32) InputBind {
	return InputBind{Tag: KeyBind, KeyVal: key}
}
func MouseButtonBinding(btn int32) InputBind {
	return InputBind{Tag: BtnBind, BtnVal: btn}
}
func ScrollWheelBinding(dir bool) InputBind {
	return InputBind{Tag: WhlBind, WhlVal: dir}
}

type InputHandler struct {
	Bindings map[Input]InputBind
}

func (handler *InputHandler) SetBinding(input Input, binding InputBind) {
	handler.Bindings[input] = binding
}

func (handler *InputHandler) SetBindingsToDefaults() {
	handler.SetBinding(CreateNode, MouseButtonBinding(rl.MouseButtonLeft))
	handler.SetBinding(ConnectNodes, MouseButtonBinding(rl.MouseButtonLeft))
	handler.SetBinding(RemoveHovered, MouseButtonBinding(rl.MouseButtonRight))
}

// Makes the input map and inits to defaults
func (handler *InputHandler) Init() {
	handler.Bindings = make(map[Input]InputBind)
	handler.SetBindingsToDefaults()
}

func (handler InputHandler) IsInput(inp Input, state string) bool {
	bind := handler.Bindings[inp]

	switch bind.Tag {

	case KeyBind:
		switch state {
		case "Pressed":
			return rl.IsKeyPressed(bind.KeyVal)
		case "Down":
			return rl.IsKeyDown(bind.KeyVal)
		case "Released":
			return rl.IsKeyReleased(bind.KeyVal)
		case "Up":
			return rl.IsKeyUp(bind.KeyVal)
		default:
			panic("Input state not specialized for keyboard key")
		}

	case BtnBind:
		switch state {
		case "Pressed":
			return rl.IsMouseButtonPressed(bind.BtnVal)
		case "Down":
			return rl.IsMouseButtonDown(bind.BtnVal)
		case "Released":
			return rl.IsMouseButtonReleased(bind.BtnVal)
		case "Up":
			return rl.IsMouseButtonUp(bind.BtnVal)
		default:
			panic("Input state not specialized for mouse button")
		}

	case WhlBind:
		wheelDir := rl.GetMouseWheelMove()
		if bind.WhlVal {
			return wheelDir > 0.0
		} else {
			return wheelDir < 0.0
		}

	default:
		panic("Binding tag not specialized")
	}
}

func (handler InputHandler) IsPressed(inp Input) bool {
	return handler.IsInput(inp, "Pressed")
}
func (handler InputHandler) IsDown(inp Input) bool {
	return handler.IsInput(inp, "Down")
}
func (handler InputHandler) IsReleased(inp Input) bool {
	return handler.IsInput(inp, "Released")
}
func (handler InputHandler) IsUp(inp Input) bool {
	return handler.IsInput(inp, "Up")
}
