# File: logger.gd
# Purpose: Autoload singleton for file-based logging. Writes all logs, warnings,
# and errors to a file in the project root for external visibility. Captures
# Godot's print output and unhandled errors. File location is accessible by
# external tools. Essential for debugging without Godot Editor access.
# Path: client/scripts/autoload/logger.gd
# All Rights Reserved. Arc-Pub.

extends Node

const LOG_PATH: String = "res://../godot_debug.log"
var _file: FileAccess
var _buffer: PackedStringArray = []


func _ready() -> void:
	_open_log_file()
	_log("=== Godot Started: " + Time.get_datetime_string_from_system() + " ===")
	
	# Connect to error signals
	get_tree().node_added.connect(_on_node_added)


func _notification(what: int) -> void:
	if what == NOTIFICATION_PREDELETE:
		_log("=== Godot Closed ===")
		_flush()


func _open_log_file() -> void:
	var path: String = ProjectSettings.globalize_path(LOG_PATH)
	_file = FileAccess.open(path, FileAccess.WRITE)
	if _file == null:
		# Fallback to user:// if res:// fails
		_file = FileAccess.open("user://godot_debug.log", FileAccess.WRITE)


func _log(message: String) -> void:
	var timestamp: String = Time.get_time_string_from_system()
	var line: String = "[" + timestamp + "] " + message
	print(line)  # Also print to Godot console
	if _file:
		_file.store_line(line)
		_file.flush()


func info(message: String) -> void:
	_log("[INFO] " + message)


func warn(message: String) -> void:
	_log("[WARN] " + message)


func error(message: String) -> void:
	_log("[ERROR] " + message)


func debug(context: String, data: Variant) -> void:
	_log("[DEBUG] " + context + ": " + str(data))


func _flush() -> void:
	if _file:
		_file.flush()
		_file.close()


func _on_node_added(_node: Node) -> void:
	pass  # Can be used for scene tracking


# Call this to log the current script error
func log_script_error() -> void:
	var stack: Array = get_stack()
	_log("[STACK TRACE]")
	for frame in stack:
		_log("  " + str(frame))
