# File: ship.gd
# Purpose: Ship component that renders 3D model (.glb) via SubViewport for
# display in 2D game. Loads model from ShipsRegistry. Model keeps its original
# textures without faction color overlay.
# Path: client/scripts/components/ship.gd
# All Rights Reserved. Arc-Pub.

extends Node2D

var ship_type: String = "betha_1"
var faction: String = "red"

@onready var sprite: Sprite2D = $Sprite2D
@onready var sub_viewport: SubViewport = $SubViewport
@onready var model_container: Node3D = $SubViewport/ModelContainer


func _ready() -> void:
	_load_ship()


func setup(p_ship_type: String, p_faction: String) -> void:
	ship_type = p_ship_type
	faction = p_faction
	if is_node_ready():
		_load_ship()


func _load_ship() -> void:
	var model_path: String = ShipsRegistry.get_model(ship_type)
	if model_path.is_empty():
		return
	
	var model: PackedScene = load(model_path)
	if model:
		for child in model_container.get_children():
			child.queue_free()
		var instance: Node3D = model.instantiate()
		model_container.add_child(instance)
	
	sprite.texture = sub_viewport.get_texture()
