# File: ships_registry.gd
# Purpose: Registry singleton containing metadata for all ship types. Provides
# ship textures, scales, and names by ship_type key. Entry point for loading
# ship assets. Extensible for new ship types without code changes.
# Path: client/scripts/autoload/ships_registry.gd
# All Rights Reserved. Arc-Pub.

extends Node

const SHIPS: Dictionary = {
	"betha_1": {
		"name": "Betha-1",
		"texture": "res://assets/ships/betha_1/base.png",
		"scale": 1.0,
		"description": "Standard starter ship"
	}
}


func get_ship(ship_type: String) -> Dictionary:
	return SHIPS.get(ship_type, {})


func get_texture(ship_type: String) -> Texture2D:
	var data: Dictionary = get_ship(ship_type)
	var path: String = data.get("texture", "")
	if path.is_empty():
		return null
	return load(path)


func get_scale(ship_type: String) -> float:
	var data: Dictionary = get_ship(ship_type)
	return data.get("scale", 1.0)
