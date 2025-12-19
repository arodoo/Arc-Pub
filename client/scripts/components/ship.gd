# File: ship.gd
# Purpose: Reusable ship component that loads texture from ShipsRegistry and
# applies faction color. Has visible fallback if texture unavailable. Handles
# shader uniform updates for faction color with graceful degradation.
# Path: client/scripts/components/ship.gd
# All Rights Reserved. Arc-Pub.

extends Sprite2D

const FACTION_COLORS: Dictionary = {
	"red": Color(0.9, 0.3, 0.3),
	"blue": Color(0.3, 0.5, 0.9),
	"green": Color(0.3, 0.8, 0.4)
}

var ship_type: String = "betha_1"
var faction: String = "red"


func _ready() -> void:
	_load_ship()
	_apply_faction_color()


func setup(p_ship_type: String, p_faction: String) -> void:
	ship_type = p_ship_type
	faction = p_faction
	_load_ship()
	_apply_faction_color()


func _load_ship() -> void:
	var tex: Texture2D = ShipsRegistry.get_texture(ship_type)
	if tex:
		texture = tex
		scale = Vector2.ONE * ShipsRegistry.get_scale(ship_type)
	else:
		# Fallback: create visible placeholder texture
		_create_fallback_texture()


func _create_fallback_texture() -> void:
	var img: Image = Image.create(64, 64, false, Image.FORMAT_RGBA8)
	img.fill(Color.WHITE)
	# Draw simple ship shape
	for y in range(64):
		for x in range(64):
			var dx: int = x - 32
			var dy: int = y - 32
			if abs(dx) < 20 - abs(dy) / 2:
				img.set_pixel(x, y, Color(0.9, 0.9, 0.9))
	texture = ImageTexture.create_from_image(img)
	scale = Vector2(2, 2)


func _apply_faction_color() -> void:
	var color: Color = FACTION_COLORS.get(faction, Color.WHITE)
	if material and material is ShaderMaterial:
		(material as ShaderMaterial).set_shader_parameter("faction_color", color)
	else:
		modulate = color
