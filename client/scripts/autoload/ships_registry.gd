# File: ships_registry.gd
# Purpose: Registry singleton containing metadata for all ship types. Provides
# ship models, scales, and names by ship_type key. Extensible for new types.
# Path: client/scripts/autoload/ships_registry.gd
# All Rights Reserved. Arc-Pub.

extends Node

const SHIPS: Dictionary = {
	"betha_1": {
		"name": "Betha-1",
		"model": "res://assets/ships/betha_1/ship.glb",
		"scale": 1.0,
		"description": "Standard starter ship"
	}
}


func get_ship(ship_type: String) -> Dictionary:
	return SHIPS.get(ship_type, {})


func get_model(ship_type: String) -> String:
	var data: Dictionary = get_ship(ship_type)
	return data.get("model", "")


func get_scale(ship_type: String) -> float:
	var data: Dictionary = get_ship(ship_type)
	return data.get("scale", 1.0)
