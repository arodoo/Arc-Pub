# File: game_data.gd
# Purpose: Singleton holding current game session data. Stores selected ship,
# faction, and map ID for use in game scene. Provides map_id based on faction
# (red→1.1, blue→2.1, green→3.1). Cleared on logout.
# Path: client/scripts/autoload/game_data.gd
# All Rights Reserved. Arc-Pub.

extends Node

const FACTION_MAPS: Dictionary = {
	"red": "1.1",
	"blue": "2.1",
	"green": "3.1"
}

var current_ship: Dictionary = {}
var faction: String = ""


func set_ship(ship: Dictionary, player_faction: String) -> void:
	current_ship = ship
	faction = player_faction


func get_map_id() -> String:
	return FACTION_MAPS.get(faction, "1.1")


func clear() -> void:
	current_ship = {}
	faction = ""
